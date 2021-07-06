package service

import (
	"fmt"
	"quickstart/common/model"
	"quickstart/common/lib/traits"
	"quickstart/common/function"
)

type Config struct {
	traits.ErrorReport
	Id          int
	Name   		string
	Value    	string
}

func (c Config) GetInfo(refresh bool) interface{} {
	model := model.Config{}
	arr, _ := model.GetInfoByAll()
	str, err := function.Cache("sysconfig", arr, 3600)
	fmt.Print(err)
	//list := arr
	//list, err := function.Cache("sysconfig", "", 0)
	//if err != nil || refresh == true {
	//	model := model.Config{}
	//	arr, _ := model.GetInfoByAll()
	//	function.Cache("sysconfig", arr, 3600)
	//	if err == nil {
	//		//function.Cache("sysconfig", list, 3600)
	//	}
	//	list = arr
	//}
	return str
}