package model

import (
	"github.com/astaxie/beego/orm"
	"quickstart/common/validate"
	"quickstart/common/lib/traits"
)

type Menu struct {
	traits.ErrorReport
	Id       	int			`orm:"column(id);auto;size(5)" description:"主键"`
	Pid    		int			`orm:"column(pid);size(5);" description:"父亲ID"`
	Name   		string		`orm:"column(name);size(32);type(varchar);" description:"菜单名称"`
	Url   		string    	`orm:"column(url);size(65);type(varchar);" description:"菜单URL"`
	Icon  		string 		`orm:"column(icon);size(32);type(varchar);default('');" description:"菜单图标"`
	Sort 		int			`orm:"column(sort);size(5);default(50);type(smallint);" description:"菜单排序"`
	State 		int			`orm:"column(state);size(3);default(1);type(tinyint);" description:"1启用0禁用"`
}

func (t *Menu) TableName() string {
	return "tp_menu"
}

// 需要在init中注册定义的model
func init() {
	orm.RegisterModel(new(Menu))
}

/**
 * 通过菜单ID获取菜单详情
 * @param int id
 * @return array
 */
func (c Menu) GetInfoById(id int) (v *Menu, err error) {
	o := orm.NewOrm()
	v = &Menu{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

/**
 * 通过url获取菜单
 * @return array
 */
func (c Menu) GetInfoByUrl(url string) (v *Menu, err error) {
	o := orm.NewOrm()
	v = &Menu{Url: url}
	if err = o.Read(v, "Url"); err == nil {
		return v, nil
	}
	return nil, err
}

/**
 * 获取所有菜单
 * @return array
 */
func (c Menu) GetMenuByIds(state int, ids interface{}) (m []*Menu, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(c)
	if state != 9 {
		qs = qs.Filter("State", state)
	}
	qs = qs.Filter("Id__in", ids)
	qs = qs.OrderBy("sort")
	qs.All(&m);
	return m, err
}

/**
 * 获取所有菜单
 * @return array
 */
func (c Menu) GetMenuByAll(state int) (m []*Menu, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(c)
	if state != 9 {
		qs = qs.Filter("State", state)
	}
	qs = qs.OrderBy("sort")
	qs.All(&m);
	return m, err
}

/**
 * 编辑菜单
 */
func (c *Menu) Edit() (bool) {
	valid := validate.Menu{
		Name: c.Name,
		Url: c.Url,
	}
	if state := valid.Valid(); state == false {
		c.SetError(valid.GetError())
		return false
	}
	o := orm.NewOrm()
	if _, err := o.Update(c); err != nil {
		c.SetError("修改数据失败")
		return false
	}
	return true
}

/**
 * 添加菜单
 */
func (c *Menu) Add() (id int, status bool) {
	valid := validate.Menu{
		Name: c.Name,
		Url: c.Url,
	}
	if state := valid.Valid(); state == false {
		c.SetError(valid.GetError())
		return id, false
	}
	var err error
	o := orm.NewOrm()
	var menuId int64
	menuId, err = o.Insert(c)
	if err != nil {
		c.SetError("插入数据失败，请联系客服")
		return id, false
	}
	return int(menuId), true
}

/**
 * 删除菜单
 */
func (c *Menu) DelInfoById(id int) (bool) {
	if id == 0 {
		c.SetError("参数错误")
		return false
	}
	o := orm.NewOrm()
	menu := Menu{Id:id}
	if err := o.Read(&menu); err != nil {
		c.SetError("未找到数据")
		return false
	}
	if _, err := o.Delete(&menu); err != nil {
		c.SetError("删除数据失败")
		return false
	}
	return true
}
