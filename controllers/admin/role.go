package admin

import (
	"quickstart/common/model"
	"quickstart/common/constant"
	"strconv"
	"strings"
)

// 角色管理
type RoleController struct {
	BaseController
	model.Role
}

/**
 * 页面
 */
func (c *RoleController) Index() {
	page , _  := c.GetInt("page", 0)
	lists, count, _ := c.GetList(9, c.pageRow, (page - 1) * c.pageRow)
	c.list(count)
	c.Data["lists"] = lists
	c.TplName = "role/index.html"
}

/**
 * 添加页面
 */
func (c *RoleController) Add() {
	c.TplName = "role/add.html"
}

/**
 * 添加数据
 */
func (c *RoleController) AddPost() {
	state , _ := c.GetInt("state", 0)
	model := model.Role{
		Name:   	c.GetString("name"),
		State:  	state,
	}
	id, status := model.Add()
	if status == false {
		c.WriteLog(constant.USER_EVENT_ADD, "尝试添加角色失败", "")
		c.ajaxError(model.GetError(), "", "")
	}
	c.WriteLog(constant.USER_EVENT_ADD, "尝试添加角色成功", strconv.Itoa(id))
	c.ajaxSuccess("新增成功！", "", "-1")
}

/**
 * 编辑页面
 */
func (c *RoleController) Edit() {
	id , err := c.GetInt("id", 0)
	if err != nil {
		c.ajaxError("ID不存在", "", "")
	}
	info, _ := c.GetInfoById(id)
	c.Data["data"] = info
	c.TplName = "role/add.html"
}

/**
 * 编辑数据
 */
func (c *RoleController) EditPost() {
	id , err := c.GetInt("id", 0)
	if err != nil {
		c.ajaxError("ID不存在", "", "")
	}
	state , _ := c.GetInt("state", 0)
	model := model.Role{
		Id:			id,
		Name:   	c.GetString("name"),
		State:  	state,
	}
	if status := model.Edit(); status == false {
		c.WriteLog(constant.USER_EVENT_EDIT, "尝试编辑角色失败", strconv.Itoa(id))
		c.ajaxError(model.GetError(), "", "")
	}
	c.WriteLog(constant.USER_EVENT_EDIT, "尝试编辑角色成功", strconv.Itoa(id))
	c.ajaxSuccess("编辑数据成功！", "", "-1")
}

type RoleInterface struct {
	RoleId    		int
	MenuId   		int
	Url   			string
}

/**
 * 添加权限页面
 */
func (c *RoleController) Access()  {
	id, _ := c.GetInt("id", 0)
	info, _ := c.GetInfoById(id)
	c.Data["data"] = info
	modelAccess := model.Access{}
	access, _ := modelAccess.GetRoleByAll(id)
	accessIds := make(map[int]string)
	for _, v := range access {
		key :=  int(v.MenuId)
		accessIds[key] = v.Url
	}
	modelMenu := model.Menu{}
	menus, _ := modelMenu.GetMenuByAll(constant.MENU_STATE_ON)
	allMunes := c.getMenu(menus, accessIds)
	c.Data["all_munes"] = allMunes
	c.TplName = "role/access.html"
}

type EditorInput struct {
	RoleId    int      `form:"role_id"`
	Rulegroup []string `form:"rule"`
}

/**
 * 添加权限
 */
func (c *RoleController) AccessPost()  {
	ei := EditorInput{}
	if err := c.ParseForm(&ei); err != nil {
		c.ajaxError("参数错误", ei, "")
	}
	delModel := model.Access{}
	if status := delModel.DelRoleAll(ei.RoleId); status == false {
		c.ajaxError(delModel.GetError(), "", "")
	}
	for _, check := range ei.Rulegroup{
		strArray := strings.Split(check,",")
		menuId, _ := strconv.Atoi(strArray[0])
		mod := model.Access{
			RoleId: ei.RoleId,
			MenuId: menuId,
			Url:	strArray[1],
		}
		if status := mod.Add(); status == false{
			c.WriteLog(constant.USER_EVENT_ACCESS, "尝试编辑角色权限失败", "")
			c.ajaxError(mod.GetError(), "", "")
		}
	}
	c.WriteLog(constant.USER_EVENT_ACCESS, "尝试编辑角色权限失败", strings.Join(ei.Rulegroup, ","))
	c.ajaxSuccess("添加角色权限成功", "", "-1")
}

/**
 * 删除
 */
func (c *RoleController) Del()  {
	id , err := c.GetInt("id", 0)
	if err != nil {
		c.ajaxError("获取参数失败", "", "")
	}
	model := model.Role{}
	if status := model.DelInfoById(id); status == false {
		c.WriteLog(constant.USER_EVENT_DROP, "尝试编辑角色失败", strconv.Itoa(id))
		c.ajaxError(model.GetError(), "", "")
	}
	c.WriteLog(constant.USER_EVENT_DROP, "尝试编辑角色成功", strconv.Itoa(id))
	c.ajaxSuccess("删除成功！", "", "-1")
}