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

	W         int
	H         int
	XWrapper  *wrapper.Wrapper `prettystring:"simple"`
	YWrapper  *wrapper.Wrapper `prettystring:"simple"`
	XWrapSafe func(i int) int
	YWrapSafe func(i int) int

	background *Sprite
	clouds     *Sprite

	grayball *Sprite
	spiral   *Sprite

	explodeetc  *Sprite
	explodeball *Sprite
	spawn       *Sprite

	scroll stroll8way.Stroll8
}

func NewViewport2d() *Viewport2d {
	vp := &Viewport2d{}
	vp.W = 1000
	vp.H = 1000
	vp.XWrapper = wrapper.New(vp.W)
	vp.YWrapper = wrapper.New(vp.H)
	vp.XWrapSafe = vp.XWrapper.GetWrapSafeFn()
	vp.YWrapSafe = vp.YWrapper.GetWrapSafeFn()

	vp.Canvas = js.Global().Get("document").Call("getElementById", "viewport2DCanvas")
	if !vp.Canvas.Truthy() {
		fmt.Printf("fail to get canvas viewport2DCanvas\n")
	}
	vp.Canvas.Set("width", vp.W)
	vp.Canvas.Set("height", vp.H)

	vp.context2d = vp.Canvas.Call("getContext", "2d")
	vp.context2d.Set("imageSmoothingEnabled", false)
	if !vp.context2d.Truthy() {
		fmt.Printf("fail to get context2d\n")
	}

	vp.background = LoadSpriteXYN("background", "bgStore", 1, 1)
	vp.clouds = LoadSpriteXYN("clouds", "cloudStore", 1, 4)
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
	vp.scroll.SetDxy(1, -1)
	return vp
}

func (app *WasmClient) drawCanvas(this js.Value, args []js.Value) interface{} {
	defer func() {
		js.Global().Call("requestAnimationFrame", js.FuncOf(app.drawCanvas))
	}()
	dispCount := app.DispInterDur.GetCount()
	_ = dispCount
	act := app.DispInterDur.BeginAct()
	defer act.End()

	app.vp.scroll.Move()

	x := app.vp.XWrapSafe(app.vp.scroll.X)
	y := app.vp.YWrapSafe(app.vp.scroll.Y)
	_ = x
	_ = y
	app.vp.background.DrawImageSlice(app.vp.context2d, x-app.vp.W, y-app.vp.H, 0)
	app.vp.clouds.DrawImageSlice(app.vp.context2d, 4, 9, 0)

	return nil
}
