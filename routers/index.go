package routers

import (
	"github.com/gin-gonic/gin"
	"wechat/routers/wechat"
)

func InitRouters (router *gin.Engine) {
	wechat.InitRouters(router)
}
