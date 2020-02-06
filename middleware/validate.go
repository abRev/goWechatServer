package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Validate() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "" {
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "未传输token",
			})
		}
		if token == "abtest" {
			c.Next()
		} else {
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "token验证失败",
			})
		}
	}
}
