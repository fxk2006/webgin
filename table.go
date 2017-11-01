package model

import (
	"webgin/global"
	"time"
	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"
)
type IDC struct {
	//idc表配置
	ID   int    `gorm:"primary_key"`
	Name string `gorm:"type:varchar(255)"`
}

func (IDC) TableName() string {
	return "idc"
}

type ServerVersion struct {
	//服务器型号表配置
	ID      int    `gorm:"primary_key"`
	Version string `gorm:"type:varchar(255)"`
}
type Dept struct {
	//部门表配置
	ID   int    `gorm:"primary_key"`
	Name string `gorm:"type:varchar(255)"`
}

type Network struct {
	ID      int    `gorm:"primary_key"`
	Segment string `gorm:"type:varchar(255)"`
	Remark  string `gorm:"type:text"`
}

type Asset struct {
	ID              int64         `gorm:"primary_key"`                                 //id
	AssetLabel      string        `gorm:"type:varchar(100);column:asset_label"`        //资产标签 运维标签
	AssetNum        string        `gorm:"type:varchar(100);column:asset_num"`          //资产编号 财务编码
	AssetSN         string        `gorm:"type:varchar(100);column:asset_sn"`           //资产sn
	ExpiredTime     time.Time     `gorm:"type:date;column:expired_time"`               //过期时间
	OnlineTime      time.Time     `gorm:"type:date;column:online_time"`                //上线时间
	Status          byte          `gorm:"column:status"`                               //状态
	AssetType       string        `gorm:"type:varchar(255);column:asset_type"`         //资产类型
	Config          string        `gorm:"type:text;column:config"`                     //配置信息
	Remark          string        `gorm:"type:text;column:remark"`                     //备注
	Position        string        `gorm:"type:varchar(100);column:position"`           //机柜位置
	Hosts           []Host        `gorm:"ForeignKey:AssetID;AssociationForeignKey:ID"` //主机
	Dept            Dept          `gorm:"ForeignKey:DeptID;AssociationForeignKey:ID"`  //部门
	DeptID          int
	ServerVersion   ServerVersion `gorm:"ForeignKey:ServerVersionID;AssociationForeignKey:ID"` //资产型号
	ServerVersionID int
	Idc             IDC           `gorm:"ForeignKey:IdcID;AssociationForeignKey:ID"` //托管idc
	IdcID           int
	CreatedAt       time.Time //创建时间
	UpdatedAt       time.Time //更新时间
}

type ListAsset []Asset

func (this ListAsset) ToListJson() (container []gin.H) {
	for _, asset := range this {
		container = append(container, asset.ToMapJson())
	}
	return
}
func (this *Asset) ToMapJson() gin.H {
	return gin.H{
		"id":            this.ID,
		"asset_label":   this.AssetLabel,
		"asset_num":     this.AssetNum,
		"asset_sn":      this.AssetSN,
		"expired_time":  this.ExpiredTime,
		"online_time":   this.OnlineTime,
		"status":        this.Status,
		"asset_type":    this.AssetType,
		"config":        this.Config,
		"Remark":        this.Remark,
		"Dept":          this.Dept.Name,
		"ServerVersion": this.ServerVersion.Version,
		"Idc":           this.Idc.Name,
	}

}

type Host struct {
	ID       int64  `gorm:"primary_key"`                    //id
	AssetID  int64  `gorm:"index"`                          //资产id
	HostName string `gorm:"type:varchar(255)"`              //主机名
	IsVm     bool                                           //是否虚拟机
	Label    string `gorm:"type:varchar(100);column:label"` //标签
	Remark   string `gorm:"type:text;column:remark"`        //备注
	IPs      []IP   `gorm:"ForeignKey:HostID"`              //Ip地址
}

type IP struct {
	ID          int64  `gorm:"primary_key"`       //id
	HostID      int64  `gorm:"index"`             //主机id
	InnerIpaddr string `gorm:"type:varchar(255)"` //内网地址
	WlanIpaddr  string `gorm:"type:varchar(255)"` //外网地址
	IsVip       bool                              //是否vip
}









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
	global.GLog.Debug("master db init success!")
}

type JsonMapObject interface {
	ToMapJson() gin.H
}
type JsonListObject interface {
	ToListJson() []gin.H
}

func init() {
	createTable()
}
