package routers

import (
	"wechat/controller/login"
	"wechat/middleware"
	"wechat/middleware/jwt"

	"github.com/gin-gonic/gin"
)

func InitRouters(router *gin.Engine) {

	router.POST("/login", middleware.IpRateLimit(), login.Login)
	router.POST("/register", middleware.IpRateLimit(), login.Register)
	InitWechatRouters(router)
	InitUserRouters(router)
	InitHomeRouters(router, middleware.IpRateLimit(), jwt.JWTAuth())
	InitSearchRouters(router)
	InitRedisRouter(router)
}
