package controllers

import (
	"encoding/json"
	"fmt"
	_ "fmt"
	models "forms/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/leekchan/accounting"
	_ "github.com/shopspring/decimal"
)

type FormsController struct {
	beego.Controller
}

type formsGroup struct {
	Form_id string
}

func (api *FormsController) GetAllForms() {
	o := orm.NewOrm()
	o.Using("default")
	sql := "select form_id from forms where transform = 'y' group by form_id"

	var formsGroup []formsGroup
	_, err := o.Raw(sql).QueryRows(&formsGroup)

	if err == nil {
		// ... handle error
		api.Data["json"] = formsGroup
	}

	api.ServeJSON()
}

func (api *FormsController) GetAllField() {
	frm := api.Ctx.Input.RequestBody
	o := orm.NewOrm()
	o.Using("default")

	u := &formsGroup{}
	json.Unmarshal(frm, u)
	fmt.Println(frm)
	query := fmt.Sprintf(`
	select * from forms where form_id = '%s' and transform = 'y'
	`, u.Form_id)

	var formsGroup []models.Forms
	_, err := o.Raw(query).QueryRows(&formsGroup)

	if err == nil {
		// ... handle error
		api.Data["json"] = formsGroup
	}

	api.ServeJSON()
}

func (api *FormsController) GetAllData() {
	o := orm.NewOrm()
	o.Using("default")
	sql := "select * from forms order by form_id"

	var formsGroup []models.Forms
	_, err := o.Raw(sql).QueryRows(&formsGroup)

	if err == nil {
		// ... handle error
		api.Data["json"] = formsGroup
	}

	api.ServeJSON()
}

func (api *FormsController) ListForms() {
	o := orm.NewOrm()
	o.Using("default")
	sql := "select form_id from forms group by form_id order by form_id"

	var formsGroup []formsGroup
	_, err := o.Raw(sql).QueryRows(&formsGroup)

	if err == nil {
		// ... handle error
		api.Data["json"] = formsGroup
	}

	api.ServeJSON()
}

func (api *FormsController) GenerateForms() {
	o := orm.NewOrm()
	o.Using("default")
	sql := "update forms set transform = 'y'"

	var formsGroup []formsGroup
	_, err := o.Raw(sql).QueryRows(&formsGroup)

	if err == nil {
		// ... handle error
		api.Data["json"] = formsGroup
	}

	api.ServeJSON()
}

func (api *FormsController) CreateDatas() {

	frm := api.Ctx.Input.RequestBody

	o := orm.NewOrm()
	o.Using("default")

	u := &models.Forms{}
	json.Unmarshal(frm, u)
	fmt.Println(frm)

	// // insert
	PostsQry := models.Forms{FormId: u.FormId, FieldName: u.FieldName, FieldHtml: u.FieldHtml}

	_, err := o.Insert(&PostsQry)

	if err != nil {
		api.Data["json"] = err.Error()
		api.ServeJSON()
	}
	api.Data["json"] = "Successfully save data "
	api.ServeJSON()
}
