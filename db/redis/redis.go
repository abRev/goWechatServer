package redis

import (
	"fmt"

	"github.com/go-redis/redis/v7"
	"github.com/spf13/viper"
)

var client *redis.Client

func init() {
	addr := fmt.Sprintf("%s:%s", viper.GetString("common.redis.host"), viper.GetString("common.redis.port"))
	client = redis.NewClient(&redis.Options{
		Addr:     addr,
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
