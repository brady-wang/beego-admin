package services

import (
	"beego-admin/formvalidate"
	"beego-admin/global"
	"beego-admin/models"
	"beego-admin/utils/page"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"net/url"
)

// AdminPeopleService struct
type AdminPeopleService struct {
	BaseService
}
type PeopleNum struct {
	FamilyId int `orm:"column(family_id);;size(18);default(0)" description:"" json:"family_id"`
	Number   int `orm:"column(number);;size(18);default(0)" description:"" json:"number"`
}

// GetAdminPeopleById 根据id获取一条admin_user数据
func (*AdminPeopleService) GetAdminPeopleById(id int) *models.AdminPeople {
	o := orm.NewOrm()
	AdminPeople := models.AdminPeople{Id: id}
	err := o.Read(&AdminPeople)
	if err != nil {
		return nil
	}
	return &AdminPeople
}

// GetPaginateData 通过分页获取adminPeople
func (aus *AdminPeopleService) GetPaginateData(listRows int, params url.Values) ([]*models.AdminPeople, page.Pagination) {
	//搜索、查询字段赋值
	aus.SearchField = append(aus.SearchField, new(models.AdminPeople).SearchField()...)
	fmt.Println(aus.SearchField)
	var adminPeople []*models.AdminPeople
	o := orm.NewOrm().QueryTable(new(models.AdminPeople))
	_, err := aus.PaginateAndScopeWhere(o, listRows, params).All(&adminPeople)
	if err != nil {
		return nil, aus.Pagination
	}
	return adminPeople, aus.Pagination
}

// Create 新增admin user用户
func (*AdminPeopleService) Create(form *formvalidate.AdminPeopleForm) int {

	adminPeople := models.AdminPeople{
		Name:       form.Name,
		Sex:        form.Sex,
		Address:    form.Address,
		IdCard:     form.IdCard,
		SocialCard: form.SocialCard,
		FamilyId:   form.FamilyId,
		Mobile:     form.Mobile,
		CreatedAt:  global.GetNowTime(),
	}
	id, err := orm.NewOrm().Insert(&adminPeople)

	if err == nil {
		return int(id)
	}
	return 0
}

func (*AdminPeopleService) GetCount(familyId int ) int {
	count, err := orm.NewOrm().QueryTable(new(models.AdminPeople)).Filter("family_id",familyId).Count()
	if err != nil {
		return 0
	}
	return int(count)
}

// IsExistName 名称验重
func (*AdminPeopleService) IsExistName(name string, id int) bool {
	if id == 0 {
		return orm.NewOrm().QueryTable(new(models.AdminPeople)).Filter("name", name).Exist()
	}
	return orm.NewOrm().QueryTable(new(models.AdminPeople)).Filter("name", name).Exclude("id", id).Exist()
}

// Update 更新用户信息
func (*AdminPeopleService) Update(form *formvalidate.AdminPeopleForm) int {
	o := orm.NewOrm()
	adminPeople := models.AdminPeople{Id: form.Id}

	if o.Read(&adminPeople) == nil {
		adminPeople.Name = form.Name
		adminPeople.Sex = form.Sex
		adminPeople.IdCard = form.IdCard
		adminPeople.Mobile = form.Mobile
		adminPeople.SocialCard = form.SocialCard
		adminPeople.FamilyId = form.FamilyId
		adminPeople.Address = form.Address
		adminPeople.UpdatedAt = global.GetNowTime()
		num, err := o.Update(&adminPeople)
		fmt.Println(err)
		if err == nil {
			return int(num)
		}
		return 0
	}
	return 0
}

// Del 删除用户
func (*AdminPeopleService) Del(ids []int) int {
	count, err := orm.NewOrm().QueryTable(new(models.AdminPeople)).Filter("id__in", ids).Delete()
	if err == nil {
		return int(count)
	}
	return 0
}

// GetExportData 获取导出数据
func (us *AdminPeopleService) GetExportData(params url.Values) []*models.AdminPeople {
	//搜索、查询字段赋值
	us.SearchField = append(us.SearchField, new(models.AdminPeople).SearchField()...)
	var adminPeople []*models.AdminPeople
	o := orm.NewOrm().QueryTable(new(models.AdminPeople))
	_, err := us.ScopeWhere(o, params).All(&adminPeople)
	if err != nil {
		return nil
	}
	return adminPeople
}