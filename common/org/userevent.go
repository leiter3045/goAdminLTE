package org

import (
	"quickstart/common/function"
	"github.com/astaxie/beego/context"
)

type UserAgent struct {
	Ctx  *context.Context
}

/**
 * 电脑唯一识别码
 * @return Ambigous <mixed, void, NULL, unknown, multitype:, string>
 */
func (c UserAgent) ComputerCookie(user_agent string, request_time string, code string) string {
	cookie_key := code + "_computer";
	cookie := c.Ctx.GetCookie(cookie_key);
	if len(cookie) == 0 {
		cookie = function.Md5(user_agent + request_time + function.GetUuid());
		c.Ctx.SetCookie(cookie_key, cookie, 10 * 365 * 24 * 3600);
	}
	return cookie;
}