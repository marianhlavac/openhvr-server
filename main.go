package main

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/mmajko/openhvr-server/controllers"
	"github.com/mmajko/openhvr-server/models"
	_ "github.com/mmajko/openhvr-server/routers"
)

const workerUpdateInterval = 100

var isReady = false

func init() {
	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "data.db")
	orm.RegisterModel(
		new(models.Device),
	)
	orm.RunCommand()

	force := false
	err := orm.RunSyncdb("default", force, true)
	if err != nil {
		fmt.Println(err)
	}

	orm.Debug = true
	isReady = true
}

func worker() {
	for {
		if isReady {
			disabled := controllers.DisableAllTimeoutedDevices()
			if disabled > 0 {
				beego.Info("Disabled", disabled, "timeouted devices")
			}
		}
		time.Sleep(time.Millisecond * workerUpdateInterval)
	}
}

func main() {
	beego.BConfig.WebConfig.StaticDir["/docs"] = "swagger"
	beego.BConfig.WebConfig.StaticDir["/"] = "static"
	go worker()
	beego.Run()
}
