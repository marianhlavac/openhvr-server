package models

import (
	"testing"

	"github.com/mmajko/openhvr-server/helpers"
	. "github.com/smartystreets/goconvey/convey"
)

func TestEffects(t *testing.T) {
	Convey("Effects Model Test", t, func() {
		manager := NewEffectManager()
		effectTypeNone := &EffectType{Id: 0, Name: "None"}
		pos123 := helpers.NewVector3(1, 2, 3)
		reqFuture := NewEffectRequest(
			effectTypeNone, 10, pos123, 10,
		)
		reqHistoric := NewEffectRequest(
			effectTypeNone, -10, pos123, 10,
		)

		Convey("Running effect correctly manages", func() {
			manager.Clear()
			err := manager.RunEffect(reqFuture)
			So(err, ShouldBeNil)
			So(manager.All(), ShouldHaveLength, 1)
			So(manager.All()[0].Request, ShouldEqual, reqFuture)
		})

		Convey("All() returns correctly all requests", func() {
			manager.Clear()
			err := manager.RunEffect(reqFuture)
			So(err, ShouldBeNil)
			err = manager.RunEffect(reqHistoric)
			So(err, ShouldBeNil)

			So(manager.All(), ShouldHaveLength, 2)
			So(manager.All()[0].Request, ShouldEqual, reqFuture)
			So(manager.All()[1].Request, ShouldEqual, reqHistoric)
		})

		Convey("AllTimeouted() returns correctly only timeouted", func() {
			manager.Clear()
			err := manager.RunEffect(reqFuture)
			So(err, ShouldBeNil)
			err = manager.RunEffect(reqHistoric)
			So(err, ShouldBeNil)

			So(manager.AllTimeouted(), ShouldHaveLength, 1)
			So(manager.AllTimeouted()[0].Request, ShouldEqual, reqHistoric)
		})
	})
}
