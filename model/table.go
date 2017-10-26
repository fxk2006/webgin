package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"fmt"
	"github.com/robfig/config"
	"log"
	"webgin/global"
)

var DB *gorm.DB


type Products struct {
	Code  string
	Price uint
}


func getMysqlUrl(conf *config.Config) (url string) {

	user, _ := conf.String("mysql", "user")
	host, _ := conf.String("mysql", "host")
	password, _ := conf.String("mysql", "password")
	database, _ := conf.String("mysql", "database")
	url = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, database)
	fmt.Println(url)
	return
}

func getSqliteUrl(conf *config.Config) (url string) {
	url, _ = conf.String("sqlite", "database")
	return
}

func connectDB(conf *config.Config) (db *gorm.DB) {

	database, err := conf.String("database", "database")
	if err != nil{
		log.Fatalln(err)
	}
	if database == "mysql" {
		url := getMysqlUrl(conf)
		db, err = gorm.Open("mysql", url)
		if err != nil {
			panic("failed to connect database")
		}
		//mysql 连接池http://jinzhu.me/gorm/advanced.html#compose-primary-key
		MaxIdleConns,_ := conf.Int("mysql","MaxIdleConns")
		MaxOpenConns,_ := conf.Int("mysql","MaxOpenConns")
		db.DB().SetMaxIdleConns(MaxIdleConns)
		db.DB().SetMaxOpenConns(MaxOpenConns)
		db.Debug()
	}
	if database == "sqlite" {
		url := getSqliteUrl(conf)
		db, err = gorm.Open("sqlite3", url)
		if err != nil {
			panic("failed to connect database")
		}
		db.Debug()
	}

	return
}
func createTable() {
	DB = connectDB(global.Config)
	DB.AutoMigrate(&Products{})
}

func init() {
	fmt.Println("go to model init")
	createTable()
}
