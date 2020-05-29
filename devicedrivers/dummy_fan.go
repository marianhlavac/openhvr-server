package devicedrivers

import (
	"github.com/astaxie/beego"
	"github.com/mmajko/openhvr-server/models"
)

// DummyFanDriver is driver for testing a simulated fan behaviour using the
// console terminal. It will output to console if the virtual fan starts
// spinning of stops.
func DummyFanDriver(device *models.Device, request *models.EffectRequest) error {
	if request != nil {
		// Filter out device if it's out of the direction spread range
		if request.Directional && !IsRotatedWithinSpread(device, *request.Direction) {
			return DeviceFilteredOut
		}
		beego.Info("Spinning up dummy fan " + string(device.Id))
		return nil
	}
	beego.Info("Spinning down dummy fan " + string(device.Id))
	return nil
}
