package main

import (
	"DownLoadWeb/models"
	_ "DownLoadWeb/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	//注册数据库
	models.RegisterDB()
}
func main() {
	//开启ORM调试模式
	orm.Debug = true
	//自动建表
	orm.RunSyncdb("default", false, true)
	//注册静态目录
	beego.SetStaticPath("/static", "static")
	//运行
	beego.Run()
}
