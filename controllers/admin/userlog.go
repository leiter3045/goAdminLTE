package admin

import (
	"quickstart/common/model"
	"quickstart/common/constant"
	"quickstart/common/function"
)

// 员工操作日志
type UserlogController struct {
	BaseController
	model.Userlog
}

func (c *UserlogController) Index() {
	page , _  := c.GetInt("page", 0)
	where := c._search_where()
	list, conut, _ := c.GetList(where, c.pageRow, (page - 1) * c.pageRow, true)
	models := new (constant.Constant)
	var lists []map[string]interface{}
	for _, v := range list {
		_, event, _ := models.UserEvent(v.EventId)
		log := make(map[string]interface{}, 0)
		log["Id"]			= v.Id
		log["Users"] 		= v.Users
		log["Event"]  		= event
		log["RelationId"] 	= v.RelationId
		log["Module"] 		= v.Module
		log["Desc"]			= v.Desc + "，引用ID：" + v.RelationId
		log["Ip"] 			= v.Ip
		log["Browser"] 		= v.Browser
		log["Cookie"] 		= v.Cookie
		log["AddTime"] 		= function.ConvertToTime(v.AddTime)
		lists = append(lists, log)
	}
	c.list(conut)
	events, _, _ := models.UserEvent(0)
	c.Data["lists"] = lists
	c.Data["events"] = events
	c.TplName = "userlog/index.html"
}

func (c *UserlogController) _search_where() map[string]interface{} {
	where := make(map[string]interface{}, 0)
	getParams := make(map[string]interface{}, 0)
	if val := c.GetString("username"); len(val) > 0 {
		model := new (model.User)
		info, _ := model.GetInfoUsername(val, false)
		where["user_id"] = info.Id
		getParams["username"] = val
	}
	if val := c.GetString("ip"); len(val) > 0 {
		where["ip"] = val
		getParams["ip"] = val
	}
	if val := c.GetString("startTime"); len(val) > 0 {
		where["add_time__gte"] = val
		getParams["startTime"] = val
	}
	if val := c.GetString("endTime"); len(val) > 0 {
		where["add_time__lte"] = val
		getParams["endTime"] = val
	}
	if val, _ := c.GetInt("event_id", 0); val > 0 {
		where["event_id"] = val
		getParams["event_id"] = val
	}
	c.Data["getParams"] = getParams
	return where
}
