package lib

import (
	cache "wechat/db/redis"

	"github.com/go-redis/redis/v7"
)

type Bloom struct {
	Key string
}

func (bl *Bloom) Add(member string) (int64, error) {
	client := cache.GetDB()
	args := make([]interface{}, 3)
	args[0] = "bf.add"
	args[1] = bl.Key
	args[2] = member

	cmd := redis.NewIntCmd(args...)
	client.Process(cmd)
	if count, err := cmd.Result(); err != nil {
		return 0, err
	} else {
		return count, nil
	}
}

func (bl *Bloom) Exists(member string) (int64, error) {
	client := cache.GetDB()
	args := make([]interface{}, 3)
	args[0] = "bf.exists"
	args[1] = bl.Key
	args[2] = member
	cmd := redis.NewIntCmd(args...)
	client.Process(cmd)
	if count, err := cmd.Result(); err != nil {
		return 0, err
	} else {
		return count, nil
	}
}

func (bl *Bloom) MAdd(members []string) ([]int64, error) {
	client := cache.GetDB()
	args := make([]interface{}, 2+len(members))
	args[0] = "BF.MADD"
	args[1] = bl.Key
	for i, val := range members {
		args[i+2] = val
	}
	cmd := redis.NewIntSliceCmd(args...)
	client.Process(cmd)
	if resultArr, err := cmd.Result(); err != nil {
		return []int64{}, err
	} else {
		return resultArr, nil
	}
}
