package controllers

import (
	"github.com/astaxie/beego"
)

type AdminTodayRechargeController struct {
	beego.Controller
}

func (this *AdminTodayRechargeController) Get() {
	sess := this.StartSession()
	username := sess.Get("username")
	if username == nil || username == "" {
		this.Ctx.Redirect(302, "/Admin/login")
		return
	}
	this.Data["name"] = username
	this.Data["logintime"] = sess.Get("logintime")
	this.Data["Content"] = "今日充值"
	this.TplName = "admintodayrecharge.tpl"
}
func (this *AdminTodayRechargeController) Post() {
	sess := this.StartSession()
	quit := this.Input().Get("quit")
	if quit == "1" {
		sess.Delete("username")
		returnData := map[string]int{"code": 1}
		this.Data["json"] = &returnData
		this.ServeJSON()
	}
}
