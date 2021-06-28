package admin

import (
	"quickstart/common/model"
	"quickstart/common/org"
	"quickstart/common/constant"
	"strconv"
)

type MenuController struct {
	BaseController
	model.Menu
}

func (c *MenuController) Index() {
	menus, _ := c.GetMenuByAll(9)
	m := org.ToHtmlTree(menus, 0, 0)
	c.Data["lists"] = m
	c.TplName = "menu/index.html"
}

func (c *MenuController) GetAll() {
	info, _ := c.GetMenuByAll(1)
	c.ajaxSuccess("获取数据成功", info, "")
}

func (c *MenuController) Add() {
	menus, _ := c.GetMenuByAll(1)
	m := org.ToHtmlTree(menus, 0, 0)
	c.Data["menualls"] = m
	c.TplName = "menu/add.html"
}

func (c *MenuController) Edit() {
	id , _ := c.GetInt("id", 0)
	info, _ := c.GetInfoById(id)
	menus, _ := c.GetMenuByAll(1)
	m := org.ToHtmlTree(menus, 0, 0)
	c.Data["menualls"] = m
	c.Data["data"] = info
	c.TplName = "menu/add.html"
}

func (c *MenuController) EditPost() {
	id , err := c.GetInt("id", 0)
	if err != nil {
		c.ajaxError("ID不存在", "", "")
	}
	sort , _ := c.GetInt("sort", 50)
	pid , _ := c.GetInt("pid", 0)
	state , _ := c.GetInt("state", 0)
	model := model.Menu{
		Id: 	id,
		Pid:    pid,
		Name:   c.GetString("name"),
		Url:	c.GetString("url"),
		Icon:   c.GetString("icon"),
		Sort:   sort,
		State:  state,
	}
	var arr []string
	if status := model.Edit(); status == false {
		c.WriteLog(constant.USER_EVENT_EDIT, "尝试编辑菜单失败", "")
		c.ajaxError(model.GetError(), arr, "")
	}
	c.WriteLog(constant.USER_EVENT_EDIT, "尝试编辑菜单成功", strconv.Itoa(id))
	c.ajaxSuccess("编辑数据成功！", arr, "-1")
}

func (c *MenuController) AddPost() {
	sort , _ := c.GetInt("sort", 50)
	pid , _ := c.GetInt("pid", 0)
	state , _ := c.GetInt("state", 0)
	model := model.Menu{
		Pid:    pid,
		Name:   c.GetString("name"),
		Url:	c.GetString("url"),
		Icon:   c.GetString("icon"),
		Sort:   sort,
		State:  state,
	}
	var arr []string
	id, status := model.Add()
	if status == false {
		c.WriteLog(constant.USER_EVENT_ADD, "尝试添加菜单失败", "")
		c.ajaxError(model.GetError(), "", "")
	}
	c.WriteLog(constant.USER_EVENT_ADD, "尝试添加菜单成功", strconv.Itoa(id))
	c.ajaxSuccess("新增数据成功！", arr, "-1")
}

func (c *MenuController) Del()  {
	if c.loginuser["Id"] != USER_AUTH_KEY {
		c.ajaxError("菜单只有超级管理才可删除", "", "")
	}
	id , err := c.GetInt("id", 0)
	if err != nil {
		c.ajaxError("获取参数失败", "", "")
	}
	model := model.Menu{}
	if status := model.DelInfoById(id); status == false {
		c.WriteLog(constant.USER_EVENT_DROP, "尝试删除菜单失败", strconv.Itoa(id))
		c.ajaxError(model.GetError(), "", "")
	}
	c.WriteLog(constant.USER_EVENT_DROP, "尝试删除菜单成功", strconv.Itoa(id))
	c.ajaxSuccess("删除成功！", "", "-1")
}
