package w2d_obj

import (
	"github.com/kasworld/gowasm2dgame/enums/acttype"
)

type Act struct {
	Act      acttype.ActType
	DstPos   [2]int
	DstObjID string
	// some more?
}

type Cloud struct {
	X  int
	Y  int
	Dx int
	Dy int
}

type Background struct {
	X  int
	Y  int
	Dx int
	Dy int
}
