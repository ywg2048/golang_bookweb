package controllers

import (
	"DownLoadWeb/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"time"
)

type UploadController struct {
	beego.Controller
}

func (c *UploadController) Get() {
	beego.SetLogFuncCall(true)
	logs.Debug("GET is start***")
	c.TplName = "upload.tpl"
}

func (c *UploadController) Post() {

	timestamp := time.Now().Unix()
	tm := time.Unix(timestamp, 0)

	f, h, err := c.GetFile("image")
	defer f.Close()
	if err != nil {
		fmt.Println("getfile err ", err)
	} else {
		err1 := c.SaveToFile("image", "static/upload/"+h.Filename) // 保存位置在 static/upload,没有文件夹要先创建
		if err1 == nil {
			logs.Debug("保存成功")
			o := orm.NewOrm()
			var uploadfile models.UploadFile
			uploadfile.Path = "static/upload/"
			uploadfile.FileName = h.Filename
			uploadfile.Size = 0
			uploadfile.Type = h.Filename
			uploadfile.UploadTime = tm
			o.Insert(&uploadfile)
		} else {
			c.Ctx.WriteString("ok")
		}
	}
	c.Ctx.WriteString("ok")
}
