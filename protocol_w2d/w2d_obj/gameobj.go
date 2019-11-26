package w2d_obj

import (
	"github.com/kasworld/gowasm2dgame/enums/teamtype"
	"github.com/kasworld/gowasm2dgame/lib/anglemove"
	"github.com/kasworld/gowasm2dgame/lib/gobase"
	"github.com/kasworld/gowasm2dgame/lib/posacc"
)

type Cloud struct {
	Pa        posacc.PosAcc
	SpriteNum int
}

type Background struct {
	Pa posacc.PosAcc
}

////////////////////

type SuperShield struct {
	GOBase gobase.GOBase
	Am     anglemove.AngleMove
}

type Shield struct {
	GOBase gobase.GOBase
	Am     anglemove.AngleMove
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
	TeamType      teamtype.TeamType
	Ball          *Ball
	Shields       []*Shield
	SuperShields  []*SuperShield
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
