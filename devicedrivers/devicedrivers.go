package devicedrivers

import (
	"errors"

	"github.com/mmajko/openhvr-server/models"
)

type deviceEffectRequest = func(*models.Device, *models.EffectRequest) error

var registeredDrivers = map[string]deviceEffectRequest{}
var DriverNotExistError = errors.New("Device driver does not exist.")
var DeviceWrongDirection = errors.New("Device is just in wrong direction")
var CommunicationNonOk = errors.New("CommunicationNonOk")

// RegisterDriver registers a new driver under specified name
func RegisterDriver(deviceTypeName string, driver deviceEffectRequest) {
	registeredDrivers[deviceTypeName] = driver
}

// HandleRequestWithDriver finds appropriate driver for the device and handles
// the request with it
func HandleRequestWithDriver(device *models.Device, request *models.EffectRequest) error {
	driver, found := registeredDrivers[device.Type]

	if !found {
		return DriverNotExistError
	}

	return driver(device, request)
}

// GetAvailableDrivers Returns names of currently available drivers
func GetAvailableDrivers() []string {
	driversNames := make([]string, 0, len(registeredDrivers))
	for k := range registeredDrivers {
		driversNames = append(driversNames, k)
	}
	return driversNames
}
