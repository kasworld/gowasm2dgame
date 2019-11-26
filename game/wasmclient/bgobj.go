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

	"github.com/kasworld/gowasm2dgame/lib/posacc"
	"github.com/kasworld/wrapper"
)

type BGObj struct {
	sp *Sprite

	bgXWrap func(i int) int
	bgYWrap func(i int) int

	Pa posacc.PosAcc
}

func NewBG() *BGObj {
	bg := &BGObj{}
	bg.sp = LoadSpriteXYN("background", "bgStore", 1, 1)
	bg.bgXWrap = wrapper.New(bg.sp.W).GetWrapSafeFn()
	bg.bgYWrap = wrapper.New(bg.sp.H).GetWrapSafeFn()
	bg.Pa.SetDxy(1, -1)
	return bg
}

func (bg *BGObj) DrawTo(ctx js.Value) {
	bg.Pa.Move()
	x := bg.bgXWrap(bg.Pa.X)
	y := bg.bgYWrap(bg.Pa.Y)

	srcx, srcy := bg.sp.GetSliceXY(0)
	ctx.Call("drawImage", bg.sp.ImgCanvas,
		srcx, srcy, bg.sp.W, bg.sp.H,
		x-bg.sp.W, y-bg.sp.H, bg.sp.W, bg.sp.H,
	)
	ctx.Call("drawImage", bg.sp.ImgCanvas,
		srcx, srcy, bg.sp.W, bg.sp.H,
		x-bg.sp.W, y, bg.sp.W, bg.sp.H,
	)
	ctx.Call("drawImage", bg.sp.ImgCanvas,
		srcx, srcy, bg.sp.W, bg.sp.H,
		x, y-bg.sp.H, bg.sp.W, bg.sp.H,
	)
	ctx.Call("drawImage", bg.sp.ImgCanvas,
		srcx, srcy, bg.sp.W, bg.sp.H,
		x, y, bg.sp.W, bg.sp.H,
	)
}
