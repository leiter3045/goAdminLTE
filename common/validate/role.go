package validate

import (
	"quickstart/common/lib/traits"
)

type Role struct {
	traits.ErrorReport
	Name    	string
}

/**
 * 验证数据
 * @return bool
 */
func (u *Role) Valid() (bool) {
	if name := u.Name; len(name) == 0 {
		u.SetError("name不得为空")
		return false
	}
	return true
}