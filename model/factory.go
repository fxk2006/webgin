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

var MasterDB *gorm.DB

func getMysqlUrl(conf *config.Config) (url string) {
	mysql := global.MYSQL
	user, _ := conf.String(mysql, "user")
	host, _ := conf.String(mysql, "host")
	password, _ := conf.String(mysql, "password")
	database, _ := conf.String(mysql, "database")
	url = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, database)
	//https://github.com/go-sql-driver/mysql#timetime-support
	fmt.Println(url)
	return
}

func getSqliteUrl(conf *config.Config) (url string) {
	sqlite := global.SQLITE
	url, _ = conf.String(sqlite, "database")
	return
}

func connectDB(conf *config.Config) (db *gorm.DB) {
	database := global.DATABASE
	sqlite := global.SQLITE
	mysql := global.MYSQL
	database, err := conf.String(database, "database")
	if err != nil {
		log.Fatalln(err)
	}
	if database == mysql {
		url := getMysqlUrl(conf)
		db, err = gorm.Open("mysql", url)
		if err != nil {
			panic("failed to connect database")
		}
		//mysql 连接池http://jinzhu.me/gorm/advanced.html#compose-primary-key
		MaxIdleConns, _ := conf.Int(mysql, "MaxIdleConns")
		MaxOpenConns, _ := conf.Int(mysql, "MaxOpenConns")
		db.DB().SetMaxIdleConns(MaxIdleConns)
		db.DB().SetMaxOpenConns(MaxOpenConns)
		db.Debug()
	}
	if database == sqlite {
		url := getSqliteUrl(conf)
		db, err = gorm.Open("sqlite3", url)
		if err != nil {
			panic("failed to connect database")
		}
		db.Debug()
	}

	return
}
