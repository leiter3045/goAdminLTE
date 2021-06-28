package main

import (
	_ "quickstart/routers"
	db "quickstart/service/databsae"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	_ "github.com/astaxie/beego/session/redis"
	_ "github.com/beego/beego/v2/server/web/session/redis"
)

func init() {
	conf, err := config.NewConfig("ini", "conf/app.conf")
	if err != nil {
		logrus.Fatalf(err.Error())
	}
	database, _ := db.NewDataBase(conf.String("db::dbType"))
	orm.RegisterDriver(database.GetDriverName(), database.GetDriver())
	orm.RegisterDataBase(database.GetAliasName(), database.GetDriverName(), database.GetStr())
}

func main() {
	//orm.Debug = true
	beego.Run()
}

