package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"message_boards/pkg/util"
)

var (
	encryptHard = 12
)

type User struct {
	gorm.Model
	UserName string `gorm:"index"`
	NickName string
	PassWord string
	Avatar   string
	Email    string
	Phone    string
}

func (u *User) EncryptPass(passWord string) error {
	pass, err := bcrypt.GenerateFromPassword([]byte(passWord), encryptHard)
	if err != nil {
		util.LogrusObj.Infoln(err)
		return err
	}
	u.PassWord = string(pass)
	return nil
}
func (u *User) CheckPass(passWord string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PassWord), []byte(passWord))
	return err == nil
}
