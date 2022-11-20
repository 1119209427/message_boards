package api

import (
	"github.com/gin-gonic/gin"
	"message_boards/pkg/util"
	"message_boards/service"
	"net/http"
	"strconv"
)

func ConversationPublish(ctx *gin.Context) {
	var ConversationService service.ConversationPublishService
	userIdStr := ctx.GetString("userId")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		util.LogrusObj.Infoln(err)
	}
	if err = ctx.ShouldBind(&ConversationService); err == nil {
		res := ConversationService.PublishService(ctx, uint(userId))
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusOK, ErrorResponse(err))
	}
}

// ConversationsGet 获取话题的一部分内容和标题，一次5条
func ConversationsGet(ctx *gin.Context) {
	var ConversationService service.ConversationsGetService
	if err := ctx.ShouldBind(&ConversationService); err == nil {
		res := ConversationService.GetManyService(ctx.Request.Context())
		ctx.JSON(http.StatusOK, res)
	} else {
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusOK, ErrorResponse(err))
	}
}
