package model


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
