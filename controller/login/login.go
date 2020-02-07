package login

import (
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
	"wechat/db/pg"
	"wechat/middleware/jwt"
	"wechat/model/user"
)

func Login(c *gin.Context) {
	liginReq := &user.UserBody{}
	if c.BindJSON(&liginReq) == nil {
		// 判断账号密码是否正确
		db := pg.GetDB()
		if db == nil {
			c.JSON(http.StatusOK, gin.H{
				"ok": false,
			})
			return
		}
		userInfo := &user.UserDB{}
		log.Println(liginReq.Phone, liginReq.Password)
		if err := db.Get(userInfo, `SELECT * FROM "user" WHERE "phone"=$1`, liginReq.Phone); err != nil {
			// 校验失败
			c.JSON(http.StatusOK, gin.H{
				"ok":  false,
				"msg": "未查询到当前用户，请注册",
			})
		} else {
			// 查询成功判断密码是否正确
			if userInfo.Password == liginReq.Password {
				// 生成token
				generateToken(c, *userInfo)
			} else {
				c.JSON(http.StatusOK, gin.H{
					"ok":  false,
					"msg": "密码错误",
				})
			}
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    "json 解析失败",
		})
	}
}

func generateToken(c *gin.Context, userInfo user.UserDB) {
	j := &jwt.JWT{
		SigningKey: []byte("newAbang"),
	}
	claims := jwt.CustomClaims{
		ID:    "01234",
		Name:  userInfo.Name,
		Phone: userInfo.Phone,
		StandardClaims: jwtgo.StandardClaims{
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
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    "登录成功！",
		"token":  token,
	})
	return
}

func Register(c *gin.Context) {

}
