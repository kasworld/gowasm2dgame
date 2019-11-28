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

	"github.com/kasworld/uuidstr"

	"github.com/kasworld/go-abs"
	"github.com/kasworld/gowasm2dgame/enums/gameobjtype"
	"github.com/kasworld/gowasm2dgame/enums/teamtype"
	"github.com/kasworld/gowasm2dgame/game/gameconst"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_obj"
)

type BallTeam struct {
	TeamType teamtype.TeamType
	Ball     *GameObj // ball is special
	Objs     []*GameObj
}

func NewBallTeam(TeamType teamtype.TeamType, x, y float64) *BallTeam {
	nowtick := time.Now().UnixNano()

	bt := &BallTeam{
		TeamType: TeamType,
		Ball: &GameObj{
			GOType:       gameobjtype.Ball,
			UUID:         uuidstr.New(),
			BirthTick:    nowtick,
			LastMoveTick: nowtick,
			X:            x,
			Y:            y,
		},
		Objs: make([]*GameObj, 0),
	}
	return bt
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

func (bt *BallTeam) Move(now int64) {
	bt.Ball.MoveStraight(now)
	bt.Ball.BounceNormalize(gameconst.StageW, gameconst.StageH)
	for _, v := range bt.Objs {
		if v.toDelete {
			continue
		}
		switch v.GOType {
		default:
		case gameobjtype.Bullet, gameobjtype.SuperBullet:
			v.MoveStraight(now)
			if !v.IsIn(gameconst.StageW, gameconst.StageH) {
				v.toDelete = true
			}
		case gameobjtype.Shield, gameobjtype.SuperShield:
			v.MoveCircular(now, bt.Ball.X, bt.Ball.Y)
		case gameobjtype.HommingShield:
		case gameobjtype.HommingBullet:
		}
		if !v.CheckLife(now) {
			v.toDelete = true
		}
	}
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

type GameObj struct {
	GOType       gameobjtype.GameObjType
	UUID         string
	BirthTick    int64
	LastMoveTick int64 // time.unixnano
	toDelete     bool

	X float64
	Y float64

	Dx float64 // move line
	Dy float64

	Angle  float64 // move circular
	AngleV float64

	DstUUID string // move to dest
}

func (o *GameObj) ToPacket() *w2d_obj.GameObj {
	return &w2d_obj.GameObj{
		GOType:       o.GOType,
		UUID:         o.UUID,
		BirthTick:    o.BirthTick,
		LastMoveTick: o.LastMoveTick,
		X:            o.X,
		Y:            o.Y,
		Dx:           o.Dx,
		Dy:           o.Dy,
		Angle:        o.Angle,
		AngleV:       o.AngleV,
		DstUUID:      o.DstUUID,
	}
}

func (o *GameObj) MoveStraight(now int64) {
	diff := float64(now-o.LastMoveTick) / float64(time.Second)
	o.LastMoveTick = now
	o.X += o.Dx * diff
	o.Y += o.Dy * diff
}

func (o *GameObj) MoveCircular(now int64, cx, cy float64) {
	diff := float64(now-o.LastMoveTick) / float64(time.Second)
	o.LastMoveTick = now
	o.Angle += o.AngleV * diff
	r := gameobjtype.Attrib[o.GOType].R
	o.X, o.Y = o.CalcCircularPos(cx, cy, r)
}

func (o *GameObj) MoveHomming(now int64, dstx, dsty float64) {
}

func (o *GameObj) CheckLife(now int64) bool {
	lifetick := gameobjtype.Attrib[o.GOType].LifeTick
	return now-o.BirthTick < lifetick
}

func (o *GameObj) IsIn(w, h float64) bool {
	return 0 <= o.X && o.X <= w && 0 <= o.Y && o.Y <= h
}

func (o *GameObj) SetDxy(dx, dy float64) {
	o.Dx = dx
	o.Dy = dy
}

func (o *GameObj) BounceNormalize(w, h float64) {
	if o.X < 0 {
		o.X = 0
		o.Dx = abs.Absf(o.Dx)
	}
	if o.Y < 0 {
		o.Y = 0
		o.Dy = abs.Absf(o.Dy)
	}

	if o.X > w {
		o.X = w
		o.Dx = -abs.Absf(o.Dx)
	}
	if o.Y > h {
		o.Y = h
		o.Dy = -abs.Absf(o.Dy)
	}
}

func (o *GameObj) Wrap(w, h float64) (float64, float64) {
	if o.X < 0 {
		o.X = w
	}
	if o.Y < 0 {
		o.Y = h
	}

	if o.X > w {
		o.X = 0
	}
	if o.Y > h {
		o.Y = 0
	}
	return o.X, o.Y
}

func (o *GameObj) CalcCircularPos(cx, cy, r float64) (float64, float64) {
	rad := o.Angle * math.Pi / 180
	dstx := cx + r*math.Cos(rad)
	dsty := cy + r*math.Sin(rad)
	return dstx, dsty
}
