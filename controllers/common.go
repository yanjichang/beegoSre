/**********************************************
** @Des: base controller
** @Author: haodaquan
** @Date:   2017-09-07 16:54:40
** @Last Modified by:   haodaquan
** @Last Modified time: 2017-09-18 10:28:01
***********************************************/
package controllers

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/Ohimma/beegoSre/libs"
	"github.com/Ohimma/beegoSre/utils"
	"golang.org/x/text/encoding/simplifiedchinese"

	"github.com/Ohimma/beegoSre/models"
	"github.com/astaxie/beego"
	cache "github.com/patrickmn/go-cache"
)

type Charset string

const (
	MSG_OK  = 0
	MSG_ERR = -1

	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
)

type BaseController struct {
	beego.Controller
	controllerName string
	actionName     string
	user           *models.Admin
	userId         int
	userName       string
	loginName      string
	pageSize       int
	allowUrl       string
}

//前期准备
func (self *BaseController) Prepare() {
	self.pageSize = 20
	controllerName, actionName := self.GetControllerAndAction()
	self.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	self.actionName = strings.ToLower(actionName)
	self.Data["version"] = beego.AppConfig.String("version")
	self.Data["siteName"] = beego.AppConfig.String("site.name")
	self.Data["curRoute"] = self.controllerName + "." + self.actionName
	self.Data["curController"] = self.controllerName
	self.Data["curAction"] = self.actionName
	// noAuth := "ajaxsave/ajaxdel/table/loginin/loginout/getnodes/start"
	// isNoAuth := strings.Contains(noAuth, self.actionName)
	fmt.Println("c conmmon preprare self.controllerName ==== ", self.controllerName)
	if (strings.Compare(self.controllerName, "apidoc")) != 0 {
		self.auth()
	}

	self.Data["loginUserId"] = self.userId
	self.Data["loginUserName"] = self.userName
}

//登录权限验证
func (self *BaseController) auth() {
	arr := strings.Split(self.Ctx.GetCookie("auth"), "|")
	self.userId = 0
	if len(arr) == 2 {
		idstr, password := arr[0], arr[1]
		userId, _ := strconv.Atoi(idstr)
		if userId > 0 {
			var err error

			cheUser, found := utils.Che.Get("uid" + strconv.Itoa(userId))
			user := &models.Admin{}
			if found && cheUser != nil { //从缓存取用户
				user = cheUser.(*models.Admin)
			} else {
				user, err = models.AdminGetById(userId)
				utils.Che.Set("uid"+strconv.Itoa(userId), user, cache.DefaultExpiration)
			}
			if err == nil && password == libs.Md5([]byte(self.getClientIp()+"|"+user.Password+user.Salt)) {
				self.userId = user.Id

				self.loginName = user.LoginName
				self.userName = user.RealName
				self.user = user
				self.AdminAuth()
			}

			isHasAuth := strings.Contains(self.allowUrl, self.controllerName+"/"+self.actionName)
			//不需要权限检查
			noAuth := "ajaxsave/ajaxdel/table/loginin/loginout/getnodes/start/show/ajaxapisave/index/group/public/env/code/apidetail"
			isNoAuth := strings.Contains(noAuth, self.actionName)
			if isHasAuth == false && isNoAuth == false {
				self.Ctx.WriteString("没有权限")
				self.ajaxMsg("没有权限", MSG_ERR)
				return
			}
		}
	}

	if self.userId == 0 && (self.controllerName != "login" && self.actionName != "loginin") {
		self.redirect(beego.URLFor("LoginController.LoginIn"))
	}
}

func (self *BaseController) AdminAuth() {
	cheMen, found := utils.Che.Get("menu" + strconv.Itoa(self.user.Id))
	fmt.Printf("c comon adminauth chemen==%v  found=%v\n", cheMen, found)

	if found && cheMen != nil { //从缓存取菜单
		fmt.Println("缓存 内调用显示菜单")
		menu := cheMen.(*CheMenu)
		self.Data["SideMenu1"] = menu.List1 //一级菜单
		self.Data["SideMenu2"] = menu.List2 //二级菜单
		self.allowUrl = menu.AllowUrl
	} else {
		// 左侧导航栏
		fmt.Println("数据库 内调用显示菜单")
		filters := make([]interface{}, 0)
		filters = append(filters, "status", 1)
		if self.userId != 1 {
			//普通管理员
			adminAuthIds, _ := models.RoleAuthGetByIds(self.user.RoleIds)
			adminAuthIdArr := strings.Split(adminAuthIds, ",")
			filters = append(filters, "id__in", adminAuthIdArr)
		}
		result, _ := models.AuthGetList(1, 1000, filters...)
		list := make([]map[string]interface{}, len(result))
		list2 := make([]map[string]interface{}, len(result))
		allow_url := ""
		i, j := 0, 0
		for _, v := range result {
			if v.AuthUrl != " " || v.AuthUrl != "/" {
				allow_url += v.AuthUrl
			}
			row := make(map[string]interface{})
			if v.Pid == 1 && v.IsShow == 1 {
				row["Id"] = int(v.Id)
				row["Sort"] = v.Sort
				row["AuthName"] = v.AuthName
				row["AuthUrl"] = v.AuthUrl
				row["Icon"] = v.Icon
				row["Pid"] = int(v.Pid)
				list[i] = row
				i++
			}
			if v.Pid != 1 && v.IsShow == 1 {
				row["Id"] = int(v.Id)
				row["Sort"] = v.Sort
				row["AuthName"] = v.AuthName
				row["AuthUrl"] = v.AuthUrl
				row["Icon"] = v.Icon
				row["Pid"] = int(v.Pid)
				list2[j] = row
				j++
			}
		}
		self.Data["SideMenu1"] = list[:i]  //一级菜单
		self.Data["SideMenu2"] = list2[:j] //二级菜单
		fmt.Println("c comon admin menu.List1 == ", self.Data["SideMenu1"])

		self.allowUrl = allow_url + "/home/index"
		cheM := &CheMenu{}
		cheM.AllowUrl = self.allowUrl
		cheM.List1 = self.Data["SideMenu1"].([]map[string]interface{})
		cheM.List2 = self.Data["SideMenu2"].([]map[string]interface{})
		utils.Che.Set("menu"+strconv.Itoa(self.user.Id), cheM, cache.DefaultExpiration)
	}

}

type CheMenu struct {
	List1    []map[string]interface{}
	List2    []map[string]interface{}
	AllowUrl string
}

// 是否POST提交
func (self *BaseController) isPost() bool {
	return self.Ctx.Request.Method == "POST"
}

//获取用户IP地址
func (self *BaseController) getClientIp() string {
	s := self.Ctx.Request.RemoteAddr
	l := strings.LastIndex(s, ":")
	return s[0:l]
}

// 重定向
func (self *BaseController) redirect(url string) {
	self.Redirect(url, 302)
	self.StopRun()
}

//加载模板
func (self *BaseController) display(tpl ...string) {
	var tplname string
	if len(tpl) > 0 {
		tplname = strings.Join([]string{tpl[0], "html"}, ".")
	} else {
		tplname = self.controllerName + "_" + self.actionName + ".html"
	}
	self.Layout = "public/layout2.html"
	self.TplName = tplname
	fmt.Println("tplname == ", tplname, "layout == ", self.Layout)
}

//ajax 返回消息和状态
func (self *BaseController) ajaxMsg(msg interface{}, msgno int) {
	out := make(map[string]interface{})
	out["status"] = msgno
	out["message"] = msg
	// out["msg"] = msg
	self.Data["json"] = out
	self.ServeJSON()
	self.StopRun()
}

//ajax 返回 code + 状态 + 数量 + 列表数据
func (self *BaseController) ajaxList(msg interface{}, msgno int, count int64, data interface{}) {
	out := make(map[string]interface{})
	out["code"] = msgno
	// out["status"] = msgno
	out["msg"] = msg
	out["count"] = count
	out["data"] = data
	self.Data["json"] = out
	self.ServeJSON()
	self.StopRun()
	fmt.Println("ajaxlist json == ", self.Data["json"], "self == ", self)
}

// 执行shell 命令

//对字符进行转码 cmd 窗口会乱码
func ConvertByte2String(byte []byte, charset Charset) string {
	var str string
	switch charset {
	case GB18030:
		var decodeBytes, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		str = string(decodeBytes)
	case UTF8:
		fallthrough
	default:
		str = string(byte)
	}
	return str
}

func (self *BaseController) exec_shell(s string) ([]string, error) {
	//函数返回一个*Cmd，用于使用给出的参数执行name指定的程序
	// cmd := exec.Command("/bin/bash", "-c", s)
	// cmd := exec.Command(s)
	cmd := exec.Command("ping", "www.baidu.com")

	fmt.Println("CmdAndChangeDir", cmd.Dir, cmd.Args)
	// fmt.Printf("执行命令: %s\n", strings.Join(cmd.Args[1:], " "))

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, "error===>", err.Error())
	}
	err = cmd.Start()
	if err != nil {
		fmt.Println("c common cmd start err == ", err)
	}

	var contentArray = make([]string, 0, 5)
	in := bufio.NewScanner(stdout)
	for in.Scan() {
		line := ConvertByte2String(in.Bytes(), "GB18030")
		fmt.Println(line)
		contentArray = append(contentArray, line)
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Println("c common cmd wait err == ", err)
	}

	return contentArray, err
}
