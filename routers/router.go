package routers

import (
	"quickstart/controllers/admin"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &admin.IndexController{}, "Get:Index")
	beego.Router("/menu", &admin.MenuController{}, "Get:Index")
	beego.Router("/menu/add", &admin.MenuController{}, "Get:Add;POST:AddPost")
	beego.Router("/menu/edit", &admin.MenuController{}, "Get:Edit;POST:EditPost")
	beego.Router("/menu/del", &admin.MenuController{}, "Get:Del")
	beego.Router("/user", &admin.UserController{}, "Get:Index")
	beego.Router("/user/add", &admin.UserController{}, "Get:Add;POST:AddPost")
	beego.Router("/user/edit", &admin.UserController{}, "Get:Edit;POST:EditPost")
	beego.Router("/user/del", &admin.UserController{}, "Get:Del")
	beego.Router("/user/ajaxpass", &admin.UserController{}, "Get:AjaxPass")
	beego.Router("/user/access", &admin.UserController{}, "Get:Access;POST:AccessPost")
	beego.Router("/role", &admin.RoleController{}, "Get:Index")
	beego.Router("/role/add", &admin.RoleController{}, "Get:Add;POST:AddPost")
	beego.Router("/role/edit", &admin.RoleController{}, "Get:Edit;POST:EditPost")
	beego.Router("/role/del", &admin.RoleController{}, "Get:Del")
	beego.Router("/role/access", &admin.RoleController{}, "Get:Access;POST:AccessPost")
	beego.Router("/userlog", &admin.UserlogController{}, "Get:Index")
	beego.Router("/login", &admin.LoginController{}, "Get:Index;POST:DoLogin")
	beego.Router("/login/ajaxout", &admin.LoginController{}, "Get:LoginOut")
}
