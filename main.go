package main

import (
	"wechat/config"
	"github.com/spf13/viper"
	"github.com/gin-gonic/gin"
	"wechat/routers"
)

func main() {
	if err := config.Init(""); err != nil {
        panic(err)
	}

	router := gin.Default()
	routers.InitRouters(router)
	router.Run(":" + viper.GetString("port"))
}

