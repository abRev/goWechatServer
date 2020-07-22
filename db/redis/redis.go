package redis

import (
	"fmt"

	_ "wechat/config"

	"github.com/go-redis/redis/v7"
	"github.com/spf13/viper"
)

var client *redis.Client

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("common.redis.addr"),
		Password: viper.GetString("common.redis.password"),
		DB:       viper.GetInt("common.redis.db"),
	})
	pong, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("redis connect. ", pong)
}

// GetDB 获取redis pool
func GetDB() *redis.Client {
	if client != nil {
		return client
	}
	return nil
}
