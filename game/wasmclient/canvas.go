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
)

type Viewport2d struct {
	Canvas    js.Value
	context2d js.Value

	background js.Value
}

func NewViewport2d(cnvid string) *Viewport2d {
	vp := &Viewport2d{}

	vp.Canvas = js.Global().Get("document").Call("getElementById", cnvid)
	if !vp.Canvas.Truthy() {
		fmt.Printf("fail to get canvas %v\n", cnvid)
	}
	vp.context2d = vp.Canvas.Call("getContext", "2d")
	vp.context2d.Set("imageSmoothingEnabled", false)
	if !vp.context2d.Truthy() {
		fmt.Printf("fail to get context2d context2d\n")
	}
	vp.background = vp.LoadImage("background")
	return vp
}

func (vp *Viewport2d) LoadImage(name string) js.Value {
	img := js.Global().Get("document").Call("getElementById", name)
	if !img.Truthy() {
		fmt.Printf("fail to get %v", name)
		return js.Null()
	}
	return img
}

func (vp *Viewport2d) DrawImage(img js.Value, dstx, dsty int) {
	vp.context2d.Call("drawImage", img, dstx, dsty)
}

func (app *WasmClient) drawCanvas(this js.Value, args []js.Value) interface{} {
	defer func() {
		js.Global().Call("requestAnimationFrame", js.FuncOf(app.drawCanvas))
	}()

	app.vp2d.DrawImage(app.vp2d.background, 0, 0)
	return nil
}
