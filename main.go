package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"wechat/config"
	"wechat/db/es"
	"wechat/db/pg"
	"wechat/db/redis"
	"wechat/routers"
)

func main() {
	if err := config.Init(""); err != nil {
		panic(err)
	}
	if err := redis.Init(); err != nil {
		panic(err)
	}
	if err := pg.Init(); err != nil {
		panic(err)
	}
	if err := es.Init(); err != nil {
		panic(err)
	}
	router := gin.Default()
	routers.InitRouters(router)
	router.Run(":" + viper.GetString("port"))
}
