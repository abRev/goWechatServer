package db

import "github.com/garyburd/redigo/redis"

// Get 查询string类型的key
func Get(key string) (string, error) {
	client := RedisPool.Get()
	defer client.Close()
	return redis.String(client.Do("GET", key))
}

// Set 设置string类型的key
func Set(key string, val interface{}) (string, error) {
	client := RedisPool.Get()
	defer client.Close()
	return redis.String(client.Do("SET", key, val))
}

// Del 删除任一key
func Del(key string) (int64, error) {
	client := RedisPool.Get()
	defer client.Close()
	return redis.Int64(client.Do("DEL", key))
}

// Incr 增加一个value为number类型 +1
func Incr(key string) (int64, error) {
	client := RedisPool.Get()
	defer client.Close()
	return redis.Int64(client.Do("INCR", key))
}

// Decr 减少一个value为number类型 -1
func Decr(key string) (int64, error) {
	client := RedisPool.Get()
	defer client.Close()
	return redis.Int64(client.Do("DECR", key))
}

// BFAdd 布隆过滤器对一个key添加一个成员
func BFAdd(key, member string) (int64, error) {
	client := RedisPool.Get()
	defer client.Close()
	return redis.Int64(client.Do("BF.ADD", key, member))
}

// BFExists 布隆过滤器对一个key查询某个成员
func BFExists(key, member string) (int64, error) {
	client := RedisPool.Get()
	defer client.Close()
	return redis.Int64(client.Do("BF.EXISTS", key, member))
}

// BFMAdd 批量添加
func BFMAdd(key string, members []string) ([]int, error) {
	client := RedisPool.Get()
	defer client.Close()
	args := make([]interface{}, 1+len(members))
	args[0] = key
	for i, v := range members {
		args[i+1] = v
	}
	return redis.Ints(client.Do("BF.MADD", args...))
}
