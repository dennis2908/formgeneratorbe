package routers

import (
	"forms/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/forms/group", &controllers.FormsController{}, "get:GetAllForms")
	beego.Router("/forms/all/field", &controllers.FormsController{}, "post:GetAllField")
	beego.Router("/forms/all/data", &controllers.FormsController{}, "get:GetAllData")
	beego.Router("/forms/all/list", &controllers.FormsController{}, "get:ListForms")
	beego.Router("/forms/generate", &controllers.FormsController{}, "get:GenerateForms")
	beego.Router("/forms", &controllers.FormsController{}, "post:CreateDatas")
}
