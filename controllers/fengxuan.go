/**********************************************
** @Des: This file ...
** @Author: haodaquan
** @Date:   2017-09-08 10:21:13
** @Last Modified by:   haodaquan
** @Last Modified time: 2017-09-09 18:04:41
***********************************************/
package controllers

import (
	"fmt"
	"github.com/Ohimma/beegoSre/models"

	"strings"
)

type FengxuanController struct {
	BaseController
}

func (self *FengxuanController) List() {
	// self.Data["pageTitle"] = "丰炫列表"
	self.display()
}

func (self *FengxuanController) Add() {
	self.Data["pageTitle"] = "新增链路"

	// 角色
	filters := make([]interface{}, 0)
	// filters = append(filters, "status", 1)
	result, _ := models.FengxuanGetList(1, 1000, filters...)
	list := make([]map[string]interface{}, len(result))

	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["fengxuan_name"] = v.Fengxuan_name
		list[k] = row
	}

	self.Data["role"] = list

	self.display()
}

func (self *FengxuanController) Edit() {
	self.Data["pageTitle"] = "编辑数据"

	id, _ := self.GetInt("id")
	Fengxuan, _ := models.FengxuanGetById(id)

	row := make(map[string]interface{})
	row["id"] = Fengxuan.Id
	row["shortname"] = Fengxuan.Shortname
	row["fengxuan_name"] = Fengxuan.Fengxuan_name
	row["liuliangNW"] = Fengxuan.LiuliangNW
	row["kuajing"] = Fengxuan.Kuajing
	row["kuajingNW"] = Fengxuan.KuajingNW
	row["haiwai"] = Fengxuan.Haiwai
	row["groupid"] = Fengxuan.Groupid
	row["delay"] = Fengxuan.Delay
	row["servstate"] = Fengxuan.Servstate

	row["type"] = Fengxuan.Type
	row["ports"] = Fengxuan.Ports
	row["secret"] = Fengxuan.Secret
	row["starospassword"] = Fengxuan.Starospassword
	row["local_port"] = Fengxuan.Local_port
	row["ssrObfs"] = Fengxuan.SsrObfs
	row["ssrProtocol"] = Fengxuan.SsrProtocol

	self.Data["fengxuan"] = row

	fmt.Println("c fengxuan edit data == ", self.Data)

	// role_ids := strings.Split(Fengxuan.RoleIds, ",")

	// filters := make([]interface{}, 0)
	// filters = append(filters, "status", 1)
	// result, _ := models.RoleGetList(1, 1000, filters...)
	// list := make([]map[string]interface{}, len(result))
	// for k, v := range result {
	// 	row := make(map[string]interface{})
	// 	row["checked"] = 0
	// 	for i := 0; i < len(role_ids); i++ {
	// 		role_id, _ := strconv.Atoi(role_ids[i])
	// 		if role_id == v.Id {
	// 			row["checked"] = 1
	// 		}
	// 		fmt.Println(role_ids[i])
	// 	}
	// 	row["id"] = v.Id
	// 	row["role_name"] = v.RoleName
	// 	list[k] = row
	// }
	// self.Data["role"] = list
	self.display()
}

func (self *FengxuanController) AjaxSave() {
	Fengxuan_id, _ := self.GetInt("id")
	fmt.Println("Fengxuan_id == ", Fengxuan_id)

	Fengxuan := new(models.Fengxuan)
	Fengxuan.Shortname = strings.TrimSpace(self.GetString("shortname"))
	Fengxuan.Fengxuan_name = strings.TrimSpace(self.GetString("fengxuan_name"))
	Fengxuan.LiuliangNW = strings.TrimSpace(self.GetString("liuliangNW"))
	Fengxuan.Kuajing = strings.TrimSpace(self.GetString("kuajing"))
	Fengxuan.KuajingNW = strings.TrimSpace(self.GetString("kuajingNW"))
	Fengxuan.Haiwai = strings.TrimSpace(self.GetString("haiwai"))

	Fengxuan.Groupid, _ = self.GetInt("groupid")
	Fengxuan.Delay, _ = self.GetInt("delay")
	Fengxuan.Servstate, _ = self.GetInt("servstate")
	Fengxuan.Ports, _ = self.GetInt("ports")
	Fengxuan.Local_port, _ = self.GetInt("local_port")

	Fengxuan.Type = strings.TrimSpace(self.GetString("type"))
	Fengxuan.Secret = strings.TrimSpace(self.GetString("secret"))
	Fengxuan.Starospassword = strings.TrimSpace(self.GetString("starospassword"))
	Fengxuan.SsrProtocol = strings.TrimSpace(self.GetString("ssrProtocol"))

	if Fengxuan_id == 0 {
		// 检查登录名是否已经存在
		fmt.Println("Fengxuan_name == ", Fengxuan.Fengxuan_name)
		fmt.Println("Fengxuan_name == ", Fengxuan)
		_, err := models.FengxuanGetByName(Fengxuan.Fengxuan_name)

		if err == nil {
			self.ajaxMsg("登录名已经存在", MSG_ERR)
		}

		_, err = models.FengxuanAdd(Fengxuan)
		if err != nil {
			self.ajaxMsg(err.Error(), MSG_ERR)
		}
		self.ajaxMsg("", MSG_OK)
	}

	//修改
	// Fengxuan, _ = models.FengxuanGetById(Fengxuan_id)
	Fengxuan.Id = Fengxuan_id

	if err := Fengxuan.Update(); err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	self.ajaxMsg("", MSG_OK)
}

// 更新多条或者一条数据
func (self *FengxuanController) AjaxDel() {

	Fengxuan_id, _ := self.GetInt("id")

	Fengxuan_servstate := -10
	Fengxuan, _ := models.FengxuanGetById(Fengxuan_id)
	Fengxuan.Servstate = Fengxuan_servstate
	Fengxuan.Id = Fengxuan_id

	fmt.Println("c fengxuan ajaxdel Fengxuan_id == ", Fengxuan_id)
	fmt.Println("c fengxuan ajaxdel Fengxuan == ", Fengxuan)

	err := Fengxuan.Update()

	if err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	self.ajaxMsg("操作成功", MSG_OK)
}

func (self *FengxuanController) Table() {
	//列表
	page, err := self.GetInt("page")
	if err != nil {
		page = 1
	}
	limit, err := self.GetInt("limit")
	if err != nil {
		limit = 30
	}

	fmt.Println("c fengxuan table limit==", limit, " page==", page, " self.pageSize===", self.pageSize)
	Fengxuan_name := strings.TrimSpace(self.GetString("fengxuan_name"))
	self.pageSize = limit

	//查询条件
	filters := make([]interface{}, 0)
	filters = append(filters, "servstate__gt", -9)

	if Fengxuan_name != "" {
		filters = append(filters, "fengxuan_name__icontains", Fengxuan_name)
	}
	result, count := models.FengxuanGetList(page, self.pageSize, filters...)
	list := make([]map[string]interface{}, len(result))

	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["shortname"] = v.Shortname
		row["groupid"] = v.Groupid

		row["fengxuan_name"] = v.Fengxuan_name
		row["liuliangNW"] = v.LiuliangNW
		row["kuajing"] = v.Kuajing
		row["kuajingNW"] = v.KuajingNW
		row["haiwai"] = v.Haiwai

		row["servstate"] = v.Servstate
		row["type"] = v.Type
		row["ports"] = v.Ports
		row["local_port"] = v.Local_port
		row["secret"] = v.Secret
		row["starospassword"] = v.Starospassword

		row["delay"] = v.Delay
		list[k] = row
	}
	self.ajaxList("成功", MSG_OK, count, list)
}

func (self *FengxuanController) Deploy_zb() {
	// self.Data["pageTitle"] = "部署zabbix输出..."

	fengxuan_name := self.GetString("fengxuan_name")
	deploy := self.GetString("deploy")

	//fengxuan_name = strings.Replace(fengxuan_name, ",", "','", -1)
	// fmt.Println("c fengxuan deploy_zb fengxuan_name == ", fengxuan_name, "deploy == ", deploy)

        cmd := "/bin/sh /etc/ansible_bak/fengxuan_deploy_zb.sh " + fengxuan_name + "  " + deploy
        fmt.Println("c fengxuan deploy_zb cmd == ", cmd)
	re, _ := self.exec_shell(cmd)
	// fmt.Printf("re t == %T   v == %v\n", re, re)

	list := make([]map[string]interface{}, len(re))

	for k, v := range re {
		row := make(map[string]interface{})
		row["line"] = v
		list[k] = row
	}

	self.Data["fengxuan_name"] = fengxuan_name
	self.Data["deploy"] = deploy
	self.Data["output"] = list
	self.display("deploy_zb")
}

