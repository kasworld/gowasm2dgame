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
	"math/rand"
	"time"

	"github.com/kasworld/gowasm2dgame/enums/gameobjtype"
	"github.com/kasworld/gowasm2dgame/enums/teamtype"
	"github.com/kasworld/gowasm2dgame/game/gameconst"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_obj"
	"github.com/kasworld/uuidstr"
)

type BallTeam struct {
	rnd *rand.Rand `prettystring:"hide"`

	TeamType    teamtype.TeamType
	IsAlive     bool
	RespawnTick int64

	Ball *GameObj // ball is special
	Objs []*GameObj
}

func NewBallTeam(TeamType teamtype.TeamType) *BallTeam {
	nowtick := time.Now().UnixNano()
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	bt := &BallTeam{
		rnd:      rnd,
		IsAlive:  true,
		TeamType: TeamType,
		Ball: &GameObj{
			teamType:     TeamType,
			GOType:       gameobjtype.Ball,
			UUID:         uuidstr.New(),
			BirthTick:    nowtick,
			LastMoveTick: nowtick,
			X:            rnd.Float64() * gameconst.StageW,
			Y:            rnd.Float64() * gameconst.StageH,
		},
		Objs: make([]*GameObj, 0),
	}
	maxv := gameobjtype.Attrib[gameobjtype.Ball].V
	dx, dy := CalcDxyFromAngelV(
		bt.rnd.Float64()*360,
		bt.rnd.Float64()*maxv,
	)
	bt.Ball.SetDxy(dx, dy)
	return bt
}

func (bt *BallTeam) RespawnBall() {
	bt.IsAlive = true
	bt.Ball.toDelete = false
	bt.Ball.X = bt.rnd.Float64() * gameconst.StageW
	bt.Ball.Y = bt.rnd.Float64() * gameconst.StageH
	bt.Ball.Dx = 0
	bt.Ball.Dy = 0
}

func (bt *BallTeam) ToPacket() *w2d_obj.BallTeam {
	rtn := &w2d_obj.BallTeam{
		TeamType: bt.TeamType,
		Ball:     bt.Ball.ToPacket(),
		Objs:     make([]*w2d_obj.GameObj, 0),
	}
	for _, v := range bt.Objs {
		if v.toDelete {
			continue
		}
		rtn.Objs = append(rtn.Objs, v.ToPacket())
	}
	return rtn
}

func (bt *BallTeam) Count(ot gameobjtype.GameObjType) int {
	rtn := 0
	for _, v := range bt.Objs {
		if v.toDelete {
			continue
		}
		if v.GOType == ot {
			rtn++
		}
	}
	return rtn
}

func (bt *BallTeam) addGObj(o *GameObj) {
	for i, v := range bt.Objs {
		if v.toDelete {
			bt.Objs[i] = o
			return
		}
	}
	bt.Objs = append(bt.Objs, o)
}

func (bt *BallTeam) AddShield(angle, anglev float64) *GameObj {
	nowtick := time.Now().UnixNano()
	o := &GameObj{
		teamType:     bt.TeamType,
		GOType:       gameobjtype.Shield,
		UUID:         uuidstr.New(),
		BirthTick:    nowtick,
		LastMoveTick: nowtick,
		Angle:        angle,
		AngleV:       anglev,
	}
	bt.addGObj(o)
	return o
}

func (bt *BallTeam) AddSuperShield(angle, anglev float64) *GameObj {
	nowtick := time.Now().UnixNano()
	o := &GameObj{
		teamType:     bt.TeamType,
		GOType:       gameobjtype.SuperShield,
		UUID:         uuidstr.New(),
		BirthTick:    nowtick,
		LastMoveTick: nowtick,
		Angle:        angle,
		AngleV:       anglev,
	}
	bt.addGObj(o)
	return o
}

func (bt *BallTeam) AddBullet(angle, anglev float64) *GameObj {
	nowtick := time.Now().UnixNano()
	dx, dy := CalcDxyFromAngelV(angle, anglev)
	o := &GameObj{
		teamType:     bt.TeamType,
		GOType:       gameobjtype.Bullet,
		UUID:         uuidstr.New(),
		BirthTick:    nowtick,
		LastMoveTick: nowtick,
		X:            bt.Ball.X,
		Y:            bt.Ball.Y,
		Dx:           dx,
		Dy:           dy,
	}
	bt.addGObj(o)
	return o
}

func (bt *BallTeam) AddSuperBullet(angle, anglev float64) *GameObj {
	nowtick := time.Now().UnixNano()
	dx, dy := CalcDxyFromAngelV(angle, anglev)
	o := &GameObj{
		teamType:     bt.TeamType,
		GOType:       gameobjtype.SuperBullet,
		UUID:         uuidstr.New(),
		BirthTick:    nowtick,
		LastMoveTick: nowtick,
		X:            bt.Ball.X,
		Y:            bt.Ball.Y,
		Dx:           dx,
		Dy:           dy,
	}
	bt.addGObj(o)
	return o
}

func (bt *BallTeam) AddHommingShield(angle, anglev float64) *GameObj {
	nowtick := time.Now().UnixNano()
	dx, dy := CalcDxyFromAngelV(angle, anglev)
	o := &GameObj{
		teamType:     bt.TeamType,
		GOType:       gameobjtype.HommingShield,
		UUID:         uuidstr.New(),
		BirthTick:    nowtick,
		LastMoveTick: nowtick,
		Angle:        angle,
		AngleV:       anglev,
		X:            bt.Ball.X + dx,
		Y:            bt.Ball.Y + dy,
		Dx:           dx,
		Dy:           dy,
	}
	bt.addGObj(o)
	return o
}

func (bt *BallTeam) AddHommingBullet(angle, anglev float64, dstid string) *GameObj {
	nowtick := time.Now().UnixNano()
	dx, dy := CalcDxyFromAngelV(angle, anglev)
	o := &GameObj{
		teamType:     bt.TeamType,
		GOType:       gameobjtype.HommingBullet,
		UUID:         uuidstr.New(),
		BirthTick:    nowtick,
		LastMoveTick: nowtick,
		Angle:        angle,
		AngleV:       anglev,
		X:            bt.Ball.X,
		Y:            bt.Ball.Y,
		Dx:           dx,
		Dy:           dy,
		DstUUID:      dstid,
	}
	bt.addGObj(o)
	return o
}
