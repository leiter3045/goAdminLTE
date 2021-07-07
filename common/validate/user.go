package validate

import (
	"quickstart/common/lib/traits"
	"quickstart/common/function"
)

type User struct {
	traits.ErrorReport
	Username    	string
	Userpass   		string
}

/**
 * 验证数据
 * @return bool
 */
func (u *User) Valid() (bool) {
	if username := u.Username; len(username) == 0 {
		u.SetError("name不得为空")
		return false
	}
	if match := function.ValidData(u.Username, "Username"); !match {
		u.SetError("用户名由6-16位字母开头的数字字母组成！")
		return false
	}
	if userpass := u.Userpass; len(userpass) == 0 {
		u.SetError("密码不得为空")
		return false
	}
	if match := function.ValidData(u.Userpass, "Password"); !match {
		u.SetError("密码由6-16位数字字母组成！")
		return false
	}
	return true
}