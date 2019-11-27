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

var Attrib = [GameObjType_Count]struct {
	Size        float64
	R           float64 // from main ball center
	V           float64 // speed pixel/sec or degree/sec
	FramePerSec float64
}{
	Ball:          {32, 0, 300, 0},
	Shield:        {16, 28, 180, 0},
	SuperShield:   {16, 48, 180, 30},
	HommingShield: {16, 0, 300, 30},
	Bullet:        {16, 0, 500, 0},
	SuperBullet:   {32, 0, 600, 30},
	HommingBullet: {16, 0, 300, 30},
}
