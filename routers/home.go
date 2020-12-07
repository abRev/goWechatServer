package routers

import (
	"wechat/controller/home"

	"github.com/gin-gonic/gin"
)

func InitHomeRouters(router *gin.Engine, middlewares ...gin.HandlerFunc) {
	homeRouter := router.Group("/home")
	for _, middleware := range middlewares {
		homeRouter.Use(middleware)
	}
	homeRouter.GET("/redis/set/:name", home.Home)
	homeRouter.GET("/redis/value", home.GetValue)
	homeRouter.GET("/stats", home.Stats)
	homeRouter.POST("/pg/user", home.SetValue)
	homeRouter.GET("/pg/users", home.GetUserList)
	homeRouter.GET("/pg/queryx", home.LearnQueryx)
	homeRouter.POST("/pg/tx", home.LearnTx)
	homeRouter.GET("/file/:filename", home.GetFile)
	homeRouter.GET("/home", home.CreateHome)
	homeRouter.GET("/homes", home.ListHome)
	homeRouter.GET("/grpc/hello", home.GrpcHello)
	homeRouter.GET("/grpc/route/feature", home.GrpcRouteFeature)
	homeRouter.POST("/grpc/route/chat", home.RunRouteChat)
}
