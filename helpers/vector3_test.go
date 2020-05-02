package helpers

import (
	"math"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDevices(t *testing.T) {
	Convey("Vector3 Dot()", t, func() {
		v1 := NewVector3(10.5, 21.75, 33)
		Convey("Dot product behaves correctly", func() {
			So(v1.Dot(NewVector3(0, 0, 0)), ShouldEqual, 0)

			So(v1.Dot(NewVector3(1, 0, 0)), ShouldEqual, 10.5)
			So(v1.Dot(NewVector3(0, 1, 0)), ShouldEqual, 21.75)
			So(v1.Dot(NewVector3(0, 0, 1)), ShouldEqual, 33)
			So(v1.Dot(NewVector3(1, 1, 1)), ShouldEqual, 65.25)

			So(v1.Dot(NewVector3(2, 0, 0)), ShouldEqual, 21)
		})

		Convey("Dot product is comutative", func() {
			So(NewVector3(0, 0, 0).Dot(v1), ShouldEqual, 0)
			So(NewVector3(2, 0, 0).Dot(v1), ShouldEqual, 21)
		})
	})

	Convey("Vector3 Magnitude()", t, func() {
		Convey("Magnitude behaves correctly", func() {
			So(NewVector3(0, 0, 0).Magnitude(), ShouldEqual, 0)
			So(NewVector3(1, 0, 0).Magnitude(), ShouldEqual, 1)
			So(NewVector3(0, 2, 0).Magnitude(), ShouldEqual, 2)
			So(NewVector3(0, 0, 3).Magnitude(), ShouldEqual, 3)
			So(NewVector3(1, 1, 0).Magnitude(), ShouldAlmostEqual, math.Sqrt(2))
		})
	})
}
