package validate

import (
	"quickstart/common/lib/traits"
)

type Access struct {
	traits.ErrorReport
	RoleId    	int
	MenuId   	int
	Url   		string
}

func (u *Access) Valid() (bool) {
	if roleId := u.RoleId; roleId == 0 {
		u.SetError("角色ID不得为空")
		return false
	}
	if menuId := u.MenuId; menuId == 0 {
		u.SetError("菜单ID不得为空")
		return false
	}
	if url := u.Url; len(url) == 0 {
		u.SetError("菜单理由地址错误")
		return false
	}
	return true
}