package constant

const (
	// 定义系统加密解密KEY
	SYSTEM_SECRET_KEY = "RSA:!@#f8fe04b3a&*()_+"
	// JWT密钥
	LCOBUCCI_JWT_SECRET = "debdc9aad00dd7ed031684f7acb6640f"
	// 电脑cookie标识KEY
	COMPUTER_COOKIE 	= "computer_cookie"
	// 用户session标识KEY
	ADMIN_SESSION 		= "computer_admin"
	USER_EVENT_LOGIN 	= 1 // 登录平台
	USER_EVENT_LOGOUT 	= 2 // 退出登录
	USER_EVENT_VIEW 	= 3	// 信息查看
	USER_EVENT_ADD 		= 4 // 信息添加
	USER_EVENT_EDIT 	= 5 // 信息编辑
	USER_EVENT_DROP 	= 6 // 数据删除
	USER_EVENT_ACCESS 	= 7 // 权限设置
	// 员工状态
	USER_STATE_ON		= 1 // 正常
	USER_STATE_OFF		= 0 // 禁用
	// 员工状态
	MENU_STATE_ON		= 1 // 正常
	MENU_STATE_OFF		= 0 // 禁用
)

type Constant struct {}

/**
 * 管理员操作事件
 * @param int eventId
 */
func (c Constant) UserEvent(eventId int) (v map[int]string, eId string, err error) {
	options := make(map[int]string)
	options[USER_EVENT_LOGIN]  = "登录平台"
	options[USER_EVENT_LOGOUT] = "退出登录"
	options[USER_EVENT_VIEW]   = "信息查看"
	options[USER_EVENT_ADD]    = "信息添加"
	options[USER_EVENT_EDIT]   = "信息编辑"
	options[USER_EVENT_DROP]   = "数据删除"
	options[USER_EVENT_ACCESS] = "权限设置"
	return options, options[eventId], err
}