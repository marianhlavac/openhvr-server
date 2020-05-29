package devicedrivers

import (
	"testing"

	"github.com/mmajko/openhvr-server/helpers"
	"github.com/mmajko/openhvr-server/models"
	. "github.com/smartystreets/goconvey/convey"
)

func TestTasmotaFanDriver(t *testing.T) {
	Convey("IsRotatedWithinSpread 45 degree spread device", t, func() {
		d := models.Device{
			Id: 0, Name: "", Type: "", EffectType: 0,
			LocationX: 1, LocationY: 2, LocationZ: 3,
			DirectionX: 1, DirectionY: 0, DirectionZ: 0,
			DirectionSpread: 45, ConnectorUri: "", ConnectorParam: "",
			TimeoutAt: 0,
		}

		Convey("Correctly confirms aligned rotations", func() {
			So(IsRotatedWithinSpread(&d, *helpers.NewVector3(1, 0, 0)), ShouldBeTrue)
			So(IsRotatedWithinSpread(&d, *helpers.NewVector3(0.9, 0, 0)), ShouldBeTrue)
			So(IsRotatedWithinSpread(&d, *helpers.NewVector3(0.9, 0.1, 0.1)), ShouldBeTrue)
			So(IsRotatedWithinSpread(&d, *helpers.NewVector3(0.927, 0.374, 0)), ShouldBeTrue)
			So(IsRotatedWithinSpread(&d, *helpers.NewVector3(0.984, -0.174, 0)), ShouldBeTrue)
		})

		Convey("Correctly rejects misaligned rotations", func() {
			So(IsRotatedWithinSpread(&d, *helpers.NewVector3(0, 1, 0)), ShouldBeFalse)
			So(IsRotatedWithinSpread(&d, *helpers.NewVector3(0, 0, 1)), ShouldBeFalse)
			So(IsRotatedWithinSpread(&d, *helpers.NewVector3(0.1, 1, 0)), ShouldBeFalse)
			So(IsRotatedWithinSpread(&d, *helpers.NewVector3(0.913, 0.406, 0)), ShouldBeFalse)
			So(IsRotatedWithinSpread(&d, *helpers.NewVector3(0.866, -0.5, 0)), ShouldBeFalse)
		})
	})
}
