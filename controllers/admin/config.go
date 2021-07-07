package admin

import (
	"quickstart/common/model"
	"quickstart/common/constant"
)

// 系统配置
type ConfigController struct {
	BaseController
}

/**
 * 页面
 */
func (c *ConfigController) Index() {
	c.TplName = "config/index.html"
}

/**
 * 编辑
 */
func (c *ConfigController) Edit() {
	model := new (model.Config)
	webname := c.GetString("webname")
	if err := model.Edit("webname", webname); err == false {
		c.ajaxError(model.GetError(), "", "")
	}
	website := c.GetString("website")
	if err := model.Edit("website", website); err == false {
		c.ajaxError(model.GetError(), "", "")
	}
	c._getConfig(true)
	c.WriteLog(constant.USER_EVENT_EDIT, "编辑系统配置成功", "")
	c.ajaxSuccess("更新成功", "", "-1")
}