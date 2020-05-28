package routers

import (
	"wechat/middleware"
	"wechat/middleware/jwt"
	"wechat/routers/home"
	"wechat/routers/login"
	"wechat/routers/search"
	"wechat/routers/user"
	"wechat/routers/wechat"

	"github.com/gin-gonic/gin"
)

func InitRouters(router *gin.Engine) {
	login.InitRouter(router)
	wechat.InitRouters(router)
	user.InitRouters(router)
	home.InitRouters(router, middleware.IpRateLimit(), jwt.JWTAuth())
	search.InitRouters(router)
}
