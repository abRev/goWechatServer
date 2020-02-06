package routers

import (
	"github.com/gin-gonic/gin"
	"wechat/middleware/jwt"
	"wechat/routers/home"
	"wechat/routers/login"
	"wechat/routers/user"
	"wechat/routers/wechat"
)

func InitRouters(router *gin.Engine) {
	login.InitRouter(router)
	wechat.InitRouters(router)
	user.InitRouters(router)
	home.InitRouters(router, jwt.JWTAuth())
}
