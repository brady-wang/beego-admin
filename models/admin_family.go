package models

import (
	"github.com/beego/beego/v2/client/orm"
)

// AdminUser struct
type AdminFamily struct {
	Id        int    `orm:"column(id);auto;size(11)" description:"表ID" json:"id"`
	Name      string `orm:"column(name);size(100)" description:"家庭名称" json:"name"`
	Number    string `orm:"column(number);size(11)" description:"家庭人数" json:"number"`
	Address   string `orm:"column(address);size(255)" description:"详细地址" json:"address"`
	ZoneId    int    `orm:"column(zone_id);size(11)" description:"区域ID" json:"zone_id"`
	CreatedAt int    `orm:"column(created_at);;size(18);default(0)" description:"创建时间" json:"created_at"`
	UpdatedAt int    `orm:"column(updated_at);;size(18);default(0)" description:"创建时间" json:"updated_at"`
}

// TableName 自定义table 名称
func (*AdminFamily) TableName() string {
	return "admin_family"
}

// SearchField 定义模型的可搜索字段
func (*AdminFamily) SearchField() []string {
	return []string{"name"}
}

// WhereField 定义模型可作为条件的字段
func (AdminFamily) WhereField() []string {
	return []string{}
}

//在init中注册定义的model
func init() {
	orm.RegisterModel(new(AdminFamily))
}
