package model

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"quickstart/common/validate"
	"quickstart/common/lib/traits"
)

type Config struct {
	traits.ErrorReport
	Id  			int   		`orm:"auto;column(id)"`
	Name    		string		`orm:"column(name);size(32);type(varchar);" description:"字段名，字母下划线"`
	Value   		string		`orm:"column(value);size(255);type(varchar);" description:"加减分项目的内容"`
}

func (t *Config) TableName() string {
	return "tp_config"
}

// 需要在init中注册定义的model
func init() {
	orm.RegisterModel(new(Config))
}

/**
 * 通过菜单ID获取详情
 * @param int id
 * @return array
 */
func (c Config) GetInfoById(Id int) (v *Config, err error) {
	o := orm.NewOrm()
	v = &Config{Id: Id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

/**
 * 通过菜单ID获取详情
 * @param int id
 * @return array
 */
func (c Config) GetInfoByAll() (m []*Config, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(c)
	qs.All(&m);
	fmt.Print(m)
	return m, err
}

/**
 * 编辑菜单
 */
func (c *Config) Edit() (bool) {
	valid := validate.Config{
		Name:  c.Name,
		Value: c.Value,
	}
	if state := valid.Valid(); state == false {
		c.SetError(valid.GetError())
		return false
	}
	o := orm.NewOrm()
	config := Config{Name: c.Name}
	if o.Read(&config) == nil {
		config.Value = c.Value
		if _, err := o.Update(&config, "Value"); err != nil {
			c.SetError("修改数据失败")
			return false
		}
	}
	return true
}
