package home

import (
	"fmt"
	"log"
	"strconv"
	"time"

	// "strconv"  // 类型转换使用
	"net/http"
	"wechat/db"
	"wechat/db/pg"
	"wechat/db/redis"
	"wechat/grpc"
	jwt "wechat/middleware/jwt"
	"wechat/model/money"
	"wechat/model/user"
	"wechat/modelgorm"
	pb "wechat/pb/helloworld"
	routeguide "wechat/pb/routeguide"

	"github.com/gin-gonic/gin"
)

type BodyJSON struct {
	From  string  `json:"from" binding:"required"`
	To    string  `json:"to" binding:"required"`
	Money float32 `json:"money" binding:"required"`
}

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

func CreateHome(c *gin.Context) {
	home := modelgorm.Home{
		Title:    "十二德堡",
		Birthday: time.Now(),
		Email:    "f@163.com",
	}
	ok := db.DB.Create(&home)
	c.JSON(http.StatusOK, gin.H{
		"ok": ok,
	})
}

func ListHome(c *gin.Context) {
	homes := []modelgorm.Home{}
	db.DB.Debug().Find(&homes)
	c.JSON(http.StatusOK, gin.H{
		"homes": homes,
	})
}

func GetValue(c *gin.Context) {
	claims := c.MustGet("claims").(*jwt.CustomClaims)
	fmt.Println(*claims)
	name := claims.Name
	client := redis.GetDB()
	if client == nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "no db",
		})
		return
	}
	val, err := client.Get(name).Result()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err,
			"extra":   *claims,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		name: val,
	})
}

func Stats(c *gin.Context) {
	db := pg.GetDB()
	status := db.Stats()
	c.JSON(http.StatusOK, gin.H{
		"status": status,
	})
}

func SetValue(c *gin.Context) {
	userJson := &user.UserBody{}
	if err := c.ShouldBindJSON(&userJson); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"ok":      false,
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
	result, err := db.Exec(`INSERT INTO "user"("name","age") VALUES($1, $2)`, name, age)
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

func GetUserList(c *gin.Context) {
	db := pg.GetDB()
	if db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ok":      false,
			"message": "数据库失联了",
		})
	}
	userDBs := &[]user.UserDB{}
	sql := `SELECT * FROM "user"`
	err := db.Select(userDBs, sql)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ok":      false,
			"message": "查询失败了",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"ok":    true,
		"users": userDBs,
	})
}

func LearnQueryx(c *gin.Context) {
	db := pg.GetDB()
	if db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ok":      false,
			"message": "数据库失联了",
		})
	}
	sql := `SELECT * FROM "user"`
	// 创建查询对象，返回sql.Rows对象 对象有多种方法
	// https://godoc.org/github.com/jmoiron/sqlx#Rows
	// https://golang.org/pkg/database/sql/#Rows
	rows, err := db.Queryx(sql)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ok":      false,
			"message": "查询失败了",
		})
	}
	// 查询表的所有列名
	strs, err := rows.Columns()
	for i, str := range strs {
		fmt.Println(i, str)
	}
	userDB := &user.UserDB{}
	for rows.Next() {
		fmt.Println("-----")
		// 依次打印所有行的内容
		if err := rows.Scan(&userDB.Id, &userDB.Age, &userDB.Name, &userDB.Phone, &userDB.Password); err == nil {
			fmt.Println("Scan: ", *userDB)
		} else {
			fmt.Println("err: ", err)
		}
	}
	if err := rows.Close(); err != nil {
		fmt.Println("Close rows err: ", err)
	}
	c.JSON(http.StatusOK, gin.H{
		"ok": true,
	})
}

func LearnTx(c *gin.Context) {
	bodyJSON := &BodyJSON{}
	if err := c.ShouldBindJSON(&bodyJSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"ok":      false,
			"message": "参数传递错误",
		})
	}
	fromUser := bodyJSON.From
	toUser := bodyJSON.To
	moneyTo := bodyJSON.Money

	db := pg.GetDB()
	if db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ok":      false,
			"message": "数据库失联了",
		})
	}
	moneyDB := &money.Money{}
	if err := db.Get(moneyDB, `SELECT * FROM "moneys" WHERE "user"=$1`, fromUser); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"ok":      false,
			"message": err,
		})
		return
	}
	// if moneyDB == nil {
	// 	if _, err := db.Exec(`INSERT INTO "moneys"("user", "count") VALUES($1,$2)`, fromUser, 0); err != nil {
	// 		fmt.Println(err)
	// 		c.JSON(http.StatusOK, gin.H{
	// 			"ok":      false,
	// 			"message": "创建用户初始化数据失败",
	// 		})
	// 	}
	// }
	if moneyDB.Count < moneyTo {
		c.JSON(http.StatusOK, gin.H{
			"ok":      false,
			"message": "余额不足",
		})
		return
	}
	// 事务开始
	tx, err := db.Beginx()
	if err != nil {
		fmt.Println("事务开始失败了")
	}
	if _, err := tx.Exec(`UPDATE "moneys" SET count = count-$1 WHERE "user"=$2`, moneyTo, fromUser); err != nil {
		fmt.Println(err)
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"ok":      false,
			"message": "操作失败请重试",
		})
		return
	}
	if _, err := tx.Exec(`UPDATE "moneys" SET count = count+$1 WHERE "user"=$2`, moneyTo, toUser); err != nil {
		fmt.Println(err)
		tx.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"ok":      false,
			"message": "操作失败请重试",
		})
		return
	}
	if err := tx.Commit(); err != nil {
		fmt.Println("事务结束失败了")
	}
	// 事务提交
	c.JSON(http.StatusOK, gin.H{
		"ok": true,
	})
}

func GetFile(c *gin.Context) {
	filename := c.Param("filename")
	fmt.Println(" : ", filename)
	if filename == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"ok": false,
		})
		return
	}
	c.String(http.StatusOK, filename)
}

func GrpcHello(c *gin.Context) {
	name := c.Query("name")
	req := &pb.HelloRequest{Name: name}
	res, err := grpc.HelloClient.SayHello(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": fmt.Sprint(res.Message),
	})
}

func GrpcRouteFeature(c *gin.Context) {
	la := c.Query("la")
	latitude, _ := strconv.ParseInt(la, 10, 32)
	lo := c.Query("lo")
	longitude, _ := strconv.ParseInt(lo, 10, 32)
	point := &routeguide.Point{
		Latitude:  int32(latitude),
		Longitude: int32(longitude),
	}
	feature, err := grpc.RouteClient.GetFeature(c, point)
	if err != nil {
		log.Fatalf("%v.GetFeatures(_) = _, %v: ", grpc.RouteClient, err)
	}
	c.JSON(http.StatusOK, gin.H{
		"feature": feature,
	})
}

type routeNode struct {
	Location *routeguide.Point `json:"location"`
	Message  string            `json:"message"`
}

func RunRouteChat(c *gin.Context) {
	notes := []*routeguide.RouteNote{}
	if err := c.ShouldBindJSON(&notes); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
	}
}
