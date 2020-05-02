package devicedrivers

import (
	"github.com/astaxie/beego"
	"github.com/mmajko/openhvr-server/models"
)

// DummyFanDriver is driver for fans connected using Tasmota-enabled relay
func DummyFanDriver(device *models.Device, request *models.EffectRequest) error {
	if request != nil {
		// Filter out device if it's out of the direction spread range
		if request.Directional && !isRotatedWithinSpread(device, *request.Direction) {
			return DeviceWrongDirection
		}
		beego.Info("Spinning up dummy fan " + string(device.Id))
		return nil
	}
	beego.Info("Spinning down dummy fan " + string(device.Id))
	return nil
}
