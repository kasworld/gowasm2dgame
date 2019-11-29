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

package stage

import (
	"github.com/kasworld/gowasm2dgame/enums/gameobjtype"
	"github.com/kasworld/gowasm2dgame/lib/quadtreef"
)

func (bt *BallTeam) AI(aienv *quadtreef.QuadTree) {
	switch bt.rnd.Intn(6) {
	case 0:
		maxv := gameobjtype.Attrib[gameobjtype.Bullet].V
		bt.AddBullet(bt.rnd.Float64()*360, maxv)
	case 1:
		maxv := gameobjtype.Attrib[gameobjtype.SuperBullet].V
		bt.AddSuperBullet(bt.rnd.Float64()*360, maxv)
	case 2:
		if bt.Count(gameobjtype.SuperShield) < 12 {
			maxv := gameobjtype.Attrib[gameobjtype.SuperShield].V
			bt.AddSuperShield(bt.rnd.Float64()*360, bt.rnd.Float64()*maxv)
		}
	case 3:
		if bt.Count(gameobjtype.HommingShield) < 12 {
			maxv := gameobjtype.Attrib[gameobjtype.HommingShield].V
			bt.AddHommingShield(bt.rnd.Float64()*360, maxv)
		}
	case 4:
		if bt.Count(gameobjtype.Shield) < 12 {
			maxv := gameobjtype.Attrib[gameobjtype.Shield].V
			bt.AddShield(bt.rnd.Float64()*360, bt.rnd.Float64()*maxv)
		}
	case 5:
		maxv := gameobjtype.Attrib[gameobjtype.Ball].V
		dx, dy := CalcDxyFromAngelV(
			bt.rnd.Float64()*360,
			bt.rnd.Float64()*maxv,
		)
		bt.Ball.AddDxy(dx, dy)
	}
}
