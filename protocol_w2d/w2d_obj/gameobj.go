package w2d_obj

import (
	"github.com/kasworld/gowasm2dgame/lib/gobase"
	"github.com/kasworld/gowasm2dgame/lib/posacc"
)

type SuperShield struct {
	GOBase gobase.GOBase
	Angle  int
	AngleV int
}

type Shield struct {
	GOBase gobase.GOBase
	Angle  int
	AngleV int
}

type HommingShield struct {
	GOBase gobase.GOBase
	Pa     posacc.PosAcc
}

type Ball struct {
	GOBase gobase.GOBase
	Pa     posacc.PosAcc
}

type BallTeam struct {
	Ball          *Ball
	Shiels        []*Shield
	SuperShiels   []*SuperShield
	HommingShiels []*HommingShield
}

type Bullet struct {
	GOBase gobase.GOBase
	Pa     posacc.PosAcc
}

type HommingBullet struct {
	GOBase gobase.GOBase
	Pa     posacc.PosAcc

	DstUUID string
}

type SuperBullet struct {
	GOBase gobase.GOBase
	Pa     posacc.PosAcc
}
