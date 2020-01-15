package home

import (
	"fmt"
	// "strconv"
	"github.com/gin-gonic/gin"
	"net/http"
	"wechat/db/redis"
	"wechat/db/pg"
	"wechat/model/user"
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

func SetValue(c *gin.Context) {
	userJson := &user.UserBody{}
	if err:= c.ShouldBindJSON(&userJson); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"ok": false,
			"message": err,
		})
		return
	}
	name := userJson.Name
	age := userJson.Age
	db := pg.GetDB()
	if db == nil {
		c.JSON(http.StatusOK, gin.H{
			"ok": false,
		})
		return
	}
	result, err:= db.Exec(`INSERT INTO "user"("name","age") VALUES($1, $2)`, name, age)
	fmt.Println(result)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"ok": true,
		})
		return
	}
	fmt.Println(err)
	c.JSON(http.StatusOK, gin.H{
		"ok": false,
	})
}