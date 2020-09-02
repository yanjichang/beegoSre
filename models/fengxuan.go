/**********************************************
** @Des: This file ...
** @Author: haodaquan
** @Date:   2017-09-14 15:24:51
** @Last Modified by:   haodaquan
** @Last Modified time: 2017-09-17 11:48:52
***********************************************/
package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Fengxuan struct {
	Id            int
	Fengxuan_name string
	LiuliangNW    string `orm:"column(liuliangNW)"`
	Kuajing       string
	KuajingNW     string `orm:"column(kuajingNW)"`
	Haiwai        string
	Shortname     string

	Type             string
	Secret           string
	Starospassword   string
	DelayCof         float32 `orm:"column(delayCof)"`
	SsrProtocol      string  `orm:"column(ssrProtocol)"`
	SsrProtocolParam string  `orm:"column(ssrProtocolParam)"`
	SsrObfs          string  `orm:"column(ssrObfs)"`
	SsrObfsParam     string  `orm:"column(ssrObfsParam)"`

	Ports      int
	Local_port int
	Delay      int
	Servstate  int
	Groupid    int
}

func (a *Fengxuan) TableName() string {
	return TableName("fengxuan_nas")
}

func FengxuanAdd(Fengxuan *Fengxuan) (int64, error) {
	id, err := orm.NewOrm().Insert(Fengxuan)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func FengxuanGetList(page, pageSize int, filters ...interface{}) ([]*Fengxuan, int64) {
	offset := (page - 1) * pageSize
	list := make([]*Fengxuan, 0)

	query := orm.NewOrm().QueryTable(TableName("fengxuan_nas"))

	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offset).All(&list)
	return list, total
}

func FengxuanGetByName(fengxuan_name string) (*Fengxuan, error) {
	a := new(Fengxuan)
	err := orm.NewOrm().QueryTable(TableName("fengxuan_nas")).Filter("fengxuan_name", fengxuan_name).One(a)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func FengxuanGetById(id int) (*Fengxuan, error) {
	r := new(Fengxuan)
	err := orm.NewOrm().QueryTable(TableName("fengxuan_nas")).Filter("id", id).One(r)
	if err != nil {
		return nil, err
	}
	fmt.Println("m fengxuan getbyids r == ", r)
	return r, nil
}

// fields 指的是只更新 fields 字段
func (r *Fengxuan) Update(fields ...string) error {

	_, err := orm.NewOrm().Update(r, fields...)
	if err != nil {
		return err
	}
	return nil
}
