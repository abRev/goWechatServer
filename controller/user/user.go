package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserInfo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "abang",
	})
}
