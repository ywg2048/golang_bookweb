package controllers

import (
	"github.com/astaxie/beego"
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
		filename := this.Ctx.Input.Param(":filename")
		this.Ctx.Output.Download("static/upload/" + filename)
	}

}
