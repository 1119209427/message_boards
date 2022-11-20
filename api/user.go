package api

import (
	"github.com/gin-gonic/gin"
	"message_boards/pkg/util"
	"message_boards/service"
	"net/http"
)

func UserRegister(ctx *gin.Context) {
	var UserService service.UserRegisterService
	err := ctx.ShouldBind(&UserService)
	if err != nil {
		ctx.JSON(http.StatusOK, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	} else {
		rep := UserService.Register(ctx.Request.Context())
		ctx.JSON(http.StatusOK, rep)
	}
}

func UserLogin(ctx *gin.Context) {
	var UserService service.UserLoginService
	err := ctx.ShouldBind(&UserService)
	if err != nil {
		ctx.JSON(http.StatusOK, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	} else {
		rep := UserService.Login(ctx.Request.Context())
		ctx.JSON(http.StatusOK, rep)
	}
}
