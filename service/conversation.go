package service

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
	"message_boards/config"
	"message_boards/dao"
	"message_boards/model"
	e2 "message_boards/pkg/e"
	util2 "message_boards/pkg/util"
	"message_boards/serialization"
	"strconv"
	"time"
)

type ConversationPublishService struct {
	Title   string `json:"title" form:"title"`     //帖子的标题
	Content string `json:"content" form:"content"` //贴子的内容
	Cancel  string `json:"cancel" form:"cancel"`   //取消发布未0，发布为1
}

type ConversationsGetService struct {
	LastTime string `json:"last_time" form:"last_time"` //获取拉取话题的时间戳
}

// GetManyService 获取话题
func (c *ConversationsGetService) GetManyService(ctx context.Context) *serialization.Response {
	code := e2.SUCCESS
	conversationDao := dao.NewConversationDao(ctx)
	log.Printf("传入的时间" + c.LastTime)
	var lastTime time.Time
	if c.LastTime != "0" {
		me, err := strconv.ParseInt(c.LastTime, 10, 64)
		if err != nil {
			util2.LogrusObj.Infoln(err)
			code = e2.ParseError
			return &serialization.Response{
				Status: code,
				Msg:    e2.GetMsg(code),
			}
		}
		lastTime = time.Unix(me, 0)
	} else {
		lastTime = time.Now()
	}
	log.Printf("获取到时间戳%v", lastTime)

	//获取话题，进行组装(获取作者信息)
	conversations, err := conversationDao.GetConversations(lastTime)
	if err != nil {
		util2.LogrusObj.Infoln(err)
		code = e2.ErrorDatabase
		return &serialization.Response{
			Status: code,
			Msg:    e2.GetMsg(code),
		}
	}
	var userService GetUserService
	cons := make([]serialization.Conversations, config.ConversationNum)
	for _, item := range conversations {
		err, user := userService.GetUser(ctx, item.UId)
		if err != nil {
			util2.LogrusObj.Infoln(err)
			//如果数据库出错
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				code = e2.ErrorDatabase
				return &serialization.Response{
					Status: code,
					Msg:    e2.GetMsg(code),
				}
			}
			con := serialization.BuildConversations(&item, nil)
			cons = append(cons, con)
		} else {
			con := serialization.BuildConversations(&item, &user)
			cons = append(cons, con)
		}
	}
	return &serialization.Response{
		Status: code,
		Data:   cons,
		Msg:    e2.GetMsg(code),
	}

}

func (c *ConversationPublishService) PublishService(ctx context.Context, Uid uint) *serialization.Response {
	code := e2.SUCCESS
	conversationDao := dao.NewConversationDao(ctx)
	//检查内容是否合规
	ok, sensitive := util2.Filter.Validate(c.Title)
	if !ok {
		code = e2.ErrorSensitive
		return &serialization.Response{
			Status: code,
			Data:   fmt.Sprintf("敏感词:%s 请删除后再发表内容", sensitive),
			Msg:    e2.GetMsg(code),
		}
	}
	ok, sensitive = util2.Filter.Validate(c.Content)
	if !ok {
		code = e2.ErrorSensitive
		return &serialization.Response{
			Status: code,
			Data:   fmt.Sprintf("敏感词:%s 请删除后再发表内容", sensitive),
			Msg:    e2.GetMsg(code),
		}
	}
	publishTime := time.Now()
	var conversation model.Conversation
	conversation.UId = Uid
	conversation.Title = c.Title
	conversation.Content = c.Content
	conversation.PublishTime = publishTime
	flag, err := strconv.Atoi(c.Cancel)
	if err != nil {
		code = e2.ParseError
		util2.LogrusObj.Infoln(err)
		return &serialization.Response{
			Status: code,
			Msg:    e2.GetMsg(code),
		}
	}
	if flag == config.NOTPublish {
		return &serialization.Response{
			Status: e2.ErrorNotPublish,
			Msg:    e2.GetMsg(code),
		}
	}
	err = conversationDao.InsertConversation(conversation)
	if err != nil {
		util2.LogrusObj.Infoln(err)
		code = e2.ErrorDatabase
		return &serialization.Response{
			Status: code,
			Msg:    e2.GetMsg(code),
		}
	}
	return &serialization.Response{
		Status: code,
		Data:   "发表评论成功",
		Msg:    e2.GetMsg(code),
	}
}
