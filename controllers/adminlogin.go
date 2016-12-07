package controllers

import (
	"DownLoadWeb/models"
	"crypto/md5"
	"encoding/hex"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type AdminLoginController struct {
	beego.Controller
}

func (this *AdminLoginController) Get() {
	sess := this.StartSession()
	username := sess.Get("username")
	logs.Debug("session username: ", username)
	if username == nil || username == "" {
		this.TplName = "adminlogin.tpl"
	} else {
		this.Ctx.Redirect(302, "/Admin/center")
	}

}

func (this *AdminLoginController) Post() {
	sess := this.StartSession()
	username := this.Input().Get("username")
	passwd := this.Input().Get("passwd")

	h := md5.New()
	h.Write([]byte(passwd)) //使用zhifeiya名字做散列值，设定后不要变
	passwd = hex.EncodeToString(h.Sum(nil))
	logs.Debug("username =", username)
	logs.Debug("passwd = ", passwd)
	o := orm.NewOrm()
	var data models.AdminUser
	err := o.QueryTable("admin_user").Filter("username__eq", username).One(&data)
	if err == nil {
		logs.Debug("data", data.PassWd)
		if data.PassWd == passwd {
			sess.Set("username", data.UserName)
			this.Ctx.Redirect(302, "/Admin/center")
		} else {
			this.Data["err"] = "密码或用户名错误"
			this.TplName = "adminlogin.tpl"
		}
	} else {
		this.Data["err"] = "密码或用户名错误"
		this.TplName = "adminlogin.tpl"
	}

}
