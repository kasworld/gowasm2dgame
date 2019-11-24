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

type Cloud struct {
	sp  *Sprite
	spn int

	bgXWrap func(i int) int
	bgYWrap func(i int) int

	X  int
	Y  int
	Dx int
	Dy int
}

func NewCloud(sp *Sprite, spn int,
	initdir direction.Direction_Type,
	x, y int,
	w, h int,
) *Cloud {
	cld := &Cloud{
		sp:  sp,
		spn: spn,
		X:   x,
		Y:   y,
	}
	cld.bgXWrap = wrapper.New(w).GetWrapSafeFn()
	cld.bgYWrap = wrapper.New(h).GetWrapSafeFn()
	cld.SetDir(initdir)
	return cld
}

func (cld *Cloud) DrawTo(ctx js.Value) {
	cld.Move()
	x := cld.bgXWrap(cld.X)
	y := cld.bgYWrap(cld.Y)
	srcx, srcy := cld.sp.GetSliceXY(cld.spn)
	ctx.Call("drawImage", cld.sp.ImgCanvas,
		srcx, srcy, cld.sp.W, cld.sp.H,
		x-cld.sp.W/2, y-cld.sp.H/2, cld.sp.W, cld.sp.H,
	)
}

func (cld *Cloud) Move() {
	cld.X += cld.Dx
	cld.Y += cld.Dy
}

func (cld *Cloud) SetDxy(dx, dy int) {
	cld.Dx = dx
	cld.Dy = dy
}

func (cld *Cloud) SetDir(dir direction.Direction_Type) {
	cld.Dx = dir.Vector()[0]
	cld.Dy = dir.Vector()[1]
}

func (cld *Cloud) AccelerateDir(dir direction.Direction_Type) {
	if dir == direction.Dir_stop {
		cld.Dx = dir.Vector()[0]
		cld.Dy = dir.Vector()[1]
	} else {
		cld.Dx += dir.Vector()[0]
		cld.Dy += dir.Vector()[1]
	}
}
