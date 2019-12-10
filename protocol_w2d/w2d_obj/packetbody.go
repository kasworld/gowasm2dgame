package w2d_obj

import (
	"github.com/kasworld/gowasm2dgame/enums/acttype"
	"github.com/kasworld/gowasm2dgame/enums/acttype_stats"
	"github.com/kasworld/gowasm2dgame/enums/teamtype"
)

type ReqInvalid_data struct {
	Dummy uint8
}
type RspInvalid_data struct {
	Dummy uint8
}

type ReqListStage_data struct {
	Dummy uint8
}
type RspListStage_data struct {
	Dummy uint8
}

type ReqEnterStage_data struct {
	Dummy uint8
}
type RspEnterStage_data struct {
	Dummy uint8
}

type ReqLeaveStage_data struct {
	Dummy uint8
}
type RspLeaveStage_data struct {
	Dummy uint8
}

type ReqChatToStage_data struct {
	Dummy uint8
}
type RspChatToStage_data struct {
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

////////////////

type NotiInvalid_data struct {
	Dummy uint8
}

type NotiStageInfo_data struct {
	Tick       int64
	Background *Background
	Teams      []*Team
	Effects    []*Effect
	Clouds     []*Cloud
}

type NotiStatsInfo_data struct {
	ActStats [teamtype.TeamType_Count]acttype_stats.ActTypeStat
}

type NotiStageBroadCast_data struct {
	Dummy uint8
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
