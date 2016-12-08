package controllers

import (
	"DownLoadWeb/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type UserCenterController struct {
	beego.Controller
}

func (this *UserCenterController) Get() {
	sess := this.StartSession()
	user_name := sess.Get("user_name")
	if user_name == "" || user_name == nil {
		this.Ctx.Redirect(302, "/User/login")
	}
	//断言
	uname, ok := user_name.(string)
	if !ok {
		logs.Error("断言失败")
	}
	//下载记录
	o := orm.NewOrm()
	var downloadrecord []models.DownLoadRecord
	_, err := o.QueryTable("down_load_record").Filter("username__eq", uname).Limit(2, 0).All(&downloadrecord)
	if err != nil {
		logs.Error("错误", err)
	}
	//登录地址记录
	var userloginip []models.UserLoginIp
	_, err1 := o.QueryTable("userloginip").Filter("username_eq", uname).Limit(2, 0).All(&userloginip)
	this.Data["downloadrecord"] = downloadrecord
	this.Data["username"] = uname
	this.TplName = "usercenter.tpl"
}

func (this *UserCenterController) Post() {

}
