package controllers

import (
	"encoding/json"
	"log"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/mmajko/openhvr-server/devicedrivers"
	"github.com/mmajko/openhvr-server/models"
)

type DevicesController struct {
	beego.Controller
}

const timeoutedDevicesUpdateInterval = 2 * time.Second
const testEffectDuration = 2

var timeoutedDevicesCache []models.Device = nil
var timeoutedDevicesLastUpdate time.Time

// @Title Get Room Devices
// @Description Return a list of devices registered to the room
// @Success 200 {object} []models.Device
// @router / [get]
func (c *DevicesController) GetAll() {
	o := orm.NewOrm()
	var devices []models.Device
	_, err := o.QueryTable("device").All(&devices)

	if err == nil {
		c.Data["json"] = devices
		c.ServeJSON()
	} else {
		log.Fatal(err)
		c.Abort("500")
	}
}

// @Title Get Room Device
// @Description Return a single device
// @Param	deviceId		path 	string	true		"the objectid you want to get"
// @Success 200 {object} []models.Device
// @router /:deviceId [get]
func (c *DevicesController) Get() {
	var device models.Device
	device.Id, _ = strconv.Atoi(c.Ctx.Input.Param(":deviceId"))

	o := orm.NewOrm()
	err := o.Read(&device)

	if err == nil {
		c.Data["json"] = device
		c.ServeJSON()
	} else {
		log.Fatal(err)
		c.Abort("500")
	}
}

// @Title Register Room Device
// @Description Register a new room device
// @Success 200 {object} models.Device
// @Param body body models.Device true
// @router / [post]
func (c *DevicesController) Post() {
	var device models.Device
	json.Unmarshal(c.Ctx.Input.RequestBody, &device)

	if isDirectionInvalid(device.DirectionX, device.DirectionY, device.DirectionZ) {
		c.CustomAbort(400, "Direction vector is invalid!")
		return
	}

	o := orm.NewOrm()
	_, err := o.Insert(&device)

	if err == nil {
		c.Data["json"] = &device
		c.ServeJSON()
	} else {
		log.Fatal(err)
		c.Abort("500")
	}
}

// @Title Update Room Device
// @Description Update a room device
// @Param	deviceId		path 	string	true		"the objectid you want to get"
// @Success 200 {object} models.Device
// @Param body body models.Device true
// @router /:deviceId [put]
func (c *DevicesController) Put() {
	var device models.Device
	json.Unmarshal(c.Ctx.Input.RequestBody, &device)
	device.Id, _ = strconv.Atoi(c.Ctx.Input.Param(":deviceId"))

	if isDirectionInvalid(device.DirectionX, device.DirectionY, device.DirectionZ) {
		c.CustomAbort(400, "Direction vector is invalid!")
		return
	}

	o := orm.NewOrm()
	num, err := o.Update(&device)

	if err == nil {
		if num == 0 {
			c.Abort("404")
		}

		c.Data["json"] = &device
		c.ServeJSON()
	} else {
		log.Fatal(err)
		c.Abort("500")
	}
}

// @Title Delete Room Device
// @Description Delete a room device
// @Param	deviceId		path 	string	true		"the objectid you want to get"
// @Success 200 {object} models.ActionResult
// @router /:deviceId [delete]
func (c *DevicesController) Delete() {
	o := orm.NewOrm()
	var device_id, _ = strconv.Atoi(c.Ctx.Input.Param(":deviceId"))
	var device = models.Device{Id: device_id}
	num, err := o.Delete(&device)

	if err == nil {
		if num == 0 {
			c.Abort("404")
		}

		c.Data["json"] = &models.ActionResults{Result: "deleted"}
		c.ServeJSON()
	} else {
		log.Fatal(err)
		c.Abort("500")
	}
}

// @Title Test Room Device
// @Description Test a room device
// @Param	deviceId		path 	string	true		"the objectid you want to get"
// @Success 200 {string}
// @router /:deviceId/test [post]
func (c *DevicesController) PostTest() {
	var device models.Device
	device.Id, _ = strconv.Atoi(c.Ctx.Input.Param(":deviceId"))

	o := orm.NewOrm()
	err := o.Read(&device)
	if err == nil {
		go RequestOnDevice(&device, models.NewUnconstrainedRequest(testEffectDuration))

		c.Data["json"] = &models.ActionResults{Result: "test sent"}
		c.ServeJSON()
	} else {
		c.Abort("500")
	}
}

// @Title Get Available Drivers
// @Description Return a list of names of currently supported drivers by the server
// @Success 200 {object} []string
// @router /drivers [get]
func (c *DevicesController) GetDrivers() {
	drivers := devicedrivers.GetAvailableDrivers()

	c.Data["json"] = drivers
	c.ServeJSON()
}

// RequestOnDevice processes an effect request on a selected device. It
// updates the timeouts and calls for correct driver to further handle
// the request.
func RequestOnDevice(device *models.Device, effectRequest *models.EffectRequest) error {
	o := orm.NewOrm()
	var timeoutsAt = time.Now().Unix() + int64(effectRequest.Duration)
	var err = devicedrivers.HandleRequestWithDriver(device, effectRequest)
	if err == nil {
		device.TimeoutAt = timeoutsAt
		o.Update(device)
	} else if err != devicedrivers.DeviceWrongDirection {
		beego.Error("Failed to request effect on device ", device.Id, "at", device.ConnectorUri)
		beego.Error("   |- ", err)
	}
	return err
}

// CancelOnDevice cancels all currently produced effects on a selected device.
func CancelOnDevice(device *models.Device) error {
	o := orm.NewOrm()
	var err = devicedrivers.HandleRequestWithDriver(device, nil)
	if err == nil {
		device.TimeoutAt = 0
		o.Update(device)
	} else if err != devicedrivers.DeviceWrongDirection {
		beego.Error("Failed to cancel effect on device ", device.Id, "at", device.ConnectorUri)
		beego.Error("   |- ", err)
	}
	return err
}

func UpdateTimeoutedDevicesCache(o orm.Ormer) {
	now := time.Now()
	o.QueryTable("device").Filter("timeout_at__lte", now.Unix()).Filter("timeout_at__gt", 0).All(&timeoutedDevicesCache)
	timeoutedDevicesLastUpdate = now
}

func removeTimeoutedDevice(s []models.Device, id int) []models.Device {
	n := 0
	for _, x := range s {
		if x.Id != id {
			s[n] = x
			n++
		}
	}
	return s[:n]
}

func DisableAllTimeoutedDevices() int {
	o := orm.NewOrm()
	if timeoutedDevicesCache == nil || timeoutedDevicesLastUpdate.Add(timeoutedDevicesUpdateInterval).Before(time.Now()) {
		UpdateTimeoutedDevicesCache(o)
	}

	for _, device := range timeoutedDevicesCache {
		err := CancelOnDevice(&device)
		if err == nil {
			timeoutedDevicesCache = removeTimeoutedDevice(timeoutedDevicesCache, device.Id)
		}
	}

	return len(timeoutedDevicesCache)
}

func isDirectionInvalid(x, y, z float32) bool {
	return x == 0 && y == 0 && z == 0
}
