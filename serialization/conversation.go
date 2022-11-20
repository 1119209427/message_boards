package serialization

import (
	"message_boards/model"
)

type Conversations struct {
	ID      uint   ` json:"id"`
	UId     uint   `json:"uid"`     //发话题用户id
	Title   string `json:"title"`   //帖子的标题
	Content string `json:"content"` //贴子的内容
	Author  User   `json:"author"`
}

// BuildConversations 序列化话题
func BuildConversations(conversation *model.Conversation, user *model.User) Conversations {
	return Conversations{
		ID:      conversation.Model.ID,
		UId:     conversation.UId,
		Title:   conversation.Title,
		Content: conversation.Content,
		Author: User{
			ID:       user.ID,
			UserName: user.UserName,
			NickName: user.NickName,
		},
	}
}
