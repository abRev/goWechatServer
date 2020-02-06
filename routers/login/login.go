package login

import (
	"github.com/gin-gonic/gin"
	"wechat/controller/login"
)

func InitRouter(router *gin.Engine) {
	loginRouter := router.Group("/login")
	loginRouter.POST("/login", login.Login)
	loginRouter.POST("/register", login.Register)
}
