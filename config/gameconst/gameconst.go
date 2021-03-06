package gameconst

import "github.com/kasworld/gowasm2dgame/lib/vector2f"

const (
	StagePerServer = 10

	StageSize         = 1000.0
	StageW            = 1000.0
	StageH            = 1000.0
	BallRespawnDurSec = 5

	MaxChatLen = 80
)

var StageRect = vector2f.Rect{
	vector2f.Vector2f{0, 0},
	vector2f.Vector2f{StageW, StageH},
}
