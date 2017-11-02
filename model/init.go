package model

import (
	"webgin/global"
	"github.com/jinzhu/gorm"
)

func createTable() {
	MasterDB = connectDB(global.Config)
	MasterDB.SingularTable(true)
	//// 全局禁用表名复数
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "t_" + defaultTableName;
	}
	//更改默认表名
	//http://gorm.book.jasperxu.com/models.html#md
	MasterDB.AutoMigrate(&Asset{}, &Host{}, &IP{}, &Network{}, &IDC{}, &ServerVersion{}, &Dept{})
	global.GLog.Debug("master db init success")
	MasterDB.LogMode(true) // 启用Logger，显示详细日志 相当于单个db.Debug().Where
	global.GLog.Debug("set master db logmode true")
	global.GLog.Error("error 测试")
}

func init() {
	createTable()
}
