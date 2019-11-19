// Copyright 2015,2016,2017,2018 SeukWon Kang (kasworld@gmail.com)

// Package stroll8way move position 8way semi autometically
package stroll8way

import (
	"fmt"

	"github.com/kasworld/direction"
)

// 8 direct move
type Stroll8 struct {
	X  int
	Y  int
	Dx int
	Dy int
}

func (s8 Stroll8) String() string {
	return fmt.Sprintf("Stroll8[(%d %d) (%d %d)]",
		s8.X, s8.Y,
		s8.Dx, s8.Dy,
	)
}

func (s8 *Stroll8) Move() {
	s8.X += s8.Dx
	s8.Y += s8.Dy

}

func (s8 *Stroll8) SetDxy(dx, dy int) {
	s8.Dx = dx
	s8.Dy = dy
}

func (s8 *Stroll8) SetDir(dir direction.Direction_Type) {
	s8.Dx = dir.Vector()[0]
	s8.Dy = dir.Vector()[1]
}

func (s8 *Stroll8) AccelerateDir(dir direction.Direction_Type) {
	if dir == direction.Dir_stop {
		s8.Dx = dir.Vector()[0]
		s8.Dy = dir.Vector()[1]
	} else {
		s8.Dx += dir.Vector()[0]
		s8.Dy += dir.Vector()[1]
	}
}
