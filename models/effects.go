package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/mmajko/openhvr-server/helpers"
)

var DefaultEffectManager *EffectManager
var registeredEffectTypes []*EffectType

type EffectType struct {
	Id   int `orm:"auto"`
	Name string
}

type EffectRequest struct {
	EffectType  *EffectType
	Duration    int
	Position    *helpers.Vector3
	Direction   *helpers.Vector3
	Range       float32
	Directional bool
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

func NewEffectRequest(effectType *EffectType, duration int, position *helpers.Vector3,
	effectRange float32) *EffectRequest {
	return &EffectRequest{effectType, duration, position,
		nil, effectRange, false}
}

func NewDirectionalRequest(effectType *EffectType, duration int,
	position, direction *helpers.Vector3, effectRange float32) *EffectRequest {
	return &EffectRequest{effectType, duration, position,
		direction, effectRange, true}
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

func (m *EffectManager) Clear() {
	m.runningEffects = nil
}

func RegisterDefaultEffectType(id int, name string) {
	registeredEffectTypes = append(registeredEffectTypes, &EffectType{Id: id, Name: name})
}

func ApplyDefaultEffectRegistration() {
	o := orm.NewOrm()
	for _, et := range registeredEffectTypes {
		o.Insert(et)
	}
}

func init() {
	DefaultEffectManager = NewEffectManager()
}
