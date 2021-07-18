package controllers

import (
	"beego-admin/services"
)

// PeopleController struct
type PeopleController struct {
	baseController
}

// Index 用户管理-首页
func (auc *PeopleController) Index() {
	var adminUserService services.AdminUserService
	data, pagination := adminUserService.GetPaginateData(admin["per_page"].(int), gQueryParams)
	auc.Data["data"] = data
	auc.Data["paginate"] = pagination

	auc.Layout = "public/base.html"
	auc.TplName = "people/index.html"
}
