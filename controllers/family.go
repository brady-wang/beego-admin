package controllers

import (
	"beego-admin/formvalidate"
	"beego-admin/global"
	"beego-admin/global/response"
	"beego-admin/services"
	"github.com/gookit/validate"
	"strconv"
	"strings"
)

// FamilyController struct
type FamilyController struct {
	baseController
}

var zoneIdList = make([]map[string]interface{}, 0)
var zoneId1 = map[string]interface{}{"Id": 1, "Name": "洛佐1社"}
var zoneId2 = map[string]interface{}{"Id": 2, "Name": "洛佐2社"}
var zoneId3 = map[string]interface{}{"Id": 3, "Name": "洛佐3社"}
var zoneId4 = map[string]interface{}{"Id": 4, "Name": "洛佐4社"}
var zoneId5 = map[string]interface{}{"Id": 5, "Name": "洛佐5社"}

func init()  {
	zoneIdList = append(zoneIdList, zoneId1, zoneId2, zoneId3, zoneId4, zoneId5)
}
// Index 家庭管理-首页
func (auc *FamilyController) Index() {
	var adminFamilyService services.AdminFamilyService
	data, pagination,totalArr:= adminFamilyService.GetPaginateData(admin["per_page"].(int), gQueryParams)
	auc.Data["data"] = data
	auc.Data["paginate"] = pagination

	auc.Data["zoneIdList"] = zoneIdList
	auc.Data["totalArr"] = totalArr

	auc.Layout = "public/base.html"
	auc.TplName = "family/index.html"
}

// Add 家庭管理-添加界面
func (auc *FamilyController) Add() {
	auc.Data["zoneIdList"] = zoneIdList

	auc.Layout = "public/base.html"
	auc.TplName = "family/add.html"
}

// Create 家庭管理-添加界面
func (auc *FamilyController) Create() {
	var adminFamilyForm formvalidate.AdminFamilyForm
	if err := auc.ParseForm(&adminFamilyForm); err != nil {
		response.ErrorWithMessage(err.Error(), auc.Ctx)
	}
	roles := make([]string, 0)
	auc.Ctx.Input.Bind(&roles, "role")

	//账号验重
	var adminFamilyService services.AdminFamilyService
	if adminFamilyService.IsExistName(strings.TrimSpace(adminFamilyForm.Name), 0) {
		response.ErrorWithMessage("家庭已经存在", auc.Ctx)
	}

	insertID := adminFamilyService.Create(&adminFamilyForm)

	url := global.URL_BACK
	if adminFamilyForm.IsCreate == 1 {
		url = global.URL_RELOAD
	}

	if insertID > 0 {
		response.SuccessWithMessageAndUrl("添加成功", url, auc.Ctx)
	} else {
		response.Error(auc.Ctx)
	}
}

// Edit 系统管理-家庭管理-修改界面
func (auc *FamilyController) Edit() {
	id, _ := auc.GetInt("id", -1)
	if id <= 0 {
		response.ErrorWithMessage("Param is error.", auc.Ctx)
	}

	var (
		adminFamilyService services.AdminFamilyService
	)

	adminFamily := adminFamilyService.GetAdminFamilyById(id)
	if adminFamily == nil {
		response.ErrorWithMessage("Not Found Info By Id.", auc.Ctx)
	}

	auc.Data["zoneIdList"] = zoneIdList

	auc.Data["data"] = adminFamily

	auc.Layout = "public/base.html"
	auc.TplName = "family/edit.html"
}

// Update 系统管理-家庭管理-修改
func (auc *FamilyController) Update() {
	var adminFamilyForm formvalidate.AdminFamilyForm
	if err := auc.ParseForm(&adminFamilyForm); err != nil {
		response.ErrorWithMessage(err.Error(), auc.Ctx)
	}

	if adminFamilyForm.Id <= 0 {
		response.ErrorWithMessage("Params is Error.", auc.Ctx)
	}

	roles := make([]string, 0)
	auc.Ctx.Input.Bind(&roles, "role")

	v := validate.Struct(adminFamilyForm)

	if !v.Validate() {
		response.ErrorWithMessage(v.Errors.One(), auc.Ctx)
	}

	//账号验重
	var adminFamilyService services.AdminFamilyService
	if adminFamilyService.IsExistName(strings.TrimSpace(adminFamilyForm.Name), adminFamilyForm.Id) {
		response.ErrorWithMessage("家庭已经存在", auc.Ctx)
	}

	num := adminFamilyService.Update(&adminFamilyForm)

	if num > 0 {
		response.Success(auc.Ctx)
	} else {
		response.Error(auc.Ctx)
	}
}

// Del 删除
func (auc *FamilyController) Del() {
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

	var adminFamilyService services.AdminFamilyService
	count := adminFamilyService.Del(idArr)

	if count > 0 {
		response.SuccessWithMessageAndUrl("操作成功", global.URL_RELOAD, auc.Ctx)
	} else {
		response.Error(auc.Ctx)
	}
}
