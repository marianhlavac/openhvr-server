package controllers

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/mmajko/openhvr-server/models"
)

const DEFAULT_ROOM_ID = 1

type RoomController struct {
	beego.Controller
}

// @Title Get Room Properties
// @Description Return the current room configuration
// @Success 200 {object} models.Room
// @router / [get]
func (c *RoomController) Get() {
	o := orm.NewOrm()
	room := models.Room{Id: DEFAULT_ROOM_ID}
	err := o.Read(&room)

	if err == orm.ErrNoRows {
		c.Abort("404")
	} else {
		c.Data["json"] = room
		c.ServeJSON()
	}
}

// @Title Update Room Properties
// @Description Update the current room configuration
// @Success 200 {object} models.Room
// @Param body body models.Room true
// @router / [put]
func (c *RoomController) Put() {
	var room models.Room
	json.Unmarshal(c.Ctx.Input.RequestBody, &room)
	room.Id = DEFAULT_ROOM_ID

	o := orm.NewOrm()
	num, err := o.Update(&room)
	if err == nil && num == 0 {
		num, err = o.Insert(&room)
	}

	if err == nil {
		c.Data["json"] = &room
		c.ServeJSON()
	} else {
		log.Fatal(err)
		c.Abort("500")
	}
}

// @Title Get Room Devices
// @Description Return a list of devices registered to the room
// @Success 200 {object} []models.RoomDevice
// @router /devices/ [get]
func (c *RoomController) DevicesGet() {
	o := orm.NewOrm()
	room := models.Room{Id: DEFAULT_ROOM_ID}
	err := o.Read(&room)

	if err == orm.ErrNoRows {
		c.Abort("404")
	} else if err == nil {
		c.Data["json"] = room.Devices
		c.ServeJSON()
	} else {
		log.Fatal(err)
		c.Abort("500")
	}
}

// @Title Register Room Device
// @Description Register a new room device
// @Success 200 {object} models.RoomDevice
// @Param body body models.RoomDevice true
// @router /devices/ [post]
func (c *RoomController) DevicesPost() {
	var device models.RoomDevice
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
// @Success 200 {object} models.RoomDevice
// @Param body body models.RoomDevice true
// @router /devices/:deviceId [put]
func (c *RoomController) DevicesPut() {
	var device models.RoomDevice
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
// @Success 200 {string}
// @router /devices/:deviceId [delete]
func (c *RoomController) DevicesDelete() {
	o := orm.NewOrm()
	var device_id, _ = strconv.Atoi(c.Ctx.Input.Param(":deviceId"))
	var device = models.RoomDevice{Id: device_id}
	num, err := o.Delete(&device)

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
