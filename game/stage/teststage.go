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
	"github.com/kasworld/gowasm2dgame/enums/effecttype"
	"github.com/kasworld/gowasm2dgame/enums/teamtype"
	"github.com/kasworld/gowasm2dgame/game/gameconst"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_obj"
)

func (stg *Stage) makeTestData() {
	stg.Background = stg.NewBackground()
	stg.Clouds = make([]*w2d_obj.Cloud, 10)
	for i := range stg.Clouds {
		stg.Clouds[i] = stg.NewCloud(i)
	}
	stg.Effects = make([]*w2d_obj.Effect, effecttype.EffectType_Count*5)
	for i := range stg.Effects {
		et := effecttype.EffectType(i % effecttype.EffectType_Count)
		stg.Effects[i] = stg.NewEffect(et,
			stg.rnd.Float64()*gameconst.StageW,
			stg.rnd.Float64()*gameconst.StageH,
		)
	}
	stg.Teams = make([]*w2d_obj.BallTeam, teamtype.TeamType_Count)
	for i := range stg.Teams {
		stg.Teams[i] = stg.NewBallTeam(
			teamtype.TeamType(i),
			stg.rnd.Float64()*gameconst.StageW,
			stg.rnd.Float64()*gameconst.StageH,
		)
	}
}
