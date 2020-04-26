package models

import "time"

var DefaultEffectManager *EffectManager

type Location struct {
	X float32
	Y float32
	Z float32
}

type EffectRequest struct {
	EffectType     int
	Duration       int
	Position       Location
	Range          float32
	Direction      float32
	DirectionBased bool
}

type EffectPerformance struct {
	Id         int64
	Request    *EffectRequest
	TimeoutsAt int64
}

type EffectManager struct {
	runningEffects []*EffectPerformance
	lastID         int64
}

func NewEffectManager() *EffectManager {
	return &EffectManager{}
}

func NewLocation(x, y, z float32) *Location {
	return &Location{x, y, z}
}

func NewEffectRequest(effectType int, duration int, position *Location, effectRange, direction float32, dirBased bool) *EffectRequest {
	return &EffectRequest{effectType, duration, *position, effectRange, direction, dirBased}
}

func (m *EffectManager) RunEffect(effectRequest *EffectRequest) error {
	m.lastID++
	var timeoutsAt = time.Now().Unix() + int64(effectRequest.Duration)

	m.runningEffects = append(m.runningEffects, &EffectPerformance{
		m.lastID, effectRequest, timeoutsAt,
	})

	return nil
}

func (m *EffectManager) All() []*EffectPerformance {
	return m.runningEffects
}

func (m *EffectManager) AllTimeouted() []*EffectPerformance {
	var now = time.Now().Unix()
	var timeouted []*EffectPerformance
	for _, effect := range m.runningEffects {
		if effect.TimeoutsAt < now {
			timeouted = append(timeouted, effect)
		}
	}
	return timeouted
}

func init() {
	DefaultEffectManager = NewEffectManager()
}
