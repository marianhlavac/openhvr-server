package models

import "github.com/mmajko/openhvr-server/helpers"

type Device struct {
	Id              int `orm:"auto"`
	Name            string
	Type            string
	EffectType      int
	LocationX       float32
	LocationY       float32
	LocationZ       float32
	DirectionX      float32
	DirectionY      float32
	DirectionZ      float32
	DirectionSpread float32
	ConnectorUri    string
	ConnectorParam  string
	TimeoutAt       int64 `orm:"null"`
}

func (d *Device) GetLocationVector() *helpers.Vector3 {
	return helpers.NewVector3(d.LocationX, d.LocationY, d.LocationZ)
}

func (d *Device) GetDirectionVector() *helpers.Vector3 {
	return helpers.NewVector3(d.DirectionX, d.DirectionY, d.DirectionZ)
}
