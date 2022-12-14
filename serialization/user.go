package serialization

import "message_boards/model"

type User struct {
	ID       uint   `json:"id"`
	UserName string `json:"user_name"`
	NickName string `json:"nick_name"`
	Email    string `json:"email"`
	Status   string `json:"status"`
	CreateAt int64  `json:"create_at"`
}

// BuildUser 序列化用户
func BuildUser(user *model.User) User {
	return User{
		ID:       user.ID,
		UserName: user.UserName,
		NickName: user.NickName,
		Email:    user.Email,
		CreateAt: user.CreatedAt.Unix(),
	}
}

func BuildUsers(items []*model.User) (users []User) {
	for _, item := range items {
		user := BuildUser(item)
		users = append(users, user)
	}
	return users
}
