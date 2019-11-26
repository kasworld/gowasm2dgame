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

package posacc

import (
	"github.com/kasworld/go-abs"
)

type PosAcc struct {
	X  float64
	Y  float64
	Dx float64
	Dy float64
}

func (pa *PosAcc) Move() {
	pa.X += pa.Dx
	pa.Y += pa.Dy
}

func (pa *PosAcc) BounceNormalize(w, h float64) {
	if pa.X < 0 {
		pa.X = 0
		pa.Dx = abs.Absf(pa.Dx)
	}
	if pa.Y < 0 {
		pa.Y = 0
		pa.Dy = abs.Absf(pa.Dy)
	}

	if pa.X > w {
		pa.X = w
		pa.Dx = -abs.Absf(pa.Dx)
	}
	if pa.Y > h {
		pa.Y = h
		pa.Dy = -abs.Absf(pa.Dy)
	}
}

func (pa *PosAcc) IsIn(w, h float64) bool {
	return 0 <= pa.X && pa.X <= w && 0 <= pa.Y && pa.Y <= h
}

func (pa *PosAcc) Wrap(w, h float64) (float64, float64) {
	if pa.X < 0 {
		pa.X = w
	}
	if pa.Y < 0 {
		pa.Y = h
	}

	if pa.X > w {
		pa.X = 0
	}
	if pa.Y > h {
		pa.Y = 0
	}
	return pa.X, pa.Y
}

func (pa *PosAcc) SetDxy(dx, dy float64) {
	pa.Dx = dx
	pa.Dy = dy
}

// func (pa *PosAcc) SetDir(dir direction.Direction_Type) {
// 	pa.Dx = dir.Vector()[0]
// 	pa.Dy = dir.Vector()[1]
// }

// func (pa *PosAcc) AccelerateDir(dir direction.Direction_Type) {
// 	pa.Dx += dir.Vector()[0]
// 	pa.Dy += dir.Vector()[1]
// }
