package models

import (
	"github.com/beego/beego/v2/client/orm"
)

// AdminUser struct
type AdminPeople struct {
	Id         int    `orm:"column(id);auto;size(11)" description:"表ID" json:"id"`
	Name       string `orm:"column(name);size(100)" description:"人员名称" json:"name"`
	Sex        string `orm:"column(sex);size(18)" description:"性别" json:"sex"`
	Address    string `orm:"column(address);size(255)" description:"详细地址" json:"address"`
	Mobile     string `orm:"column(mobile);size(18)" description:"手机号" json:"mobile"`
	IdCard     string `orm:"column(id_card);size(25)" description:"身份证" json:"id_card"`
	SocialCard string `orm:"column(social_card);size(25)" description:"社保卡" json:"social_card"`
	FamilyId   int `orm:"column(family_id);size(25)" description:"家庭" json:"family_id"`
	CreatedAt  int    `orm:"column(created_at);;size(18);default(0)" description:"创建时间" json:"created_at"`
	UpdatedAt  int    `orm:"column(updated_at);;size(18);default(0)" description:"创建时间" json:"updated_at"`
}

// TableName 自定义table 名称
func (*AdminPeople) TableName() string {
	return "admin_people"
}

// SearchField 定义模型的可搜索字段
func (*AdminPeople) SearchField() []string {
	return []string{"name","family_id"}
}

// WhereField 定义模型可作为条件的字段
func (AdminPeople) WhereField() []string {
	return []string{}
}

//在init中注册定义的model
func init() {
	orm.RegisterModel(new(AdminPeople))
}
