package main

import (
	"wechat/routers"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	router := gin.Default()
	routers.InitRouters(router)
	router.Run(":" + viper.GetString("port"))
}
