package controllers

import (
	"DownLoadWeb/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"time"
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
	//User信息
	o := orm.NewOrm()
	var user models.User
	o.QueryTable("user").Filter("username__eq", uname).One(&user)

	VIP := "否"
	var VIPExprice time.Time
	VIPType := 0
	var GetVIPTime time.Time
	if user.IsVip {
		VIP = "是"
		if user.VipType == 1 {
			//年费会员
			VIPType = 1
			baseTime := user.GetVipTime
			date := baseTime.Add(365 * 24 * time.Hour)
			VIPExprice = date
		} else if user.VipType == 2 {
			//终身会员
			VIPType = 1
			baseTime := user.GetVipTime
			date := baseTime.Add(365 * 24 * time.Hour)
			VIPExprice = date
			GetVIPTime = user.GetVipTime
		}

	}
	//下载记录
	var downloadrecord []models.DownLoadRecord
	_, err := o.QueryTable("down_load_record").Filter("username__eq", uname).Limit(2, 0).All(&downloadrecord)
	if err != nil {
		logs.Error("错误", err)
	}
	//登录地址记录
	var userloginip []models.UserLoginIp
	o.QueryTable("user_login_ip").Filter("username__eq", uname).Limit(10, 0).All(&userloginip)
	this.Data["userloginip"] = userloginip
	this.Data["downloadrecord"] = downloadrecord
	this.Data["username"] = uname
	this.Data["vip"] = VIP
	this.Data["VIPExprice"] = VIPExprice
	this.Data["VIPType"] = VIPType
	this.Data["GetVIPTime"] = GetVIPTime
	this.TplName = "usercenter.tpl"
}

func (this *UserCenterController) Post() {

}
