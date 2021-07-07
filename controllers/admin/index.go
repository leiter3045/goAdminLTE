package admin

type IndexController struct {
	BaseController
}

// 系统首页
func (c *IndexController) Index() {
	c.TplName = "index/index.html"
}


