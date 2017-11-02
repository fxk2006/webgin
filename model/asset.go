package model

import (
	"time"
	"github.com/gin-gonic/gin"
)

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

type AssetJson struct {
	AssetLabel    string
	AssetNum      string    //资产编号 财务编码
	AssetSN       string    //资产sn
	ExpiredTime   time.Time //过期时间
	OnlineTime    time.Time //上线时间
	Status        byte      //状态
	AssetType     string    //资产类型
	Config        string    //配置信息
	Remark        string    //备注
	Position      string    //机柜位置
	Dept          string
	ServerVersion string
	Idc           string
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
		"AssetLabel":    this.AssetLabel,
		"AssetNum":      this.AssetNum,
		"AssetSN":       this.AssetSN,
		"ExpiredTime":   this.ExpiredTime,
		"OnlineTime":    this.OnlineTime,
		"Status":        this.Status,
		"AssetType":     this.AssetType,
		"Config":        this.Config,
		"Remark":        this.Remark,
		"Dept":          this.Dept.Name,
		"ServerVersion": this.ServerVersion.Version,
		"Idc":           this.Idc.Name,
	}
}
