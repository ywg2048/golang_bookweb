package routers

import (
	"DownLoadWeb/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/upload", &controllers.UploadController{})
	beego.Router("/Add", &controllers.AddController{})
	beego.Router("/Admin/login", &controllers.AdminLoginController{})
	beego.Router("/Admin/center", &controllers.AdminCenterController{})
	beego.Router("/Admin/list", &controllers.AdminListController{})
	beego.Router("/Admin/todayincome", &controllers.AdminTodayIncomeController{})
	beego.Router("/Admin/todayrecharge", &controllers.AdminTodayRechargeController{})
	beego.Router("/Admin/todaydownload", &controllers.AdminTodayDownloadController{})
	beego.Router("/Admin/totalincome", &controllers.AdminTotalIncomeController{})

	beego.Router("/User/register", &controllers.UserRegisterController{})
	beego.Router("/User/registerok", &controllers.UserRegisterokController{})
	beego.Router("/User/login", &controllers.UserLoginController{})
	beego.Router("/User/center", &controllers.UserCenterController{})

	beego.Router("/DownLoad/:filename", &controllers.DownLoadController{})

}
