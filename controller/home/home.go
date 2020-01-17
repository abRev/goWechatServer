package home

import (
	"fmt"
	// "strconv"  // 类型转换使用
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

func GetUserList(c *gin.Context) {
	db:= pg.GetDB();
	if db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ok": false,
			"message": "数据库失联了",
		})
	}
	userDBs := &[]user.UserDB{}
	sql := `SELECT * FROM "user"`
	err := db.Select(userDBs, sql)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ok": false,
			"message": "查询失败了",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"ok": true,
		"users": userDBs,
	})
}

func LearnQueryx(c *gin.Context) {
	db := pg.GetDB()
	if db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ok": false,
			"message": "数据库失联了",
		})
	}
	sql := `SELECT * FROM "user"`
	// 创建查询对象，返回sql.Rows对象 对象有多种方法
	// https://godoc.org/github.com/jmoiron/sqlx#Rows
	// https://golang.org/pkg/database/sql/#Rows
	rows,err := db.Queryx(sql)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ok": false,
			"message": "查询失败了",
		})
	}
	// 查询表的所有列名
	strs,err := rows.Columns()
	for i, str := range strs {
		fmt.Println(i, str)
	}
	userDB := &user.UserDB{}
	for rows.Next() {
		// 依次打印所有行的内容
		if err:= rows.Scan(&userDB.Name, &userDB.Age); err == nil {
			fmt.Println("Scan: ", *userDB)
		}
	}
	if err:= rows.Close(); err != nil {
		fmt.Println("Close rows err: ", err)
	}
	c.JSON(http.StatusOK, gin.H{
		"ok": true,
	})
}

func LearnTx(c *gin.Context) {
	db := pg.GetDB()
	if db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ok": false,
			"message": "数据库失联了",
		})
	}
	tx,err := db.Beginx()
	if err != nil {
		fmt.Println("事务开始失败了")
	}
	if result,err:= tx.Exec(`UPDATE "weather" SET temp_lo=temp_lo+1 WHERE temp_lo<40`); err !=nil {
		fmt.Println("修改失败了",result);
	}
	if result,err:= tx.Exec(`UPDATE "weather" SET temp_lo=temp_lo+1 WHERE temp_lo<40`); err !=nil {
		fmt.Println("修改失败了",result);
	}
	if err := tx.Commit(); err != nil {
		fmt.Println("事务结束失败了")
	}
	c.JSON(http.StatusOK, gin.H{
		"ok": true,
	})
}

