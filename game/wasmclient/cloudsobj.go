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
	"github.com/kasworld/gowasm2dgame/lib/posacc"
	"github.com/kasworld/wrapper"
)

type Cloud struct {
	sp  *Sprite
	spn int

	bgXWrap func(i int) int
	bgYWrap func(i int) int

	Pa posacc.PosAcc
}

func NewCloud(sp *Sprite, spn int,
	initdir direction.Direction_Type,
	x, y int,
	w, h int,
) *Cloud {
	cld := &Cloud{
		sp:  sp,
		spn: spn,
	}
	cld.bgXWrap = wrapper.New(w).GetWrapSafeFn()
	cld.bgYWrap = wrapper.New(h).GetWrapSafeFn()
	cld.Pa = posacc.PosAcc{
		X: x,
		Y: y,
	}
	cld.Pa.SetDir(initdir)
	return cld
}

func (cld *Cloud) DrawTo(ctx js.Value) {
	cld.Pa.Move()
	x := cld.bgXWrap(cld.Pa.X)
	y := cld.bgYWrap(cld.Pa.Y)
	srcx, srcy := cld.sp.GetSliceXY(cld.spn)
	ctx.Call("drawImage", cld.sp.ImgCanvas,
		srcx, srcy, cld.sp.W, cld.sp.H,
		x-cld.sp.W/2, y-cld.sp.H/2, cld.sp.W, cld.sp.H,
	)
}
