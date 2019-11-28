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
	"math"
	"time"

	"github.com/kasworld/gowasm2dgame/enums/effecttype"
	"github.com/kasworld/gowasm2dgame/game/gameconst"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_obj"
)

func CalcDxyFromAngelV(angle float64, speed float64) (float64, float64) {
	rad := angle * math.Pi / 180
	dx := speed * math.Cos(rad)
	dy := speed * math.Sin(rad)
	return dx, dy
}

func (stg *Stage) NewBackground() *w2d_obj.Background {
	dx, dy := CalcDxyFromAngelV(
		stg.rnd.Float64()*360,
		stg.rnd.Float64()*300,
	)
	return &w2d_obj.Background{
		LastMoveTick: time.Now().UnixNano(),
		Dx:           dx,
		Dy:           dy,
	}
}

func (stg *Stage) NewCloud(i int) *w2d_obj.Cloud {
	dx, dy := CalcDxyFromAngelV(
		stg.rnd.Float64()*360,
		stg.rnd.Float64()*300,
	)
	return &w2d_obj.Cloud{
		SpriteNum:    i,
		X:            stg.rnd.Float64() * gameconst.StageW,
		Y:            stg.rnd.Float64() * gameconst.StageH,
		Dx:           dx,
		Dy:           dy,
		LastMoveTick: time.Now().UnixNano(),
	}
}

func (stg *Stage) NewEffect(
	et effecttype.EffectType, x, y float64) *w2d_obj.Effect {
	return &w2d_obj.Effect{
		EffectType: et,
		BirthTick:  time.Now().UnixNano(),
		X:          x,
		Y:          y,
	}
}
