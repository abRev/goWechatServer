package routers

import (
	"github.com/gin-gonic/gin"
	"wechat/routers/home"
	"wechat/routers/user"
	"wechat/routers/wechat"
)

func InitRouters(router *gin.Engine) {
	wechat.InitRouters(router)
	home.InitRouters(router)
	user.InitRouters(router)
}
