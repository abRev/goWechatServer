package routers

import (
	"wechat/controller/search"

	"github.com/gin-gonic/gin"
)

func InitSearchRouters(router *gin.Engine) {
	searchRouter := router.Group("/search")
	searchRouter.GET("/info", search.Info)
}
