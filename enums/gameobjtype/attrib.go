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

package gameobjtype

import (
	"math"
	"time"
)

const LongLife = int64(time.Second) * 3600 * 24 * 365

var Attrib = [GameObjType_Count]struct {
	Size        float64
	R           float64 // from main ball center
	V           float64 // speed pixel/sec or rad/sec
	FramePerSec float64
	LifeTick    int64
}{
	Ball:          {32, 0, 300, 0, LongLife},
	Shield:        {16, 28, 180 * math.Pi / 180, 0, LongLife},
	SuperShield:   {16, 48, 180 * math.Pi / 180, 30, int64(time.Second) * 60},
	HommingShield: {16, 0, 50, 30, int64(time.Second) * 60},
	Bullet:        {16, 0, 500, 0, LongLife},
	SuperBullet:   {32, 0, 600, 30, LongLife},
	HommingBullet: {16, 0, 300, 30, int64(time.Second) * 60},
}

var collisionRule = [GameObjType_Count][]GameObjType{
	Ball:          {Ball, Shield, SuperShield, Bullet, HommingShield, SuperBullet, HommingBullet},
	Shield:        {Ball, Shield, SuperShield, Bullet, HommingShield, SuperBullet, HommingBullet},
	SuperShield:   {SuperShield, SuperBullet, HommingBullet},
	Bullet:        {Ball, Shield, SuperShield, Bullet, HommingShield, SuperBullet, HommingBullet},
	HommingShield: {Ball, Shield, SuperShield, Bullet, HommingShield, SuperBullet, HommingBullet},
	SuperBullet:   {SuperShield, SuperBullet, HommingBullet},
	HommingBullet: {SuperShield, SuperBullet, HommingBullet},
}

func CollisionTo(srcType, dstType GameObjType) bool {
	for _, v := range collisionRule[srcType] {
		if v == dstType {
			return true
		}
	}
	return false
}
