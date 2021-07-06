package validate

import (
	"quickstart/common/lib/traits"
)

type Config struct {
	traits.ErrorReport
	Name    	string
	Value   	string
}

func (u *Config) Valid() (bool) {
	if name := u.Name; len(name) == 0 {
		u.SetError("配置名称不得为空")
		return false
	}
	if value := u.Value; len(value) == 0 {
		u.SetError("配置内容不得为空")
		return false
	}
	return true
}