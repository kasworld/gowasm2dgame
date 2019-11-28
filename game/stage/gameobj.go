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

	"github.com/kasworld/go-abs"
	"github.com/kasworld/gowasm2dgame/enums/gameobjtype"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_obj"
)

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
	diff := float64(now-o.LastMoveTick) / float64(time.Second)
	o.LastMoveTick = now
	o.X += o.Dx * diff
	o.Y += o.Dy * diff

	maxv := gameobjtype.Attrib[o.GOType].V
	dx := dstx - o.X
	dy := dsty - o.Y
	l := math.Sqrt(dx*dx + dy*dy)
	o.Dx += dx / l * maxv
	o.Dy += dy / l * maxv
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
