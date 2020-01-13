package wechat

import (
	"github.com/gin-gonic/gin"
	"wechat/controller/wechat"
)

func InitRouters(router *gin.Engine) {
	wechatRouter := router.Group("/official")
	wechatRouter.Any("/",wechat.Hello)
}
