## beego AdminLTE-2.4.13

基于beego, jquery ,bootstrap, AdminLTE-2.4.13的一个后台管理系统

VERSION = "0.1.1"

## 获取安装

执行以下命令，就能够在你的`GOPATH/src` 目录下发现beego admin
```bash
$ go get github.com/beego
```

## 初次使用
需要使用mysql跟session，session基于Redis的所以需要安装redis
```bash
$ go get  github.com/astaxie/beego/session/redis
$ go get github.com/beego/beego/v2/server/web/session/redis
```
### 创建应用
首先,使用bee工具创建一个应用程序，参考[`http://beego.me/quickstart`](beego的入门)

```go
import (
	"github.com/astaxie/beego"  //beego 包
	"github.com/beego/admin"  //admin 包
)

```
引入代码，再`init`函数中使用它
```go
func init() {
	beego.Router("/", &admin.IndexController{})
}
```
### 配置文件

数据库目前仅支持MySQL后续会添加更多的数据库支持。

系统配置文件
```
appname = quickstart
httpport = 8080
runmode = dev
EnableAdmin = false
url = 192.168.33.20:8080
view = default
copyrequestbody = true
limit = 10
title = Go Houtai
```
数据库redis配置信息
```
# 设置session存储引擎
sessionon = true
sessionprovider = "redis"
sessionproviderconfig = "127.0.0.1:6379"

[db]
dbType = mysql
dbUser = root
dbPass = d75eb65d3524ddb7
dbHost = 127.0.0.1
dbPort = 3306
dbName = gohoutai

[redis]
rHost = 127.0.0.1
rPort = 6379
```
以上配置信息都需要加入到conf/app.conf文件中。

### 编译项目

全部做好了以后。就可以编译了,进入hello目录
```
$ bee go
```
好了，现在可以通过浏览器地址访问了[`http://localhost:8080/`](http://localhost:8080/)

默认得用户名密码是
admin 123456

