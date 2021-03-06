package controllers

import (
	"beego-admin/formvalidate"
	"beego-admin/global"
	"beego-admin/global/response"
	"beego-admin/services"
	"beego-admin/utils/exceloffice"
	"beego-admin/utils/template"
	"fmt"
	"github.com/gookit/validate"
	"strconv"
	"strings"
	"time"
)

// PeopleController struct
type PeopleController struct {
	baseController
}

// Index 人员管理-首页
func (auc *PeopleController) Index() {

	var adminFamilyService services.AdminFamilyService
	userLevel := adminFamilyService.GetAllFamily()
	familyMap := make(map[int]string)
	for _, item := range userLevel {
		familyMap[item.Id] = item.Name
	}
	auc.Data["familyMap"] = familyMap

	fmt.Println(gQueryParams)
	var adminPeopleService services.AdminPeopleService
	data, pagination := adminPeopleService.GetPaginateData(admin["per_page"].(int), gQueryParams)
	auc.Data["data"] = data
	fmt.Println(data)
	auc.Data["paginate"] = pagination

	auc.Layout = "public/base.html"
	auc.TplName = "people/index.html"
}

// Add 人员管理-添加界面
func (auc *PeopleController) Add() {

	var adminFamilyService services.AdminFamilyService
	familyList := adminFamilyService.GetAllFamily()
	auc.Data["familyList"] = familyList

	auc.Layout = "public/base.html"
	auc.TplName = "people/add.html"
}

// Create 人员管理-添加界面
func (auc *PeopleController) Create() {
	var adminPeopleForm formvalidate.AdminPeopleForm
	if err := auc.ParseForm(&adminPeopleForm); err != nil {
		response.ErrorWithMessage(err.Error(), auc.Ctx)
	}
	roles := make([]string, 0)
	auc.Ctx.Input.Bind(&roles, "role")

	//账号验重
	var adminPeopleService services.AdminPeopleService
	if adminPeopleService.IsExistName(strings.TrimSpace(adminPeopleForm.Name), 0) {
		response.ErrorWithMessage("人员已经存在", auc.Ctx)
	}

	insertID := adminPeopleService.Create(&adminPeopleForm)

	url := global.URL_BACK
	if adminPeopleForm.IsCreate == 1 {
		url = global.URL_RELOAD
	}

	if insertID > 0 {
		response.SuccessWithMessageAndUrl("添加成功", url, auc.Ctx)
	} else {
		response.Error(auc.Ctx)
	}
}

// Edit 系统管理-人员管理-修改界面

func (auc *PeopleController) Edit() {
	var adminFamilyService services.AdminFamilyService
	var adminPeopleService services.AdminPeopleService

	familyList := adminFamilyService.GetAllFamily()
	auc.Data["familyList"] = familyList
	fmt.Println(familyList)
	fmt.Println(familyList)
	id, _ := auc.GetInt("id", -1)
	if id <= 0 {
		response.ErrorWithMessage("Param is error.", auc.Ctx)
	}

	adminPeople := adminPeopleService.GetAdminPeopleById(id)
	if adminPeople == nil {
		response.ErrorWithMessage("Not Found Info By Id.", auc.Ctx)
	}

	auc.Data["data"] = adminPeople
	auc.Layout = "public/base.html"
	auc.TplName = "people/edit.html"
}

// Update 系统管理-人员管理-修改
func (auc *PeopleController) Update() {
	var adminPeopleForm formvalidate.AdminPeopleForm
	if err := auc.ParseForm(&adminPeopleForm); err != nil {
		response.ErrorWithMessage(err.Error(), auc.Ctx)
	}

	if adminPeopleForm.Id <= 0 {
		response.ErrorWithMessage("Params is Error.", auc.Ctx)
	}

	roles := make([]string, 0)
	auc.Ctx.Input.Bind(&roles, "role")

	v := validate.Struct(adminPeopleForm)

	if !v.Validate() {
		response.ErrorWithMessage(v.Errors.One(), auc.Ctx)
	}

	//账号验重
	var adminPeopleService services.AdminPeopleService
	if adminPeopleService.IsExistName(strings.TrimSpace(adminPeopleForm.Name), adminPeopleForm.Id) {
		response.ErrorWithMessage("人员已经存在", auc.Ctx)
	}

	num := adminPeopleService.Update(&adminPeopleForm)

	if num > 0 {
		response.Success(auc.Ctx)
	} else {
		response.Error(auc.Ctx)
	}
}

// Del 删除
func (auc *PeopleController) Del() {
	idStr := auc.GetString("id")
	ids := make([]int, 0)
	var idArr []int

	if idStr == "" {
		auc.Ctx.Input.Bind(&ids, "id")
	} else {
		id, _ := strconv.Atoi(idStr)
		idArr = append(idArr, id)
	}

	if len(ids) > 0 {
		idArr = ids
	}

	if len(idArr) == 0 {
		response.ErrorWithMessage("参数id错误.", auc.Ctx)
	}

	var adminPeopleService services.AdminPeopleService
	count := adminPeopleService.Del(idArr)

	if count > 0 {
		response.SuccessWithMessageAndUrl("操作成功", global.URL_RELOAD, auc.Ctx)
	} else {
		response.Error(auc.Ctx)
	}
}


// Export 导出
func (uc *PeopleController) Export() {
	exportData := uc.GetString("export_data")
	if exportData == "1" {
		var adminPeopleService services.AdminPeopleService

		data := adminPeopleService.GetExportData(gQueryParams)
		header := []string{"ID", "姓名", "性别", "手机号", "地址", "身份证", "社保卡","家庭ID","创建时间","更新时间"}
		body := [][]string{}
		for _, item := range data {
			record := []string{
				strconv.Itoa(item.Id),
				item.Name,
			}

			record = append(record, item.Sex)
			record = append(record, item.Mobile)
			record = append(record, item.Address)
			record = append(record, item.IdCard)
			record = append(record, item.SocialCard)
			record = append(record, strconv.Itoa(item.FamilyId))
			var createdAt string
			var updatedAt string
			if item.CreatedAt >0 {
				createdAt = template.UnixTimeForFormat(item.CreatedAt)
			}
			if item.UpdatedAt >0 {
				updatedAt = template.UnixTimeForFormat(item.UpdatedAt)
			}
			record = append(record, createdAt)
			record = append(record, updatedAt)
			body = append(body, record)
		}
		uc.Ctx.ResponseWriter.Header().Set("a", "b")
		exceloffice.ExportData(header, body, "people-"+time.Now().Format("2006-01-02-15-04-05"), "", "", uc.Ctx.ResponseWriter)
	}

	response.Error(uc.Ctx)
}