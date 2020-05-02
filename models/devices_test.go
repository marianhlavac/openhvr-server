package models

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDevices(t *testing.T) {
	Convey("Devices Model Test", t, func() {
		d := Device{
			Id: 0, Name: "", Type: "", EffectType: 0,
			LocationX: 1, LocationY: 2, LocationZ: 3,
			DirectionX: 0.4, DirectionY: 0.5, DirectionZ: 0.6,
			DirectionSpread: 0, ConnectorUri: "", ConnectorParam: "",
			TimeoutAt: 0,
		}

		Convey("GetLocationVector returns correct location", func() {
			lv := d.GetLocationVector()
			So(lv.X, ShouldEqual, d.LocationX)
			So(lv.Y, ShouldEqual, d.LocationY)
			So(lv.Z, ShouldEqual, d.LocationZ)
		})

		Convey("GetLocationVector returns correct direction", func() {
			dv := d.GetDirectionVector()
			So(dv.X, ShouldEqual, d.DirectionX)
			So(dv.Y, ShouldEqual, d.DirectionY)
			So(dv.Z, ShouldEqual, d.DirectionZ)
		})
	})
}
