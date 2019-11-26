// Copyright 2015,2016,2017,2018,2019 SeukWon Kang (kasworld@gmail.com)
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package wasmclient

import (
	"github.com/kasworld/gowasm2dgame/enums/teamtype"
	"github.com/kasworld/gowasm2dgame/lib/anglemove"
	"github.com/kasworld/gowasm2dgame/lib/posacc"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_obj"
)

func (vp *Viewport2d) makeTestObj() {
	vp.bgObj = &w2d_obj.Background{}
	vp.bgObj.Pa.SetDxy(vp.rnd.Float64()*5-2, vp.rnd.Float64()*5-2)

	vp.cloudObjs = make([]*w2d_obj.Cloud, 10)
	for i := range vp.cloudObjs {
		vp.cloudObjs[i] = &w2d_obj.Cloud{
			SpriteNum: i,
			Pa: posacc.PosAcc{
				X:  vp.rnd.Float64() * vp.W,
				Y:  vp.rnd.Float64() * vp.H,
				Dx: vp.rnd.Float64()*5 - 2,
				Dy: vp.rnd.Float64()*5 - 2,
			},
		}
	}
	vp.ballTeams = make([]*w2d_obj.BallTeam, teamtype.TeamType_Count)
	for i := range vp.ballTeams {
		vp.ballTeams[i] = vp.newBallTeam(teamtype.TeamType(i))
	}
}

func (vp *Viewport2d) newBallTeam(TeamType teamtype.TeamType) *w2d_obj.BallTeam {
	bl := &w2d_obj.BallTeam{
		TeamType: TeamType,
		Ball: &w2d_obj.Ball{
			Pa: posacc.PosAcc{
				X:  vp.rnd.Float64() * vp.W,
				Y:  vp.rnd.Float64() * vp.H,
				Dx: vp.rnd.Float64()*5 - 2,
				Dy: vp.rnd.Float64()*5 - 2,
			},
		},
		Shields:      make([]*w2d_obj.Shield, 12),
		SuperShields: make([]*w2d_obj.SuperShield, 12),
	}
	for i := range bl.Shields {
		bl.Shields[i] = &w2d_obj.Shield{
			Am: anglemove.AngleMove{
				Angle:  vp.rnd.Float64() * 360,
				AngleV: vp.rnd.Float64()*3 - 1,
			},
		}
	}
	for i := range bl.SuperShields {
		bl.SuperShields[i] = &w2d_obj.SuperShield{
			Am: anglemove.AngleMove{
				Angle:  vp.rnd.Float64() * 360,
				AngleV: vp.rnd.Float64()*3 - 1,
			},
		}
	}
	return bl
}
