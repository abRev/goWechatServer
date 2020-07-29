package pg

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var DbCon *sqlx.DB

func init() {
	dialect := viper.GetString("common.pg.dialect")
	host := viper.GetString("common.pg.host")
	port := viper.GetString("common.pg.port")
	user := viper.GetString("common.pg.username")
	password := viper.GetString("common.pg.password")
	database := viper.GetString("common.pg.database")
	max := viper.GetInt("common.pg.max")
	idle := viper.GetInt("common.pg.idle")
	dataSourceName := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable", dialect, user, password, host, port, database)
	db, err := sqlx.Connect(dialect, dataSourceName)
	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(idle)
	db.SetMaxOpenConns(max)
	DbCon = db
}

// GetDB 获取数据库链接
func GetDB() *sqlx.DB {
	if DbCon != nil {
		return DbCon
	}
	return nil
}
