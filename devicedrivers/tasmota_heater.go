package devicedrivers

import "github.com/mmajko/openhvr-server/models"

// TasmotaHeaterDriver is driver for heaters connected using Tasmota-enabled relay
func TasmotaHeaterDriver(device *models.Device, request *models.EffectRequest) error {
	relayName := ReadRelayNameFromParamDefault(device.ConnectorParam)
	if request != nil {
		// Match the effect type
		if !MatchEffectType(device, request) {
			return DeviceFilteredOut
		}
		return TasmotaHTTPSendCommand(device.ConnectorUri, relayName+"%201")
	}
	return TasmotaHTTPSendCommand(device.ConnectorUri, relayName+"%200")
}
