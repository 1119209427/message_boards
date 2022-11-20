package route

import (
	"github.com/gin-gonic/gin"
	"message_boards/api"
	"message_boards/middlerware/jwt"
)

func InitRoute(r *gin.Engine) {

	apiRoute := r.Group("/message_boards")
	//User baseApi
	apiRoute.POST("/user/register/", api.UserRegister)
	apiRoute.POST("/user/login/", api.UserLogin)

	//conversation baseApi
	apiRoute.POST("/conversation/publish/", jwt.Auth(), api.ConversationPublish)
	apiRoute.GET("/conversation/get/", jwt.AuthWithoutLogin(), api.ConversationsGet) //获取话题，限制5条

	//comment baseApi
	apiRoute.POST("/comment/publish/", api.CommentPublish)

}
