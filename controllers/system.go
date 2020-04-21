package controllers

import (
	"github.com/astaxie/beego"
)

type SystemController struct {
	beego.Controller
}

// @Title Get Status
// @Description Return the current system status
// @Success 200 {string}
// @router /status [get]
func (c *SystemController) GetStatus() {
	c.Data["json"] = map[string]string { "result": "done" }
	c.ServeJSON()
}
