package pg

import (
	"fmt"
	_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

var db *sqlx.DB

func GetDB() *sqlx.DB {
	if db != nil {
		return db
	}
	host := viper.GetString("common.pg.host")
	port := viper.GetString("common.pg.port")
	user := viper.GetString("common.pg.username")
	password := viper.GetString("common.pg.password")
	database := viper.GetString("common.pg.databse")
	max := viper.GetInt("common.pg.max")
	idle := viper.GetInt("common.pg.idle")
	dataSourceName := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, database)
	db = sqlx.MustConnect("postgres", dataSourceName)
	db.SetMaxIdleConns(idle)
	db.SetMaxOpenConns(max)
	return db
}