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

func (sp *Sprite) GetSliceCount() int {
	return sp.XCount * sp.YCount
}

// GetSliceXY return nth slice pos
func (sp *Sprite) GetSliceXY(n int) (int, int) {
	srcxn := n % sp.XCount
	srcyn := (n / sp.XCount) % sp.YCount
	return sp.W * srcxn, sp.H * srcyn
}
func (sp *Sprite) CalcAlignDstTopLeft(dstx, dsty int) (int, int) {
	return dstx, dsty
}

func (sp *Sprite) CalcAlignDstCenter(dstx, dsty int) (int, int) {
	return dstx - sp.W/2, dsty - sp.H/2
}

func (sp *Sprite) Filter(index int, value int) {
	imgData := sp.ImgCtx.Call("getImageData", 0, 0, sp.W*sp.XCount, sp.H*sp.YCount)
	pixels := imgData.Get("data") // Uint8ClampedArray
	l := sp.W * sp.XCount * sp.H * sp.YCount
	for i := 0; i < l; i++ {
		pixels.SetIndex(i*4+index, value)
	}
	sp.ImgCtx.Call("putImageData", imgData, 0, 0)
}

// LoadSpriteXYN load multi image sprite
func LoadSpriteXYN(
	srcImageID string, dstCanvasID string,
	xn, yn int) *Sprite {
	img, srcw, srch := getImgWH(srcImageID)
	dstcnv, dstctx := getCnv2dCtx(dstCanvasID)

	dstcnv.Set("width", srcw)
	dstcnv.Set("height", srch)
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

func LoadSpriteXYNResize(
	srcImageID string, dstCanvasID string,
	xn, yn int,
	xSliceW, ySliceH int,
) *Sprite {
	img, srcw, srch := getImgWH(srcImageID)
	dstcnv, dstctx := getCnv2dCtx(dstCanvasID)

	dstW := xn * xSliceW
	dstH := yn * ySliceH
	dstcnv.Set("width", dstW)
	dstcnv.Set("height", dstH)
	dstctx.Call("clearRect", 0, 0, dstW, dstH)
	dstctx.Call("drawImage", img,
		0, 0, srcw, srch,
		0, 0, dstW, dstH)
	return &Sprite{
		W:         xSliceW,
		H:         ySliceH,
		XCount:    xn,
		YCount:    yn,
		ImgCanvas: dstcnv,
		ImgCtx:    dstctx,
	}
}

// LoadSpriteRotate load a image and make multi rotated image sptite
func LoadSpriteRotate(
	srcImageID string, dstCanvasID string,
	start, end, step float64) *Sprite {
	img, srcw, srch := getImgWH(srcImageID)
	dstcnv, dstctx := getCnv2dCtx(dstCanvasID)

	xn := int((end - start) / step)
	dstcnv.Set("width", srcw*xn)
	dstcnv.Set("height", srch)
	dstctx.Call("clearRect", 0, 0, srcw*xn, srch)
	for i := 0; i < xn; i++ {
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
		XCount:    xn,
		YCount:    1,
	}
}

func LoadSpriteRotateResize(
	srcImageID string, dstCanvasID string,
	start, end, step float64,
	xSliceW, ySliceH int,
) *Sprite {
	img, srcw, srch := getImgWH(srcImageID)
	dstcnv, dstctx := getCnv2dCtx(dstCanvasID)
	_ = srcw
	_ = srch
	xn := int((end - start) / step)
	yn := 1
	dstW := xn * xSliceW
	dstH := yn * ySliceH

	dstcnv.Set("width", dstW)
	dstcnv.Set("height", dstH)
	dstctx.Call("clearRect", 0, 0, dstW, dstH)
	for i := 0; i < xn; i++ {
		dstctx.Call("save")
		dstctx.Call("translate", xSliceW*i+xSliceW/2, ySliceH/2)
		dstctx.Call("rotate", float64(i)*step*math.Pi/180)
		dstctx.Call("drawImage", img,
			0, 0, srcw, srch,
			-xSliceW/2, -ySliceH/2, xSliceW, ySliceH)
		dstctx.Call("restore")
	}
	return &Sprite{
		W:         xSliceW,
		H:         ySliceH,
		ImgCanvas: dstcnv,
		ImgCtx:    dstctx,
		XCount:    xn,
		YCount:    1,
	}
}
