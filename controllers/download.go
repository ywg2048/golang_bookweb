package controllers

import (
	"DownLoadWeb/models"
	"DownLoadWeb/toolfunction"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"time"
)

type DownLoadController struct {
	beego.Controller
}

func (this *DownLoadController) Get() {
	sess := this.StartSession()
	user_name := sess.Get("user_name")
	if user_name == nil || user_name == "" {
		//用户未登录
		this.Ctx.Redirect(302, "/User/login")
	} else {
		//登录可以暂时可以下载
		timestamp := time.Now().Unix()
		tm := time.Unix(timestamp, 0)

		filename := this.Ctx.Input.Param(":filename")
		value, ok := user_name.(string)
		if !ok {
			logs.Debug("类型断言失败")
		}
		o := orm.NewOrm()
		downloadrecord := new(models.DownLoadRecord)
		downloadrecord.UserName = value
		downloadrecord.FileName = filename
		downloadrecord.Ip = toolfunction.Getip()
		downloadrecord.DownLoadTime = tm
		o.Insert(downloadrecord)

		//username断言
		uname, ok := user_name.(string)
		if !ok {
			logs.Error("断言失败")
		}
		var user models.User
		o.QueryTable("user").Filter("username__eq", uname).One(&user)
		user.DownLoadCounts++
		o.Update(&user)
		this.Ctx.Output.Download("static/upload/" + filename)
	}

}
