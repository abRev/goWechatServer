package search

import (
	"github.com/gin-gonic/gin"
	"wechat/controller/search"
)

func InitRouters(router *gin.Engine) {
	searchRouter := router.Group("/search")
	searchRouter.GET("/info", search.Info)
}
