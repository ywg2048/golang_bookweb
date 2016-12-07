package controllers

import (
	"DownLoadWeb/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type AdminListController struct {
	beego.Controller
}

func (this *AdminListController) Get() {
	sess := this.StartSession()
	username := sess.Get("username")
	if username == nil || username == "" {
		this.Ctx.Redirect(302, "/Admin/login")
		return
	}

	//拉取用户列表
	o := orm.NewOrm()
	var userlist []models.User
	_, err := o.QueryTable("user").All(&userlist)
	if err != nil {
		logs.Error("错误", err)
	}

	this.Data["name"] = username
	this.Data["logintime"] = sess.Get("logintime")
	this.Data["Content"] = userlist
	this.TplName = "adminlist.tpl"
}
func (this *AdminListController) Post() {
	sess := this.StartSession()
	quit := this.Input().Get("quit")
	if quit == "1" {
		sess.Delete("username")
		returnData := map[string]int{"code": 1}
		this.Data["json"] = &returnData
		this.ServeJSON()
	}
}
