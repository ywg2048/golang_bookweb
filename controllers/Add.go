package controllers

import (
	"github.com/astaxie/beego"
)

type AddController struct {
	beego.Controller
}

type User struct {
	Id    int         `form:"-"`
	Name  interface{} `form:"username"`
	Age   int         `form:"age,text,年龄："`
	Sex   string
	Intro string `form:",textarea"`
}

func (this *AddController) Get() {
	this.Data["Form"] = &User{}
	this.TplName = "bind.tpl"
}
