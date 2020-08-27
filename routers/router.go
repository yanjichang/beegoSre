package routers

import (
	"github.com/astaxie/beego"
	"github.com/george518/PPGo_ApiAdmin/controllers"
)

func init() {
	// 默认登录
	beego.Router("/doc", &controllers.ApiDocController{}, "*:Index")
	beego.AutoRouter(&controllers.ApiDocController{})

	beego.Router("/", &controllers.LoginController{}, "*:LoginIn")
	beego.Router("/login", &controllers.LoginController{}, "*:LoginIn")
	beego.Router("/login_out", &controllers.LoginController{}, "*:LoginOut")
	beego.Router("/no_auth", &controllers.LoginController{}, "*:NoAuth")

	beego.Router("/home", &controllers.HomeController{}, "*:Index")
	beego.Router("/home/start", &controllers.HomeController{}, "*:Start")

	// zabbix 管理

	// API 管理
	beego.AutoRouter(&controllers.ApiController{})
	beego.AutoRouter(&controllers.ApiSourceController{})

	// beego.AutoRouter(&controllers.ApiMonitorController{})

	// 基础设置
	beego.AutoRouter(&controllers.GroupController{})
	beego.AutoRouter(&controllers.EnvController{})
	beego.AutoRouter(&controllers.ApiPublicController{})
	beego.AutoRouter(&controllers.CodeController{})
	beego.AutoRouter(&controllers.TemplateController{})

	// 权限设置
	beego.AutoRouter(&controllers.AuthController{})
	beego.AutoRouter(&controllers.RoleController{})
	beego.AutoRouter(&controllers.AdminController{})
	// 资料修改
	beego.AutoRouter(&controllers.UserController{})

}
