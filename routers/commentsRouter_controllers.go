package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/mmajko/openhvr-server/controllers:DevicesController"] = append(beego.GlobalControllerRouter["github.com/mmajko/openhvr-server/controllers:DevicesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/mmajko/openhvr-server/controllers:DevicesController"] = append(beego.GlobalControllerRouter["github.com/mmajko/openhvr-server/controllers:DevicesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/mmajko/openhvr-server/controllers:DevicesController"] = append(beego.GlobalControllerRouter["github.com/mmajko/openhvr-server/controllers:DevicesController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:deviceId`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/mmajko/openhvr-server/controllers:DevicesController"] = append(beego.GlobalControllerRouter["github.com/mmajko/openhvr-server/controllers:DevicesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:deviceId`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/mmajko/openhvr-server/controllers:DevicesController"] = append(beego.GlobalControllerRouter["github.com/mmajko/openhvr-server/controllers:DevicesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:deviceId`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/mmajko/openhvr-server/controllers:DevicesController"] = append(beego.GlobalControllerRouter["github.com/mmajko/openhvr-server/controllers:DevicesController"],
        beego.ControllerComments{
            Method: "PostTest",
            Router: `/:deviceId/test`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/mmajko/openhvr-server/controllers:DevicesController"] = append(beego.GlobalControllerRouter["github.com/mmajko/openhvr-server/controllers:DevicesController"],
        beego.ControllerComments{
            Method: "GetDrivers",
            Router: `/drivers`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/mmajko/openhvr-server/controllers:EffectsController"] = append(beego.GlobalControllerRouter["github.com/mmajko/openhvr-server/controllers:EffectsController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/mmajko/openhvr-server/controllers:EffectsController"] = append(beego.GlobalControllerRouter["github.com/mmajko/openhvr-server/controllers:EffectsController"],
        beego.ControllerComments{
            Method: "DeleteAll",
            Router: `/`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/mmajko/openhvr-server/controllers:EffectsController"] = append(beego.GlobalControllerRouter["github.com/mmajko/openhvr-server/controllers:EffectsController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:effectId`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/mmajko/openhvr-server/controllers:EffectsController"] = append(beego.GlobalControllerRouter["github.com/mmajko/openhvr-server/controllers:EffectsController"],
        beego.ControllerComments{
            Method: "GetTypes",
            Router: `/types`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/mmajko/openhvr-server/controllers:SystemController"] = append(beego.GlobalControllerRouter["github.com/mmajko/openhvr-server/controllers:SystemController"],
        beego.ControllerComments{
            Method: "GetStatus",
            Router: `/status`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
