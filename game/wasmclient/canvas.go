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

	"github.com/kasworld/gowasm2dgame/enums/effecttype"
	"github.com/kasworld/gowasm2dgame/enums/teamtype"

	"github.com/kasworld/gowasm2dgame/enums/gameobjtype"
	"github.com/kasworld/gowasm2dgame/game/gameconst"
	"github.com/kasworld/gowasm2dgame/lib/aniframe"
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

func (vp *Viewport2d) draw(now int64) {
	si := vp.stageInfo
	if si == nil {
		return
	}
	si.Background.Pa.Move(now)
	for _, bt := range si.Teams {
		bt.Ball.MoveStraight(now)
		for _, v := range bt.Objs {
			switch v.GOType {
			default:
			case gameobjtype.Bullet, gameobjtype.SuperBullet:
				v.MoveStraight(now)
			case gameobjtype.Shield, gameobjtype.SuperShield:
				v.MoveCircular(now, bt.Ball.X, bt.Ball.Y)
			case gameobjtype.HommingShield:
				v.MoveHomming(now, bt.Ball.X, bt.Ball.Y)
			case gameobjtype.HommingBullet:
			}
		}
	}
	for _, cld := range si.Clouds {
		cld.Pa.Move(now)
	}
	vp.drawBG()
	for _, v := range si.Teams {
		vp.drawBallTeam(v, now)
	}
	for _, v := range si.Effects {
		vp.drawEffect(v, now)
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

func (vp *Viewport2d) drawEffect(eff *w2d_obj.Effect, now int64) {
	x, y := eff.Pa.X, eff.Pa.Y
	sp := gSprites.EffectSprite[eff.EffectType]
	frame := aniframe.CalcCurrentFrame(now-eff.BirthTick,
		effecttype.Attrib[eff.EffectType].FramePerSec,
	)
	srcx, srcy := sp.GetSliceXY(frame)
	vp.context2d.Call("drawImage", sp.ImgCanvas,
		srcx, srcy, sp.W, sp.H,
		x-sp.W/2, y-sp.H/2, sp.W, sp.H,
	)
}

func (vp *Viewport2d) drawBallTeam(bl *w2d_obj.BallTeam, now int64) {
	vp.drawGameObj(bl.TeamType, bl.Ball, now)
	for _, v := range bl.Objs {
		vp.drawGameObj(bl.TeamType, v, now)
	}
}

func (vp *Viewport2d) drawGameObj(
	teamtype teamtype.TeamType, v *w2d_obj.GameObj, now int64) {
	dispSize := gameobjtype.Attrib[v.GOType].Size
	sp := gSprites.BallSprites[teamtype][v.GOType]
	frame := aniframe.CalcCurrentFrame(now-v.BirthTick,
		gameobjtype.Attrib[v.GOType].FramePerSec,
	)
	srcx, srcy := sp.GetSliceXY(frame)
	vp.context2d.Call("drawImage", sp.ImgCanvas,
		srcx, srcy, dispSize, dispSize,
		v.X-dispSize/2, v.Y-dispSize/2, dispSize, dispSize,
	)
}
