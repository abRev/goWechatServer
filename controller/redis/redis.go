package redis

import (
	"fmt"
	"net/http"
	cache "wechat/db/redis"

	"wechat/lib"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

type zsetBody struct {
	Key     string    `json:"key"`
	Members []redis.Z `json:"members"`
}

// ZsetSet 设置zset数据
func ZsetSet(c *gin.Context) {
	body := &zsetBody{}
	c.ShouldBindJSON(&body)
	client := cache.GetDB()
	res := client.ZAdd(body.Key, body.Members...)
	if intCmd, err := res.Result(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(intCmd)
	}
	c.JSON(http.StatusOK, gin.H{
		"ok": true,
	})
}

type getByScore struct {
	Key       string         `json:"key"`
	Order     bool           `json:"order"`
	Condition redis.ZRangeBy `json:"condition"`
}

// GetZsetByScore 查询
func GetZsetByScore(c *gin.Context) {
	body := &getByScore{}
	c.ShouldBindJSON(&body)
	client := cache.GetDB()
	var res *redis.ZSliceCmd
	if body.Order {
		res = client.ZRangeByScoreWithScores(body.Key, body.Condition)
	} else {
		res = client.ZRevRangeByScoreWithScores(body.Key, body.Condition)
	}
	if data, err := res.Result(); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"ok":  true,
			"msg": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"ok":   true,
			"data": data,
		})
	}

}

type getByRank struct {
	Key   string `json:"key"`
	Order bool   `json:"order" default:true`
	Start int64  `json:"start"`
	Stop  int64  `json:"stop"`
}

// GetZsetByRank 通过排名来查询
func GetZsetByRank(c *gin.Context) {
	body := &getByRank{}
	c.ShouldBindJSON(&body)
	client := cache.GetDB()
	var res *redis.ZSliceCmd
	if body.Order {
		res = client.ZRangeWithScores(body.Key, body.Start, body.Stop)
	} else {
		res = client.ZRevRangeWithScores(body.Key, body.Start, body.Stop)
	}
	if data, err := res.Result(); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"ok":  true,
			"msg": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"ok":   true,
			"data": data,
		})
	}
}

type remByRank struct {
	Key   string `json:"key"`
	Start int64  `json:"start"`
	Stop  int64  `json:"stop"`
}

// RemZsetByRank 通过排名区间来删除
func RemZsetByRank(c *gin.Context) {
	body := &remByRank{}
	c.ShouldBindJSON(&body)
	client := cache.GetDB()
	var res *redis.IntCmd
	res = client.ZRemRangeByRank(body.Key, body.Start, body.Stop)
	if data, err := res.Result(); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"ok":  true,
			"msg": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"ok":   true,
			"data": data,
		})
	}
}

type remByScore struct {
	Key string `json:"key"`
	Min string `json:"min"`
	Max string `json:"max"`
}

// RemZsetByScore 通过分数范围来删除
func RemZsetByScore(c *gin.Context) {
	body := &remByScore{}
	c.ShouldBindJSON(&body)
	client := cache.GetDB()
	res := client.ZRemRangeByScore(body.Key, body.Min, body.Max)
	if data, err := res.Result(); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"ok":  true,
			"msg": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"ok":   true,
			"data": data,
		})
	}
}

type pfAdd struct {
	Key     string   `json:"key"`
	Members []string `json:"members"`
}

// PFAdd 增加hyperLogLog的成员
func PFAdd(c *gin.Context) {
	body := &pfAdd{}
	c.ShouldBindJSON(&body)
	client := cache.GetDB()
	res := client.PFAdd(body.Key, body.Members)
	if count, err := res.Result(); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"ok":  true,
			"err": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"ok":    true,
			"count": count,
		})
	}
}

// PFCount 查询统计人数
func PFCount(c *gin.Context) {
	key := c.Query("key")
	client := cache.GetDB()
	res := client.PFCount(key)
	if count, err := res.Result(); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"ok":  true,
			"err": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"ok":    true,
			"count": count,
		})
	}
}

type pfMerge struct {
	Dest    string   `json:"dest"`
	Sources []string `json:"sources"`
}

// PFMerge 合并HyperLogLog
func PFMerge(c *gin.Context) {
	body := &pfMerge{}
	c.ShouldBindJSON(&body)
	client := cache.GetDB()
	res := client.PFMerge(body.Dest, body.Sources...)
	if count, err := res.Result(); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"ok":  true,
			"err": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"ok":    true,
			"count": count,
		})
	}
}

type bladd struct {
	Key    string `json:"key"`
	Member string `json:"member"`
}

// BloomAdd bloom过滤器添加元素
func BloomAdd(c *gin.Context) {
	body := &bladd{}
	c.ShouldBindJSON(&body)
	userBL := &lib.Bloom{
		Key: body.Key,
	}
	members := []string{
		body.Member,
	}

	if count, err := userBL.MAdd(members); err != nil {
		fmt.Println("err:", err)
		c.JSON(http.StatusOK, gin.H{
			"ok":  true,
			"err": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"ok":    true,
			"count": count,
		})
	}
}

type blExists struct {
	Key    string `json:"key"`
	Member string `json:"member"`
}

func BloomExists(c *gin.Context) {
	body := &blExists{}
	c.ShouldBindJSON(&body)
	userBL := &lib.Bloom{
		Key: body.Key,
	}
	if count, err := userBL.Exists(body.Member); err != nil {
		fmt.Println("err:", err)
		c.JSON(http.StatusOK, gin.H{
			"ok":  true,
			"err": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"ok":    true,
			"count": count,
		})
	}
}
