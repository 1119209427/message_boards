package dao

import (
	"context"
	"gorm.io/gorm"
	"message_boards/model"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{NewDBClient(ctx)}
}

func NewUserDaoByDB(db *gorm.DB) *UserDao {
	return &UserDao{db}
}

func (u *UserDao) GetUserByUserName(userName string) (error, model.User) {
	var user model.User
	if err := u.DB.Model(&model.User{}).Where("user_name", userName).First(&user).Error; err != nil {
		return err, user
	}
	return nil, user
}

func (u *UserDao) GetUserByUid(uid uint) (error, model.User) {
	var user model.User
	if err := u.DB.Model(&model.User{}).Where("id", uid).First(&user).Error; err != nil {
		return err, user
	}
	return nil, user
}

// CheckUser 用户名是否重复，重复返回true，不重复返回false
func (u *UserDao) CheckUser(userName string) (error, bool) {
	var cnt int64
	err := u.DB.Model(&model.User{}).Where("user_name", userName).Count(&cnt).Error
	if cnt == 0 {
		return err, false
	}
	return err, true
}

// InsertUser 将用户存入数据库
func (u *UserDao) InsertUser(user model.User) error {
	return u.DB.Model(&model.User{}).Create(&user).Error
}
