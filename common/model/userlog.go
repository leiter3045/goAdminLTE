package model

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"quickstart/common/validate"
	"quickstart/common/lib/traits"
	"time"
)

type Userlog struct {
	traits.ErrorReport
	Id       	int			`orm:"column(id);auto;size(5)" description:"主键"`
	Users    	*User		`orm:"column(user_id);rel(one);" description:"管理员ID"`
	EventId   	int			`orm:"column(event_id);size(3);type(int);" description:"操作事件ID"`
	RelationId  string    	`orm:"column(relation_id);size(255);type(varchar);default('');" description:"关联信息ID，可能多个"`
	Module   	string    	`orm:"column(module);size(30);type(varchar);default('');" description:"所属模型"`
	Desc  		string 		`orm:"column(desc);size(127);type(varchar);default('');" description:"操作描述"`
	Ip 			string		`orm:"column(ip);size(15);type(char);default('');" description:"操作用户IP"`
	//Os 			string		`orm:"column(os);size(32);;type(varchar);default('');" description:"操作系统"`
	Browser 	string		`orm:"column(browser);size(500);;type(varchar);default('');" description:"浏览器"`
	Cookie 		string		`orm:"column(cookie);size(32);type(varchar);default('');" description:"电脑识别码"`
	AddTime 	int64		`orm:"column(add_time);size(10);type(int);" description:"操作时间"`
}

func (t *Userlog) TableName() string {
	return "tp_user_log"
}

// 需要在init中注册定义的model
func init() {
	orm.RegisterModel(new(Userlog))
}

/**
 * 通过菜单ID获取详情
 * @param int id
 * @return array
 */
func (c Userlog) GetInfoById(id int) (v *Userlog, err error) {
	o := orm.NewOrm()
	v = &Userlog{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

/**
 * 写入管理员操作日志
 * @return boolean
 */
func (c *Userlog) Write() (id int, status bool) {
	valid := validate.Userlog{
		EventId: c.EventId,
		Desc: 	 c.Desc,
		Ip:		 c.Ip,
	}
	if state := valid.Valid(); state == false {
		c.SetError(valid.GetError())
		return id, false
	}
	c.AddTime = time.Now().Unix()
	var err error
	o := orm.NewOrm()
	var userlogId int64
	userlogId, err = o.Insert(c)
	if err != nil {
		c.SetError("插入数据失败，请联系客服")
		return id, false
	}
	return int(userlogId), true
}

/**
 * 获取所有员工
 * @return array
 */
func (c Userlog) GetList(where map[string]interface{}, limit int, offset int, relation bool) (m []*Userlog, count int, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(c)
	var qscount int64
	qscount, err = qs.Count()
	if err != nil {
		return m, count, errors.New("未获取到数据")
	}
	if where != nil {
		for k, v := range where {
			qs = qs.Filter(k, v)
		}
	}
	qs = qs.OrderBy("-id")
	qs = qs.Limit(limit, offset)
	qs.All(&m)
	if relation {
		for _, v := range m {
			o.LoadRelated(v, "Users")
		}
	}
	return m, int(qscount), err
}

