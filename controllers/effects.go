package controllers

import (
	"encoding/json"
	"log"

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

	for _, device := range GetDevicesInRequestRange(o, effectRequest) {
		go RequestOnDevice(device, &effectRequest)
	}

	c.Data["json"] = &models.ActionResults{Result: "requested"}
	c.ServeJSON()
}

// @Title Immediately Cancel All
// @Description Immediately cancel all effects
// @Success 200 {object} models.ActionResult
// @router / [delete]
func (c *EffectsController) DeleteAll() {
	o := orm.NewOrm()
	var devices []*models.Device

	o.QueryTable("device").All(&devices)

	for _, device := range devices {
		go CancelOnDevice(device)
	}

	c.Data["json"] = &models.ActionResults{Result: "cancelled all"}
	c.ServeJSON()
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
// @Success 200 {object} []models.EffectType
// @router /types [get]
func (c *EffectsController) GetTypes() {
	o := orm.NewOrm()
	var types []models.EffectType
	_, err := o.QueryTable("effect_type").All(&types)

	if err == nil {
		c.Data["json"] = types
		c.ServeJSON()
	} else {
		log.Fatal(err)
		c.Abort("500")
	}
}
