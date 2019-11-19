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
	"math"
	"syscall/js"
)

type Sprite struct {
	W       int
	H       int
	ImgData js.Value
}

func LoadSprite(name string, angle float64) *Sprite {
	img := js.Global().Get("document").Call("getElementById", name)
	if !img.Truthy() {
		fmt.Printf("fail to get %v", name)
		return nil
	}
	w := img.Get("naturalWidth")
	h := img.Get("naturalHeight")
	cnv := js.Global().Get("document").Call("getElementById", "hiddenCanvas")
	if !cnv.Truthy() {
		fmt.Printf("fail to get canvas\n")
	}
	cnv.Set("width", w)
	cnv.Set("height", h)
	ctx := cnv.Call("getContext", "2d")
	if !ctx.Truthy() {
		fmt.Printf("fail to get context\n")
	}
	ctx.Set("imageSmoothingEnabled", false)
	ctx.Call("save")
	ctx.Call("clearRect", 0, 0, w, h)
	ctx.Call("translate", w.Int()/2, h.Int()/2)
	ctx.Call("rotate", angle*math.Pi/180)
	ctx.Call("drawImage", img, -w.Int()/2, -h.Int()/2)
	ctx.Call("restore")
	imgData := ctx.Call("getImageData", 0, 0, w, h)
	return &Sprite{
		W:       w.Int(),
		H:       h.Int(),
		ImgData: imgData,
	}
}

func (sp *Sprite) PutImageData(dst js.Value, x, y int) {
	dst.Call("putImageData", sp.ImgData, x, y)
}
