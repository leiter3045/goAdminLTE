package model

import (
	"github.com/astaxie/beego/orm"
	"quickstart/common/validate"
	"quickstart/common/lib/traits"
)

type Access struct {
	traits.ErrorReport
	Id  			int   		`orm:"auto;column(id)"`
	RoleId    		int			`orm:"column(role_id);size(10);type(int);" description:"角色ID"`
	MenuId   		int			`orm:"column(menu_id);size(10);type(int);" description:"菜单ID"`
	Url   			string    	`orm:"column(url);size(65);type(varchar);" description:"菜单URL"`
}

func (t *Access) TableName() string {
	return "tp_access"
}

// 需要在init中注册定义的model
func init() {
	orm.RegisterModel(new(Access))
}

/**
 * 通过ID获取详情
 * @param int roleId
 * @return Access
 */
func (c Access) GetInfoById(roleId int) (v *Access, err error) {
	o := orm.NewOrm()
	v = &Access{RoleId: roleId}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

/**
 * 获取详情
 * @param int roleId
 * @param int menuId
 * @return Access
 */
func (c Access) GetMenuIdsByRole(roleId int, menuId int) (v *Access, err error) {
	o := orm.NewOrm()
	v = &Access{RoleId: roleId, MenuId: menuId}
	if err = o.Read(v, "RoleId", "MenuId"); err == nil {
		return v, nil
	}
	return nil, err
}

/**
 * 获取所有菜单
 * @return array
 */
func (c Access) GetMenuByAll(state int) (m []*Access, err error) {
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
func (c *Access) Edit() (bool) {
	valid := validate.Access{
		RoleId: c.RoleId,
		MenuId: c.MenuId,
		Url: 	c.Url,
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
func (c *Access) Add() (bool) {
	valid := validate.Access{
		RoleId: c.RoleId,
		MenuId: c.MenuId,
		Url: 	c.Url,
	}
	if state := valid.Valid(); state == false {
		c.SetError(valid.GetError())
		return false
	}
	o := orm.NewOrm()
	var err error
	_, err = o.Insert(c)
	if err != nil {
		c.SetError("插入数据失败，请联系客服")
		return false
	}
	return true
}

/**
 * 通过RoleId获取数据
 */
func (c Access) GetRoleByAll(roleId int) (m []*Access, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(c)
	qs = qs.Filter("RoleId", roleId)
	qs.All(&m);
	return m, err
}

/**
 * 删除菜单
 */
func (c *Access) DelRoleAll(roleId int) (bool) {
	if roleId == 0 {
		c.SetError("参数错误")
		return false
	}
	o := orm.NewOrm()
	qs := o.QueryTable(c)
	qs = qs.Filter("RoleId", roleId)
	if _, err := qs.Delete(); err != nil {
		c.SetError("删除数据失败")
		return false
	}
	return true
}
