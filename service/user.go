package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
	"message_boards/config"
	"message_boards/dao"
	"message_boards/model"
	e2 "message_boards/pkg/e"
	"message_boards/pkg/util"
	"message_boards/serialization"
	"strconv"
	"time"
)

type UserRegisterService struct {
	UserName string `json:"user_name" form:"user_name" binding:"required"`
	PassWord string `json:"pass_word" form:"pass_word" binding:"required"`
	NickName string `json:"nick_name" form:"nick_name"`
}

type UserLoginService struct {
	UserName string `json:"user_name" form:"user_name" binding:"required"`
	PassWord string `json:"pass_word" form:"pass_word" binding:"required"`
}

type GetUserService struct {
}

// GetUser 通过id获得用户信息
func (u *GetUserService) GetUser(ctx context.Context, uid uint) (error, model.User) {
	userDao := dao.NewUserDao(ctx)

	err, user := userDao.GetUserByUid(uid)
	if err != nil {
		util.LogrusObj.Infoln(err)
		return err, user
	}

	return nil, user
}

func (u *UserLoginService) Login(ctx context.Context) *serialization.Response {
	userDao := dao.NewUserDao(ctx)
	code := e2.SUCCESS
	err, user := userDao.GetUserByUserName(u.UserName)
	if err != nil {
		util.LogrusObj.Infoln(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			code = e2.ErrorUserNotFound
			return &serialization.Response{
				Status: code,
				Data:   nil,
				Msg:    e2.GetMsg(code),
				Error:  "",
			}
		}
		code = e2.ErrorDatabase
		return &serialization.Response{
			Status: code,
			Msg:    e2.GetMsg(code),
		}
	}
	//检查密码
	flag := user.CheckPass(u.PassWord)
	if !flag {
		code = e2.ErrorNotCompare
		return &serialization.Response{
			Status: code,
			Msg:    e2.GetMsg(code),
		}
	}
	//生成token
	token := GenerateToken(user)
	return &serialization.Response{
		Status: code,
		Msg:    e2.GetMsg(code),
		Data:   serialization.TokenData{User: serialization.BuildUser(&user), Token: token},
	}

}

func (u *UserRegisterService) Register(ctx context.Context) *serialization.Response {
	userDao := dao.NewUserDao(ctx)
	var user model.User
	code := e2.SUCCESS
	err, flag := userDao.CheckUser(u.UserName)
	if err != nil {

		util.LogrusObj.Infoln(err)
		code = e2.ErrorDatabase
		return &serialization.Response{
			Status: code,
			Msg:    e2.GetMsg(code),
		}
	}
	if flag {
		code = e2.ErrorExistUser
		return &serialization.Response{
			Status: code,
			Msg:    e2.GetMsg(code),
		}
	}
	err = user.EncryptPass(u.PassWord)
	if err != nil {
		util.LogrusObj.Infoln(err)
		code = e2.ErrorFailEncryption
		return &serialization.Response{
			Status: code,
			Msg:    e2.GetMsg(code),
		}
	}
	user.UserName = u.UserName
	user.NickName = u.NickName
	err = userDao.InsertUser(user)
	if err != nil {
		util.LogrusObj.Infoln(err)
		code = e2.ErrorDatabase
		return &serialization.Response{
			Status: code,
			Msg:    e2.GetMsg(code),
		}
	}
	return &serialization.Response{
		Status: code,
		Msg:    e2.GetMsg(code),
		Data:   "注册成功",
	}
}

// GenerateToken 根据username生成一个token
func GenerateToken(user model.User) string {
	fmt.Printf("generatetoken %v", user)
	token := NewToken(user)
	println(token)
	return token
}

// NewToken 生成token
func NewToken(u model.User) string {
	expiresTime := time.Now().Unix() + int64(config.OneDayOfHours)
	fmt.Printf("expiresTime %d \n", expiresTime)
	id := u.ID
	fmt.Printf("id %d \n", id)
	claim := jwt.StandardClaims{
		Audience:  u.UserName,
		ExpiresAt: expiresTime,
		Id:        strconv.FormatUint(uint64(id), 10),
		IssuedAt:  time.Now().Unix(),
		Issuer:    "zhen_xi",
		NotBefore: time.Now().Unix(),
		Subject:   "token",
	}
	var jwtSecret = []byte(config.Secret)
	tokenClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	if token, err := tokenClaim.SignedString(jwtSecret); err == nil {
		token = "Bearer " + token
		println("generate token success \n")
		return token
	} else {
		util.LogrusObj.Infoln(err)
		return "fail"
	}
}
