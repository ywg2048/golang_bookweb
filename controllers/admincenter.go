package controllers

import (
	"DownLoadWeb/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type AdminCenterController struct {
	beego.Controller
}

func (this *AdminCenterController) Get() {
	sess := this.StartSession()
	username := sess.Get("username")
	if username == nil || username == "" {
		this.Ctx.Redirect(302, "/Admin/login")
		return
	}
	o := orm.NewOrm()
	var data models.AdminUser
	err := o.QueryTable("admin_user").Filter("username__eq", username).One(&data)
	if err != nil {
		logs.Error("找不到用户", err)
	}
	sess.Set("logintime", data.LoginTime.Format("2006-01-02 15:04:05"))
	this.Data["name"] = username
	this.Data["logintime"] = data.LoginTime.Format("2006-01-02 15:04:05")
	this.TplName = "admincenter.tpl"
}

func (this *AdminCenterController) Post() {
	sess := this.StartSession()
	quit := this.Input().Get("quit")
	if quit == "1" {
		sess.Delete("username")
		returnData := map[string]int{"code": 1}
		this.Data["json"] = &returnData
		this.ServeJSON() 
	}
}
