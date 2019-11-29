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
	"github.com/kasworld/gowasm2dgame/enums/acttype"
	"github.com/kasworld/gowasm2dgame/enums/gameobjtype"
	"github.com/kasworld/gowasm2dgame/lib/quadtreef"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_obj"
)

func (stg *Stage) AI(bt *BallTeam, aienv *quadtreef.QuadTree) *w2d_obj.Act {
	switch bt.rnd.Intn(10) {
	default:
		//pass
	case 0:
		maxv := gameobjtype.Attrib[gameobjtype.Bullet].V
		return &w2d_obj.Act{
			Act:    acttype.Bullet,
			Angle:  bt.rnd.Float64() * 360,
			AngleV: maxv,
		}
	case 1:
		maxv := gameobjtype.Attrib[gameobjtype.SuperBullet].V
		return &w2d_obj.Act{
			Act:    acttype.SuperBullet,
			Angle:  bt.rnd.Float64() * 360,
			AngleV: maxv,
		}
	case 2:
		maxv := gameobjtype.Attrib[gameobjtype.HommingBullet].V
		dstteam := stg.Teams[bt.rnd.Intn(len(stg.Teams))]
		if dstteam != bt && dstteam.IsAlive {
			return &w2d_obj.Act{
				Act:      acttype.HommingBullet,
				Angle:    bt.rnd.Float64() * 360,
				AngleV:   maxv,
				DstObjID: dstteam.Ball.UUID,
			}
		}
	case 3:
		if bt.Count(gameobjtype.Shield) < 12 {
			maxv := gameobjtype.Attrib[gameobjtype.Shield].V
			return &w2d_obj.Act{
				Act:    acttype.Shield,
				Angle:  bt.rnd.Float64() * 360,
				AngleV: bt.rnd.Float64() * maxv,
			}
		}
	case 4:
		if bt.Count(gameobjtype.SuperShield) < 12 && bt.rnd.Intn(10) == 0 {
			maxv := gameobjtype.Attrib[gameobjtype.SuperShield].V
			return &w2d_obj.Act{
				Act:    acttype.SuperShield,
				Angle:  bt.rnd.Float64() * 360,
				AngleV: bt.rnd.Float64() * maxv,
			}
		}
	case 5:
		if bt.Count(gameobjtype.HommingShield) < 6 && bt.rnd.Intn(10) == 0 {
			maxv := gameobjtype.Attrib[gameobjtype.HommingShield].V
			return &w2d_obj.Act{
				Act:    acttype.HommingShield,
				Angle:  bt.rnd.Float64() * 360,
				AngleV: maxv,
			}
		}
	case 6:
		maxv := gameobjtype.Attrib[gameobjtype.Ball].V
		return &w2d_obj.Act{
			Act:    acttype.Accel,
			Angle:  bt.rnd.Float64() * 360,
			AngleV: bt.rnd.Float64() * maxv,
		}
	}
	return &w2d_obj.Act{
		Act: acttype.Nothing,
	}
}
