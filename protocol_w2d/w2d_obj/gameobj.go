package w2d_obj

import "time"

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

	X  int
	Y  int
	Dx int
	Dy int
}

type Ball struct {
	TeamUUID  string
	UUID      string
	BirthTime time.Time

	X  int
	Y  int
	Dx int
	Dy int
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

	X  int
	Y  int
	Dx int
	Dy int
}

type HommingBullet struct {
	TeamUUID  string
	UUID      string
	BirthTime time.Time

	X  int
	Y  int
	Dx int
	Dy int

	DstUUID string
}

type SuperBullet struct {
	TeamUUID  string
	UUID      string
	BirthTime time.Time

	X  int
	Y  int
	Dx int
	Dy int
}
