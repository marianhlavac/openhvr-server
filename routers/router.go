// @APIVersion 1.0.0
// @Title OpenHVR Server
// @Description OpenHVR Server is one of the OpenHVR components to provide real-world haptic feedback effects for VR applications
// @Contact m@marianhlavac.cz
package routers

import (
	"github.com/astaxie/beego"
	"github.com/mmajko/openhvr-server/controllers"
)

func init() {
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
