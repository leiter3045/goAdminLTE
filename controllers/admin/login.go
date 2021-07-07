package admin

import (
	"quickstart/common/service"
	"quickstart/common/function"
	"quickstart/common/constant"
)

// 员工登录
type LoginController struct {
	BaseController
}

func (c *LoginController) Index() {
	c.TplName = "login/index.html"
}

/**
 * 登录
 */
func (c *LoginController) DoLogin() {
	username :=	c.GetString("Username")
	userpass := c.GetString("Userpass")
	model := service.User{
		Username: username,
		Userpass: userpass,
	}
	resurl, status := model.DoLogin()
	if status == false {
		c.ajaxError(model.GetError(), "", "")
	}
	jsonStr, _ := function.JsonEncode(resurl)
	c.SetSession(constant.ADMIN_SESSION, jsonStr)
	c.ajaxSuccess("登录成功", "", "/")
}

func (c *LoginController) LoginOut() {
	c.WriteLog(constant.USER_EVENT_LOGOUT, "员工退出登录成功", "")
	c.DestroySession()
	c.ajaxSuccess("退出成功", "", "/login")
}