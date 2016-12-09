package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type AdminUser struct {
	Id        int64 `orm:"pk;auto"`
	UserName  string
	PassWd    string
	LoginTime time.Time
}
type User struct {
	Id             int64 `orm:"pk;auto"`
	UserName       string
	Email          string
	PassWd         string
	IsVip          bool
	GetVipTime     time.Time
	VipType        int32 //1:年费会员，2:终身会员
	DownLoadCounts int64
	LoginTime      time.Time
	RegisterTime   time.Time
}
type UserLoginIp struct {
	Id        int64 `orm:"pk;auto"`
	UserName  string
	Ip        string
	Country   string
	Region    string
	City      string
	LoginTime time.Time
}
type UploadFile struct {
	Id         int64 `orm:"pk;auto"`
	Path       string
	FileName   string
	Size       int64
	Type       string
	UploadTime time.Time
}
type DownLoadRecord struct {
	Id           int64 `orm:"pk;auto"`
	UserName     string
	FileName     string
	DownLoadTime time.Time
	Ip           string
}

func RegisterDB() {
	//注册model
	orm.RegisterModel(new(AdminUser), new(User), new(UserLoginIp), new(UploadFile), new(DownLoadRecord))
	//注册驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//注册默认数据库
	orm.RegisterDataBase("default", "mysql", "root:@/bookweb?charset=utf8&loc=Local")

}
