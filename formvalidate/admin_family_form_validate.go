package formvalidate

import "github.com/gookit/validate"

// AdminUserForm admin_user 表单
type AdminFamilyForm struct {
	Id       int    `form:"id"`
	Name     string `form:"name" validate:"required"`
	Number   string `form:"number" `
	ZoneId   int    `form:"zone_id" `
	Address  string `form:"address" `
	IsCreate int    `form:"_create"`
}

// Messages 自定义验证返回消息
func (f AdminFamilyForm) Messages() map[string]string {
	return validate.MS{
		"Name.required": "请填写家庭名称.",
	}
}
