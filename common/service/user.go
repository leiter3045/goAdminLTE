package service

import (
	"quickstart/common/model"
	"quickstart/common/lib/traits"
	"quickstart/common/function"
	"quickstart/common/constant"
)

type User struct {
	traits.ErrorReport
	Id          int
	Username   	string
	Userpass    string
}

/**
 * 用户登录
 * @return interface{}，bool
 */
func (c *User) DoLogin() (m interface{}, bool bool) {
	model := new (model.User)
	model.Username =  c.Username
	info, err := model.GetInfoUsername(c.Username, true)
	if err != nil {
		c.SetError("账号或者密码错误")
		return m, false
	}
	if info.State != constant.USER_STATE_ON {
		c.SetError("员工状态错误，登录失败")
		return m, false
	}
	passRemote := info.Userpass
	passLocal := function.PasswordHash(c.Userpass)
	if passRemote != passLocal {
		c.SetError("密码错误")
		return m, false
	}
	auth := make(map[string]interface{}, 0)
	auth["Id"] = info.Id
	auth["Username"] = info.Username
	auth["Roles"] = info.Roles
	return auth, true
}



