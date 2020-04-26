package test

import (
	"testing"

	"github.com/mmajko/openhvr-server/models"
)

func TestNewManager(t *testing.T) {
	var m = models.NewEffectManager()

	if m == nil {
		t.Error("New EffectManager was not created.")
	}
}

func TestNewRequest(t *testing.T) {
	var m = models.NewEffectManager()
	var position = models.NewLocation(0.5, 1.5, -2.0)
	var request = models.NewEffectRequest(0, 10, position, 1.5, 1.44, false)

	if m.RunEffect(request) != nil {
		t.Error("Error when running effect")
	}
	if len(m.All()) == 0 {
		t.Error("Manager is empty after creating effect")
	}
	if m.All()[0].Request != request {
		t.Error("Different request has been saved")
	}
}

func TestTimeouted(t *testing.T) {
	var m = models.NewEffectManager()
	var position = models.NewLocation(0.5, 1.5, -2.0)
	var requestFuture = models.NewEffectRequest(0, 10, position, 1.5, 1.44, false)
	var requestHistory = models.NewEffectRequest(0, -10, position, 1.5, 1.44, false)

	if m.RunEffect(requestFuture) != nil {
		t.Error("Error when running effect")
	}
	if m.RunEffect(requestHistory) != nil {
		t.Error("Error when running effect")
	}

	if len(m.All()) != 2 {
		t.Error("Manager didn't create all objects")
	}
	if len(m.AllTimeouted()) != 1 {
		t.Error("Timeouted objects are missing")
	}
	var timeoutedFirst = m.AllTimeouted()[0]
	if timeoutedFirst.Request != requestHistory {
		t.Error("Wrongly determined timeouted requests")
	}
}
