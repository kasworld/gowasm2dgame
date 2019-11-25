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
	"syscall/js"

	"github.com/kasworld/direction"
	"github.com/kasworld/gowasm2dgame/enums/gameobjtype"
	"github.com/kasworld/wrapper"
)

type SuperShield struct {
	sp     *Sprite
	angle  int
	angleV int
	frame  int
}

type Shield struct {
	sp     *Sprite
	angle  int
	angleV int
}

type HommingShield struct {
	sp *Sprite
	X  int
	Y  int
	Dx int
	Dy int
}

type Bullet struct {
	sp *Sprite
	X  int
	Y  int
	Dx int
	Dy int
}

type HommingBullet struct {
	sp    *Sprite
	X     int
	Y     int
	Dx    int
	Dy    int
	DstID int
}

type SuperBullet struct {
	sp *Sprite
	X  int
	Y  int
	Dx int
	Dy int
}

type BallTeam struct {
	sp *Sprite

	shiels        []*Shield
	superShields  []*SuperShield
	hommingShiels []*HommingShield

	bgXWrap func(i int) int
	bgYWrap func(i int) int
	BorderW int
	BorderH int

	X  int
	Y  int
	Dx int
	Dy int
}

func NewBallTeam(
	sp [gameobjtype.GameObjType_Count]*Sprite,
	initdir direction.Direction_Type,
	x, y int,
	w, h int,
) *BallTeam {
	bl := &BallTeam{
		sp:      sp[gameobjtype.Ball],
		X:       x,
		Y:       y,
		BorderW: w,
		BorderH: h,
	}
	bl.bgXWrap = wrapper.New(w).GetWrapSafeFn()
	bl.bgYWrap = wrapper.New(h).GetWrapSafeFn()
	bl.SetDir(initdir)

	bl.shiels = make([]*Shield, 24)
	for i := range bl.shiels {
		av := 1
		if i%2 == 0 {
			av = -1
		}
		bl.shiels[i] = &Shield{
			sp:     sp[gameobjtype.Shield],
			angle:  i * 15,
			angleV: av,
		}
	}

	bl.superShields = make([]*SuperShield, 24)
	for i := range bl.superShields {
		av := 1
		if i%2 == 0 {
			av = -1
		}
		bl.superShields[i] = &SuperShield{
			sp:     sp[gameobjtype.SuperShield],
			angle:  15 + i*15,
			angleV: av,
			frame:  i * 3,
		}
	}
	return bl
}

func (bl *BallTeam) DrawTo(ctx js.Value) {
	bl.Move()
	dispSize := gameobjtype.Attrib[gameobjtype.Ball].Size
	srcx, srcy := bl.sp.GetSliceXY(0)
	dstx, dsty := bl.X-dispSize/2, bl.Y-dispSize/2
	ctx.Call("drawImage", bl.sp.ImgCanvas,
		srcx, srcy, dispSize, dispSize,
		dstx, dsty, dispSize, dispSize,
	)
	dispSize = gameobjtype.Attrib[gameobjtype.Shield].Size
	dispR := gameobjtype.Attrib[gameobjtype.Shield].R
	for _, v := range bl.shiels {
		v.angle += v.angleV
		srcx, srcy := v.sp.GetSliceXY(0)
		x, y := calcCircularPos(bl.X, bl.Y, v.angle, dispR)
		ctx.Call("drawImage", v.sp.ImgCanvas,
			srcx, srcy, dispSize, dispSize,
			x-dispSize/2, y-dispSize/2, dispSize, dispSize,
		)
	}
	dispSize = gameobjtype.Attrib[gameobjtype.SuperShield].Size
	dispR = gameobjtype.Attrib[gameobjtype.SuperShield].R
	for _, v := range bl.superShields {
		v.angle += v.angleV
		v.frame++
		srcx, srcy := v.sp.GetSliceXY(v.frame)
		x, y := calcCircularPos(bl.X, bl.Y, v.angle, dispR)
		ctx.Call("drawImage", v.sp.ImgCanvas,
			srcx, srcy, dispSize, dispSize,
			x-dispSize/2, y-dispSize/2, dispSize, dispSize,
		)
	}
}

// check bounce
func (bl *BallTeam) Normalize() {
	if bl.X < 0 {
		bl.X = 0
		bl.Dx = bl.GetAbsDx()
	}
	if bl.Y < 0 {
		bl.Y = 0
		bl.Dy = bl.GetAbsDy()
	}

	if bl.X > bl.BorderW {
		bl.X = bl.BorderW
		bl.Dx = -bl.GetAbsDx()
	}
	if bl.Y > bl.BorderH {
		bl.Y = bl.BorderH
		bl.Dy = -bl.GetAbsDy()
	}
}

func (bl *BallTeam) Move() {
	bl.X += bl.Dx
	bl.Y += bl.Dy
	bl.Normalize()
}

func (bl *BallTeam) SetDxy(dx, dy int) {
	bl.Dx = dx
	bl.Dy = dy
}

func (bl *BallTeam) SetDir(dir direction.Direction_Type) {
	bl.Dx = dir.Vector()[0]
	bl.Dy = dir.Vector()[1]
}

func (bl *BallTeam) AccelerateDir(dir direction.Direction_Type) {
	if dir == direction.Dir_stop {
		bl.Dx = dir.Vector()[0]
		bl.Dy = dir.Vector()[1]
	} else {
		bl.Dx += dir.Vector()[0]
		bl.Dy += dir.Vector()[1]
	}
}

func (bl *BallTeam) GetAbsDx() int {
	if bl.Dx < 0 {
		return -bl.Dx
	}
	return bl.Dx
}
func (bl *BallTeam) GetAbsDy() int {
	if bl.Dy < 0 {
		return -bl.Dy
	}
	return bl.Dy
}
