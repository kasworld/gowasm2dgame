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

	"github.com/kasworld/gowasm2dgame/enums/teamtype"

	"github.com/kasworld/go-abs"
	"github.com/kasworld/gowasm2dgame/enums/gameobjtype"
	"github.com/kasworld/gowasm2dgame/lib/vector2f"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_obj"
)

func (o *GameObj) GetUUID() string {
	return o.UUID
}
func (o *GameObj) GetRect() vector2f.Rect {
	r := gameobjtype.Attrib[o.GOType].Size
	return vector2f.Rect{
		o.PosVt.X - r/2, o.PosVt.Y - r/2, r, r,
	}
}

type GameObj struct {
	teamType     teamtype.TeamType
	GOType       gameobjtype.GameObjType
	UUID         string
	BirthTick    int64
	LastMoveTick int64 // time.unixnano
	toDelete     bool

	PosVt vector2f.Vector2f
	MvVt  vector2f.Vector2f

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
		X:            o.PosVt.X,
		Y:            o.PosVt.Y,
		Dx:           o.MvVt.X,
		Dy:           o.MvVt.Y,
		Angle:        o.Angle,
		AngleV:       o.AngleV,
		DstUUID:      o.DstUUID,
	}
}

func (o *GameObj) MoveStraight(now int64) {
	diff := float64(now-o.LastMoveTick) / float64(time.Second)
	o.LastMoveTick = now
	o.PosVt = o.PosVt.Add(o.MvVt.MulF(diff))
}

func (o *GameObj) MoveCircular(now int64, ctvt vector2f.Vector2f) {
	diff := float64(now-o.LastMoveTick) / float64(time.Second)
	o.LastMoveTick = now
	o.Angle += o.AngleV * diff
	r := gameobjtype.Attrib[o.GOType].R
	o.PosVt = o.CalcCircularPos(ctvt, r)
}

func (o *GameObj) MoveHommingShield(now int64, dstPosVt vector2f.Vector2f) {
	diff := float64(now-o.LastMoveTick) / float64(time.Second)
	o.LastMoveTick = now
	o.PosVt = o.PosVt.Add(o.MvVt.MulF(diff))

	maxv := gameobjtype.Attrib[o.GOType].V

	dxyVt := dstPosVt.Sub(o.PosVt)
	o.MvVt = o.MvVt.Add(dxyVt.Normalize().MulF(maxv))
}

func (o *GameObj) MoveHommingBullet(now int64, dstPosVt vector2f.Vector2f) {
	diff := float64(now-o.LastMoveTick) / float64(time.Second)
	o.LastMoveTick = now
	o.PosVt = o.PosVt.Add(o.MvVt.MulF(diff))

	maxv := gameobjtype.Attrib[o.GOType].V
	dxyVt := dstPosVt.Sub(o.PosVt)
	o.MvVt = o.MvVt.Add(dxyVt.Normalize().MulF(maxv))
	o.LimitDxy()
}

func (o *GameObj) CheckLife(now int64) bool {
	lifetick := gameobjtype.Attrib[o.GOType].LifeTick
	return now-o.BirthTick < lifetick
}

func (o *GameObj) IsIn(w, h float64) bool {
	return 0 <= o.PosVt.X && o.PosVt.X <= w && 0 <= o.PosVt.Y && o.PosVt.Y <= h
}

func (o *GameObj) SetDxy(vt vector2f.Vector2f) {
	o.MvVt = vt
	o.LimitDxy()
}

func (o *GameObj) AddDxy(vt vector2f.Vector2f) {
	o.MvVt = o.MvVt.Add(vt)
	o.LimitDxy()
}

func (o *GameObj) LimitDxy() {
	maxv := gameobjtype.Attrib[o.GOType].V
	if o.MvVt.Abs() > maxv {
		o.MvVt = o.MvVt.Normalize().MulF(maxv)
	}
}

func (o *GameObj) BounceNormalize(w, h float64) {
	if o.PosVt.X < 0 {
		o.PosVt.X = 0
		o.MvVt.X = abs.Absf(o.MvVt.X)
	}
	if o.PosVt.Y < 0 {
		o.PosVt.Y = 0
		o.MvVt.Y = abs.Absf(o.MvVt.Y)
	}

	if o.PosVt.X > w {
		o.PosVt.X = w
		o.MvVt.X = -abs.Absf(o.MvVt.X)
	}
	if o.PosVt.Y > h {
		o.PosVt.Y = h
		o.MvVt.Y = -abs.Absf(o.MvVt.Y)
	}
}

func (o *GameObj) Wrap(w, h float64) vector2f.Vector2f {
	if o.PosVt.X < 0 {
		o.PosVt.X = w
	}
	if o.PosVt.Y < 0 {
		o.PosVt.Y = h
	}

	if o.PosVt.X > w {
		o.PosVt.X = 0
	}
	if o.PosVt.Y > h {
		o.PosVt.Y = 0
	}
	return o.PosVt
}

func (o *GameObj) CalcCircularPos(center vector2f.Vector2f, r float64) vector2f.Vector2f {

	rpos := vector2f.Vector2f{r * math.Cos(o.Angle), r * math.Sin(o.Angle)}
	return center.Add(rpos)
}

// CalcLenChange calc two gameobj change of len with time
// return current len , len change with time
// currentlen adjust with obj size
func (o *GameObj) CalcLenChange(dsto *GameObj) (float64, float64) {
	r1 := gameobjtype.Attrib[o.GOType].Size / 2
	r2 := gameobjtype.Attrib[dsto.GOType].Size / 2
	curLen := dsto.PosVt.Sub(o.PosVt).Abs()
	nextLen := dsto.PosVt.Add(dsto.MvVt).Sub(
		o.PosVt.Add(o.MvVt),
	).Abs()
	lenChange := nextLen - curLen
	return curLen - r1 - r2, lenChange
}
