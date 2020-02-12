package home

import (
	"github.com/gin-gonic/gin"
	"wechat/controller/home"
)

func InitRouters(router *gin.Engine, middleware gin.HandlerFunc) {
	homeRouter := router.Group("/home")
	homeRouter.Use(middleware)
	homeRouter.GET("/redis/set/:name", home.Home)
	homeRouter.GET("/redis/value", home.GetValue)
	homeRouter.GET("/stats", home.Stats)
	homeRouter.POST("/pg/user", home.SetValue)
	homeRouter.GET("/pg/users", home.GetUserList)
	homeRouter.GET("/pg/queryx", home.LearnQueryx)
	homeRouter.POST("/pg/tx", home.LearnTx)
	homeRouter.GET("/file/:filename", home.GetFile)
}
