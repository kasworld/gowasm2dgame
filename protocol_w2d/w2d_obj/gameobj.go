package w2d_obj

import (
	"github.com/kasworld/gowasm2dgame/enums/effecttype"
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

type Effect struct {
	EffectType effecttype.EffectType
	Pa         posacc.PosAcc
	Frame      int
}

////////////////////

type SuperShield struct {
	GOBase gobase.GOBase
	Am     anglemove.AngleMove
	Frame  int
}

type Shield struct {
	GOBase gobase.GOBase
	Am     anglemove.AngleMove
}

type HommingShield struct {
	GOBase gobase.GOBase
	Pa     posacc.PosAcc
	Frame  int
}

type Ball struct {
	GOBase gobase.GOBase
	Pa     posacc.PosAcc
}

type BallTeam struct {
	TeamType       teamtype.TeamType
	Ball           *Ball
	Shields        []*Shield
	SuperShields   []*SuperShield
	HommingShields []*HommingShield
	Bullets        []*Bullet
	SuperBullets   []*SuperBullet
	HommingBullets []*HommingBullet
}

type Bullet struct {
	GOBase gobase.GOBase
	Pa     posacc.PosAcc
}

type HommingBullet struct {
	GOBase  gobase.GOBase
	Pa      posacc.PosAcc
	Frame   int
	DstUUID string
}

type SuperBullet struct {
	GOBase gobase.GOBase
	Pa     posacc.PosAcc
	Frame  int
}
