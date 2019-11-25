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
	"syscall/js"

	"github.com/kasworld/direction"
	"github.com/kasworld/gowasm2dgame/enums/gameobjtype"
	"github.com/kasworld/gowasm2dgame/enums/teamtype"
	"github.com/kasworld/gowasm2dgame/lib/anglemove"
	"github.com/kasworld/gowasm2dgame/lib/posacc"
	"github.com/kasworld/wrapper"
)

type SuperShield struct {
	Am    anglemove.AngleMove
	frame int
}

type Shield struct {
	Am anglemove.AngleMove
}

type HommingShield struct {
	Pa posacc.PosAcc
}

type Bullet struct {
	Pa posacc.PosAcc
}

type HommingBullet struct {
	Pa    posacc.PosAcc
	DstID int
}

type SuperBullet struct {
	Pa posacc.PosAcc
}

type BallTeam struct {
	TeamType teamtype.TeamType

	shiels        []*Shield
	superShields  []*SuperShield
	hommingShiels []*HommingShield

	bgXWrap func(i int) int
	bgYWrap func(i int) int
	BorderW int
	BorderH int

	Pa posacc.PosAcc
}

func NewBallTeam(
	TeamType teamtype.TeamType,
	initdir direction.Direction_Type,
	x, y int,
	w, h int,
) *BallTeam {
	bl := &BallTeam{
		TeamType: TeamType,
		BorderW:  w,
		BorderH:  h,
	}
	bl.bgXWrap = wrapper.New(w).GetWrapSafeFn()
	bl.bgYWrap = wrapper.New(h).GetWrapSafeFn()
	bl.Pa = posacc.PosAcc{
		X: x,
		Y: y,
	}
	bl.Pa.SetDir(initdir)

	bl.shiels = make([]*Shield, 24)
	for i := range bl.shiels {
		av := 1
		if i%2 == 0 {
			av = -1
		}
		bl.shiels[i] = &Shield{
			Am: anglemove.AngleMove{
				Angle:  i * 15,
				AngleV: av,
			},
		}
	}

	bl.superShields = make([]*SuperShield, 24)
	for i := range bl.superShields {
		av := 1
		if i%2 == 0 {
			av = -1
		}
		bl.superShields[i] = &SuperShield{
			Am: anglemove.AngleMove{
				Angle:  15 + i*15,
				AngleV: av,
			},
			frame: i * 3,
		}
	}
	return bl
}

func (bl *BallTeam) DrawTo(ctx js.Value) {
	bl.Pa.Move()
	for _, v := range bl.shiels {
		v.Am.Move()
	}
	for _, v := range bl.superShields {
		v.Am.Move()
	}

	bl.Pa.BounceNormalize(bl.BorderW, bl.BorderH)
	dispSize := gameobjtype.Attrib[gameobjtype.Ball].Size
	sp := gSprites.BallSprites[bl.TeamType][gameobjtype.Ball]
	srcx, srcy := sp.GetSliceXY(0)
	dstx, dsty := bl.Pa.X-dispSize/2, bl.Pa.Y-dispSize/2
	ctx.Call("drawImage", sp.ImgCanvas,
		srcx, srcy, dispSize, dispSize,
		dstx, dsty, dispSize, dispSize,
	)
	dispSize = gameobjtype.Attrib[gameobjtype.Shield].Size
	dispR := gameobjtype.Attrib[gameobjtype.Shield].R
	sp = gSprites.BallSprites[bl.TeamType][gameobjtype.Shield]
	for _, v := range bl.shiels {
		srcx, srcy := sp.GetSliceXY(0)
		x, y := calcCircularPos(bl.Pa.X, bl.Pa.Y, v.Am.Angle, dispR)
		ctx.Call("drawImage", sp.ImgCanvas,
			srcx, srcy, dispSize, dispSize,
			x-dispSize/2, y-dispSize/2, dispSize, dispSize,
		)
	}
	dispSize = gameobjtype.Attrib[gameobjtype.SuperShield].Size
	dispR = gameobjtype.Attrib[gameobjtype.SuperShield].R
	sp = gSprites.BallSprites[bl.TeamType][gameobjtype.SuperShield]
	for _, v := range bl.superShields {
		v.frame++
		srcx, srcy := sp.GetSliceXY(v.frame)
		x, y := calcCircularPos(bl.Pa.X, bl.Pa.Y, v.Am.Angle, dispR)
		ctx.Call("drawImage", sp.ImgCanvas,
			srcx, srcy, dispSize, dispSize,
			x-dispSize/2, y-dispSize/2, dispSize, dispSize,
		)
	}
}
