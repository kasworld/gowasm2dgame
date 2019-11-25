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
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_obj"
)

func NewBallTeam(TeamType teamtype.TeamType,
	initdir direction.Direction_Type, x, y int,
) *w2d_obj.BallTeam {
	bl := &w2d_obj.BallTeam{
		TeamType: TeamType,
		Ball:     &w2d_obj.Ball{},
	}
	bl.Ball.Pa = posacc.PosAcc{
		X: x,
		Y: y,
	}
	bl.Ball.Pa.SetDir(initdir)

	bl.Shields = make([]*w2d_obj.Shield, 24)
	for i := range bl.Shields {
		av := 1
		if i%2 == 0 {
			av = -1
		}
		bl.Shields[i] = &w2d_obj.Shield{
			Am: anglemove.AngleMove{
				Angle:  i * 15,
				AngleV: av,
			},
		}
	}

	bl.SuperShields = make([]*w2d_obj.SuperShield, 24)
	for i := range bl.SuperShields {
		av := 1
		if i%2 == 0 {
			av = -1
		}
		bl.SuperShields[i] = &w2d_obj.SuperShield{
			Am: anglemove.AngleMove{
				Angle:  15 + i*15,
				AngleV: av,
			},
			// frame: i * 3,
		}
	}
	return bl
}

func DrawTo(bl *w2d_obj.BallTeam, ctx js.Value) {
	dispSize := gameobjtype.Attrib[gameobjtype.Ball].Size
	sp := gSprites.BallSprites[bl.TeamType][gameobjtype.Ball]
	srcx, srcy := sp.GetSliceXY(0)
	dstx, dsty := bl.Ball.Pa.X-dispSize/2, bl.Ball.Pa.Y-dispSize/2
	ctx.Call("drawImage", sp.ImgCanvas,
		srcx, srcy, dispSize, dispSize,
		dstx, dsty, dispSize, dispSize,
	)
	dispSize = gameobjtype.Attrib[gameobjtype.Shield].Size
	dispR := gameobjtype.Attrib[gameobjtype.Shield].R
	sp = gSprites.BallSprites[bl.TeamType][gameobjtype.Shield]
	for _, v := range bl.Shields {
		srcx, srcy := sp.GetSliceXY(0)
		x, y := calcCircularPos(bl.Ball.Pa.X, bl.Ball.Pa.Y, v.Am.Angle, dispR)
		ctx.Call("drawImage", sp.ImgCanvas,
			srcx, srcy, dispSize, dispSize,
			x-dispSize/2, y-dispSize/2, dispSize, dispSize,
		)
	}
	dispSize = gameobjtype.Attrib[gameobjtype.SuperShield].Size
	dispR = gameobjtype.Attrib[gameobjtype.SuperShield].R
	sp = gSprites.BallSprites[bl.TeamType][gameobjtype.SuperShield]
	for _, v := range bl.SuperShields {
		// v.frame++
		srcx, srcy := sp.GetSliceXY(0)
		x, y := calcCircularPos(bl.Ball.Pa.X, bl.Ball.Pa.Y, v.Am.Angle, dispR)
		ctx.Call("drawImage", sp.ImgCanvas,
			srcx, srcy, dispSize, dispSize,
			x-dispSize/2, y-dispSize/2, dispSize, dispSize,
		)
	}
}

func Move(bl *w2d_obj.BallTeam, w, h int) {
	bl.Ball.Pa.Move()
	bl.Ball.Pa.BounceNormalize(w, h)
	for _, v := range bl.Shields {
		v.Am.Move()
	}
	for _, v := range bl.SuperShields {
		v.Am.Move()
	}
}
