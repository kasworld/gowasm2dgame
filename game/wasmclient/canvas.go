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
	"math/rand"
	"syscall/js"
	"time"

	"github.com/kasworld/gowasm2dgame/enums/gameobjtype"
	"github.com/kasworld/gowasm2dgame/game/gameconst"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_obj"
)

type Viewport2d struct {
	Canvas    js.Value
	context2d js.Value
	rnd       *rand.Rand

	W float64
	H float64

	stageInfo *w2d_obj.NotiStageInfo_data
}

func NewViewport2d() *Viewport2d {
	vp := &Viewport2d{
		rnd: rand.New(rand.NewSource(time.Now().UnixNano())),
		W:   gameconst.StageW,
		H:   gameconst.StageH,
	}

	vp.Canvas, vp.context2d = getCnv2dCtx("viewport2DCanvas")
	vp.Canvas.Set("width", vp.W)
	vp.Canvas.Set("height", vp.H)
	return vp
}

func (vp *Viewport2d) draw() {
	si := vp.stageInfo
	if si == nil {
		return
	}
	si.Background.Pa.Move()
	for _, bt := range si.Teams {
		bt.Ball.Pa.Move()
		for _, v := range bt.Shields {
			v.Am.Move()
		}
		for _, v := range bt.SuperShields {
			v.Am.Move()
		}
		for _, v := range bt.HommingShields {
			v.Pa.Move()
		}
		for _, v := range bt.Bullets {
			v.Pa.Move()
		}
		for _, v := range bt.SuperBullets {
			v.Pa.Move()
		}
		for _, v := range bt.HommingBullets {
			v.Pa.Move()
		}
	}
	for _, cld := range si.Clouds {
		cld.Pa.Move()
	}
	vp.drawBG()
	for _, v := range si.Teams {
		vp.drawBallTeam(v)
	}
	for _, v := range si.Clouds {
		vp.drawCloud(v)
	}
}

func (vp *Viewport2d) drawBG() {
	si := vp.stageInfo
	sp := gSprites.BGSprite
	x, y := si.Background.Pa.X, si.Background.Pa.Y
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

func (vp *Viewport2d) drawCloud(cld *w2d_obj.Cloud) {
	x, y := cld.Pa.X, cld.Pa.Y
	sp := gSprites.CloudSprite
	srcx, srcy := sp.GetSliceXY(cld.SpriteNum)
	vp.context2d.Call("drawImage", sp.ImgCanvas,
		srcx, srcy, sp.W, sp.H,
		x-sp.W/2, y-sp.H/2, sp.W, sp.H,
	)
}

func (vp *Viewport2d) drawBallTeam(bl *w2d_obj.BallTeam) {
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
		x, y := v.Am.CalcCircularPos(bl.Ball.Pa.X, bl.Ball.Pa.Y, dispR)
		vp.context2d.Call("drawImage", sp.ImgCanvas,
			srcx, srcy, dispSize, dispSize,
			x-dispSize/2, y-dispSize/2, dispSize, dispSize,
		)
	}

	dispSize = gameobjtype.Attrib[gameobjtype.SuperShield].Size
	dispR = gameobjtype.Attrib[gameobjtype.SuperShield].R
	sp = gSprites.BallSprites[bl.TeamType][gameobjtype.SuperShield]
	for _, v := range bl.SuperShields {
		srcx, srcy := sp.GetSliceXY(v.Frame)
		x, y := v.Am.CalcCircularPos(bl.Ball.Pa.X, bl.Ball.Pa.Y, dispR)
		vp.context2d.Call("drawImage", sp.ImgCanvas,
			srcx, srcy, dispSize, dispSize,
			x-dispSize/2, y-dispSize/2, dispSize, dispSize,
		)
	}

	dispSize = gameobjtype.Attrib[gameobjtype.HommingShield].Size
	sp = gSprites.BallSprites[bl.TeamType][gameobjtype.HommingShield]
	for _, v := range bl.HommingShields {
		srcx, srcy := sp.GetSliceXY(v.Frame)
		x, y := v.Pa.X, v.Pa.Y
		vp.context2d.Call("drawImage", sp.ImgCanvas,
			srcx, srcy, dispSize, dispSize,
			x-dispSize/2, y-dispSize/2, dispSize, dispSize,
		)
	}

	dispSize = gameobjtype.Attrib[gameobjtype.Bullet].Size
	sp = gSprites.BallSprites[bl.TeamType][gameobjtype.Bullet]
	for _, v := range bl.Bullets {
		srcx, srcy := sp.GetSliceXY(0)
		x, y := v.Pa.X, v.Pa.Y
		vp.context2d.Call("drawImage", sp.ImgCanvas,
			srcx, srcy, dispSize, dispSize,
			x-dispSize/2, y-dispSize/2, dispSize, dispSize,
		)
	}

	dispSize = gameobjtype.Attrib[gameobjtype.SuperBullet].Size
	sp = gSprites.BallSprites[bl.TeamType][gameobjtype.SuperBullet]
	for _, v := range bl.SuperBullets {
		srcx, srcy := sp.GetSliceXY(v.Frame)
		x, y := v.Pa.X, v.Pa.Y
		vp.context2d.Call("drawImage", sp.ImgCanvas,
			srcx, srcy, dispSize, dispSize,
			x-dispSize/2, y-dispSize/2, dispSize, dispSize,
		)
	}

	dispSize = gameobjtype.Attrib[gameobjtype.HommingBullet].Size
	sp = gSprites.BallSprites[bl.TeamType][gameobjtype.HommingBullet]
	for _, v := range bl.HommingBullets {
		srcx, srcy := sp.GetSliceXY(v.Frame)
		x, y := v.Pa.X, v.Pa.Y
		vp.context2d.Call("drawImage", sp.ImgCanvas,
			srcx, srcy, dispSize, dispSize,
			x-dispSize/2, y-dispSize/2, dispSize, dispSize,
		)
	}
}
