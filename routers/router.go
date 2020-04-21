// @APIVersion 1.0.0
// @Title OpenHVR Server
// @Description OpenHVR Server is a TBD
// @Contact m@marianhlavac.cz
package routers

import (
	"github.com/mmajko/openhvr-server/controllers"
	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/room",
			beego.NSInclude(&controllers.RoomController{}),
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
