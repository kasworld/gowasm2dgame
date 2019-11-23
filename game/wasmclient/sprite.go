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
	ImgCanvas js.Value
	ImgCtx    js.Value
	// one sprite slice size
	W int
	H int
	// image count x, y
	XCount int
	YCount int
}

// DrawImageSlice draw nth sprite image
func (sp *Sprite) DrawImageSlice(dstctx js.Value, dstx, dsty int, n int) {
	srcxn := n % sp.XCount
	srcyn := (n / sp.XCount) % sp.YCount
	dstctx.Call("drawImage",
		sp.ImgCanvas,
		sp.W*srcxn, sp.H*srcyn, sp.W, sp.H,
		dstx, dsty, sp.W, sp.H,
	)
}

// LoadSpriteXYN load multi image sprite
func LoadSpriteXYN(srcImageID string, dstCanvasID string, xn, yn int) *Sprite {
	img := js.Global().Get("document").Call("getElementById", srcImageID)
	if !img.Truthy() {
		fmt.Printf("fail to get %v", srcImageID)
		return nil
	}
	srcw := img.Get("naturalWidth").Int()
	srch := img.Get("naturalHeight").Int()

	dstcnv := js.Global().Get("document").Call("getElementById", dstCanvasID)
	if !dstcnv.Truthy() {
		fmt.Printf("fail to get canvas\n")
	}
	dstcnv.Set("width", srcw)
	dstcnv.Set("height", srch)
	dstctx := dstcnv.Call("getContext", "2d")
	if !dstctx.Truthy() {
		fmt.Printf("fail to get context\n")
	}
	dstctx.Set("imageSmoothingEnabled", false)
	dstctx.Call("clearRect", 0, 0, srcw, srch)
	dstctx.Call("drawImage", img, 0, 0)
	return &Sprite{
		W:         srcw / xn,
		H:         srch / yn,
		XCount:    xn,
		YCount:    yn,
		ImgCanvas: dstcnv,
		ImgCtx:    dstctx,
	}
}

// LoadSpriteRotate load a image and make multi rotated image sptite
func LoadSpriteRotate(srcImageID string, dstCanvasID string, start, end, step float64) *Sprite {
	dstcount := int((end - start) / step)

	img := js.Global().Get("document").Call("getElementById", srcImageID)
	if !img.Truthy() {
		fmt.Printf("fail to get %v", srcImageID)
		return nil
	}
	srcw := img.Get("naturalWidth").Int()
	srch := img.Get("naturalHeight").Int()

	dstcnv := js.Global().Get("document").Call("getElementById", dstCanvasID)
	if !dstcnv.Truthy() {
		fmt.Printf("fail to get canvas\n")
	}
	dstcnv.Set("width", srcw*dstcount)
	dstcnv.Set("height", srch)
	dstctx := dstcnv.Call("getContext", "2d")
	if !dstctx.Truthy() {
		fmt.Printf("fail to get context\n")
	}
	dstctx.Set("imageSmoothingEnabled", false)
	dstctx.Call("clearRect", 0, 0, srcw*dstcount, srch)
	for i := 0; i < dstcount; i++ {
		dstctx.Call("save")
		dstctx.Call("translate", srcw*i+srcw/2, srch/2)
		dstctx.Call("rotate", float64(i)*step*math.Pi/180)
		dstctx.Call("drawImage", img, -srcw/2, -srch/2)
		dstctx.Call("restore")
	}
	return &Sprite{
		W:         srcw,
		H:         srch,
		ImgCanvas: dstcnv,
		ImgCtx:    dstctx,
		XCount:    dstcount,
		YCount:    1,
	}
}
