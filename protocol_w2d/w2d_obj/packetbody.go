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

type ReqEnterStage_data struct {
	StageUUID string // may be not same to req stage
	NickToUse string
}
type RspEnterStage_data struct {
	StageUUID string // may be not same to req stage
	NickToUse string // may be not same to req nick
}

type ReqChatToStage_data struct {
	Chat string
}
type RspChatToStage_data struct {
	Dummy uint8
}

type ReqHeartbeat_data struct {
	Tick int64
}
type RspHeartbeat_data struct {
	Tick int64 // same req tick , to calc round trip time
}

//////////////////////////////////////////////////////////////////////////////

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
	UUID  string
	Stats []TeamStat
}

type NotiStageBroadCast_data struct {
	SenderNick string
	Chat       string
}

//////////////////////////////////////////////////////////////////////////////

type TeamStat struct {
	UUID     string
	Alive    bool
	TeamType teamtype.TeamType
	ActStats acttype_stats.ActTypeStat
}

type Act struct {
	Act acttype.ActType

	// accel, fire bullet
	Angle  float64 // degree
	AngleV float64 // pixel /sec

	// homming
	DstObjID string
}
