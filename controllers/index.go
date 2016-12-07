package controllers

import (
	"DownLoadWeb/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"net"
	"os"
	"strconv"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	logs.Debug("Index is starting....")
	sess := this.StartSession()
	username := sess.Get("user_name")
	logs.Debug("user_name:", username)
	if username != "" && username != nil {
		this.Data["username"] = username
		this.Data["islogin"] = true
		logs.Debug("已经登录")
	} else {
		this.Data["islogin"] = false
		logs.Debug("没有登录")
	}

	o := orm.NewOrm()
	var files []models.UploadFile

	//获取页面显示数量
	PageSize, _ := beego.AppConfig.Int("pagesize")
	//获取记录总数
	o.QueryTable("upload_file").Filter("id__gt", 0).All(&files)
	TotalRecord := len(files)
	//获取最大页数
	TotalPages := TotalRecord / PageSize
	if TotalRecord%PageSize > 0 {
		TotalPages = TotalRecord/PageSize + 1
	}
	//索引
	Index := 0
	if this.Input().Get("Index") != "" {
		Index, _ = strconv.Atoi(this.Input().Get("Index"))
		if Index >= TotalPages {
			Index = TotalPages - 1
		}
	}
	//上一页
	PrePage := Index - 1
	if PrePage <= 0 {
		PrePage = 0
	}
	//下一页
	Nextpage := Index + 1
	if Nextpage >= TotalPages {
		Nextpage = TotalPages - 1
	}

	ip := getip()
	port := beego.AppConfig.String("httpport")
	pagelistsize, _ := beego.AppConfig.Int("pagelistsize")

	o.QueryTable("upload_file").Filter("id__gt", 0).Limit(PageSize, Index*PageSize).All(&files)
	for i := 0; i < len(files); i++ {
		files[i].Path = "http://" + ip + ":" + port + "/DownLoad/" + files[i].FileName
	}
	this.Data["totalpages"] = TotalPages
	this.Data["currentpage"] = Index + 1
	this.Data["nextpage"] = Nextpage
	this.Data["prepage"] = PrePage
	this.Data["files"] = files
	this.Data["lastpage"] = TotalPages - 1
	this.Data["pagelistsize"] = pagelistsize
	this.TplName = "index.tpl"
}
func (this *MainController) Post() {
	sess := this.StartSession()
	quit := this.Input().Get("quit")
	if quit == "1" {
		sess.Delete("user_name")
		returnData := map[string]int{"code": 1}
		this.Data["json"] = &returnData
		this.ServeJSON()
	}
}
func getip() string {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		logs.Error(err)
		os.Exit(1)
	}

	for _, address := range addrs {

		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}

		}
	}
	return "127.0.0.1"
}
