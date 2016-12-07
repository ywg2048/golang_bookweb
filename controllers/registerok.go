package controllers

import (
	"github.com/astaxie/beego"
)

type UserRegisterokController struct {
	beego.Controller
}

func (this *UserRegisterokController) Get() {

	this.TplName = "registerok.tpl"
}
