package home

import (
	"github.com/gin-gonic/gin"
	"wechat/controller/home"
)

func InitRouters(router *gin.Engine) {
	homeRouter := router.Group("/home")
	homeRouter.GET("/set/:name", home.Home)
	homeRouter.GET("/set/:name/value", home.GetValue)
	homeRouter.POST("/pgset", home.SetValue)
}