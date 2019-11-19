// Copyright 2015,2016,2017,2018 SeukWon Kang (kasworld@gmail.com)

// Package bounce bound objrect in borderrect
package bounce

import (
	"fmt"

	"github.com/kasworld/rect"
)

type Bounce struct {
	Border rect.Rect
	Object rect.Rect
	Dx     int
	Dy     int
}

func New(border, obj rect.Rect, dx, dy int) *Bounce {
	bn := &Bounce{
		Border: border,
		Object: obj,
		Dx:     dx,
		Dy:     dy,
	}
	bn.IsValid()
	return bn
}

func (bn *Bounce) IsValid() error {
	if bn.Object.W > bn.Border.W {
		return fmt.Errorf("invalid %v", bn)
	}
	if bn.Object.H > bn.Border.H {
		return fmt.Errorf("invalid %v", bn)
	}
	bn.Normalize()
	return nil
}

func (bn *Bounce) Normalize() {
	if bn.Object.X < bn.Border.X {
		bn.Object.X = bn.Border.X
		bn.Dx = bn.GetAbsDx()
	}
	if bn.Object.Y < bn.Border.Y {
		bn.Object.Y = bn.Border.Y
		bn.Dy = bn.GetAbsDy()
	}

	if bn.Object.X+bn.Object.W > bn.Border.X+bn.Border.W {
		bn.Object.X = bn.Border.X + bn.Border.W - bn.Object.W
		bn.Dx = -bn.GetAbsDx()
	}
	if bn.Object.Y+bn.Object.H > bn.Border.Y+bn.Border.H {
		bn.Object.Y = bn.Border.Y + bn.Border.H - bn.Object.H
		bn.Dy = -bn.GetAbsDy()
	}
}

func (bn *Bounce) GetBorder() rect.Rect {
	return bn.Border
}
func (bn *Bounce) SetBorder(border rect.Rect) {
	bn.Border = border
	bn.IsValid()
}

func (bn *Bounce) GetObject() rect.Rect {
	return bn.Object
}
func (bn *Bounce) SetObject(obj rect.Rect) {
	bn.Object = obj
	bn.IsValid()
}

func (bn *Bounce) SetDx(dx int) {
	bn.Dx = dx
}
func (bn *Bounce) GetDx() int {
	return bn.Dx
}
func (bn *Bounce) GetAbsDx() int {
	if bn.Dx < 0 {
		return -bn.Dx
	}
	return bn.Dx
}

func (bn *Bounce) SetDy(dy int) {
	bn.Dy = dy
}
func (bn *Bounce) GetAbsDy() int {
	if bn.Dy < 0 {
		return -bn.Dy
	}
	return bn.Dy
}
func (bn *Bounce) GetDy() int {
	return bn.Dy
}

func (bn *Bounce) Move() {
	bn.Object.X += bn.Dx
	bn.Object.Y += bn.Dy
	bn.Normalize()
}
