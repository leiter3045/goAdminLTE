package validate

import (
"quickstart/common/lib/traits"
)

type Userlog struct {
	traits.ErrorReport
	EventId    	int
	Desc   		string
	Ip			string
}

func (u *Userlog) Valid() (bool) {
	if eventId := u.EventId; eventId == 0 {
		u.SetError("事件ID不得为空")
		return false
	}
	if desc := u.Desc; len(desc) == 0 {
		u.SetError("事件描述不得为空")
		return false
	}
	if ip := u.Ip; len(ip) == 0 {
		u.SetError("IP地址不得为空")
		return false
	}
	return true
}