package controllers

import (
	"encoding/json"
	"log"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/mmajko/openhvr-server/models"
)

type EffectsController struct {
	beego.Controller
}

// Returns all devices in range requested by the effect request
func getDevicesInRequestRange(o orm.Ormer, req models.EffectRequest) []*models.Device {
	var devices []*models.Device
	o.QueryTable("device").
		Filter("location_x__gte", req.Position.X-req.Range).
		Filter("location_x__lte", req.Position.X+req.Range).
		Filter("location_y__lte", req.Position.Y+req.Range).
		Filter("location_y__gte", req.Position.Y-req.Range).
		Filter("location_z__lte", req.Position.Z+req.Range).
		Filter("location_z__gte", req.Position.Z-req.Range).
		All(&devices)
	return devices
}

// @Title Request Effect Performance
// @Description Requests effect performance
// @Param body body models.EffectRequest true
// @Success 200 {object} models.ActionResult
// @router / [post]
func (c *EffectsController) Post() {
	o := orm.NewOrm()

	var effectRequest models.EffectRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &effectRequest)
	var timeoutsAt = time.Now().Unix() + int64(effectRequest.Duration)

	var allErr error = nil

	for _, device := range getDevicesInRequestRange(o, effectRequest) {
		var err = SendDeviceCommand(device, "Power%20On")
		if err == nil {
			device.TimeoutAt = timeoutsAt
			o.Update(device)
		} else {
			beego.Error("Failed to enable device", device.Id, "at", device.ConnectorUri, "via", device.ConnectorType)
			beego.Error("   |- ", err)
			allErr = err
		}
	}

	if allErr == nil {
		c.Data["json"] = &models.ActionResults{Result: "requested"}
		c.ServeJSON()
	} else {
		c.Abort("500")
	}
}

// @Title Immediately Cancel All
// @Description Immediately cancel all effects
// @Success 200 {object} models.ActionResult
// @router / [delete]
func (c *EffectsController) DeleteAll() {
	o := orm.NewOrm()
	var devices []*models.Device

	o.QueryTable("device").All(&devices)
	var allErr error = nil

	for _, device := range devices {
		var err = SendDeviceCommand(device, "Power%20Off")
		if err == nil {
			device.TimeoutAt = 0
			o.Update(device)
		} else {
			beego.Error("Failed to disable device", device.Id, "at", device.ConnectorUri, "via", device.ConnectorType)
			beego.Error("   |- ", err)
			allErr = err
		}
	}

	if allErr == nil {
		c.Data["json"] = &models.ActionResults{Result: "cancelled all"}
		c.ServeJSON()
	} else {
		c.Abort("500")
	}
}

// @Title Immediately Cancel Specific
// @Description Immediately cancel effect
// @Param	effectId		path 	string	true		"the objectid you want to get"
// @Success 200 {object} models.ActionResult
// @router /:effectId [delete]
func (c *EffectsController) Delete() {
	c.Abort("501")
}

// @Title Get Effect Types
// @Description Return the list of available effect types
// @Success 200 {string}
// @router /types [get]
func (c *EffectsController) GetTypes() {
	o := orm.NewOrm()
	var types orm.ParamsList
	_, err := o.QueryTable("device").Distinct().ValuesFlat(&types, "effect_type")

	if err == nil {
		c.Data["json"] = types
		c.ServeJSON()
	} else {
		log.Fatal(err)
		c.Abort("500")
	}
}
