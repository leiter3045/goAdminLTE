package org

import (
	"quickstart/common/model"
	"strings"
	"reflect"
)

type Tree struct {
	Id     int
	Name   string
	Pid    int
	Sort   int
	Level  int
	Url    string
	Icon   string
	Html   string
	State  int
	Active bool
	Status bool
	Child  []Tree
}

/**
 * 格式化数组
 * @param array cate
 * @param string html
 * @param number pid
 * @param number level
 * @param string path
 */
func ToHtmlTree(cate []*model.Menu, pid int, level int) []Tree {
	var arr []Tree
	html := "├─"
	for _, v := range cate {
		if pid == v.Pid {
			ctree := Tree{}
			ctree.Id = v.Id
			ctree.Pid = v.Pid
			ctree.Name = v.Name
			ctree.Sort = v.Sort
			ctree.Icon = v.Icon
			ctree.Html = strings.Repeat(html, level)
			ctree.Level = level
			ctree.Url = v.Url
			ctree.State = v.State
			arr = append(arr, ctree)
			sonCate := ToHtmlTree(cate, v.Id, level+1)
			arr = append(arr, sonCate...)
		}
	}
	return arr;
}

func InArray(need interface{}, needArr []interface{}) bool {
	for _,v := range needArr{
		if need == v{
			return true
		}
	}
	return false
}

// IsArray is 判断是否数组或者切片
func IsArray(v interface{}) bool {
	rv := reflect.ValueOf(v)
	return rv.Kind() == reflect.Array || rv.Kind() == reflect.Slice
}

/**
 * 格式化数组
 * @param array cate
 * @param string name
 * @param number access
 * @param number pid
 */
func ToArrayTree(cate []Tree, name string, pid int, url string) (resListArr []Tree) {
	for _, v := range cate {
		if v.Pid == pid {
			topMenu := Tree{}
			topMenu.Id = v.Id
			topMenu.Pid = v.Pid
			topMenu.Name = v.Name
			topMenu.Icon = v.Icon
			topMenu.Url = v.Url
			topMenu.Status = v.Status
			if v.Url == url {
				topMenu.Active = true
			}
			topMenu.Child = ToArrayTree(cate, name, v.Id, url)
			resListArr = append(resListArr, topMenu)
		}
	}
	return resListArr;
}