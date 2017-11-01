package model

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

