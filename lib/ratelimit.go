package lib

import (
	"fmt"
	"time"
	"wechat/db/redis"

	Redis "github.com/go-redis/redis/v7"
)

type Bucket struct {
	Name     string
	Rate     int
	Duration time.Duration
}

func (bucket *Bucket) Limit(key string) bool {
	client := redis.GetDB()
	pipliner := client.Pipeline()
	bucketKey := fmt.Sprintf("%s:%s", bucket.Name, key)
	pipliner.SetNX(bucketKey, bucket.Rate, bucket.Duration)
	pipliner.Decr(bucketKey)
	pipliner.PTTL(bucketKey)
	res, _ := pipliner.Exec()
	if intCmd, ok := res[1].(*Redis.IntCmd); ok == true {
		if intCmd.Val() <= 0 {
			return false
		}
	}
	if durationCmd, ok := res[2].(*Redis.DurationCmd); ok == true {
		if durationCmd.Val().Nanoseconds() <= 0 {
			return false
		}
	}
	return true
}
