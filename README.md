## beego AdminLTE-2.4.13

基于beego, jquery ,bootstrap, AdminLTE-2.4.13的一个后台管理系统。
易用、易扩展、界面友好的轻量级功能权限管理系统。前端框架基于AdminLTE2进行资源整合，包含了多款优秀的插件，是笔者对多年后台管理系统开发经验精华的萃取。 本系统非常适合进行后台管理系统开发，统一的代码和交互给二次开发带来极大的方便，在没有前端工程师参与的情况下就可以进行快速的模块式开发，并保证用户使用的友好性和易用性。

## 特点
```
1.分页列表页面的搜索条件、搜索面板、Page、当前页数、显示/隐藏列在变化时自动保存，页面刷新后、重新进入时，这些状态依然保持；
2.TreeTabe列表节点展开/收缩状态、滚动条位置时自动保存，页面刷新后、重新进入时，这些状态依然保持；
3.编辑分页列表、TreeTabe列表中数据后，当前数据行背景闪烁，如果当前数据行由于顺序变化跳出可视区域，则滚动条自动滚动，将当前数据行移动至可视区域；
4.精确至Action的轻量级功能权限控制，后台用户与角色、角色与资源（菜单、按钮）都是多对多关系，可以灵活配置用户可访问的资源。
```
## 分页
作者根据PHP，thinkPHP3.2分页类，采用go语言重写，方法调用与使用都跟thinkPHP3.2分页是一样的
## 众多的可用函数
作者封装了大量的数据处理函数，很多PHP里面字符串处理函数，采用go语言重写，并且集成过来
## 获取安装

执行以下命令，就能够在你的`GOPATH/src` 目录下发现beego
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
)
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

### 数据库文档
数据库暂时只有mysql链接
库表在根目录sql里面

### 编译项目

全部做好了以后。就可以启动
```
$ bee go
```
好了，现在可以通过浏览器地址访问了[`http://localhost:8080/`](http://localhost:8080/)

默认得用户名密码是
admin 123456

