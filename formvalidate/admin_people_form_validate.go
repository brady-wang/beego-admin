package formvalidate

import "github.com/gookit/validate"

// AdminUserForm admin_user 表单
type AdminPeopleForm struct {
	Id         int    `form:"id"`
	Name       string `form:"name" validate:"required"`
	Sex        string `form:"sex" `
	Mobile     string `form:"mobile" `
	Address    string `form:"address" `
	IdCard     string `form:"id_card" `
	FamilyId   int    `form:"family_id" validate:"required" `
	SocialCard string `form:"social_card" `
	IsCreate   int    `form:"_create"`
}

// Messages 自定义验证返回消息
func (f AdminPeopleForm) Messages() map[string]string {
	return validate.MS{
		"Name.required": "请填写人员名称.",
	}
}
