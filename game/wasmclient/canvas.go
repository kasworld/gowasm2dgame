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
	"math/rand"
	"syscall/js"
	"time"

	"github.com/kasworld/gowasm2dgame/enums/gameobjtype"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_obj"
	"github.com/kasworld/wrapper"
)

type Viewport2d struct {
	Canvas    js.Value
	context2d js.Value
	rnd       *rand.Rand

	W     int
	H     int
	XWrap func(i int) int
	YWrap func(i int) int

	bgObj     *w2d_obj.Background
	ballTeams []*w2d_obj.BallTeam
	effects   []*w2d_obj.Effect
	cloudObjs []*w2d_obj.Cloud
}

func NewViewport2d() *Viewport2d {
	vp := &Viewport2d{
		rnd: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
	vp.W = 1000
	vp.H = 1000
	vp.XWrap = wrapper.New(vp.W).GetWrapSafeFn()
	vp.YWrap = wrapper.New(vp.H).GetWrapSafeFn()

	vp.Canvas, vp.context2d = getCnv2dCtx("viewport2DCanvas")
	if !vp.context2d.Truthy() {
		fmt.Printf("fail to get context2d\n")
	}
	vp.Canvas.Set("width", vp.W)
	vp.Canvas.Set("height", vp.H)

	vp.makeTestObj()
	return vp
}

func (vp *Viewport2d) move() {
	vp.bgObj.Pa.Move()
	for _, bt := range vp.ballTeams {
		bt.Ball.Pa.Move()
		bt.Ball.Pa.BounceNormalize(vp.W, vp.H)
		for _, sld := range bt.Shields {
			sld.Am.Move()
		}
		for _, ssld := range bt.SuperShields {
			ssld.Am.Move()
		}
	}
	for _, cld := range vp.cloudObjs {
		cld.Pa.Move()
	}
}

func (vp *Viewport2d) draw(frame int) {
	vp.DrawBG()
	for _, v := range vp.ballTeams {
		vp.DrawBallTeam(v, frame)
	}
	for _, v := range vp.cloudObjs {
		vp.DrawCloud(v)
	}
}

func (vp *Viewport2d) DrawBG() {
	x := gSprites.BGXWrap(vp.bgObj.Pa.X)
	y := gSprites.BGYWrap(vp.bgObj.Pa.Y)

	sp := gSprites.BGSprite
	srcx, srcy := sp.GetSliceXY(0)
	vp.context2d.Call("drawImage", sp.ImgCanvas,
		srcx, srcy, sp.W, sp.H,
		x-sp.W, y-sp.H, sp.W, sp.H,
	)
	vp.context2d.Call("drawImage", sp.ImgCanvas,
		srcx, srcy, sp.W, sp.H,
		x-sp.W, y, sp.W, sp.H,
	)
	vp.context2d.Call("drawImage", sp.ImgCanvas,
		srcx, srcy, sp.W, sp.H,
		x, y-sp.H, sp.W, sp.H,
	)
	vp.context2d.Call("drawImage", sp.ImgCanvas,
		srcx, srcy, sp.W, sp.H,
		x, y, sp.W, sp.H,
	)
}

func (vp *Viewport2d) DrawCloud(cld *w2d_obj.Cloud) {
	x := vp.XWrap(cld.Pa.X)
	y := vp.YWrap(cld.Pa.Y)
	sp := gSprites.CloudSprite
	srcx, srcy := sp.GetSliceXY(cld.SpriteNum)
	vp.context2d.Call("drawImage", sp.ImgCanvas,
		srcx, srcy, sp.W, sp.H,
		x-sp.W/2, y-sp.H/2, sp.W, sp.H,
	)
}

func (vp *Viewport2d) DrawBallTeam(bl *w2d_obj.BallTeam, frame int) {
	dispSize := gameobjtype.Attrib[gameobjtype.Ball].Size
	sp := gSprites.BallSprites[bl.TeamType][gameobjtype.Ball]
	srcx, srcy := sp.GetSliceXY(0)
	dstx, dsty := bl.Ball.Pa.X-dispSize/2, bl.Ball.Pa.Y-dispSize/2
	vp.context2d.Call("drawImage", sp.ImgCanvas,
		srcx, srcy, dispSize, dispSize,
		dstx, dsty, dispSize, dispSize,
	)

	dispSize = gameobjtype.Attrib[gameobjtype.Shield].Size
	dispR := gameobjtype.Attrib[gameobjtype.Shield].R
	sp = gSprites.BallSprites[bl.TeamType][gameobjtype.Shield]
	srcx, srcy = sp.GetSliceXY(0)
	for _, v := range bl.Shields {
		x, y := calcCircularPos(bl.Ball.Pa.X, bl.Ball.Pa.Y, v.Am.Angle, dispR)
		vp.context2d.Call("drawImage", sp.ImgCanvas,
			srcx, srcy, dispSize, dispSize,
			x-dispSize/2, y-dispSize/2, dispSize, dispSize,
		)
	}

	dispSize = gameobjtype.Attrib[gameobjtype.SuperShield].Size
	dispR = gameobjtype.Attrib[gameobjtype.SuperShield].R
	sp = gSprites.BallSprites[bl.TeamType][gameobjtype.SuperShield]
	for _, v := range bl.SuperShields {
		srcx, srcy := sp.GetSliceXY(frame)
		x, y := calcCircularPos(bl.Ball.Pa.X, bl.Ball.Pa.Y, v.Am.Angle, dispR)
		vp.context2d.Call("drawImage", sp.ImgCanvas,
			srcx, srcy, dispSize, dispSize,
			x-dispSize/2, y-dispSize/2, dispSize, dispSize,
		)
	}
}
