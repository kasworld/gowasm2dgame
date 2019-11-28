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

	"github.com/kasworld/gowasm2dgame/enums/effecttype"

	"github.com/kasworld/gowasm2dgame/enums/gameobjtype"
	"github.com/kasworld/gowasm2dgame/enums/teamtype"
	"github.com/kasworld/gowasm2dgame/game/gameconst"
	"github.com/kasworld/gowasm2dgame/lib/w2dlog"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_obj"
)

type Stage struct {
	rnd *rand.Rand      `prettystring:"hide"`
	log *w2dlog.LogBase `prettystring:"hide"`

	Background *w2d_obj.Background
	Clouds     []*w2d_obj.Cloud

	Effects []*w2d_obj.Effect
	Teams   []*BallTeam
}

func New(l *w2dlog.LogBase) *Stage {
	stg := &Stage{
		log: l,
		rnd: rand.New(rand.NewSource(time.Now().UnixNano())),
	}

	stg.Background = stg.NewBackground()
	stg.Clouds = make([]*w2d_obj.Cloud, 10)
	for i := range stg.Clouds {
		stg.Clouds[i] = stg.NewCloud(i)
	}
	stg.Teams = make([]*BallTeam, teamtype.TeamType_Count)
	for i := range stg.Teams {
		o := NewBallTeam(
			teamtype.TeamType(i),
			stg.rnd.Float64()*gameconst.StageW,
			stg.rnd.Float64()*gameconst.StageH,
		)
		maxv := gameobjtype.Attrib[gameobjtype.Ball].V
		dx, dy := CalcDxyFromAngelV(
			stg.rnd.Float64()*360,
			stg.rnd.Float64()*maxv,
		)
		o.Ball.SetDxy(dx, dy)
		stg.Teams[i] = o
	}
	return stg
}

func (stg *Stage) Turn() {
	now := time.Now().UnixNano()
	stg.move(now)
}

func (stg *Stage) move(now int64) {
	stg.Background.Move(now)
	stg.Background.Wrap(gameconst.StageW*2, gameconst.StageH*2)
	for _, bt := range stg.Teams {
		toDelList := bt.Move(now)
		bt.AI()
		for _, v := range toDelList {
			switch v.GOType {
			case gameobjtype.Bullet, gameobjtype.HommingBullet, gameobjtype.Shield, gameobjtype.SuperShield, gameobjtype.HommingShield:
				// small effect
				stg.AddEffect(effecttype.ExplodeSmall, v.X, v.Y)
			case gameobjtype.SuperBullet:
				// big effect
				stg.AddEffect(effecttype.ExplodeBig, v.X, v.Y)
			}
		}
	}
	toDelList := stg.checkCollision()
	for _, v := range toDelList {
		switch v.GOType {
		case gameobjtype.Bullet, gameobjtype.HommingBullet, gameobjtype.Shield, gameobjtype.SuperShield, gameobjtype.HommingShield:
			// small effect
			stg.AddEffect(effecttype.ExplodeSmall, v.X, v.Y)
		case gameobjtype.SuperBullet:
			// big effect
			stg.AddEffect(effecttype.ExplodeBig, v.X, v.Y)
		}
	}

	for _, cld := range stg.Clouds {
		cld.Move(now)
		cld.Wrap(gameconst.StageW, gameconst.StageH)
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
	for _, v := range stg.Teams {
		rtn.Teams = append(rtn.Teams, v.ToPacket())
	}
	return rtn
}
