// @APIVersion 1.0.0
// @Title OpenHVR Server
// @Description OpenHVR Server is a TBD
// @Contact m@marianhlavac.cz
package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/mmajko/openhvr-server/controllers"
)

func init() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))

	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/devices",
			beego.NSInclude(&controllers.DevicesController{}),
		),
		beego.NSNamespace("/effects",
			beego.NSInclude(&controllers.EffectsController{}),
		),
		beego.NSNamespace("/system",
			beego.NSInclude(&controllers.SystemController{}),
		),
	)
	beego.AddNamespace(ns)
}
