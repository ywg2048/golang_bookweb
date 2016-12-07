package controllers

import (
	"github.com/astaxie/beego"
)

type AdminTotalIncomeController struct {
	beego.Controller
}

func (this *AdminTotalIncomeController) Get() {
	sess := this.StartSession()
	username := sess.Get("username")
	if username == nil || username == "" {
		this.Ctx.Redirect(302, "/Admin/login")
		return
	}
	this.Data["name"] = username
	this.Data["logintime"] = sess.Get("logintime")
	this.Data["Content"] = "收入总计"
	this.TplName = "admintodayrecharge.tpl"
}
