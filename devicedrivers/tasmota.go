package devicedrivers

import (
	"net/http"
	"strings"
	"time"

	"github.com/astaxie/beego"
)

var client = http.Client{Timeout: time.Second * time.Duration(beego.AppConfig.DefaultInt("http_timeout", 2))}

// TasmotaHTTPSendCommand sends a command using HTTP API
func TasmotaHTTPSendCommand(uri, command string) error {
	resp, err := client.Get("http://" + uri + "/cm?cmnd=" + command)
	if err == nil {
		if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
			return nil
		}
		return CommunicationNonOk
	}
	return err
}

func ReadRelayNameFromParamDefault(param string) string {
	if strings.HasPrefix(param, "Power") {
		return param
	}
	return "Power"
}
