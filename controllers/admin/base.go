package admin

import (
	"errors"
	"github.com/astaxie/beego"
	"quickstart/common/model"
	"quickstart/common/org"
	"quickstart/common/constant"
	"quickstart/common/function"
	"strings"
)

const USER_AUTH_KEY = 1

type BaseController struct {
	beego.Controller
	Userlog 	model.Userlog
	pageRow 	int
	RequestTime string
	UserCookie 	string
	loginuser 	map[string]interface{}
	loginuserId int
	userrole 	map[string]interface{}
	UserAent 	string
}

/**
 * 初始化
 */
func (c *BaseController) Prepare() {
	// 检查登录状态
	if c.ChaekLogin() == false {
		c.Redirect("/login", 302)
	} else {
		if c.loginuser != nil {
			if c._checkRole() == false {
				c.ajaxError("亲，不要乱来，你没有权限~~", "", "")
			}
			c._dataInit()
		}
	}
}

func (c *BaseController) list(totalRow int)  {
	model := org.Page{Request: c.Ctx.Request}
	model.NewPage().SetParams(totalRow, c.pageRow)
	model.SetConfig("prev", "上一页")
	model.SetConfig("next", "下一页")
	model.SetConfig("theme", " %HEADER%<li class='disabled'><a>%NOW_PAGE%/%TOTAL_PAGE% 页</a></li>%UP_PAGE%%FIRST%%LINK_PAGE%%DOWN_PAGE%")
	str := model.Show()
	c.Data["page_count"] = totalRow
	c.Data["pages"] = str
}

func (c *BaseController) ajaxError(msg string, data interface{}, url string) {
	response := make(map[string]interface{})
	response["info"] = msg
	response["status"] = 0
	response["data"] = data
	response["url"] = url
	c.Data["json"] = response
	c.ServeJSON()
	c.StopRun()
}

func (c *BaseController) ajaxSuccess(msg string, data interface{}, url string) {
	response := make(map[string]interface{})
	response["info"] = msg
	response["status"] = 1
	response["data"] = data
	response["url"] = url
	c.Data["json"] = response
	c.ServeJSON()
	c.StopRun()
}

/**
 * 数据初始化
 * @return bool
 */
func (c *BaseController) _dataInit() {
	c._getTreeMenu()
	c.pageRow = 15
	c.RequestTime = c.Ctx.Input.Header("request-time")
	c.UserAent = c.Ctx.Input.Header("user-agent")
	userAgent := org.UserAgent{Ctx: c.Ctx}
	c.UserCookie = userAgent.ComputerCookie(c.UserAent, c.RequestTime, constant.COMPUTER_COOKIE)
}

/**
 * 数据初始化
 * @return bool
 */
func (c *BaseController) _getTreeMenu() {
	if c.loginuserId == USER_AUTH_KEY {
		model := new (model.Menu)
		menus, _ := model.GetMenuByAll(constant.MENU_STATE_ON)
		arr := c.getMenu(menus, make(map[int]string))
		c.Data["munes"] = arr
	} else {
		accessModel := model.Access{}
		access, _ := accessModel.GetRoleByAll(int(c.userrole["id"].(float64)))
		var arrIds []int
		for _, v := range access {
			arrIds = append(arrIds, v.MenuId)
		}
		menuModel := model.Menu{}
		menus, _ := menuModel.GetMenuByIds(constant.MENU_STATE_ON, arrIds)
		arr := c.getMenu(menus, make(map[int]string))
		c.Data["munes"] = arr
	}
}

/**
 * 获取URL权限
 */
func (c *BaseController) _checkRole() bool {
	if c.loginuserId == USER_AUTH_KEY {
		return true
	}
	url := c.Ctx.Input.URL()
	if url == "/login" {
		return true
	}
	pos := strings.Index(url, "ajax")
	if pos > 0 {
		return true
	}
	roleMenu := model.Menu{}
	info, err := roleMenu.GetInfoByUrl(c.Ctx.Input.URL())
	if err != nil {
		return false
	}
	hasAccess := model.Access{}
	if _, err = hasAccess.GetMenuIdsByRole(c.loginuserId, info.Id); err != nil {
		return false
	}
	return true
}

/**
 * 检查登录状态
 * @return bool
 */
func (c *BaseController) ChaekLogin() bool {
	url := c.Ctx.Input.URL()
	if url == "/login" {
		return true
	}
	loginJson := c.GetSession(constant.ADMIN_SESSION)
	if loginJson == nil {
		return  false
	} else {
		json, _ := function.JsonDecode(loginJson.(string))
		c.loginuser = json
		c.loginuserId = int(c.loginuser["Id"].(float64))
		c.userrole  = json["Roles"].(map[string]interface{})
		c.Data["auth"] = c.loginuser
	}
	return true
}

/**
 * 写日志
 * @param int rventId
 * @return bool
 */
func (c *BaseController) WriteLog(erventId int,  desc string, ids string) (err error) {
	user := model.User{Id: c.loginuserId}
	log := model.Userlog{
		Users:     	&user,
		EventId: 	erventId,
		RelationId: ids,
		Module: 	c.Ctx.Input.URL(),
		Desc : 		desc,
		Ip: 		c.Ctx.Input.IP(),
		Browser:	c.UserAent,
		Cookie: 	c.UserCookie,
	}
	if _, status := log.Write(); status == false {
		return errors.New(log.GetError());
	}
	return nil
}

/**
 * 格式化菜单列表
 * @param unknown $menus
 * @return multitype:unknown
 */
func (c *BaseController) getMenu(menus []*model.Menu, m map[int]string) []org.Tree {
	var url string
	var arr []org.Tree
	var mod []org.Tree
	//controllerName, actionName := c.GetControllerAndAction()
	//url = controllerName + "." + actionName
	//url = c.URLFor(url)
	url = c.Ctx.Input.URL()
	for _, v := range menus {
		ctree := org.Tree{}
		if len(m[v.Id]) > 0 {
			ctree.Id = v.Id
			ctree.Pid = v.Pid
			ctree.Name = v.Name
			ctree.Url = v.Url
			ctree.Icon = v.Icon
			ctree.State = v.State
			ctree.Status = true
		} else {
			ctree.Id = v.Id
			ctree.Pid = v.Pid
			ctree.Name = v.Name
			ctree.Url = v.Url
			ctree.Icon = v.Icon
			ctree.State = v.State
			ctree.Status = false
		}
		mod = append(mod, ctree)
	}
	arr = c.filterMenu(mod, url)
	return arr
}

/**
 * 根据权限过滤菜单
 * @param unknown $menus
 * @return multitype:unknown
 */
func (c *BaseController) filterMenu(menus []org.Tree, url string) ([]org.Tree) {
	// 先格式化TREE
	menusArray := org.ToArrayTree(menus, "Child", 0, url)
	var arr []org.Tree
	for _, v := range menusArray {
		if strings.ToLower(v.Url) == strings.ToLower(url) {
			v.Active = true
		}
		for _, v1 := range v.Child {
			if v1.Active {
				v.Active = true
			}
		}
		arr = append(arr, v)
	}
	return arr;
}
