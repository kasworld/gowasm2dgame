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

package stage

import (
	"time"

	"github.com/kasworld/gowasm2dgame/lib/gobase"

	"github.com/kasworld/gowasm2dgame/enums/effecttype"
	"github.com/kasworld/gowasm2dgame/enums/teamtype"
	"github.com/kasworld/gowasm2dgame/game/gameconst"
	"github.com/kasworld/gowasm2dgame/lib/anglemove"
	"github.com/kasworld/gowasm2dgame/lib/posacc"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_obj"
)

func (stg *Stage) makeTestData() {
	nowtick := time.Now().UnixNano()
	stg.Background = &w2d_obj.Background{
		Pa: posacc.PosAcc{
			LastMoveTick: nowtick,
			Dx:           stg.rnd.Float64()*500 - 250,
			Dy:           stg.rnd.Float64()*500 - 250,
		},
	}
	stg.Clouds = make([]*w2d_obj.Cloud, 10)
	for i := range stg.Clouds {
		stg.Clouds[i] = &w2d_obj.Cloud{
			SpriteNum: i,
			Pa: posacc.PosAcc{
				X:            stg.rnd.Float64() * gameconst.StageW,
				Y:            stg.rnd.Float64() * gameconst.StageH,
				Dx:           stg.rnd.Float64()*500 - 250,
				Dy:           stg.rnd.Float64()*500 - 250,
				LastMoveTick: nowtick,
			},
		}
	}
	stg.Effects = make([]*w2d_obj.Effect, effecttype.EffectType_Count*5)
	for i := range stg.Effects {
		stg.Effects[i] = &w2d_obj.Effect{
			EffectType: effecttype.EffectType(i % effecttype.EffectType_Count),
			BirthTick:  nowtick,
			Pa: posacc.PosAcc{
				X: stg.rnd.Float64() * gameconst.StageW,
				Y: stg.rnd.Float64() * gameconst.StageH,
			},
		}
	}
	stg.Teams = make([]*w2d_obj.BallTeam, teamtype.TeamType_Count)
	for i := range stg.Teams {
		stg.Teams[i] = stg.newBallTeam(teamtype.TeamType(i), nowtick)
	}
}

func (stg *Stage) newBallTeam(TeamType teamtype.TeamType, nowtick int64) *w2d_obj.BallTeam {
	bl := &w2d_obj.BallTeam{
		TeamType: TeamType,
		Ball: &w2d_obj.Ball{
			GOBase: gobase.GOBase{
				BirthTick: nowtick,
			},
			Pa: posacc.PosAcc{
				X:            stg.rnd.Float64() * gameconst.StageW,
				Y:            stg.rnd.Float64() * gameconst.StageH,
				Dx:           stg.rnd.Float64()*500 - 250,
				Dy:           stg.rnd.Float64()*500 - 250,
				LastMoveTick: nowtick,
			},
		},
		Shields:      make([]*w2d_obj.Shield, 12),
		SuperShields: make([]*w2d_obj.SuperShield, 12),
	}
	for i := range bl.Shields {
		bl.Shields[i] = &w2d_obj.Shield{
			GOBase: gobase.GOBase{
				BirthTick: nowtick,
			},
			Am: anglemove.AngleMove{
				Angle:        stg.rnd.Float64() * 360,
				AngleV:       stg.rnd.Float64()*300 - 150,
				LastMoveTick: nowtick,
			},
		}
	}
	for i := range bl.SuperShields {
		bl.SuperShields[i] = &w2d_obj.SuperShield{
			GOBase: gobase.GOBase{
				BirthTick: nowtick,
			},
			Am: anglemove.AngleMove{
				Angle:        stg.rnd.Float64() * 360,
				AngleV:       stg.rnd.Float64()*300 - 150,
				LastMoveTick: nowtick,
			},
		}
	}
	return bl
}
