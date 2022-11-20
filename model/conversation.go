package model

import (
	"gorm.io/gorm"
	"time"
)

type Conversation struct {
	gorm.Model
	UId         uint      //发话题用户id
	Title       string    //帖子的标题
	Content     string    //贴子的内容
	PublishTime time.Time //发布的时间
}
