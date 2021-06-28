package model

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"quickstart/common/validate"
	"quickstart/common/lib/traits"
	"quickstart/common/function"
)

type User struct {
	traits.ErrorReport
	Id       		int			`orm:"column(id);auto;size(5)" description:"主键"`
	Username    	string		`orm:"column(username);size(32);" description:"员工名称"`
	Userpass   		string		`orm:"column(userpass);size(32);type(char);" description:"员工密码"`
	LastLoginIp   	string    	`orm:"column(last_login_ip);size(15);type(char);" description:"最后一次登录IP"`
	LastLoginTime  	string 		`orm:"column(last_login_time);size(10);type(int);default(0);" description:"最后一次登录时间"`
	LoginTimes 		int			`orm:"column(login_times);size(10);default(0);type(int);" description:"登录次数"`
	State 			int			`orm:"column(state);size(3);default(1);type(tinyint);" description:"1启用0禁用"`
	AddTime 		int64		`orm:"column(add_time);size(10);type(int);" description:"添加时间"`
	Roles    		*Role     	`orm:"column(role_id);rel(one);default(0);"`
}

func (t *User) TableName() string {
	return "tp_user"
}

// 需要在init中注册定义的model
func init() {
	orm.RegisterModel(new(User))
}

/**
 * 通过菜单ID获取详情
 * @param int id
 * @return array
 */
func (c User) GetInfoById(id int, relation bool) (v *User, err error) {
	o := orm.NewOrm()
	v = &User{Id: id}
	if err = o.Read(v); err == nil {
		if relation {
			o.LoadRelated(v, "Roles")
		}
		return v, nil
	}
	return nil, err
}

func (c User) GetInfoUsername(username string, relation bool) (v User, err error) {
	user := User{Username: username}
	o := orm.NewOrm()
	if err = o.Read(&user, "Username"); err == nil {
		if relation {
			o.LoadRelated(&user, "Roles")
		}
		return user, nil
	}
	return user, errors.New("找不到数据")
}

/**
 * 获取所有员工
 * @return array
 */
func (c User) GetList(state int, limit int, offset int, relation bool) (m []*User, count int, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(c)
	if state != 9 {
		qs = qs.Filter("State", state)
	}
	qs = qs.Exclude("Username", "admin")
	var qscount int64
	qscount, err = qs.Count()
	if err != nil {
		return m, count, err
	}
	qs = qs.Limit(limit, offset)
	qs.All(&m)
	if relation {
		for _, v := range m {
			o.LoadRelated(v, "Roles")
		}
	}
	return m, int(qscount), err
}

/**
 * 编辑菜单
 */
func (c *User) EditRole(roleId int) (bool) {
	if c.Id == 0 || roleId == 0 {
		c.SetError("参数错误")
		return false
	}
	o := orm.NewOrm()
	user := User{Id: c.Id}
	role := Role{Id:roleId}
	o.Read(&role)
	user.Roles = &role
	if _, err := o.Update(&user, "Roles"); err != nil {
		c.SetError("员工角色添加失败")
		return false
	}
	return true
}

/**
 * 编辑菜单
 */
func (c *User) Edit() (bool) {
	valid := validate.User{
		Username: c.Username,
		Userpass: c.Userpass,
	}
	if state := valid.Valid(); state == false {
		c.SetError(valid.GetError())
		return false
	}
	c.Userpass = function.PasswordHash(c.Userpass)
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
func (c *User) Add() (id int, status bool) {
	valid := validate.User{
		Username: c.Username,
		Userpass: c.Userpass,
	}
	if state := valid.Valid(); state == false {
		c.SetError(valid.GetError())
		return id, false
	}
	var err error
	c.Userpass = function.PasswordHash(c.Userpass)
	o := orm.NewOrm()
	var userId int64
	userId, err = o.Insert(c)
	if err != nil {
		c.SetError("插入数据失败，请联系客服")
		return id, false
	}
	return int(userId), true
}

/**
 * 删除菜单
 */
func (c *User) DelInfoById(id int) (bool) {
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
