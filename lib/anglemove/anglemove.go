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

package anglemove

import (
	"math"
	"time"
)

type AngleMove struct {
	LastMoveTick int64 // time.unixnano
	Angle        float64
	AngleV       float64
}

func (am *AngleMove) Move() {
	now := time.Now().UnixNano()
	diff := float64(now-am.LastMoveTick) / float64(time.Second)
	am.LastMoveTick = now
	am.Angle += am.AngleV * diff
}

func (am *AngleMove) CalcCircularPos(cx, cy, r float64) (float64, float64) {
	rad := am.Angle * math.Pi / 180
	dstx := cx + r*math.Cos(rad)
	dsty := cy + r*math.Sin(rad)
	return dstx, dsty
}
