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
	"fmt"
	"syscall/js"

	"github.com/kasworld/gowasm2dgame/lib/stroll8way"
	"github.com/kasworld/wrapper"
)

type Viewport2d struct {
	Canvas    js.Value
	context2d js.Value

	W       int
	H       int
	bgXWrap func(i int) int
	bgYWrap func(i int) int

	background *Sprite
	clouds     *Sprite

	grayball *Sprite
	spiral   *Sprite

	explodeetc  *Sprite
	explodeball *Sprite
	spawn       *Sprite

	bgscroll stroll8way.Stroll8
}

func NewViewport2d() *Viewport2d {
	vp := &Viewport2d{}
	vp.W = 1000
	vp.H = 1000
	vp.bgXWrap = wrapper.New(vp.W).GetWrapSafeFn()
	vp.bgYWrap = wrapper.New(vp.H).GetWrapSafeFn()

	vp.Canvas, vp.context2d = getCnv2dCtx("viewport2DCanvas")
	if !vp.Canvas.Truthy() {
		fmt.Printf("fail to get canvas viewport2DCanvas\n")
	}
	if !vp.context2d.Truthy() {
		fmt.Printf("fail to get context2d\n")
	}
	vp.Canvas.Set("width", vp.W)
	vp.Canvas.Set("height", vp.H)

	vp.background = LoadSpriteXYN("background", "bgStore", 1, 1)
	vp.bgXWrap = wrapper.New(vp.background.W).GetWrapSafeFn()
	vp.bgYWrap = wrapper.New(vp.background.H).GetWrapSafeFn()

	vp.clouds = LoadSpriteXYN("clouds", "cloudStore", 1, 4)
	vp.spawn = LoadSpriteXYN("spawn", "spawnStore", 1, 6)
	vp.explodeetc = LoadSpriteXYN("explodeetc", "explodeetcStore", 1, 8)
	vp.explodeball = LoadSpriteXYN("explodeball", "explodeballStore", 8, 1)
	vp.spiral = LoadSpriteRotate("spiral", "spiralStore", 0, 360, 10)

	/*
		vp.grayball = vp.LoadImage("grayball") // color change
		vp.spiral = vp.LoadImage("spiral")     // color change, rotate (0, 360, 10)
		// ('bounceball', "grayball.png", None),
		// ('bullet', "grayball.png", None),
		// ('hommingbullet', "spiral.png", (0, 360, 10)),
		// ('superbullet', "spiral.png", (0, 360, 10)),
		// ('circularbullet', "grayball.png", None),
		// ('shield', "grayball.png", None),
		// ('supershield', "spiral.png", (0, 360, 10))

		vp.clouds = vp.LoadImage("clouds")           // slicearg=(1, 4
		vp.explodeetc = vp.LoadImage("explodeetc")   // slicearg=(1, 8, spriteexplosioneffect
		vp.explodeball = vp.LoadImage("explodeball") // slicearg=(8, 1 ballexplosioneffect
		vp.spawn = vp.LoadImage("spawn")             // slicearg=(1, 6, reverse spawneffect
	*/
	vp.bgscroll.SetDxy(1, -1)
	return vp
}

func (vp *Viewport2d) drawBG() {
	vp.bgscroll.Move()
	x := vp.bgXWrap(vp.bgscroll.X)
	y := vp.bgYWrap(vp.bgscroll.Y)
	sp := vp.background
	sp.DrawImageSlice(vp.context2d, x-sp.W, y-sp.H, 0)
	sp.DrawImageSlice(vp.context2d, x-sp.W, y, 0)
	sp.DrawImageSlice(vp.context2d, x, y-sp.H, 0)
	sp.DrawImageSlice(vp.context2d, x, y, 0)
}

func (vp *Viewport2d) drawObj() {
	vp.clouds.DrawImageSliceAlignCenter(vp.context2d, vp.W/2, vp.H/2, 0)
}

func (app *WasmClient) drawCanvas(this js.Value, args []js.Value) interface{} {
	defer func() {
		js.Global().Call("requestAnimationFrame", js.FuncOf(app.drawCanvas))
	}()
	dispCount := app.DispInterDur.GetCount()
	_ = dispCount
	act := app.DispInterDur.BeginAct()
	defer act.End()

	app.vp.drawBG()
	app.vp.drawObj()

	return nil
}
