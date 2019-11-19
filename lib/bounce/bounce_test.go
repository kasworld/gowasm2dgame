// Copyright 2015,2016,2017,2018 SeukWon Kang (kasworld@gmail.com)

package bounce

import (
	"testing"

	"github.com/kasworld/go-sdl2/sdl"
)

func TestBounce_Move(t *testing.T) {
	bn := New(
		sdl.Rect{0, 0, 20, 20},
		sdl.Rect{0, 5, 10, 10},
		1, 1)
	t.Logf("%v", bn)
	for i := 0; i < 40; i++ {
		bn.Move()
		t.Logf("%v", bn)
	}
}
