package models

type Room struct {
	Id                 int `orm:"auto"`
	RegionTopLeftX     float32
	RegionTopLeftY     float32
	RegionBottomRightX float32
	RegionBottomRightY float32
	Devices            []*RoomDevice `orm:"reverse(many)"`
}

type RoomDevice struct {
	Id            int `orm:"auto"`
	LocationX     float32
	LocationY     float32
	Rotation      float32
	Room          *Room `orm:"rel(fk)"`
	EffectType    string
	ConnectorType string
	ConnectorURI  string
}
