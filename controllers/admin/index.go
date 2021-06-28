package admin

type IndexController struct {
	BaseController
}

func (c *IndexController) Index() {
	c.TplName = "index/index.html"
}


