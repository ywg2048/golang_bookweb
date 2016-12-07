package controllers

import (
	"DownLoadWeb/models"
	"crypto/md5"
	"encoding/hex"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"time"
)

type UserLoginController struct {
	beego.Controller
}

func (this *UserLoginController) Get() {
	sess := this.StartSession()
	user_name := sess.Get("user_name")
	logs.Debug("user_name:", user_name)
	if user_name != nil && user_name != "" {
		this.Ctx.Redirect(302, "/")
	} else {
		this.TplName = "userlogin.tpl"
	}
}
func (this *UserLoginController) Post() {
	sess := this.StartSession()
	username := this.Input().Get("username")
	passwd := this.Input().Get("passwd")

	timestamp := time.Now().Unix()
	tm := time.Unix(timestamp, 0)

	h := md5.New()
	h.Write([]byte(passwd)) //使用zhifeiya名字做散列值，设定后不要变
	passwd = hex.EncodeToString(h.Sum(nil))

	o := orm.NewOrm()
	var data models.User
	err := o.QueryTable("user").Filter("username__eq", username).Filter("passwd__eq", passwd).One(&data)
	if err != nil {
		logs.Error("没有这个用户", err)
		this.Data["err"] = "用户名或密码错误"
		this.TplName = "userlogin.tpl"
	} else {
		data.LoginTime = tm
		o.Update(&data)
		sess.Set("user_name", username)
		this.Ctx.Redirect(302, "/")
	}
}
