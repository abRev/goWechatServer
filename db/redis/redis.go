package redis

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/go-redis/redis"
)

var client *redis.Client

func Init() error {
	client = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("common.redis.addr"),
		Password: viper.GetString("common.redis.password"),
		DB:       viper.GetInt("common.redis.db"),
	})
	pong,err := client.Ping().Result()
	if err != nil {
		return err
	}
	fmt.Println("redis connect. ",pong)
	return nil
}

func GetDB() *redis.Client {
	if client != nil {
		return client
	}
	return nil
}