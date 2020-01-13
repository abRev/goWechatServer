package home

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wechat/db/redis"
)

func Home(c *gin.Context) {
	name := c.Param("name")
	value := c.Query("value")
	client := redis.GetDB()
	if client == nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "no db",
		})
	}
	err := client.Set(name, value, 0).Err()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "save",
	})
}


func GetValue(c *gin.Context) {
	name := c.Param("name")
	client := redis.GetDB()
	if client == nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "no db",
		})
	}
	val, err := client.Get(name).Result()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		name: val,
	})
}
