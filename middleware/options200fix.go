package middleware

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

// Options200Fix makes sure that CORS OPTIONS pre-flight requests aren't
// served with 404 Not Found response. This patches the unsatisfactory
// functionality of the beego.plugins.cors
func Options200Fix() beego.FilterFunc {
	return func(ctx *context.Context) {
		if ctx.Input.Method() == "OPTIONS" {
			ctx.ResponseWriter.WriteHeader(204)
			ctx.WriteString("")
			return
		}
	}
}
