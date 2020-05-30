package middleware

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

// AllowWhitelistedOnly will allow only clients with IP address included in the
// IP whitelist to perform requests to the server.
func AllowWhitelistedOnly() beego.FilterFunc {
	return func(ctx *context.Context) {
		var whitelistStr = beego.AppConfig.Strings("ip_whitelist")
		var clientIP = ctx.Input.IP()

		for _, ip := range whitelistStr {
			if ip == "*" || ip == clientIP {
				return
			}
		}
		beego.Warning("Blocked request, as this IP is not whitelisted: ", clientIP)
		ctx.ResponseWriter.WriteHeader(403)
		ctx.WriteString(fmt.Sprintf("IP addres %s is not on the whitelist", clientIP))
		return
	}
}
