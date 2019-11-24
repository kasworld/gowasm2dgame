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
	"github.com/kasworld/wrapper"
)

type Ball struct {
	sp  *Sprite
	spn int

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
		DispW:   64,
		DispH:   64,
	}
	bl.bgXWrap = wrapper.New(w).GetWrapSafeFn()
	bl.bgYWrap = wrapper.New(h).GetWrapSafeFn()
	bl.SetDir(initdir)
	return bl
}

func (bl *Ball) DrawTo(ctx js.Value) {
	bl.Move()
	// x := bl.bgXWrap(bl.X)
	// y := bl.bgYWrap(bl.Y)
	srcx, srcy := bl.sp.GetSliceXY(bl.spn)
	dstx, dsty := bl.X-bl.DispW/2, bl.Y-bl.DispH/2
	ctx.Call("drawImage",
		bl.sp.ImgCanvas,
		srcx, srcy, bl.sp.W, bl.sp.H,
		dstx, dsty, bl.DispW, bl.DispH,
	)

	// bl.sp.drawImageSliceAlignCenter(ctx, x, y, bl.spn)
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
