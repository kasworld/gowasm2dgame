package w2d_obj

import (
	"time"
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

type ReqState_data struct {
	Dummy uint8
}
type RspState_data struct {
	Time       time.Time // unixnano
	Teams      []*BallTeam
	Clouds     []*Cloud
	Background *Background
}

type ReqHeartbeat_data struct {
	Time time.Time
}
type RspHeartbeat_data struct {
	Time time.Time
}

type NotiInvalid_data struct {
	Dummy uint8
}
