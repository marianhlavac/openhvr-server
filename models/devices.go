package models

type Device struct {
	Id            int `orm:"auto"`
	LocationX     float32
	LocationY     float32
	LocationZ     float32
	Rotation      float32
	EffectType    string
	ConnectorType string
	ConnectorUri  string
	TimeoutAt     int64 `orm:"null"`
}
