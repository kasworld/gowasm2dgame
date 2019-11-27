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

	"github.com/kasworld/gowasm2dgame/enums/gameobjtype"

	"github.com/kasworld/gowasm2dgame/game/gameconst"
	"github.com/kasworld/gowasm2dgame/lib/w2dlog"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_obj"
)

type Stage struct {
	rnd *rand.Rand      `prettystring:"hide"`
	log *w2dlog.LogBase `prettystring:"hide"`

	Background *w2d_obj.Background
	Teams      []*w2d_obj.BallTeam
	Effects    []*w2d_obj.Effect
	Clouds     []*w2d_obj.Cloud
}

func New(l *w2dlog.LogBase) *Stage {
	stg := &Stage{
		log: l,
		rnd: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
	stg.makeTestData()
	return stg
}

func (stg *Stage) Turn() {
	now := time.Now().UnixNano()
	stg.move(now)
}

func (stg *Stage) move(now int64) {
	stg.Background.Pa.Move(now)
	stg.Background.Pa.Wrap(gameconst.StageW*2, gameconst.StageH*2)
	for _, bt := range stg.Teams {
		stg.moveTeam(bt, now)
	}
	for _, cld := range stg.Clouds {
		cld.Pa.Move(now)
		cld.Pa.Wrap(gameconst.StageW, gameconst.StageH)
	}
}

func (stg *Stage) moveTeam(bt *w2d_obj.BallTeam, now int64) {
	bt.Bullets = append(bt.Bullets,
		stg.NewBullet(bt.Ball.Pa.X, bt.Ball.Pa.Y),
	)
	bt.Ball.Pa.Move(now)
	bt.Ball.Pa.BounceNormalize(gameconst.StageW, gameconst.StageH)
	for _, v := range bt.Shields {
		v.Am.Move(now)
	}

	lifetick := gameobjtype.Attrib[gameobjtype.SuperShield].LifeTick
	newSuperShields := make([]*w2d_obj.SuperShield, 0, len(bt.SuperShields))
	for _, v := range bt.SuperShields {
		v.Am.Move(now)
		if now-v.GOBase.BirthTick < lifetick {
			newSuperShields = append(newSuperShields, v)
		}
	}
	bt.SuperShields = newSuperShields

	lifetick = gameobjtype.Attrib[gameobjtype.HommingShield].LifeTick
	newHommingShields := make([]*w2d_obj.HommingShield, 0, len(bt.HommingShields))
	for _, v := range bt.HommingShields {
		v.Pa.Move(now)
		if now-v.GOBase.BirthTick < lifetick {
			newHommingShields = append(newHommingShields, v)
		}
	}
	bt.HommingShields = newHommingShields

	newBullets := make([]*w2d_obj.Bullet, 0, len(bt.Bullets))
	for _, v := range bt.Bullets {
		v.Pa.Move(now)
		if v.Pa.IsIn(gameconst.StageW, gameconst.StageH) {
			newBullets = append(newBullets, v)
		}
	}
	bt.Bullets = newBullets

	newSuperBullets := make([]*w2d_obj.SuperBullet, 0, len(bt.SuperBullets))
	for _, v := range bt.SuperBullets {
		v.Pa.Move(now)
		if v.Pa.IsIn(gameconst.StageW, gameconst.StageH) {
			newSuperBullets = append(newSuperBullets, v)
		}
	}
	bt.SuperBullets = newSuperBullets

	lifetick = gameobjtype.Attrib[gameobjtype.HommingBullet].LifeTick
	newHommingBullets := make([]*w2d_obj.HommingBullet, 0, len(bt.HommingBullets))
	for _, v := range bt.HommingBullets {
		v.Pa.Move(now)
		if now-v.GOBase.BirthTick < lifetick {
			newHommingBullets = append(newHommingBullets, v)
		}
	}
	bt.HommingBullets = newHommingBullets

}

func (stg *Stage) ToStageInfo() *w2d_obj.NotiStageInfo_data {
	rtn := &w2d_obj.NotiStageInfo_data{
		Time:       time.Now(),
		Background: stg.Background,
		Teams:      stg.Teams,
		Effects:    stg.Effects,
		Clouds:     stg.Clouds,
	}
	return rtn
}
