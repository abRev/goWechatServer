package search

import (
	"fmt"
	"net/http"
	"wechat/db"

	"github.com/gin-gonic/gin"
)

func Info(c *gin.Context) {
	res, err := db.Es.Info()
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
