package redis

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	cacheDb "github.com/astaxie/beego/cache"
	"quickstart/common/lib/cache"
	_ "github.com/astaxie/beego/cache/redis"
	_ "github.com/gomodule/redigo/redis"
	"time"
)

var redisCache cacheDb.Cache

/**
 * 初始化缓存
 */
func init() {
	redisHost := beego.AppConfig.String("redis::rHost")
	dataBase := beego.AppConfig.String("redis::rDatabase")
	password := beego.AppConfig.String("redis::rPassword")
	config := fmt.Sprintf(`{"key":"demo","conn":"%s","dbNum":"%s","password":"%s"}`, redisHost, dataBase, password)
	var err error
	redisCache, err = cacheDb.NewCache("redis", config)
	if err != nil {
		errMsg := "failed to init redis"
		beego.Error(errMsg, err)
		panic(errMsg)
	}
	cache.Register("redis", NewRedisCache)
}

type Cache struct {}

/**
 * 构造函数
 */
func NewRedisCache() cache.Cache {
	return &Cache{}
}

/**
 * 设置缓存
 * @param string key
 * @param interface{} value
 * @param time.Duration time
 * @return error
 */
func (c *Cache) SetStr(key string, value interface{}, time time.Duration) (err error) {
	jsons, _ := json.Marshal(value)
	err = redisCache.Put(key, string(jsons), time)
	if err != nil {
		beego.Error("set key:", key, ",value:", value, err)
	}
	return
}

/**
 * 获取缓存
 * @param string key
 * @return data， error
 */
func (c *Cache) GetStr(key string) (data string, err error) {
	v := redisCache.Get(key)
	if v == nil {
		return data, errors.New("cache no")
	}
	value := string(v.([]byte)) //这里的转换很重要，Get返回的是interface
	return value, err
}

/**
 * 删除缓存
 * @param string key
 * @return error
 */
func (c *Cache) DelKey(key string) (err error) {
	err = redisCache.Delete(key)
	return
}