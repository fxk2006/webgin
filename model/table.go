package model

import (
	"webgin/global"
	"time"
	"github.com/jinzhu/gorm"
)

const (
	SSJ  = iota
	KN
	JR
	XD
	DSJ
	YYXT
)

const (
	Online  = iota
	Offline
)

type Network struct {
	ID      int    `gorm:"primary_key"`
	Segment string `gorm:"type:varchar(255)"`
	Remark  string `gorm:"type:text"`
}

type Asset struct {
	ID            int64     `gorm:"primary_key"`
	AssetLabel    string    `gorm:"type:varchar(100);column:asset_label"`
	AssetNum      string    `gorm:"type:varchar(100);column:asset_num"`
	AssetSN       string    `gorm:"type:varchar(100);column:asset_sn"`
	Dept          string    `gorm:"type:varchar(255);column:dept"`
	ExpiredTime   time.Time `gorm:"type:date;column:expired_time"`
	OnlineTime    time.Time `gorm:"type:date;column:online_time"`
	Status        byte      `gorm:"column:status"`
	AssetType     string    `gorm:"type:varchar(255);column:asset_type"`
	ServerVersion string    `gorm:"type:varchar(255);column:server_version"`
	Config        string    `gorm:"type:text;column:config"`
	Remark        string    `gorm:"type:text;column:remark"`
	Idc           string    `gorm:"type:varchar(100);column:idc"`
	Position      string    `gorm:"type:varchar(100);column:position"`
	Hosts         []Host
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
type Host struct {
	ID       int64  `gorm:"primary_key"`
	AssetID  int64  `gorm:"index"`
	HostName string `gorm:"type:varchar(255)"`
	IPs      []IP
}

type IP struct {
	ID          int64  `gorm:"primary_key"`
	HostID      int64  `gorm:"index"`
	InnerIpaddr string `gorm:"type:varchar(255)"`
	WlanIpaddr  string `gorm:"type:varchar(255)"`
	IsVip       bool
}

func createTable() {
	DB = connectDB(global.Config)
	DB.SingularTable(true)
	//// 全局禁用表名复数
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "t_" + defaultTableName;
	}
	DB.Debug()
	//更改默认表名
	//http://gorm.book.jasperxu.com/models.html#md
	DB.AutoMigrate(&Asset{}, &Host{}, &IP{}, &Network{})
}

func init() {
	createTable()
}
