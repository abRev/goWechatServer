package home

import (
	"github.com/gin-gonic/gin"
	"wechat/controller/home"
)

func InitRouters(router *gin.Engine) {
	homeRouter := router.Group("/home")
	homeRouter.GET("/redis/set/:name", home.Home)
	homeRouter.GET("/redis/set/:name/value", home.GetValue)
	homeRouter.POST("/pg/user", home.SetValue)
	homeRouter.GET("/pg/users", home.GetUserList)
	homeRouter.GET("/pg/queryx", home.LearnQueryx)
	homeRouter.POST("/pg/tx", home.LearnTx)
}