package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/mmajko/openhvr-server/models"
	_ "github.com/mmajko/openhvr-server/routers"
)

func init() {
	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "data.db")
	orm.RegisterModel(
		new(models.Room),
		new(models.RoomDevice),
	)
	orm.RunCommand()
	orm.Debug = true
}

func main() {
	beego.BConfig.WebConfig.StaticDir["/docs"] = "swagger"
	beego.Run()
}
