package middleware

import (
	"net/http"
	"wechat/lib"

	"github.com/gin-gonic/gin"

	"time"
)

func IpRateLimit() gin.HandlerFunc {
	bucket := lib.Bucket{
		Name: "IP",
		Rate: 10,
		Duration: 10000 * time.Millisecond,
	}
	return func(c *gin.Context) {
		ip:= c.ClientIP()
		status := bucket.Limit(ip)
		if status {
			c.Next()
		} else {
			c.Abort()
			c.JSON(http.StatusTooEarly, gin.H{
				"message": "too frequently！",
			})
		}
	}
}

