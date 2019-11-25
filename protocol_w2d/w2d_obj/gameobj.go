package w2d_obj

import (
	"time"

	"github.com/kasworld/gowasm2dgame/lib/posacc"
)

type SuperShield struct {
	TeamUUID  string
	UUID      string
	BirthTime time.Time

	Angle  int
	AngleV int
}

type Shield struct {
	TeamUUID  string
	UUID      string
	BirthTime time.Time

	Angle  int
	AngleV int
}

type HommingShield struct {
	TeamUUID  string
	UUID      string
	BirthTime time.Time

	Pa posacc.PosAcc
}

type Ball struct {
	TeamUUID  string
	UUID      string
	BirthTime time.Time

	Pa posacc.PosAcc
}

type BallTeam struct {
	Ball          *Ball
	Shiels        []*Shield
	SuperShiels   []*SuperShield
	HommingShiels []*HommingShield
}

type Bullet struct {
	TeamUUID  string
	UUID      string
	BirthTime time.Time

	Pa posacc.PosAcc
}

type HommingBullet struct {
	TeamUUID  string
	UUID      string
	BirthTime time.Time

	Pa posacc.PosAcc

	DstUUID string
}

type SuperBullet struct {
	TeamUUID  string
	UUID      string
	BirthTime time.Time

	Pa posacc.PosAcc
}
