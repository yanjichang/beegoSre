/**********************************************
** @Des: login
** @Author: haodaquan
** @Date:   2017-09-07 16:30:10
** @Last Modified by:   haodaquan
** @Last Modified time: 2017-09-17 11:55:21
***********************************************/
package controllers

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Ohimma/beegoSre/libs"
	"github.com/Ohimma/beegoSre/models"
	"github.com/Ohimma/beegoSre/utils"
	"github.com/astaxie/beego"
	cache "github.com/patrickmn/go-cache"
)

type LoginController struct {
	BaseController
}

//登录 TODO:XSRF过滤
func (self *LoginController) LoginIn() {
	fmt.Println("self.userId = ", self.userId)
	if self.userId > 0 {
		self.redirect(beego.URLFor("HomeController.Index"))
	}
	fmt.Println("self.Controller = ", self.Controller)
	beego.ReadFromRequest(&self.Controller)
	if self.isPost() {

		username := strings.TrimSpace(self.GetString("username"))
		password := strings.TrimSpace(self.GetString("password"))

		if username != "" && password != "" {
			user, err := models.AdminGetByName(username)
			fmt.Println("login user = ", user)
			flash := beego.NewFlash()
			errorMsg := ""
			if err != nil || user.Password != libs.Md5([]byte(password+user.Salt)) {
				errorMsg = "帐号或密码错误"
			} else if user.Status == 0 {
				errorMsg = "该帐号已禁用"
			} else {
				user.LastIp = self.getClientIp()
				user.LastLogin = time.Now().Unix()
				user.Update()
				utils.Che.Set("uid"+strconv.Itoa(user.Id), user, cache.DefaultExpiration)
				authkey := libs.Md5([]byte(self.getClientIp() + "|" + user.Password + user.Salt))
				self.Ctx.SetCookie("auth", strconv.Itoa(user.Id)+"|"+authkey, 7*86400)

				self.redirect(beego.URLFor("HomeController.Index"))
			}
			flash.Error(errorMsg)
			flash.Store(&self.Controller)
			self.redirect(beego.URLFor("LoginController.LoginIn"))
		}
	}
	fmt.Println("self.tpname = ", self.TplName)
	self.TplName = "public/login.html"

}

//登出
func (self *LoginController) LoginOut() {
	self.Ctx.SetCookie("auth", "")
	self.redirect(beego.URLFor("LoginController.LoginIn"))
}

func (self *LoginController) NoAuth() {
	self.Ctx.WriteString("没有权限")
}
