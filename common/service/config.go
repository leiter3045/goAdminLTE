package service

import (
	"encoding/json"
	"quickstart/common/model"
	"quickstart/common/lib/traits"
	"quickstart/common/function"
	"time"
)

type Config struct {
	traits.ErrorReport
	Id          int
	Name   		string
	Value    	string
}

/**
 * 获取配置数据支持缓存
 * @param bool refresh
 * @return array
 */
func (c Config) GetInfo(refresh bool) interface{} {
	str, err := function.Cache("sysconfig", "", 0)
	if err != nil || refresh == true {
		model := model.Config{}
		dataSlice, errs := model.GetInfoByAll()
		if errs == nil {
			function.Cache("sysconfig", dataSlice, 1000*time.Second)
		}
		str, _ = function.Cache("sysconfig", "", 0)
	}
	var dat []interface{}
	err = json.Unmarshal([]byte(str.(string)), &dat)
	return dat
}