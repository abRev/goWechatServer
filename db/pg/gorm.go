package pg

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
)

var (
	DB *gorm.DB
)

func init() {
	dialect := viper.GetString("common.pg.dialect")
	host := viper.GetString("common.pg.host")
	port := viper.GetString("common.pg.port")
	user := viper.GetString("common.pg.username")
	password := viper.GetString("common.pg.password")
	database := viper.GetString("common.pg.database")
	max := viper.GetInt("common.pg.max")
	idle := viper.GetInt("common.pg.idle")
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, database, password)
	db, err := gorm.Open(dialect, dataSourceName)
	if err != nil {
		panic(err)
	}
	if goenv := os.Getenv("GO_ENV"); goenv == "development" {
		db.LogMode(true)
	}
	db.DB().SetMaxOpenConns(max)
	db.DB().SetMaxIdleConns(idle)
	DB = db
	log.Println("pg: gorm: " + host + ":" + port)
}
