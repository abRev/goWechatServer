package login

import (
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
	"wechat/middleware/jwt"
)

type User struct {
	Phone string `json:"mobile"`
	Pwd   string `json:"password"`
}

type LoginResult struct {
	Token string `json:"token"`
	User
}

func Login(c *gin.Context) {
	var liginReq User
	if c.BindJSON(&liginReq) == nil {
		// 判断账号密码是否正确
		// 生成token
		// 校验成功
		generateToken(c, liginReq)
		// 校验失败
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    "json 解析失败",
		})
	}
}

func generateToken(c *gin.Context, user User) {
	j := &jwt.JWT{
		[]byte("newAbang"),
	}
	claims := jwt.CustomClaims{
		"01234",
		"ab",
		user.Phone,
		jwtgo.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600), // 签名失效时间 1小时
			Issuer:    "abang",                         // 签名发行者
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    err.Error(),
		})
		return
	}
	log.Println(claims, token)
	data := LoginResult{
		User:  user,
		Token: token,
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    "登录成功！",
		"data":   data,
	})
	return
}

func Register(c *gin.Context) {

}
