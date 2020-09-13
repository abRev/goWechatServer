package routers

import (
	"wechat/controller/redis"

	"github.com/gin-gonic/gin"
)

func InitRedisRouter(router *gin.Engine) {
	redisGroup := router.Group("/redis")
	redisGroup.POST("/zset", redis.ZsetSet)
	redisGroup.POST("/zsetbyscore", redis.GetZsetByScore)
	redisGroup.POST("/zsetbyrank", redis.GetZsetByRank)
	redisGroup.DELETE("/zsetrembyrank", redis.RemZsetByRank)
	/*----------------------------------------*/
	redisGroup.POST("/pfadd", redis.PFAdd)
	redisGroup.GET("/pfcount", redis.PFCount)
	redisGroup.POST("/pfmerge", redis.PFMerge)
	/*----------------------------------------*/
	redisGroup.POST("/bfadd", redis.BloomAdd)
	redisGroup.POST("/bfexists", redis.Blo-omExists)
}
