package devicedrivers

import "github.com/mmajko/openhvr-server/models"

// TasmotaHeaterDriver is driver for heaters connected using Tasmota-enabled relay
func TasmotaHeaterDriver(device *models.Device, request *models.EffectRequest) error {
	relayName := ReadRelayNameFromParamDefault(device.ConnectorParam)
	if request != nil {
		return TasmotaHTTPSendCommand(device.ConnectorUri, relayName+"%201")
	}
	return TasmotaHTTPSendCommand(device.ConnectorUri, relayName+"%200")
}
