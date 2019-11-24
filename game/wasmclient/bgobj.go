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

type BGObj struct {
	sp *Sprite

	bgXWrap func(i int) int
	bgYWrap func(i int) int

	X  int
	Y  int
	Dx int
	Dy int
}

func NewBG() *BGObj {
	bg := &BGObj{}
	bg.sp = LoadSpriteXYN("background", "bgStore", 1, 1)
	bg.bgXWrap = wrapper.New(bg.sp.W).GetWrapSafeFn()
	bg.bgYWrap = wrapper.New(bg.sp.H).GetWrapSafeFn()
	bg.SetDxy(1, -1)
	return bg
}

func (bg *BGObj) DrawTo(ctx js.Value) {
	bg.Move()
	x := bg.bgXWrap(bg.X)
	y := bg.bgYWrap(bg.Y)
	bg.sp.drawImageSlice(ctx, x-bg.sp.W, y-bg.sp.H, 0)
	bg.sp.drawImageSlice(ctx, x-bg.sp.W, y, 0)
	bg.sp.drawImageSlice(ctx, x, y-bg.sp.H, 0)
	bg.sp.drawImageSlice(ctx, x, y, 0)
}

func (bg *BGObj) Move() {
	bg.X += bg.Dx
	bg.Y += bg.Dy
}

func (bg *BGObj) SetDxy(dx, dy int) {
	bg.Dx = dx
	bg.Dy = dy
}

func (bg *BGObj) SetDir(dir direction.Direction_Type) {
	bg.Dx = dir.Vector()[0]
	bg.Dy = dir.Vector()[1]
}

func (bg *BGObj) AccelerateDir(dir direction.Direction_Type) {
	if dir == direction.Dir_stop {
		bg.Dx = dir.Vector()[0]
		bg.Dy = dir.Vector()[1]
	} else {
		bg.Dx += dir.Vector()[0]
		bg.Dy += dir.Vector()[1]
	}
}
