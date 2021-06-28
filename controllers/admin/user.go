package admin

import (
	"quickstart/common/model"
	"quickstart/common/constant"
	"strconv"
	"time"
)

type UserController struct {
	BaseController
	model.User
}

func (c *UserController) Index() {
	page , _  := c.GetInt("page", 0)
	lists, conut, _ := c.GetList(9, c.pageRow, (page - 1) * c.pageRow, true)
	c.list(conut)
	c.Data["lists"] = lists
	c.TplName = "user/index.html"
}

func (c *UserController) Add() {
	c.TplName = "user/add.html"
}

func (c *UserController) AddPost() {
	state , _ := c.GetInt("state", 0)
	userpass := c.GetString("userpass")
	userpassConfirm := c.GetString("userpass_confirm")
	var arr []string
	if (userpass != userpassConfirm) {
		c.ajaxError("两次输入的密码不一致", arr, "")
	}
	role := new(model.Role)
	model := model.User{
		Username:   c.GetString("username"),
		Userpass:	userpass,
		State:  	state,
		AddTime:    time.Now().Unix(),
		Roles: 		role,
	}
	id, status := model.Add()
	if status != true {
		c.WriteLog(constant.USER_EVENT_ADD, "尝试添加员工失败", "")
		c.ajaxError(model.GetError(), arr, "")
	}
	c.WriteLog(constant.USER_EVENT_ADD, "尝试添加员工失败", strconv.Itoa(id))
	c.ajaxSuccess("新增成功！", arr, "-1")
}

func (c *UserController) Edit() {
	id , err := c.GetInt("id", 0)
	if err != nil {
		c.ajaxError("ID不存在", "", "")
	}
	info, _ := c.GetInfoById(id, false)
	c.Data["data"] = info
	c.TplName = "user/add.html"
}

func (c *UserController) EditPost() {
	id , err := c.GetInt("id", 0)
	if err != nil {
		c.ajaxError("ID不存在", "", "")
	}
	state , _ := c.GetInt("state", 0)
	userpass := c.GetString("userpass")
	userpassConfirm := c.GetString("userpass_confirm")
	var arr []string
	if (userpass != userpassConfirm) {
		c.ajaxError("两次输入的密码不一致", arr, "")
	}
	role := new(model.Role)
	model := model.User{
		Id: 		id,
		Username:   c.GetString("username"),
		Userpass:	userpass,
		State:  	state,
		Roles: 		role,
	}
	if status := model.Edit(); status == false {
		c.WriteLog(constant.USER_EVENT_EDIT, "尝试编辑员工失败", strconv.Itoa(id))
		c.ajaxError(model.GetError(), arr, "")
	}
	c.WriteLog(constant.USER_EVENT_EDIT, "尝试编辑员工成功", strconv.Itoa(id))
	c.ajaxSuccess("编辑数据成功！", arr, "-1")
}

func (c *UserController) Access()  {
	id, _ := c.GetInt("id", 0)
	info, _ := c.GetInfoById(id, false)
	c.Data["user"] = info
	model := model.Role{}
	roles, _ := model.GetRoleByAll(1)
	c.Data["all_role"] = roles
	c.TplName = "user/access.html"
}

func (c *UserController) AccessPost()  {
	userId, _ := c.GetInt("id", 0)
	roleId, _ := c.GetInt("role_id", 0)
	user := model.User{Id: userId}
	if status := user.EditRole(roleId); status == false {
		c.WriteLog(constant.USER_EVENT_ACCESS, "尝试编辑员工权限失败", strconv.Itoa(userId))
		c.ajaxError("权限添加失败，请联系技术", "", "")
	}
	c.WriteLog(constant.USER_EVENT_ACCESS, "尝试编辑员工权限成功", strconv.Itoa(userId))
	c.ajaxSuccess("添加员工角色成功", "", "-1")
}

func (c *UserController) AjaxPass() {
	c.TplName = "user/useredit.html"
}

func (c *UserController) Del()  {
	id , err := c.GetInt("id", 0)
	if err != nil {
		c.ajaxError("获取参数失败", "", "")
	}
	if id == USER_AUTH_KEY {
		c.ajaxError("超级管理不允许删除", "", "")
	}
	model := model.User{}
	if status := model.DelInfoById(id); status == false {
		c.WriteLog(constant.USER_EVENT_ACCESS, "尝试删除员工失败", strconv.Itoa(id))
		c.ajaxError(model.GetError(), "", "")
	}
	c.WriteLog(constant.USER_EVENT_ACCESS, "尝试删除员工成功", strconv.Itoa(id))
	c.ajaxSuccess("删除成功！", "", "-1")
}
