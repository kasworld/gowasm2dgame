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

	"github.com/kasworld/gowasm2dgame/enum/gameobjtype"
	"github.com/kasworld/gowasm2dgame/lib/imagecanvas"
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
	srcImageID string,
	xn, yn int) *Sprite {

	imgcnv := imagecanvas.NewByID(srcImageID)
	imgcnv.Cnv.Set("width", imgcnv.W)
	imgcnv.Cnv.Set("height", imgcnv.H)
	imgcnv.Ctx.Call("clearRect", 0, 0, imgcnv.W, imgcnv.H)
	imgcnv.Ctx.Call("drawImage", imgcnv.Img, 0, 0)
	return &Sprite{
		W:         float64(imgcnv.W) / float64(xn),
		H:         float64(imgcnv.H) / float64(yn),
		XCount:    xn,
		YCount:    yn,
		ImgCanvas: imgcnv.Cnv,
		ImgCtx:    imgcnv.Ctx,
	}
}

func LoadSpriteXYNResize(
	srcImageID string,
	xn, yn int,
	xSliceW, ySliceH float64,
) *Sprite {
	imgcnv := imagecanvas.NewByID(srcImageID)
	dstW := float64(xn) * xSliceW
	dstH := float64(yn) * ySliceH
	imgcnv.Cnv.Set("width", dstW)
	imgcnv.Cnv.Set("height", dstH)
	imgcnv.Ctx.Call("clearRect", 0, 0, dstW, dstH)
	imgcnv.Ctx.Call("drawImage", imgcnv.Img,
		0, 0, imgcnv.W, imgcnv.H,
		0, 0, dstW, dstH)
	return &Sprite{
		W:         xSliceW,
		H:         ySliceH,
		XCount:    xn,
		YCount:    yn,
		ImgCanvas: imgcnv.Cnv,
		ImgCtx:    imgcnv.Ctx,
	}
}

// LoadSpriteRotate load a image and make multi rotated image sptite
func LoadSpriteRotate(
	srcImageID string,
	start, end, step float64) *Sprite {

	imgcnv := imagecanvas.NewByID(srcImageID)

	xn := int((end - start) / step)
	imgcnv.Cnv.Set("width", imgcnv.W*xn)
	imgcnv.Cnv.Set("height", imgcnv.H)
	imgcnv.Ctx.Call("clearRect", 0, 0, imgcnv.W*xn, imgcnv.H)
	for i := 0; i < xn; i++ {
		imgcnv.Ctx.Call("save")
		imgcnv.Ctx.Call("translate", imgcnv.W*i+imgcnv.W/2, imgcnv.H/2)
		imgcnv.Ctx.Call("rotate", float64(i)*step*math.Pi/180)
		imgcnv.Ctx.Call("drawImage", imgcnv.Img, -imgcnv.W/2, -imgcnv.H/2)
		imgcnv.Ctx.Call("restore")
	}
	return &Sprite{
		W:         float64(imgcnv.W),
		H:         float64(imgcnv.H),
		ImgCanvas: imgcnv.Cnv,
		ImgCtx:    imgcnv.Ctx,
		XCount:    xn,
		YCount:    1,
	}
}

func LoadSpriteRotateResize(
	srcImageID string,
	start, end, step float64,
	xSliceW, ySliceH float64,
) *Sprite {

	imgcnv := imagecanvas.NewByID(srcImageID)
	xn := (end - start) / step
	yn := 1.0
	dstW := xn * xSliceW
	dstH := yn * ySliceH

	imgcnv.Cnv.Set("width", dstW)
	imgcnv.Cnv.Set("height", dstH)
	imgcnv.Ctx.Call("clearRect", 0, 0, dstW, dstH)
	for i := 0; i < int(xn); i++ {
		imgcnv.Ctx.Call("save")
		imgcnv.Ctx.Call("translate", xSliceW*float64(i)+xSliceW/2, ySliceH/2)
		imgcnv.Ctx.Call("rotate", float64(i)*step*math.Pi/180)
		imgcnv.Ctx.Call("drawImage", imgcnv.Img,
			0, 0, imgcnv.W, imgcnv.H,
			-xSliceW/2, -ySliceH/2, xSliceW, ySliceH)
		imgcnv.Ctx.Call("restore")
	}
	return &Sprite{
		W:         float64(xSliceW),
		H:         float64(ySliceH),
		ImgCanvas: imgcnv.Cnv,
		ImgCtx:    imgcnv.Ctx,
		XCount:    int(xn),
		YCount:    1,
	}
}

func LoadBallSprite(teamname string) [gameobjtype.GameObjType_Count]*Sprite {
	var rtn [gameobjtype.GameObjType_Count]*Sprite
	rtn[gameobjtype.Ball] = LoadSpriteXYNResize(
		"grayball",
		1, 1,
		gameobjtype.Attrib[gameobjtype.Ball].Size,
		gameobjtype.Attrib[gameobjtype.Ball].Size,
	)

	rtn[gameobjtype.Shield] = LoadSpriteXYNResize(
		"grayball",
		1, 1,
		gameobjtype.Attrib[gameobjtype.Shield].Size,
		gameobjtype.Attrib[gameobjtype.Shield].Size,
	)

	rtn[gameobjtype.SuperShield] = LoadSpriteRotateResize(
		"spiral",
		0, 360, 10,
		gameobjtype.Attrib[gameobjtype.SuperShield].Size,
		gameobjtype.Attrib[gameobjtype.SuperShield].Size,
	)
	rtn[gameobjtype.HommingShield] = LoadSpriteRotateResize(
		"spiral",
		0, 360, 10,
		gameobjtype.Attrib[gameobjtype.HommingShield].Size,
		gameobjtype.Attrib[gameobjtype.HommingShield].Size,
	)
	rtn[gameobjtype.Bullet] = LoadSpriteXYNResize(
		"grayball",
		1, 1,
		gameobjtype.Attrib[gameobjtype.Bullet].Size,
		gameobjtype.Attrib[gameobjtype.Bullet].Size,
	)
	rtn[gameobjtype.SuperBullet] = LoadSpriteRotateResize(
		"spiral",
		0, 360, 10,
		gameobjtype.Attrib[gameobjtype.SuperBullet].Size,
		gameobjtype.Attrib[gameobjtype.SuperBullet].Size,
	)
	rtn[gameobjtype.HommingBullet] = LoadSpriteRotateResize(
		"spiral",
		0, 360, 10,
		gameobjtype.Attrib[gameobjtype.HommingBullet].Size,
		gameobjtype.Attrib[gameobjtype.HommingBullet].Size,
	)
	return rtn
}
