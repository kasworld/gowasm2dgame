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

	"github.com/kasworld/gowasm2dgame/enums/gameobjtype"
	"github.com/kasworld/gowasm2dgame/enums/teamtype"

	"github.com/kasworld/gowasm2dgame/enums/effecttype"
	"github.com/kasworld/gowasm2dgame/game/gameconst"
	"github.com/kasworld/gowasm2dgame/lib/anglemove"
	"github.com/kasworld/gowasm2dgame/lib/gobase"
	"github.com/kasworld/gowasm2dgame/lib/posacc"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_obj"
	"github.com/kasworld/uuidstr"
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
		Pa: posacc.PosAcc{
			LastMoveTick: time.Now().UnixNano(),
			Dx:           dx,
			Dy:           dy,
		},
	}
}

func (stg *Stage) NewCloud(i int) *w2d_obj.Cloud {
	dx, dy := CalcDxyFromAngelV(
		stg.rnd.Float64()*360,
		stg.rnd.Float64()*300,
	)
	return &w2d_obj.Cloud{
		SpriteNum: i,
		Pa: posacc.PosAcc{
			X:            stg.rnd.Float64() * gameconst.StageW,
			Y:            stg.rnd.Float64() * gameconst.StageH,
			Dx:           dx,
			Dy:           dy,
			LastMoveTick: time.Now().UnixNano(),
		},
	}
}

func (stg *Stage) NewEffect(
	et effecttype.EffectType, x, y float64) *w2d_obj.Effect {
	return &w2d_obj.Effect{
		EffectType: et,
		BirthTick:  time.Now().UnixNano(),
		Pa: posacc.PosAcc{
			X: x,
			Y: y,
		},
	}
}

func (stg *Stage) NewBallTeam(
	TeamType teamtype.TeamType,
	x, y float64) *w2d_obj.BallTeam {
	bl := &w2d_obj.BallTeam{
		TeamType:     TeamType,
		Ball:         stg.NewBall(x, y),
		Shields:      make([]*w2d_obj.Shield, 12),
		SuperShields: make([]*w2d_obj.SuperShield, 12),
	}
	for i := range bl.Shields {
		bl.Shields[i] = stg.NewShield()
	}
	for i := range bl.SuperShields {
		bl.SuperShields[i] = stg.NewSuperShield()
	}
	return bl
}

func (stg *Stage) NewBall(x, y float64) *w2d_obj.Ball {
	nowtick := time.Now().UnixNano()
	maxv := gameobjtype.Attrib[gameobjtype.Ball].V
	dx, dy := CalcDxyFromAngelV(
		stg.rnd.Float64()*360,
		stg.rnd.Float64()*maxv,
	)
	return &w2d_obj.Ball{
		GOBase: gobase.GOBase{
			UUID:      uuidstr.New(),
			BirthTick: nowtick,
		},
		Pa: posacc.PosAcc{
			X:            x,
			Y:            y,
			Dx:           dx,
			Dy:           dy,
			LastMoveTick: nowtick,
		},
	}
}

func (stg *Stage) NewShield() *w2d_obj.Shield {
	nowtick := time.Now().UnixNano()
	anglevmax := gameobjtype.Attrib[gameobjtype.Shield].V
	return &w2d_obj.Shield{
		GOBase: gobase.GOBase{
			UUID:      uuidstr.New(),
			BirthTick: nowtick,
		},
		Am: anglemove.AngleMove{
			Angle:        stg.rnd.Float64() * 360,
			AngleV:       stg.rnd.Float64()*anglevmax*2 - anglevmax,
			LastMoveTick: nowtick,
		},
	}
}

func (stg *Stage) NewSuperShield() *w2d_obj.SuperShield {
	nowtick := time.Now().UnixNano()
	anglevmax := gameobjtype.Attrib[gameobjtype.SuperShield].V
	return &w2d_obj.SuperShield{
		GOBase: gobase.GOBase{
			UUID:      uuidstr.New(),
			BirthTick: nowtick,
		},
		Am: anglemove.AngleMove{
			Angle:        stg.rnd.Float64() * 360,
			AngleV:       stg.rnd.Float64()*anglevmax*2 - anglevmax,
			LastMoveTick: nowtick,
		},
	}
}

func (stg *Stage) NewBullet(x, y float64) *w2d_obj.Bullet {
	nowtick := time.Now().UnixNano()
	maxv := gameobjtype.Attrib[gameobjtype.Bullet].V
	dx, dy := CalcDxyFromAngelV(
		stg.rnd.Float64()*360,
		maxv,
	)
	return &w2d_obj.Bullet{
		GOBase: gobase.GOBase{
			UUID:      uuidstr.New(),
			BirthTick: nowtick,
		},
		Pa: posacc.PosAcc{
			X:            x,
			Y:            y,
			Dx:           dx,
			Dy:           dy,
			LastMoveTick: nowtick,
		},
	}
}
