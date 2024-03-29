package devicedrivers

import (
	"math"

	"github.com/mmajko/openhvr-server/helpers"
	"github.com/mmajko/openhvr-server/models"
)

// IsRotatedWithinSpread checks if requested direction is aligned with the
// configured direction of the device.
func IsRotatedWithinSpread(device *models.Device, requestDirection helpers.Vector3) bool {
	deviceRot := device.GetDirectionVector()
	dot := requestDirection.Dot(deviceRot)
	angle := math.Acos(dot / requestDirection.Magnitude() * deviceRot.Magnitude())
	return angle <= float64(device.DirectionSpread*(math.Pi/180)/2)
}

// TasmotaFanDriver is driver for fans connected using Tasmota-enabled relay
func TasmotaFanDriver(device *models.Device, request *models.EffectRequest) error {
	relayName := ReadRelayNameFromParamDefault(device.ConnectorParam)
	if request != nil {
		// Filter out device if it's out of the direction spread range
		if request.Directional && !IsRotatedWithinSpread(device, *request.Direction) {
			return DeviceFilteredOut
		}
		// Match the effect type
		if !MatchEffectType(device, request) {
			return DeviceFilteredOut
		}
		return TasmotaHTTPSendCommand(device.ConnectorUri, relayName+"%201")
	}
	return TasmotaHTTPSendCommand(device.ConnectorUri, relayName+"%200")
}
