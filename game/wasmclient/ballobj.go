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
	"math"
	"syscall/js"

	"github.com/kasworld/direction"
	"github.com/kasworld/wrapper"
)

func calcCircularPos(cx, cy int, angle, r int) (int, int) {
	rad := float64(angle) * math.Pi / 180
	dstx := float64(cx) + float64(r)*math.Cos(rad)
	dsty := float64(cy) + float64(r)*math.Sin(rad)
	return int(dstx), int(dsty)
}

type SuperShield struct {
	sp     *Sprite
	DispW  int
	DispH  int
	r      int
	angle  int
	angleV int
	frame  int
}

type Shield struct {
	sp     *Sprite
	DispW  int
	DispH  int
	r      int
	angle  int
	angleV int
}

type HommingShield struct {
	sp    *Sprite
	DispW int
	DispH int
	X     int
	Y     int
	Dx    int
	Dy    int
}

type Bullet struct {
	sp    *Sprite
	DispW int
	DispH int
	X     int
	Y     int
	Dx    int
	Dy    int
}

type HommingBullet struct {
	sp    *Sprite
	DispW int
	DispH int
	X     int
	Y     int
	Dx    int
	Dy    int
	DstID int
}

type SuperBullet struct {
	sp    *Sprite
	DispW int
	DispH int
	X     int
	Y     int
	Dx    int
	Dy    int
}

type Ball struct {
	sp  *Sprite
	spn int

	shiels      []*Shield
	superShiels []*SuperShield

	bgXWrap func(i int) int
	bgYWrap func(i int) int
	BorderW int
	BorderH int

	DispW int
	DispH int

	X  int
	Y  int
	Dx int
	Dy int
}

func NewBall(sp *Sprite, spn int,
	ssp *Sprite,
	initdir direction.Direction_Type,
	x, y int,
	w, h int,
) *Ball {
	bl := &Ball{
		sp:      sp,
		spn:     spn,
		X:       x,
		Y:       y,
		BorderW: w,
		BorderH: h,
		DispW:   32,
		DispH:   32,
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
			sp:     sp,
			DispW:  16,
			DispH:  16,
			r:      28,
			angle:  i * 15,
			angleV: av,
		}
	}

	bl.superShiels = make([]*SuperShield, 24)
	for i := range bl.superShiels {
		av := 1
		if i%2 == 0 {
			av = -1
		}
		bl.superShiels[i] = &SuperShield{
			sp:     ssp,
			DispW:  16,
			DispH:  16,
			r:      48,
			angle:  15 + i*15,
			angleV: av,
			frame:  i * 3,
		}
	}
	return bl
}

func (bl *Ball) DrawTo(ctx js.Value) {
	bl.Move()
	srcx, srcy := bl.sp.GetSliceXY(bl.spn)
	dstx, dsty := bl.X-bl.DispW/2, bl.Y-bl.DispH/2
	ctx.Call("drawImage", bl.sp.ImgCanvas,
		srcx, srcy, bl.sp.W, bl.sp.H,
		dstx, dsty, bl.DispW, bl.DispH,
	)
	for _, v := range bl.shiels {
		v.angle += v.angleV
		srcx, srcy := v.sp.GetSliceXY(0)
		x, y := calcCircularPos(bl.X, bl.Y, v.angle, v.r)
		ctx.Call("drawImage", v.sp.ImgCanvas,
			srcx, srcy, v.sp.W, v.sp.H,
			x-v.DispW/2, y-v.DispH/2, v.DispW, v.DispH,
		)
	}
	for _, v := range bl.superShiels {
		v.angle += v.angleV
		v.frame++
		srcx, srcy := v.sp.GetSliceXY(v.frame)
		x, y := calcCircularPos(bl.X, bl.Y, v.angle, v.r)
		ctx.Call("drawImage", v.sp.ImgCanvas,
			srcx, srcy, v.sp.W, v.sp.H,
			x-v.DispW/2, y-v.DispH/2, v.DispW, v.DispH,
		)
	}
}

// check bounce
func (bl *Ball) Normalize() {
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

func (bl *Ball) Move() {
	bl.X += bl.Dx
	bl.Y += bl.Dy
	bl.Normalize()
}

func (bl *Ball) SetDxy(dx, dy int) {
	bl.Dx = dx
	bl.Dy = dy
}

func (bl *Ball) SetDir(dir direction.Direction_Type) {
	bl.Dx = dir.Vector()[0]
	bl.Dy = dir.Vector()[1]
}

func (bl *Ball) AccelerateDir(dir direction.Direction_Type) {
	if dir == direction.Dir_stop {
		bl.Dx = dir.Vector()[0]
		bl.Dy = dir.Vector()[1]
	} else {
		bl.Dx += dir.Vector()[0]
		bl.Dy += dir.Vector()[1]
	}
}

func (bl *Ball) GetAbsDx() int {
	if bl.Dx < 0 {
		return -bl.Dx
	}
	return bl.Dx
}
func (bl *Ball) GetAbsDy() int {
	if bl.Dy < 0 {
		return -bl.Dy
	}
	return bl.Dy
}
