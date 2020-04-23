package models

import "time"

var DefaultEffectManager *EffectManager

type EffectTransform struct {
	X         float32
	Y         float32
	Z         float32
	Range     float32
	Direction float32
}

type EffectRequest struct {
	EffectType     string
	Duration       int
	Transform      EffectTransform
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

func NewEffectTransform(x, y, z, effectRange, direction float32) *EffectTransform {
	return &EffectTransform{x, y, z, effectRange, direction}
}

func NewEffectRequest(effectType string, duration int, transform *EffectTransform, dirBased bool) *EffectRequest {
	return &EffectRequest{effectType, duration, *transform, dirBased}
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
