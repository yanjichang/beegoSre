package routers

import (
	"github.com/Ohimma/beegoSre/controllers"
	"github.com/astaxie/beego"
)

func init() {
	// 默认登录
	// beego.Router("/doc", &controllers.ApiDocController{}, "*:Index")

	beego.Router("/", &controllers.LoginController{}, "*:LoginIn")
	beego.Router("/login", &controllers.LoginController{}, "*:LoginIn")
	beego.Router("/login_out", &controllers.LoginController{}, "*:LoginOut")
	beego.Router("/no_auth", &controllers.LoginController{}, "*:NoAuth")

	beego.Router("/home", &controllers.HomeController{}, "*:Index")
	beego.Router("/home/start", &controllers.HomeController{}, "*:Start")

	// spn 管理
	beego.AutoRouter(&controllers.FengxuanController{})

	// 权限设置
	beego.AutoRouter(&controllers.AuthController{})
	beego.AutoRouter(&controllers.RoleController{})
	beego.AutoRouter(&controllers.AdminController{})

	// 资料修改
	beego.AutoRouter(&controllers.UserController{})

}
