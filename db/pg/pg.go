package pg

import (
	"fmt"
	_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

var dbCon *sqlx.DB

func Init() error {
	host := viper.GetString("common.pg.host")
	port := viper.GetString("common.pg.port")
	user := viper.GetString("common.pg.username")
	password := viper.GetString("common.pg.password")
	database := viper.GetString("common.pg.databse")
	max := viper.GetInt("common.pg.max")
	idle := viper.GetInt("common.pg.idle")
	dataSourceName := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, database)
	db,err := sqlx.Connect("postgres", dataSourceName)
	if err != nil {
		return err
	}
	dbCon = db
	db.SetMaxIdleConns(idle)
	db.SetMaxOpenConns(max)
	return nil
}

func GetDB() *sqlx.DB {
	if dbCon != nil {
		return dbCon
	}
	return nil
}