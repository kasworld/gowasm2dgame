package w2d_obj

import (
	"math"
	"time"

	"github.com/kasworld/gowasm2dgame/enums/effecttype"
	"github.com/kasworld/gowasm2dgame/enums/gameobjtype"
	"github.com/kasworld/gowasm2dgame/enums/teamtype"
)

type Cloud struct {
	SpriteNum    int
	LastMoveTick int64 // time.unixnano
	X            float64
	Y            float64
	Dx           float64
	Dy           float64
}

func (o *Cloud) Move(now int64) {
	diff := float64(now-o.LastMoveTick) / float64(time.Second)
	o.LastMoveTick = now
	o.X += o.Dx * diff
	o.Y += o.Dy * diff
}

func (o *Cloud) Wrap(w, h float64) (float64, float64) {
	if o.X < 0 {
		o.X = w
	}
	if o.Y < 0 {
		o.Y = h
	}

	if o.X > w {
		o.X = 0
	}
	if o.Y > h {
		o.Y = 0
	}
	return o.X, o.Y
}

type Background struct {
	LastMoveTick int64 // time.unixnano
	X            float64
	Y            float64
	Dx           float64
	Dy           float64
}

func (o *Background) Move(now int64) {
	diff := float64(now-o.LastMoveTick) / float64(time.Second)
	o.LastMoveTick = now
	o.X += o.Dx * diff
	o.Y += o.Dy * diff
}
func (o *Background) Wrap(w, h float64) (float64, float64) {
	if o.X < 0 {
		o.X = w
	}
	if o.Y < 0 {
		o.Y = h
	}

	if o.X > w {
		o.X = 0
	}
	if o.Y > h {
		o.Y = 0
	}
	return o.X, o.Y
}

type Effect struct {
	EffectType   effecttype.EffectType
	BirthTick    int64
	LastMoveTick int64 // time.unixnano
	X            float64
	Y            float64
	Dx           float64
	Dy           float64
}

func (o *Effect) CheckLife(now int64) bool {
	lifetick := effecttype.Attrib[o.EffectType].LifeTick
	return now-o.BirthTick < lifetick
}

func (o *Effect) Move(now int64) {
	diff := float64(now-o.LastMoveTick) / float64(time.Second)
	o.LastMoveTick = now
	o.X += o.Dx * diff
	o.Y += o.Dy * diff
	o.Dx *= 0.9
	o.Dy *= 0.9
}

////////////////////

type BallTeam struct {
	TeamType teamtype.TeamType
	Ball     *GameObj
	Objs     []*GameObj
}

type GameObj struct {
	GOType       gameobjtype.GameObjType
	UUID         string
	BirthTick    int64
	LastMoveTick int64 // time.unixnano
	X            float64
	Y            float64
	Dx           float64 // move line
	Dy           float64
	Angle        float64 // move circular
	AngleV       float64
	DstUUID      string // move to dest
}

func (o *GameObj) MoveStraight(now int64) {
	diff := float64(now-o.LastMoveTick) / float64(time.Second)
	o.LastMoveTick = now
	o.X += o.Dx * diff
	o.Y += o.Dy * diff
}

func (o *GameObj) MoveCircular(now int64, cx, cy float64) {
	diff := float64(now-o.LastMoveTick) / float64(time.Second)
	o.LastMoveTick = now
	o.Angle += o.AngleV * diff
	r := gameobjtype.Attrib[o.GOType].R
	o.X, o.Y = o.CalcCircularPos(cx, cy, r)
}

func (o *GameObj) CalcCircularPos(cx, cy, r float64) (float64, float64) {
	rad := o.Angle //* math.Pi / 180
	dstx := cx + r*math.Cos(rad)
	dsty := cy + r*math.Sin(rad)
	return dstx, dsty
}

func (o *GameObj) MoveHomming(now int64, dstx, dsty float64) {
	diff := float64(now-o.LastMoveTick) / float64(time.Second)
	o.LastMoveTick = now
	o.X += o.Dx * diff
	o.Y += o.Dy * diff

	maxv := gameobjtype.Attrib[o.GOType].V
	dx := dstx - o.X
	dy := dsty - o.Y
	l := math.Sqrt(dx*dx + dy*dy)
	o.Dx += dx / l * maxv
	o.Dy += dy / l * maxv
}
