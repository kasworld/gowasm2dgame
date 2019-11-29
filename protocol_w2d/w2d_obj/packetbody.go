package w2d_obj

import (
	"github.com/kasworld/gowasm2dgame/enums/acttype"
)

type ReqInvalid_data struct {
	Dummy uint8
}
type RspInvalid_data struct {
	Dummy uint8
}

type ReqMakeTeam_data struct {
	TeamName string
}
type RspMakeTeam_data struct {
	TeamName string
	TeamID   int
}

type ReqAct_data struct {
	TeamID string
	Acts   []Act
}
type RspAct_data struct {
	Dummy uint8
}

type ReqHeartbeat_data struct {
	Tick int64
}
type RspHeartbeat_data struct {
	Tick int64
}

type NotiInvalid_data struct {
	Dummy uint8
}

type NotiStageInfo_data struct {
	Tick       int64
	Background *Background
	Teams      []*BallTeam
	Effects    []*Effect
	Clouds     []*Cloud
}

////////////////

type Act struct {
	Act acttype.ActType

	// accel, fire bullet
	Angle  float64 // degree
	AngleV float64 // pixel /sec

	// homming
	DstObjID string
}
