// Copyright 2015,2016,2017,2018,2019,2020 SeukWon Kang (kasworld@gmail.com)
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

	"github.com/kasworld/gowasm2dgame/enums/gameobjtype"
)

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

type Sprite struct {
	ImgCanvas js.Value
	ImgCtx    js.Value
	// one sprite slice size
	W float64
	H float64
	// image count x, y
	XCount int
	YCount int
}

func (sp *Sprite) GetSliceCount() int {
	return sp.XCount * sp.YCount
}

// GetSliceXY return nth slice pos
func (sp *Sprite) GetSliceXY(n int) (float64, float64) {
	srcxn := n % sp.XCount
	srcyn := (n / sp.XCount) % sp.YCount
	return sp.W * float64(srcxn), sp.H * float64(srcyn)
}
func (sp *Sprite) CalcAlignDstTopLeft(dstx, dsty float64) (float64, float64) {
	return dstx, dsty
}

func (sp *Sprite) CalcAlignDstCenter(dstx, dsty float64) (float64, float64) {
	return dstx - sp.W/2, dsty - sp.H/2
}

func (sp *Sprite) Filter(index int, value int) {
	imgData := sp.ImgCtx.Call("getImageData",
		0, 0, sp.W*float64(sp.XCount), sp.H*float64(sp.YCount))
	pixels := imgData.Get("data") // Uint8ClampedArray
	l := int(sp.W) * sp.XCount * int(sp.H) * sp.YCount
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
		W:         srcw / float64(xn),
		H:         srch / float64(yn),
		XCount:    xn,
		YCount:    yn,
		ImgCanvas: dstcnv,
		ImgCtx:    dstctx,
	}
}

func LoadSpriteXYNResize(
	srcImageID string, dstCanvasID string,
	xn, yn int,
	xSliceW, ySliceH float64,
) *Sprite {
	img, srcw, srch := getImgWH(srcImageID)
	dstcnv, dstctx := getCnv2dCtx(dstCanvasID)

	dstW := float64(xn) * xSliceW
	dstH := float64(yn) * ySliceH
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
	dstcnv.Set("width", srcw*float64(xn))
	dstcnv.Set("height", srch)
	dstctx.Call("clearRect", 0, 0, srcw*float64(xn), srch)
	for i := 0; i < xn; i++ {
		dstctx.Call("save")
		dstctx.Call("translate", srcw*float64(i)+srcw/2, srch/2)
		dstctx.Call("rotate", float64(i)*step*math.Pi/180)
		dstctx.Call("drawImage", img, -srcw/2, -srch/2)
		dstctx.Call("restore")
	}
	return &Sprite{
		W:         float64(srcw),
		H:         float64(srch),
		ImgCanvas: dstcnv,
		ImgCtx:    dstctx,
		XCount:    xn,
		YCount:    1,
	}
}

func LoadSpriteRotateResize(
	srcImageID string, dstCanvasID string,
	start, end, step float64,
	xSliceW, ySliceH float64,
) *Sprite {
	img, srcw, srch := getImgWH(srcImageID)
	dstcnv, dstctx := getCnv2dCtx(dstCanvasID)
	_ = srcw
	_ = srch
	xn := (end - start) / step
	yn := 1.0
	dstW := xn * xSliceW
	dstH := yn * ySliceH

	dstcnv.Set("width", dstW)
	dstcnv.Set("height", dstH)
	dstctx.Call("clearRect", 0, 0, dstW, dstH)
	for i := 0; i < int(xn); i++ {
		dstctx.Call("save")
		dstctx.Call("translate", xSliceW*float64(i)+xSliceW/2, ySliceH/2)
		dstctx.Call("rotate", float64(i)*step*math.Pi/180)
		dstctx.Call("drawImage", img,
			0, 0, srcw, srch,
			-xSliceW/2, -ySliceH/2, xSliceW, ySliceH)
		dstctx.Call("restore")
	}
	return &Sprite{
		W:         float64(xSliceW),
		H:         float64(ySliceH),
		ImgCanvas: dstcnv,
		ImgCtx:    dstctx,
		XCount:    int(xn),
		YCount:    1,
	}
}

func LoadBallSprite(teamname string) [gameobjtype.GameObjType_Count]*Sprite {
	var rtn [gameobjtype.GameObjType_Count]*Sprite
	rtn[gameobjtype.Ball] = LoadSpriteXYNResize(
		"grayball", teamname+"_ball",
		1, 1,
		gameobjtype.Attrib[gameobjtype.Ball].Size,
		gameobjtype.Attrib[gameobjtype.Ball].Size,
	)

	rtn[gameobjtype.Shield] = LoadSpriteXYNResize(
		"grayball", teamname+"_shield",
		1, 1,
		gameobjtype.Attrib[gameobjtype.Shield].Size,
		gameobjtype.Attrib[gameobjtype.Shield].Size,
	)

	rtn[gameobjtype.SuperShield] = LoadSpriteRotateResize(
		"spiral", teamname+"_supershield",
		0, 360, 10,
		gameobjtype.Attrib[gameobjtype.SuperShield].Size,
		gameobjtype.Attrib[gameobjtype.SuperShield].Size,
	)
	rtn[gameobjtype.HommingShield] = LoadSpriteRotateResize(
		"spiral", teamname+"_hommingshield",
		0, 360, 10,
		gameobjtype.Attrib[gameobjtype.HommingShield].Size,
		gameobjtype.Attrib[gameobjtype.HommingShield].Size,
	)
	rtn[gameobjtype.Bullet] = LoadSpriteXYNResize(
		"grayball", teamname+"_bullet",
		1, 1,
		gameobjtype.Attrib[gameobjtype.Bullet].Size,
		gameobjtype.Attrib[gameobjtype.Bullet].Size,
	)
	rtn[gameobjtype.SuperBullet] = LoadSpriteRotateResize(
		"spiral", teamname+"_superbullet",
		0, 360, 10,
		gameobjtype.Attrib[gameobjtype.SuperBullet].Size,
		gameobjtype.Attrib[gameobjtype.SuperBullet].Size,
	)
	rtn[gameobjtype.HommingBullet] = LoadSpriteRotateResize(
		"spiral", teamname+"_hommingbullet",
		0, 360, 10,
		gameobjtype.Attrib[gameobjtype.HommingBullet].Size,
		gameobjtype.Attrib[gameobjtype.HommingBullet].Size,
	)
	return rtn
}
