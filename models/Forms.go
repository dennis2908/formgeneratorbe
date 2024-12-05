package models

import (
	"github.com/astaxie/beego/orm"
)

type Forms struct {
	Id        int `orm:"auto;pk;index"`
	FormId    string
	FieldName string `orm:"size(200)"`
	FieldHtml string `orm:"size(100)"`
	Transform string
}

func (a *Forms) TableName() string {
	return "forms"
}

func init() {
	orm.RegisterModel(new(Forms))
}
