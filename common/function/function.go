package function

import (
	"encoding/json"
	"fmt"
	"log"
	"quickstart/common/constant"
	"quickstart/common/lib/cache"
	_ "quickstart/common/lib/cache/redis"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"regexp"
	"strings"
	"time"
)

/**
 * 时间戳转时间
 * @param int64 int_time
 * @return string
 */
func ConvertToTime(int_time int64) string {
	return time.Unix(int_time, 0).Format("2006-01-02 15:04")
}

/**
 * 获取唯一uid
 * @return string
 */
func GetUuid() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid
}

/**
 * 设置缓存通用函数，删除缓存,读取缓存
 * @param string key
 * @param interface{} value
 * @param time.Duration 缓存时间
 * @return interface{}，error
 */
func Cache(key string, value interface{}, time time.Duration) (val interface{}, err error) {
	instance, errs := cache.GetInstance()
	if errs != nil {
		return val, errs
	}
	if value == "" || time == -1 {
		if time == -1 {
			instance.DelKey(key)
			return
		} else {
			val, err = instance.GetStr(key)
			return val, err
		}
	} else {
		instance.SetStr(key, value, time)
		return
	}
}

/**
 * 时间转时间戳
 * @param string str
 * @return int64
 */
func Strtotime(str string) (int64) {
	layout := "2006-01-02 15:04:05"
	t, err := time.Parse(layout, str)
	if err != nil {
		return 0
	}
	return t.Unix()
}

/**
 * 字符串转数组
 * @param string delimiter
 * @param string text
 * @return []string
 */
func Explode(delimiter, text string) []string {
	if len(delimiter) > len(text) {
		return strings.Split(delimiter, text)
	} else {
		return strings.Split(text, delimiter)
	}
}

/**
 * 数组转字符串
 * @param string glue
 * @param []string pieces
 * @return string
 */
func Implode(glue string, pieces []string) string {
	return strings.Join(pieces, glue)
}

/**
 * 二维数组取值转一维数组
 * @param map[string]map[string]interface{} input
 * @param string columnKey
 * @return []interface{}
 */
func ArrayColumn(input map[string]map[string]interface{}, columnKey string) []interface{} {
	columns := make([]interface{}, 0, len(input))
	for _, val := range input {
		if v, ok := val[columnKey]; ok {
			columns = append(columns, v)
		}
	}
	return columns
}

/**
 * 数据转json
 * @param interface{} data
 * @return string
 */
func JsonEncode(data interface{}) (string, error) {
	jsons, err := json.Marshal(data)
	return string(jsons), err
}

/**
 * json转数组
 * @param string data
 * @return map[string]interface{}
 */
func JsonDecode(data string) (map[string]interface{}, error) {
	var dat map[string]interface{}
	err := json.Unmarshal([]byte(data), &dat)
	return dat, err
}

/**
 * 判断数组是否存在
 * @param interface needle
 * @param interface hystack
 * @return bool
 */
func InArray(needle interface{}, hystack interface{}) bool {
	switch key := needle.(type) {
	case string:
		for _, item := range hystack.([]string) {
			if key == item {
				return true
			}
		}
	case int:
		for _, item := range hystack.([]int) {
			if key == item {
				return true
			}
		}
	case int64:
		for _, item := range hystack.([]int64) {
			if key == item {
				return true
			}
		}
	default:
		return false
	}
	return false
}

/**
 * md5加密
 * @param string src
 * @return string
 */
func Md5(src string) string {
	m := md5.New()
	m.Write([]byte(src))
	res := hex.EncodeToString(m.Sum(nil))
	return res
}

/**
 * 密码加密算法
 * @param string password
 * @return string
 */
func PasswordHash(password string) string {
	return Md5(constant.SYSTEM_SECRET_KEY + password + constant.SYSTEM_SECRET_KEY)
}

/**
 * 使用正则验证数据
 * @param string $value
 * @param string $rule
 * @return boolean
 */
func ValidData(value string, rule string) bool {
	validate := make(map[string]string)
	validate["Require"] 	= "/.+/"
	validate["Email"] 		= "/^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$/"
	validate["Mobile"] 		= "/^(1[3456789])\\d{9}$/"
	validate["Domain"] 		= "/^http(s?):\\/\\/(?:[A-Za-z0-9-]+\\.)+[A-Za-z]{2,4}$/"
	validate["Url"] 		= "/^http(s?):\\/\\/(?:[A-Za-z0-9-]+\\.)+[A-Za-z]{2,4}(?:[\\/\\?#][\\/=\\?%\\-&~`@[\\]\\':+!\\.#\\w]*)?$/"
	validate["Number"] 		= "/^\\d+$/"
	validate["Integer"] 	= "/^[-\\+]?\\d+$/"
	validate["Double"] 		= "/^[-\\+]?\\d+(\\.\\d+)?$/"
	validate["English"] 	= "/^[A-Za-z]+$/"
	validate["Chinese"] 	= "/^[\\x{4e00}-\\x{9fa5}]{1,}$/u"
	validate["Username"] 	= "^[a-zA-Z][a-zA-Z0-9]{5,15}$"
	validate["Password"] 	= "^[a-zA-Z0-9]{5,15}$"
	validate["Iprefer"] 	= "/^\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}$/"

	// 检查是否有内置的正则表达式
	if(validate[rule] != ""){
		rule = validate[rule];
	}
	if ok, _ := regexp.MatchString(rule, value); !ok {
		return  false
	}
	return true;
}