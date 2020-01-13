package home

import (
	"github.com/gin-gonic/gin"
	"wechat/controller/home"
)

func InitRouters(router *gin.Engine) {
	homeRouter := router.Group("/home")
	homeRouter.GET("/:name", home.Home)
	homeRouter.GET("/:name/value", home.GetValue)
}