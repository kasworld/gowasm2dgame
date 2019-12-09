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
	"math/rand"
	"time"

	"github.com/kasworld/gowasm2dgame/lib/vector2f"

	"github.com/kasworld/gowasm2dgame/enums/effecttype"

	"github.com/kasworld/gowasm2dgame/enums/gameobjtype"
	"github.com/kasworld/gowasm2dgame/enums/teamtype"
	"github.com/kasworld/gowasm2dgame/game/gameconst"
	"github.com/kasworld/gowasm2dgame/lib/quadtreef"
	"github.com/kasworld/gowasm2dgame/lib/w2dlog"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_obj"
)

type Stage struct {
	rnd *rand.Rand      `prettystring:"hide"`
	log *w2dlog.LogBase `prettystring:"hide"`

	Background *w2d_obj.Background
	Clouds     []*w2d_obj.Cloud

	Effects   []*w2d_obj.Effect
	Teams     []*Team
	StageRect vector2f.Rect
}

func New(l *w2dlog.LogBase) *Stage {
	stg := &Stage{
		log: l,
		rnd: rand.New(rand.NewSource(time.Now().UnixNano())),
		StageRect: vector2f.Rect{
			0, 0,
			gameconst.StageW, gameconst.StageH,
		},
	}

	stg.Background = stg.NewBackground()
	stg.Clouds = make([]*w2d_obj.Cloud, 10)
	for i := range stg.Clouds {
		stg.Clouds[i] = stg.NewCloud(i)
	}
	stg.Teams = make([]*Team, teamtype.TeamType_Count)
	for i := range stg.Teams {
		stg.Teams[i] = NewTeam(l, teamtype.TeamType(i))
	}
	return stg
}

func (stg *Stage) Turn() {
	now := time.Now().UnixNano()

	// respawn dead team
	for _, bt := range stg.Teams {
		if !bt.IsAlive && bt.RespawnTick < now {
			bt.RespawnBall(now)
			stg.AddEffect(effecttype.Spawn, bt.Ball.PosVt, vector2f.VtZero)
		}
	}

	aienv := stg.move(now)
	for _, bt := range stg.Teams {
		if !bt.IsAlive {
			continue
		}
		actObj := stg.AI(bt, now, aienv)
		if bt.GetRemainAct(now, actObj.Act) > 0 {
			bt.ApplyAct(actObj)
		} else {
			stg.log.Fatal("OverAct %v %v", bt, actObj)
		}
	}
}

func (stg *Stage) move(now int64) *quadtreef.QuadTree {
	stg.Background.Move(now)
	stg.Background.PosVt = vector2f.Rect{
		0, 0, gameconst.StageW * 2, gameconst.StageH * 2,
	}.WrapVector(stg.Background.PosVt)
	// stg.Background.Wrap(gameconst.StageW*2, gameconst.StageH*2)

	for _, bt := range stg.Teams {
		if !bt.IsAlive {
			continue
		}
		toDelList := stg.MoveTeam(bt, now)
		for _, v := range toDelList {
			stg.AddEffectByGameObj(v)
		}
	}
	toDelList, aienv := stg.checkCollision()
	for _, v := range toDelList {
		stg.AddEffectByGameObj(v)
		if v.GOType == gameobjtype.Ball {
			stg.handleBallKilled(now, v)
		}
	}

	for _, eff := range stg.Effects {
		eff.Move(now)
	}
	for _, cld := range stg.Clouds {
		cld.Move(now)
		gameconst.StageRect.WrapVector(cld.PosVt)
	}
	return aienv
}

func (stg *Stage) handleBallKilled(now int64, gobj *GameObj) {
	for _, bt := range stg.Teams {
		// find ballteam
		if bt.Ball.UUID == gobj.UUID {
			bt.IsAlive = false
			// regist respawn
			bt.RespawnTick = now + int64(time.Second)*gameconst.BallRespawnDurSec

			// add effect
			stg.AddEffectByGameObj(bt.Ball)
			for _, v := range bt.Objs {
				if v.toDelete {
					continue
				}
				v.toDelete = true
				stg.AddEffectByGameObj(v)
			}
			return
		}
	}
	stg.log.Fatal("ball not in ballteam? %v", gobj)
}

func (stg *Stage) MoveTeam(bt *Team, now int64) []*GameObj {
	toDeleteList := make([]*GameObj, 0)
	bt.Ball.MoveStraight(now)
	bt.Ball.BounceNormalize(gameconst.StageW, gameconst.StageH)
	for _, v := range bt.Objs {
		if v.toDelete {
			continue
		}
		switch v.GOType {
		default:
		case gameobjtype.Bullet, gameobjtype.SuperBullet:
			v.MoveStraight(now)
			if !v.IsIn(gameconst.StageW, gameconst.StageH) {
				v.toDelete = true
				toDeleteList = append(toDeleteList, v)
			}
		case gameobjtype.Shield, gameobjtype.SuperShield:
			v.MoveCircular(now, bt.Ball.PosVt)
		case gameobjtype.HommingShield:
			v.MoveHommingShield(now, bt.Ball.PosVt)
		case gameobjtype.HommingBullet:
			findDst := false
			for _, dstbt := range stg.Teams {
				if !dstbt.IsAlive {
					continue
				}
				if dstbt.Ball.UUID == v.DstUUID {
					findDst = true
					v.MoveHommingBullet(now, dstbt.Ball.PosVt)
					break
				}
			}
			if !findDst {
				v.MoveStraight(now)
				if !v.IsIn(gameconst.StageW, gameconst.StageH) {
					v.toDelete = true
					toDeleteList = append(toDeleteList, v)
				}
			}
		}
		if !v.toDelete && !v.CheckLife(now) {
			v.toDelete = true
			toDeleteList = append(toDeleteList, v)
		}
	}
	return toDeleteList
}

func (stg *Stage) AddEffectByGameObj(gobj *GameObj) {
	switch gobj.GOType {
	case gameobjtype.Bullet, gameobjtype.HommingBullet, gameobjtype.Shield, gameobjtype.SuperShield, gameobjtype.HommingShield:
		// small effect
		stg.AddEffect(effecttype.ExplodeSmall, gobj.PosVt, gobj.MvVt)
	case gameobjtype.Ball, gameobjtype.SuperBullet:
		// big effect
		stg.AddEffect(effecttype.ExplodeBig, gobj.PosVt, gobj.MvVt)
	}
}

func (stg *Stage) ToStageInfo() *w2d_obj.NotiStageInfo_data {
	now := time.Now().UnixNano()
	rtn := &w2d_obj.NotiStageInfo_data{
		Tick:       time.Now().UnixNano(),
		Background: stg.Background,
		Clouds:     stg.Clouds,
	}
	for _, v := range stg.Effects {
		if v.CheckLife(now) {
			rtn.Effects = append(rtn.Effects, v)
		}
	}
	for _, bt := range stg.Teams {
		if !bt.IsAlive {
			continue
		}
		rtn.Teams = append(rtn.Teams, bt.ToPacket())
	}
	return rtn
}

func (stg *Stage) ToStatsInfo() *w2d_obj.NotiStatsInfo_data {
	rtn := &w2d_obj.NotiStatsInfo_data{}
	for i, bt := range stg.Teams {
		rtn.ActStats[i] = bt.ActStats
	}
	return rtn
}
