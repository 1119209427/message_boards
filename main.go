package main

import (
	"github.com/gin-gonic/gin"
	"message_boards/conf"
	"message_boards/pkg/util"
	"message_boards/route"
)

func main() {
	util.InitFilter()
	conf.Init()
	r := gin.Default()
	route.InitRoute(r)
	_ = r.Run(conf.HttpPort)
}
