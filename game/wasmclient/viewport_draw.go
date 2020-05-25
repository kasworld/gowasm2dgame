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
	"github.com/kasworld/gowasm2dgame/enum/effecttype"
	"github.com/kasworld/gowasm2dgame/enum/gameobjtype"
	"github.com/kasworld/gowasm2dgame/enum/teamtype"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_obj"
)

func (vp *Viewport) draw(now int64) {
	si := vp.stageInfo
	if si == nil {
		return
	}
	si.Background.Move(now)
	for _, bt := range si.Teams {
		bt.Ball.MoveStraight(now)
		for _, v := range bt.Objs {
			switch v.GOType {
			default:
			case gameobjtype.Bullet, gameobjtype.SuperBullet:
				v.MoveStraight(now)
			case gameobjtype.Shield, gameobjtype.SuperShield:
				v.MoveCircular(now, bt.Ball.PosVt)
			case gameobjtype.HommingShield:
				v.MoveHomming(now, bt.Ball.PosVt)
			case gameobjtype.HommingBullet:
			}
		}
	}
	for _, eff := range si.Effects {
		eff.Move(now)
	}
	for _, cld := range si.Clouds {
		cld.Move(now)
	}
	vp.drawBG()
	for _, v := range si.Teams {
		vp.drawTeam(v, now)
	}
	for _, v := range si.Effects {
		vp.drawEffect(v, now)
	}
	for _, v := range si.Clouds {
		vp.drawCloud(v)
	}
}

func (vp *Viewport) drawBG() {
	si := vp.stageInfo
	sp := gSprites.BGSprite
	x, y := si.Background.PosVt[0], si.Background.PosVt[1]
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

func (vp *Viewport) drawCloud(cld *w2d_obj.Cloud) {
	x, y := cld.PosVt[0], cld.PosVt[1]
	sp := gSprites.CloudSprite
	srcx, srcy := sp.GetSliceXY(cld.SpriteNum)
	vp.context2d.Call("drawImage", sp.ImgCanvas,
		srcx, srcy, sp.W, sp.H,
		x-sp.W/2, y-sp.H/2, sp.W, sp.H,
	)
}

func (vp *Viewport) drawEffect(eff *w2d_obj.Effect, now int64) {
	x, y := eff.PosVt[0], eff.PosVt[1]
	sp := gSprites.EffectSprite[eff.EffectType]
	lifeTick := int(effecttype.Attrib[eff.EffectType].LifeTick)

	frame := int(now-eff.BirthTick) * sp.GetSliceCount() / lifeTick

	srcx, srcy := sp.GetSliceXY(frame)
	vp.context2d.Call("drawImage", sp.ImgCanvas,
		srcx, srcy, sp.W, sp.H,
		x-sp.W/2, y-sp.H/2, sp.W, sp.H,
	)
}

func (vp *Viewport) drawTeam(bl *w2d_obj.Team, now int64) {
	vp.drawGameObj(bl.TeamType, bl.Ball, now)
	for _, v := range bl.Objs {
		vp.drawGameObj(bl.TeamType, v, now)
	}
}

func (vp *Viewport) drawGameObj(
	teamtype teamtype.TeamType, v *w2d_obj.GameObj, now int64) {
	dispSize := gameobjtype.Attrib[v.GOType].Radius
	sp := gSprites.BallSprites[teamtype][v.GOType]
	frame := CalcCurrentFrame(now-v.BirthTick,
		gameobjtype.Attrib[v.GOType].FramePerSec,
	)
	srcx, srcy := sp.GetSliceXY(frame)
	vp.context2d.Call("drawImage", sp.ImgCanvas,
		srcx, srcy, dispSize, dispSize,
		v.PosVt[0]-dispSize/2, v.PosVt[1]-dispSize/2, dispSize, dispSize,
	)
}
