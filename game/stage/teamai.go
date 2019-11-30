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
	"github.com/kasworld/gowasm2dgame/enums/teamtype"
	"github.com/kasworld/gowasm2dgame/lib/quadtreef"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_obj"
)

func (stg *Stage) SelectRandomTeam(me *BallTeam) *BallTeam {
	for i := 0; i < teamtype.TeamType_Count; i++ {
		dstteam := stg.Teams[stg.rnd.Intn(len(stg.Teams))]
		if dstteam != me && dstteam.IsAlive {
			return dstteam
		}
	}
	return nil
}

func (stg *Stage) AI(bt *BallTeam, now int64, aienv *quadtreef.QuadTree) *w2d_obj.Act {
	switch bt.rnd.Intn(10) {
	default:
		//pass
	case 0:
		actt := acttype.Bullet
		objt := gameobjtype.Bullet
		if bt.GetRemainAct(now, actt) <= 0 {
			break
		}
		dstteam := stg.SelectRandomTeam(bt)
		if dstteam == nil {
			break
		}
		angle, v := bt.CalcAimAngleAndV(objt, dstteam.Ball)
		return &w2d_obj.Act{
			Act:    actt,
			Angle:  angle,
			AngleV: v,
		}
	case 1:
		actt := acttype.SuperBullet
		objt := gameobjtype.SuperBullet
		if bt.GetRemainAct(now, actt) <= 0 {
			break
		}
		dstteam := stg.SelectRandomTeam(bt)
		if dstteam == nil {
			break
		}
		angle, v := bt.CalcAimAngleAndV(objt, dstteam.Ball)
		return &w2d_obj.Act{
			Act:    actt,
			Angle:  angle,
			AngleV: v,
		}
	case 2:
		actt := acttype.HommingBullet
		objt := gameobjtype.HommingBullet
		if bt.GetRemainAct(now, actt) <= 0 {
			break
		}
		dstteam := stg.SelectRandomTeam(bt)
		if dstteam == nil {
			break
		}
		maxv := gameobjtype.Attrib[objt].V
		if dstteam != bt && dstteam.IsAlive {
			return &w2d_obj.Act{
				Act:      actt,
				Angle:    bt.rnd.Float64() * 360,
				AngleV:   maxv,
				DstObjID: dstteam.Ball.UUID,
			}
		}
	case 3:
		actt := acttype.Shield
		objt := gameobjtype.Shield
		if bt.GetRemainAct(now, actt) <= 0 {
			break
		}
		if bt.Count(objt) < 12 {
			maxv := gameobjtype.Attrib[objt].V
			return &w2d_obj.Act{
				Act:    actt,
				Angle:  bt.rnd.Float64() * 360,
				AngleV: bt.rnd.Float64() * maxv,
			}
		}
	case 4:
		actt := acttype.SuperShield
		objt := gameobjtype.SuperShield
		if bt.GetRemainAct(now, actt) <= 0 {
			break
		}
		if bt.Count(objt) < 12 && bt.rnd.Intn(10) == 0 {
			maxv := gameobjtype.Attrib[objt].V
			return &w2d_obj.Act{
				Act:    actt,
				Angle:  bt.rnd.Float64() * 360,
				AngleV: bt.rnd.Float64() * maxv,
			}
		}
	case 5:
		actt := acttype.HommingShield
		objt := gameobjtype.HommingShield
		if bt.GetRemainAct(now, actt) <= 0 {
			break
		}
		if bt.Count(objt) < 6 && bt.rnd.Intn(10) == 0 {
			maxv := gameobjtype.Attrib[objt].V
			return &w2d_obj.Act{
				Act:    actt,
				Angle:  bt.rnd.Float64() * 360,
				AngleV: maxv,
			}
		}
	case 6:
		actt := acttype.Accel
		objt := gameobjtype.HommingShield
		if bt.GetRemainAct(now, actt) <= 0 {
			break
		}
		if bt.GetRemainAct(now, actt) > 0 {
			maxv := gameobjtype.Attrib[objt].V
			return &w2d_obj.Act{
				Act:    actt,
				Angle:  bt.rnd.Float64() * 360,
				AngleV: bt.rnd.Float64() * maxv,
			}
		}
	}
	return &w2d_obj.Act{
		Act: acttype.Nothing,
	}
}
