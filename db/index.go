package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
)

var (
	RedisPool *redis.Pool
	DB        *gorm.DB
)

func initRedis() {
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

func initPg() {
	dialect := viper.GetString("common.pg.dialect")
	host := viper.GetString("common.pg.host")
	port := viper.GetString("common.pg.port")
	user := viper.GetString("common.pg.username")
	password := viper.GetString("common.pg.password")
	database := viper.GetString("common.pg.database")
	max := viper.GetInt("common.pg.max")
	idle := viper.GetInt("common.pg.idle")
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, database, password)
	db, err := gorm.Open(dialect, dataSourceName)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	if goenv := os.Getenv("GO_ENV"); goenv == "development" {
		db.LogMode(true)
	}
	db.DB().SetMaxOpenConns(max)
	db.DB().SetMaxIdleConns(idle)
	DB = db
	log.Println("pg: gorm: " + host + ":" + port)
}

func init() {
	initPg()
	initRedis()
}
