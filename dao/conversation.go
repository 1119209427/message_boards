package dao

import (
	"context"
	"gorm.io/gorm"
	"message_boards/config"
	"message_boards/model"
	"time"
)

type ConversationDao struct {
	*gorm.DB
}

func NewConversationDao(ctx context.Context) *ConversationDao {
	return &ConversationDao{NewDBClient(ctx)}
}

func NewConversationDB(db *gorm.DB) *ConversationDao {
	return &ConversationDao{db}
}

func (c *ConversationDao) InsertConversation(conversation model.Conversation) error {
	return c.DB.Model(&model.Conversation{}).Create(&conversation).Error

}

// GetConversations 根据时间戳获得话题
func (c *ConversationDao) GetConversations(lastTime time.Time) ([]model.Conversation, error) {
	conversation := make([]model.Conversation, config.ConversationNum)
	return conversation, c.DB.Where("publish_time <", lastTime).Order("publish_time desc").Limit(config.ConversationNum).Find(&conversation).Error
}
