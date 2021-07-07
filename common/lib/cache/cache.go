package cache

import (
	"github.com/astaxie/beego"
	"errors"
	"time"
)

type Cache interface {
	SetStr(key string, value interface{}, time time.Duration) error
	GetStr(key string) (string, error)
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

func GetInstance( ) (adapter Cache, err error) {
	var adapterName string
	adapterName = beego.AppConfig.String("cachename")
	if adapterName == "" {
		return adapter, errors.New("cache error: Can't find configuration")
	}
	if instanceFunc, ok := adapters[adapterName]; ok == true {
		return instanceFunc(), nil
	}
	return adapter, errors.New("cache: unknown:" + adapterName)
}

