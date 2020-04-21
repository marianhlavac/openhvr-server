package controllers

import (
	"log"
	"net/http"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/mmajko/openhvr-server/models"
)

type EffectsController struct {
	beego.Controller
}

// @Title Request Effect Performance
// @Description Requests effect performance
// @Success 200 {object} models.ActionResult
// @router / [post]
func (c *EffectsController) Post() {
	http.Get("http://192.168.0.66/cm?cmnd=Power%20Toggle")
	c.Data["json"] = &models.ActionResults{Result: "requested"}
	c.ServeJSON()
}

// @Title Immediately Cancel All
// @Description Immediately cancel all effects
// @Success 200 {object} models.ActionResult
// @router / [delete]
func (c *EffectsController) DeleteAll() {
	http.Get("http://192.168.0.66/cm?cmnd=Power%20Off")
	c.Data["json"] = &models.ActionResults{Result: "cancelled"}
	c.ServeJSON()
}

// @Title Immediately Cancel Specific
// @Description Immediately cancel effect
// @Param	effectId		path 	string	true		"the objectid you want to get"
// @Success 200 {object} models.ActionResult
// @router /:effectId [delete]
func (c *EffectsController) Delete() {
	http.Get("http://192.168.0.66/cm?cmnd=Power%20Toggle")
	c.Data["json"] = &models.ActionResults{Result: "cancelled"}
	c.ServeJSON()
}

// @Title Get Effect Types
// @Description Return the list of available effect types
// @Success 200 {string}
// @router /types [get]
func (c *EffectsController) GetTypes() {
	o := orm.NewOrm()
	var types orm.ParamsList
	_, err := o.QueryTable("room_device").Distinct().ValuesFlat(&types, "effect_type")

	if err == nil {
		c.Data["json"] = types
		c.ServeJSON()
	} else {
		log.Fatal(err)
		c.Abort("500")
	}
}
