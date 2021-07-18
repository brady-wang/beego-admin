package controllers

import (
	"beego-admin/formvalidate"
	"beego-admin/global"
	"beego-admin/global/response"
	"beego-admin/services"
	"beego-admin/utils/exceloffice"
	"beego-admin/utils/template"
	"github.com/gookit/validate"
	"strconv"
	"strings"
	"time"
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

func init() {
	zoneIdList = append(zoneIdList, zoneId1, zoneId2, zoneId3, zoneId4, zoneId5)
}

// Index 家庭管理-首页
func (uc *FamilyController) Index() {
	var adminFamilyService services.AdminFamilyService
	data, pagination, totalArr := adminFamilyService.GetPaginateData(admin["per_page"].(int), gQueryParams)
	uc.Data["data"] = data
	uc.Data["paginate"] = pagination

	uc.Data["zoneIdList"] = zoneIdList
	uc.Data["totalArr"] = totalArr

	uc.Layout = "public/base.html"
	uc.TplName = "family/index.html"
}

// Add 家庭管理-添加界面
func (uc *FamilyController) Add() {
	uc.Data["zoneIdList"] = zoneIdList

	uc.Layout = "public/base.html"
	uc.TplName = "family/add.html"
}

// Create 家庭管理-添加界面
func (uc *FamilyController) Create() {
	var adminFamilyForm formvalidate.AdminFamilyForm
	if err := uc.ParseForm(&adminFamilyForm); err != nil {
		response.ErrorWithMessage(err.Error(), uc.Ctx)
	}
	roles := make([]string, 0)
	uc.Ctx.Input.Bind(&roles, "role")

	//账号验重
	var adminFamilyService services.AdminFamilyService
	if adminFamilyService.IsExistName(strings.TrimSpace(adminFamilyForm.Name), 0) {
		response.ErrorWithMessage("家庭已经存在", uc.Ctx)
	}

	insertID := adminFamilyService.Create(&adminFamilyForm)

	url := global.URL_BACK
	if adminFamilyForm.IsCreate == 1 {
		url = global.URL_RELOAD
	}

	if insertID > 0 {
		response.SuccessWithMessageAndUrl("添加成功", url, uc.Ctx)
	} else {
		response.Error(uc.Ctx)
	}
}

// Edit 系统管理-家庭管理-修改界面
func (uc *FamilyController) Edit() {
	id, _ := uc.GetInt("id", -1)
	if id <= 0 {
		response.ErrorWithMessage("Param is error.", uc.Ctx)
	}

	var (
		adminFamilyService services.AdminFamilyService
	)

	adminFamily := adminFamilyService.GetAdminFamilyById(id)
	if adminFamily == nil {
		response.ErrorWithMessage("Not Found Info By Id.", uc.Ctx)
	}

	uc.Data["zoneIdList"] = zoneIdList

	uc.Data["data"] = adminFamily

	uc.Layout = "public/base.html"
	uc.TplName = "family/edit.html"
}

// Update 系统管理-家庭管理-修改
func (uc *FamilyController) Update() {
	var adminFamilyForm formvalidate.AdminFamilyForm
	if err := uc.ParseForm(&adminFamilyForm); err != nil {
		response.ErrorWithMessage(err.Error(), uc.Ctx)
	}

	if adminFamilyForm.Id <= 0 {
		response.ErrorWithMessage("Params is Error.", uc.Ctx)
	}

	roles := make([]string, 0)
	uc.Ctx.Input.Bind(&roles, "role")

	v := validate.Struct(adminFamilyForm)

	if !v.Validate() {
		response.ErrorWithMessage(v.Errors.One(), uc.Ctx)
	}

	//账号验重
	var adminFamilyService services.AdminFamilyService
	if adminFamilyService.IsExistName(strings.TrimSpace(adminFamilyForm.Name), adminFamilyForm.Id) {
		response.ErrorWithMessage("家庭已经存在", uc.Ctx)
	}

	num := adminFamilyService.Update(&adminFamilyForm)

	if num > 0 {
		response.Success(uc.Ctx)
	} else {
		response.Error(uc.Ctx)
	}
}

// Del 删除
func (uc *FamilyController) Del() {
	idStr := uc.GetString("id")
	ids := make([]int, 0)
	var idArr []int

	if idStr == "" {
		uc.Ctx.Input.Bind(&ids, "id")
	} else {
		id, _ := strconv.Atoi(idStr)
		idArr = append(idArr, id)
	}

	if len(ids) > 0 {
		idArr = ids
	}

	if len(idArr) == 0 {
		response.ErrorWithMessage("参数id错误.", uc.Ctx)
	}

	var adminPeopleService services.AdminPeopleService

	for _, item := range idArr {
		peoples := adminPeopleService.GetCount(item)
		if peoples > 0 {
			response.ErrorWithMessage("ID为 "+strconv.Itoa(item)+"的家庭有人员存在 不能删除", uc.Ctx)
		}
	}

	var adminFamilyService services.AdminFamilyService
	count := adminFamilyService.Del(idArr)

	if count > 0 {
		response.SuccessWithMessageAndUrl("操作成功", global.URL_RELOAD, uc.Ctx)
	} else {
		response.Error(uc.Ctx)
	}
}

// Export 导出
func (uc *FamilyController) Export() {
	exportData := uc.GetString("export_data")
	if exportData == "1" {
		var adminFamilyService services.AdminFamilyService

		data := adminFamilyService.GetExportData(gQueryParams)
		header := []string{"家庭ID", "家庭名称", "人数", "区域", "详细地址", "创建时间", "更新时间"}
		body := [][]string{}
		for _, item := range data {
			record := []string{
				strconv.Itoa(item.Id),
				item.Name,
			}
			//获取总数
			var adminPeopleService services.AdminPeopleService
			number := adminPeopleService.GetCount(item.Id)

			record = append(record, strconv.Itoa(number))
			var zoneIdName string
			if item.ZoneId == 0 {
				zoneIdName = "未知"
			} else{
				zoneIdName = "洛佐"+strconv.Itoa(item.ZoneId)+"社"
			}
			record = append(record, zoneIdName)
			record = append(record, item.Address)

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
		exceloffice.ExportData(header, body, "family-"+time.Now().Format("2006-01-02-15-04-05"), "", "", uc.Ctx.ResponseWriter)
	}

	response.Error(uc.Ctx)
}
