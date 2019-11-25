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
	"github.com/kasworld/gowasm2dgame/enums/gameobjtype"
	"github.com/kasworld/gowasm2dgame/enums/teamtype"
)

var gSprites *Sprites

type Sprites struct {
	BallSprites       [teamtype.TeamType_Count][gameobjtype.GameObjType_Count]*Sprite
	ExplodeetcSprite  *Sprite
	ExplodeballSprite *Sprite
	SpawnSprite       *Sprite
	CloudSprite       *Sprite
}

func LoadSprites() *Sprites {
	sps := &Sprites{}
	sps.SpawnSprite = LoadSpriteXYN("spawn", "spawnStore", 1, 6)
	sps.ExplodeetcSprite = LoadSpriteXYN("explodeetc", "explodeetcStore", 1, 8)
	sps.ExplodeballSprite = LoadSpriteXYN("explodeball", "explodeballStore", 8, 1)
	sps.CloudSprite = LoadSpriteXYN("clouds", "cloudStore", 1, 4)
	// load team sprite
	teamAttrib := teamtype.SpriteFilter
	for i := 0; i < teamtype.TeamType_Count; i++ {
		sps.BallSprites[i] = LoadBallSprite(teamAttrib[i].Name)
		for j := range sps.BallSprites[i] {
			for _, x := range teamAttrib[i].IV {
				sps.BallSprites[i][j].Filter(x.Index, x.Value)
			}
		}
	}
	return sps
}
