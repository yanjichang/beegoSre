/**********************************************
** @Des: This file ...
** @Author: haodaquan
** @Date:   2017-09-16 15:42:43
** @Last Modified by:   haodaquan
** @Last Modified time: 2017-09-24 11:48:17
***********************************************/
package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Group struct {
	Id           int
	GroupName    string
	Detail       string
	ApiPublicIds string
	CodeIds      string
	EnvIds       string
	Status       int
	CreateId     int
	UpdateId     int
	CreateTime   int64
	UpdateTime   int64
}

func (a *Group) TableName() string {
	return TableName("set_group")
}

func GroupAdd(a *Group) (int64, error) {
	return orm.NewOrm().Insert(a)
}

func GroupGetByName(groupName string) (*Group, error) {
	a := new(Group)
	err := orm.NewOrm().QueryTable(TableName("set_group")).Filter("group_name", groupName).One(a)
	if err != nil {
		return nil, err
	}
	return a, nil
}

// 返回 切片指针 list =  [0xc000332f80 0xc000333000 0xc000333100] total =  3
func GroupGetList(page, pageSize int, filters ...interface{}) ([]*Group, int64) {
	offset := (page - 1) * pageSize
	list := make([]*Group, 0)

	// 初始化数据库，并查询 qq_set_group 表，返回 QuerySeter 中用于描述字段和 sql 操作符，使用简单的 expr 查询方法
	query := orm.NewOrm().QueryTable(TableName("set_group"))

	fmt.Println("list = ", list)
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			// filters[k].(string) =  status ， filters[k+1] 1
			// fmt.Println("filters[k].(string) = ", filters[k].(string), "filters[k+1]", filters[k+1])
			query = query.Filter(filters[k].(string), filters[k+1])
			// fmt.Println("query = ", query)
		}
	}
	total, _ := query.Count()
	// 在 orm.NewOrm().QueryTable() 返回 expr， expr 前使用减号 - 表示 DESC 的排列
	query.OrderBy("-id").Limit(pageSize, offset).All(&list)

	fmt.Println("list = ", list, "total = ", total)
	return list, total
}

func GroupGetById(id int) (*Group, error) {
	r := new(Group)
	err := orm.NewOrm().QueryTable(TableName("set_group")).Filter("id", id).One(r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (a *Group) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(a, fields...); err != nil {
		return err
	}
	return nil
}
