package routers

import (
	"fmt"
	"net/http"
	"wechat/controller/user"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// 校验token，不太好用
func validate() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		fmt.Println(session)
		v := session.Get("count")
		if v == nil {
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "session 验证失败",
			})
		} else {
			c.Next()
		}
	}
}
func InitUserRouters(router *gin.Engine) {
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))
	userRouter := router.Group("/user")
	userRouter.Use(validate())
	userRouter.GET("/info", user.GetUserInfo)
}
