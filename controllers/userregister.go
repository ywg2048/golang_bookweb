package controllers

import (
	"DownLoadWeb/models"
	"crypto/md5"
	"encoding/hex"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"math/rand"
	"net/smtp"
	"strconv"
	"strings"
	"time"
)

type UserRegisterController struct {
	beego.Controller
}

func (this *UserRegisterController) Get() {

	this.TplName = "register.tpl"
}

func (this *UserRegisterController) Post() {
	sess := this.StartSession()
	var returnData map[string]int
	username := this.Input().Get("username")
	email := this.Input().Get("email")
	checkcode := this.Input().Get("checkcode")
	passwd := this.Input().Get("passwd")

	if username != "" && passwd == "" && email == "" && checkcode == "" {
		logs.Debug("username = ", username)
		//判断用户是否已经注册
		o := orm.NewOrm()
		var data models.User
		err := o.QueryTable("user").Filter("username__eq", username).One(&data)
		logs.Debug("错误：", err)
		if err != nil {
			returnData = map[string]int{"code": 0}
		} else {
			returnData = map[string]int{"code": 1}
		}
		this.Data["json"] = &returnData
		this.ServeJSON()

	} else if username != "" && passwd != "" && email != "" && checkcode != "" {
		sess_code := sess.Get("code")
		if sess_code != checkcode {
			logs.Debug("验证码错误")
			this.Data["errtext"] = "验证码错误！"
			this.TplName = "register.tpl"
		} else {
			timestamp := time.Now().Unix()
			tm := time.Unix(timestamp, 0)

			h := md5.New()
			h.Write([]byte(passwd)) //使用zhifeiya名字做散列值，设定后不要变
			passwd = hex.EncodeToString(h.Sum(nil))

			o := orm.NewOrm()
			var user models.User
			user.UserName = username
			user.Email = email
			user.PassWd = passwd
			user.IsVip = false
			user.RegisterTime = tm
			user.GetVipTime = tm
			user.VipType = 0
			user.DownLoadCounts = 0
			user.LoginTime = tm
			id, err := o.Insert(&user)
			logs.Debug("id = ", id)
			logs.Debug("err = ", err)
			//删除code session
			sess.Delete("code")
			sess.Set("user_name", username)
			this.Ctx.Redirect(302, "/User/registerok")
		}

	}
	getcode := this.Input().Get("getcode")

	if getcode == "1" {
		code := RandInt()
		logs.Debug("code is ", code)
		err := SendMail("1428261722@qq.com", "2563890789a", "smtp.qq.com:25", email, "验证码", "您的验证码是"+code, "html")
		if err != nil {
			logs.Debug("发送失败", err)
		}
		sess.Set("code", code)
		returnData = map[string]int{"code": 1}
		this.Data["json"] = &returnData
		this.ServeJSON()
	}

}
func RandInt() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := ""
	for i := 0; i < 4; i++ {
		code += strconv.Itoa(r.Intn(9))
	}
	return code
}
func SendMail(user, password, host, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}
	msg := []byte("To: " + to + "\r\nFrom: " + user + "<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}
