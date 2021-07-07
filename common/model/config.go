package model

import (
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
 * 通过ID获取详情
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
 * 获取所有数据
 * @return array
 */
func (c Config) GetInfoByAll() (m []*Config, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(c)
	qs.All(&m);
	return m, err
}

/**
 * 编辑
 */
func (c *Config) Edit(name, value string) (bool) {
	valid := validate.Config{
		Name:  name,
		Value: value,
	}
	if state := valid.Valid(); state == false {
		c.SetError(valid.GetError())
		return false
	}
	o := orm.NewOrm()
	config := Config{Name: name}
	if o.Read(&config, "Name") != nil {
		c.SetError("未获取的数据")
		return false
	}
	config.Value = value
	if _, err := o.Update(&config, "Value"); err != nil {
		c.SetError("修改数据失败")
		return false
	}
	return true
}
