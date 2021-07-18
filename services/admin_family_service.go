package services

import (
	"beego-admin/formvalidate"
	"beego-admin/global"
	"beego-admin/models"
	"beego-admin/utils/page"
	"github.com/beego/beego/v2/client/orm"
	"net/url"
)

// AdminFamilyService struct
type AdminFamilyService struct {
	BaseService
}

// GetAdminFamilyById 根据id获取一条admin_user数据
func (*AdminFamilyService) GetAdminFamilyById(id int) *models.AdminFamily {
	o := orm.NewOrm()
	AdminFamily := models.AdminFamily{Id: id}
	err := o.Read(&AdminFamily)
	if err != nil {
		return nil
	}
	return &AdminFamily
}

// GetPaginateData 通过分页获取adminFamily
func (aus *AdminFamilyService) GetPaginateData(listRows int, params url.Values) ([]*models.AdminFamily, page.Pagination) {
	//搜索、查询字段赋值
	aus.SearchField = append(aus.SearchField, new(models.AdminFamily).SearchField()...)

	var adminFamily []*models.AdminFamily
	o := orm.NewOrm().QueryTable(new(models.AdminFamily))
	_, err := aus.PaginateAndScopeWhere(o, listRows, params).All(&adminFamily)
	if err != nil {
		return nil, aus.Pagination
	}
	return adminFamily, aus.Pagination
}

// Create 新增admin user用户
func (*AdminFamilyService) Create(form *formvalidate.AdminFamilyForm) int {

	adminFamily := models.AdminFamily{
		Name:    form.Name,
		Number:  form.Number,
		Address: form.Address,
		ZoneId:  form.ZoneId,
		CreatedAt: global.GetNowTime(),
	}
	id, err := orm.NewOrm().Insert(&adminFamily)

	if err == nil {
		return int(id)
	}
	return 0
}

// IsExistName 名称验重
func (*AdminFamilyService) IsExistName(name string, id int) bool {
	if id == 0 {
		return orm.NewOrm().QueryTable(new(models.AdminFamily)).Filter("name", name).Exist()
	}
	return orm.NewOrm().QueryTable(new(models.AdminFamily)).Filter("name", name).Exclude("id", id).Exist()
}

// Update 更新用户信息
func (*AdminFamilyService) Update(form *formvalidate.AdminFamilyForm) int {
	o := orm.NewOrm()
	adminFamily := models.AdminFamily{Id: form.Id}

	if o.Read(&adminFamily) == nil {
		adminFamily.Name = form.Name
		adminFamily.Number = form.Number
		adminFamily.ZoneId = form.ZoneId
		adminFamily.Address = form.Address
		adminFamily.UpdatedAt = global.GetNowTime()
		num, err := o.Update(&adminFamily)
		if err == nil {
			return int(num)
		}
		return 0
	}
	return 0
}

// Del 删除用户
func (*AdminFamilyService) Del(ids []int) int {
	count, err := orm.NewOrm().QueryTable(new(models.AdminFamily)).Filter("id__in", ids).Delete()
	if err == nil {
		return int(count)
	}
	return 0
}
