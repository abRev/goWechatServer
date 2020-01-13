package routers

import (
	"github.com/gin-gonic/gin"
	"wechat/routers/wechat"
	"wechat/routers/home"
)

func InitRouters (router *gin.Engine) {
	wechat.InitRouters(router)
	home.InitRouters(router)
}
