package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
)

var DB *gorm.DB

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func connectMysql(url string) (db *gorm.DB) {
	db, err := gorm.Open("mysql",url)
	if err != nil {
		panic("failed to connect database")
	}
	DB = db
	db.Debug()
	return
}
func createTable(){
	db := connectMysql("admin:yangyang123@tcp(192.168.53.132:3306)/gorm?charset=utf8&parseTime=True&loc=Local")
	db.AutoMigrate(&Product{})
}

func init(){
	fmt.Println("go to model init")
	createTable()
}
