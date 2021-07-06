package service

import (
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

func (c Config) GetInfo(refresh bool) interface{} {
	list, err := function.Cache("sysconfig", "", 0)
	if err != nil || refresh == true {
		model := model.Config{}
		arr, errs := model.GetInfoByAll()
		if errs == nil {
			function.Cache("sysconfig", arr, 1000*time.Second)
		}
		list = arr
	}
	return list
}