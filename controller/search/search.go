package search

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"wechat/db/es"
)

func Info(c *gin.Context) {
	es := es.GetES()
	if es == nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"msg":    "es链接失败",
		})
	}
	res, err := es.Info()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"err":    err.Error(),
		})
		return
	}
	fmt.Println(": ", res)
	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"info":   res,
	})

}
