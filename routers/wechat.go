package routers

import (
	"wechat/controller/wechat"

	"github.com/gin-gonic/gin"
)

func InitWechatRouters(router *gin.Engine) {
	// wechatRouter := router.Group("/official")
	router.Any("/official", wechat.Hello)
}
