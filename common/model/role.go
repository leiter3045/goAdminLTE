package model

import (
	"github.com/astaxie/beego/orm"
	"quickstart/common/validate"
	"quickstart/common/lib/traits"
)

type Role struct {
	traits.ErrorReport
	Id       		int			`orm:"column(id);auto;size(5)" description:"主键"`
	Name    		string		`orm:"column(name);size(20);" description:"员工名称"`
	State 			int			`orm:"column(state);size(3);default(1);type(tinyint);" description:"1启用0禁用"`
}

func (t *Role) TableName() string {
	return "tp_role"
}

// 需要在init中注册定义的model
func init() {
	orm.RegisterModel(new(Role))
}

/**
 * 通过菜单ID获取详情
 * @param int id
 * @return array
 */
func (c Role) GetInfoById(id int) (v *Role, err error) {
	o := orm.NewOrm()
	v = &Role{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

/**
 * 获取所有员工
 * @return array
 */
func (c Role) GetList(state int, limit int, offset int) (m []*Role, count int, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(c)

	if state != 9 {
		qs = qs.Filter("State", state)
	}
	var qscount int64
	qscount, err = qs.Count()
	if err != nil {
		return m, count, err
	}
	qs = qs.Limit(limit, offset)
	qs.All(&m)
	return m, int(qscount), err
}

/**
 * 获取所有员工
 * @return array
 */
func (c Role) GetRoleByAll(state int) (m []*Role, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(c)
	if state != 9 {
		qs = qs.Filter("State", state)
	}
	qs.All(&m)
	return m, err
}

/**
 * 编辑菜单
 */
func (c *Role) Edit() (bool) {
	valid := validate.Role{
		Name: c.Name,
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
func (c *Role) Add() (id int, status bool) {
	valid := validate.Role{
		Name: c.Name,
	}
	if state := valid.Valid(); state == false {
		c.SetError(valid.GetError())
		return id, false
	}
	var err error
	o := orm.NewOrm()
	var roleId int64
	roleId, err = o.Insert(c)
	if err != nil {
		c.SetError("插入数据失败，请联系客服")
		return id, false
	}
	return int(roleId), true
}

/**
 * 删除菜单
 */
func (c *Role) DelInfoById(id int) (bool) {
	if id == 0 {
		c.SetError("参数错误")
		return false
	}
	o := orm.NewOrm()
	user := User{Id:id}
	if err := o.Read(&user); err != nil {
		c.SetError("未找到数据")
		return false
	}
	if _, err := o.Delete(&user); err != nil {
		c.SetError("删除数据失败")
		return false
	}
	return true
}
