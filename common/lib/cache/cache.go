package cache

import (
	"errors"
	"github.com/astaxie/beego"
)

type Cache interface {
	SetStr(key, value string, time int64) error
	GetStr(key string) string
	DelKey(key string) error
}

type Instance func() Cache

var adapters = make(map[string]Instance)

func Register(name string, adapter Instance) {
	if adapter == nil {
		panic("cache: Register adapter is nil")
	}
	if _, ok := adapters[name]; ok {
		panic("cache: Register called twice for adapter " + name)
	}
	adapters[name] = adapter
}

func GetInstance(adapterName string) (adapter Cache, err error) {
	if adapterName == "" {
		adapterName = beego.AppConfig.String("cachename")
	}

	if instanceFunc, ok := adapters[adapterName]; ok == true {
		return instanceFunc(), nil
	}
	return adapter, errors.New("cache: unknown:" + adapterName)
}

