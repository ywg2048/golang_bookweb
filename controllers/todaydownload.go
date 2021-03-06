package controllers

import (
	"DownLoadWeb/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"time"
)

type AdminTodayDownloadController struct {
	beego.Controller
}

func (this *AdminTodayDownloadController) Get() {
	sess := this.StartSession()
	username := sess.Get("username")
	if username == nil || username == "" {
		this.Ctx.Redirect(302, "/Admin/login")
		return
	}

	nTime := time.Now()
	yesTime := nTime.AddDate(0, 0, -1)
	logDay := yesTime.Format("2006-01-02 15:04:05")

	o := orm.NewOrm()
	var downloadrecord []models.DownLoadRecord
	_, err := o.QueryTable("down_load_record").Filter("downloadtime__gt", logDay).All(&downloadrecord)
	if err != nil {
		logs.Error("错误", err)
	}
	logs.Debug("data= ", downloadrecord)

	this.Data["downloadrecord"] = downloadrecord
	this.Data["name"] = username
	this.Data["logintime"] = sess.Get("logintime")
	this.Data["Content"] = "今日下载"
	this.TplName = "admintodaydown.tpl"
}
func (this *AdminTodayDownloadController) Post() {
	sess := this.StartSession()
	quit := this.Input().Get("quit")
	if quit == "1" {
		sess.Delete("username")
		returnData := map[string]int{"code": 1}
		this.Data["json"] = &returnData
		this.ServeJSON()
	}
}
