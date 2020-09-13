package db

import "github.com/garyburd/redigo/redis"

// PoolStats 查询连接池状态
func PoolStats() redis.PoolStats {
	return RedisPool.Stats()
}

// exnx ----------------------------------------

// SetExNx string 如果设置成功则返回nil
func SetExNx(key string, val interface{}, expireTime int) (string, error) {
	client := RedisPool.Get()
	defer client.Close()
	return redis.String(client.Do("SET", key, val, "EX", expireTime, "NX"))
}

// SetEx 设置自动过期key
func SetEx(key string, val interface{}, expireTime int) (string, error) {
	client := RedisPool.Get()
	defer client.Close()
	return redis.String(client.Do("SET", key, val, "EX", expireTime))
}

// strings ------------------------------------------

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

// list --------------------------------------------

func LPush(key string, value interface{}) (int, error) {
	client := RedisPool.Get()
	defer client.Close()
	return redis.Int(client.Do("LPUSH", key, value))
}

func RPush(key string, value interface{}) (int, error) {
	client := RedisPool.Get()
	defer client.Close()
	return redis.Int(client.Do("RPUSH", key, value))
}

func LRange(key string, start, end int) ([]string, error) {
	client := RedisPool.Get()
	defer client.Close()
	return redis.Strings(client.Do("LRANGE", key, start, end))
}

// hash --------------------------------------------

func HMSet(key string, v map[string]interface{}) (string, error) {
	client := RedisPool.Get()
	defer client.Close()
	args := make([]interface{}, 1+len(v)*2)
	args[0] = key
	index := 1
	for key, value := range v {
		args[index] = key
		args[index+1] = value
		index = index + 2
	}
	return redis.String(client.Do("HMSET", key))
}

func HGet(key string, name string) (string, error) {
	client := RedisPool.Get()
	defer client.Close()
	return redis.String(client.Do("HGET", key, name))
}

func Hincrby(key string, name string, incr float64) (float64, error) {
	client := RedisPool.Get()
	defer client.Close()
	return redis.Float64(client.Do("HINCRBY", key, name, incr))
}

func HExists(key, name string) (int, error) {
	client := RedisPool.Get()
	defer client.Close()
	return redis.Int(client.Do("HEXISTS", key, name))
}

// zset --------------------------------------------

// ZRemRangeByScore 根据分数区间删除zset成员
func ZRemRangeByScore(key, min, max string) (int64, error) {
	client := RedisPool.Get()
	defer client.Close()
	return redis.Int64(client.Do("zremrangebyscore", key, min, max))
}

// bloomFilter --------------------------------------------

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

// hyperloglog  -------------------------------------------

// PfAdd 添加
func PfAdd(key string, members []string) (int64, error) {
	client := RedisPool.Get()
	defer client.Close()
	args := make([]interface{}, 1+len(members))
	args[0] = key
	for i, v := range members {
		args[i+1] = v
	}
	return redis.Int64(client.Do("PFADD", args...))
}

// PfCount 统计
func PfCount(key string) (int64, error) {
	client := RedisPool.Get()
	defer client.Close()
	return redis.Int64(client.Do("PFCOUNT", key))
}

// PfMerge 合并多个hll到一个hll里面
func PfMerge(dest string, sources []string) (string, error) {
	client := RedisPool.Get()
	defer client.Close()
	args := make([]interface{}, 1+len(sources))
	args[0] = dest
	for i, v := range sources {
		args[i+1] = v
	}
	return redis.String(client.Do("PFMERGE", args...))
}
