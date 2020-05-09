// Copyright 2015,2016,2017,2018,2019,2020 SeukWon Kang (kasworld@gmail.com)
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package w2d_obj

import (
	"github.com/kasworld/gowasm2dgame/enum/acttype"
	"github.com/kasworld/gowasm2dgame/enum/acttype_vector"
	"github.com/kasworld/gowasm2dgame/enum/teamtype"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_idcmd"
)

type ReqInvalid_data struct {
	Dummy uint8
}
type RspInvalid_data struct {
	Dummy uint8
}

type ReqLogin_data struct {
	SessionKey string
	NickName   string
	AuthKey    string
}
type RspLogin_data struct {
	Version         string
	ProtocolVersion string
	DataVersion     string

	SessionKey string
	NickName   string
	CmdList    [w2d_idcmd.CommandID_Count]bool
}

type ReqHeartbeat_data struct {
	Tick int64
}
type RspHeartbeat_data struct {
	Tick int64
}

type ReqChat_data struct {
	Chat string
}
type RspChat_data struct {
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

type NotiStageChat_data struct {
	SenderNick string
	Chat       string
}

//////////////////////////////////////////////////////////////////////////////

type TeamStat struct {
	UUID     string
	Alive    bool
	TeamType teamtype.TeamType
	ActStats acttype_vector.ActTypeVector
}

type Act struct {
	Act acttype.ActType

	// accel, fire bullet
	Angle  float64 // degree
	AngleV float64 // pixel /sec

	// homming
	DstObjID string
}
