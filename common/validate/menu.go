package validate

import (
	"quickstart/common/lib/traits"
)

type Menu struct {
	traits.ErrorReport
	Name  	string		
	Url   	string 	
}

/**
 * 验证数据
 * @return bool
 */
func (u *Menu) Valid() (bool) {
    if name := u.Name; len(name) == 0 {
		u.SetError("name不得为空")
		return false
	}
	if url := u.Url; len(url) == 0 {
		u.SetError("url不得为空")
		return false
	}
	return true
}