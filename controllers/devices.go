package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/mmajko/openhvr-server/models"
	"github.com/pkg/errors"
)

type DevicesController struct {
	beego.Controller
}

const timeoutedDevicesUpdateInterval = 2 * time.Second

var CommunicationNonOk = errors.New("CommunicationNonOk")
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
	json.Unmarshal(c.Ctx.Input.RequestBody, &device)
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

func SendDeviceCommand(device *models.Device, command string) error {
	resp, err := http.Get("http://" + device.ConnectorUri + "/cm?cmnd=" + command)
	if err == nil {
		if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
			return nil
		} else {
			return CommunicationNonOk
		}
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
		_, err := http.Get("http://" + device.ConnectorUri + "/cm?cmnd=Power%20Off")
		if err == nil {
			device.TimeoutAt = 0
			timeoutedDevicesCache = removeTimeoutedDevice(timeoutedDevicesCache, device.Id)
			o.Update(&device)
		} else {
			beego.Error("Failed to disable device", device.Id, "at", device.ConnectorUri, "via", device.ConnectorType)
			beego.Error("   |- ", err)
		}
	}

	return len(timeoutedDevicesCache)
}
