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
	"github.com/kasworld/direction"
	"github.com/kasworld/go-abs"
)

type PosAcc struct {
	X  int
	Y  int
	Dx int
	Dy int
}

func (pa *PosAcc) Move() {
	pa.X += pa.Dx
	pa.Y += pa.Dy
}

func (pa *PosAcc) BounceNormalize(w, h int) {
	if pa.X < 0 {
		pa.X = 0
		pa.Dx = abs.Absi(pa.Dx)
	}
	if pa.Y < 0 {
		pa.Y = 0
		pa.Dy = abs.Absi(pa.Dy)
	}

	if pa.X >= w {
		pa.X = w - 1
		pa.Dx = -abs.Absi(pa.Dx)
	}
	if pa.Y >= h {
		pa.Y = h - 1
		pa.Dy = -abs.Absi(pa.Dy)
	}
}

func (pa *PosAcc) IsIn(w, h int) bool {
	return 0 <= pa.X && pa.X < w && 0 <= pa.Y && pa.Y < h
}

func (pa *PosAcc) SetDxy(dx, dy int) {
	pa.Dx = dx
	pa.Dy = dy
}

func (pa *PosAcc) SetDir(dir direction.Direction_Type) {
	pa.Dx = dir.Vector()[0]
	pa.Dy = dir.Vector()[1]
}

func (pa *PosAcc) AccelerateDir(dir direction.Direction_Type) {
	pa.Dx += dir.Vector()[0]
	pa.Dy += dir.Vector()[1]
}
