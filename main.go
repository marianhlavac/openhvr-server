package main

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"
	_ "github.com/mattn/go-sqlite3"
	"github.com/mmajko/openhvr-server/controllers"
	"github.com/mmajko/openhvr-server/devicedrivers"
	"github.com/mmajko/openhvr-server/models"
	_ "github.com/mmajko/openhvr-server/routers"
)

const workerUpdateInterval = 100

var isReady = false

func init() {
	// Register models
	orm.RegisterModel(new(models.Device), new(models.EffectType))

	// Register device drivers
	devicedrivers.RegisterDriver("tasmota_fan", devicedrivers.TasmotaFanDriver)
	devicedrivers.RegisterDriver("tasmota_heater", devicedrivers.TasmotaHeaterDriver)
	devicedrivers.RegisterDriver("dummy_fan", devicedrivers.DummyFanDriver)

	// Register default effect types
	models.RegisterDefaultEffectType(1, "Generic")
	models.RegisterDefaultEffectType(2, "Wind")
	models.RegisterDefaultEffectType(3, "Heat")

	// Prepare to run
	prepareDatabase()
	orm.Debug = false
	isReady = true
}

func main() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))

	beego.InsertFilter("*", beego.BeforeRouter, options200Fix())

	beego.BConfig.WebConfig.StaticDir["/docs"] = "swagger"
	beego.BConfig.WebConfig.StaticDir["/"] = "configurator-app/public"
	go worker()
	beego.Run()
}

// prepareDatabase makes database ready to run
func prepareDatabase() {
	// Register database
	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "data.db")

	// Migrate schema
	if err := orm.RunSyncdb("default", false, false); err != nil {
		fmt.Println(err)
	}

	// Apply registrations
	models.ApplyDefaultEffectRegistration()
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

// Options 200 OK workaround hack (cors plugin sucks!)
func options200Fix() beego.FilterFunc {
	return func(ctx *context.Context) {
		if ctx.Input.Method() == "OPTIONS" {
			ctx.ResponseWriter.WriteHeader(204)
			ctx.WriteString("")
			return
		}
	}
}
