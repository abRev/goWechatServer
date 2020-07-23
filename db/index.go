package db

import (
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/spf13/viper"
)

var RedisPool *redis.Pool

func init() {
	host := viper.GetString("common.redis.host")
	port := viper.GetString("common.redis.port")
	db := viper.GetInt("common.redis.db")
	maxidle := viper.GetInt("common.redis.maxidle")
	maxactive := viper.GetInt("common.redis.maxactive")
	idleTimeout := viper.GetDuration("common.redis.idleTimeout")
	// password := viper.GetString("common.redis.password")
	conURL := fmt.Sprintf("%s:%s", host, port)
	RedisPool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", conURL)
			if err != nil {
				return nil, err
			}
			// if _, err := c.Do("Auth", password); err != nil {
			// 	c.Close()
			// 	return nil, err
			// }
			if _, err := c.Do("SELECT", db); err != nil {
				c.Close()
				return nil, err
			}
			return c, nil
		},
		MaxIdle:     maxidle,
		MaxActive:   maxactive,
		IdleTimeout: time.Second * idleTimeout,
		Wait:        true,
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}
