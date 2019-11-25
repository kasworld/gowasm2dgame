// Copyright 2015,2016,2017,2018,2019 SeukWon Kang (kasworld@gmail.com)
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package wasmclient

import (
	"fmt"
	"math/rand"
	"syscall/js"
	"time"

	"github.com/kasworld/direction"
	"github.com/kasworld/gowasm2dgame/enums/teamtype"
)

type Viewport2d struct {
	Canvas    js.Value
	context2d js.Value
	rnd       *rand.Rand

	W int
	H int

	background *BGObj
	cloudObjs  []*Cloud
	ballTeams  []*BallTeam
}

func NewViewport2d() *Viewport2d {
	vp := &Viewport2d{
		rnd: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
	vp.W = 1000
	vp.H = 1000

	vp.Canvas, vp.context2d = getCnv2dCtx("viewport2DCanvas")
	if !vp.context2d.Truthy() {
		fmt.Printf("fail to get context2d\n")
	}
	vp.Canvas.Set("width", vp.W)
	vp.Canvas.Set("height", vp.H)
	vp.background = NewBG()

	vp.cloudObjs = make([]*Cloud, 10)
	for i := range vp.cloudObjs {
		vp.cloudObjs[i] = NewCloud(gSprites.CloudSprite, i,
			direction.Direction_Type(i%direction.Direction_Count),
			vp.rnd.Intn(vp.W), vp.rnd.Intn(vp.H),
			vp.W, vp.H,
		)
	}

	vp.ballTeams = make([]*BallTeam, teamtype.TeamType_Count)
	for i := range vp.ballTeams {
		vp.ballTeams[i] = NewBallTeam(
			teamtype.TeamType(i),
			direction.Direction_Type(i%direction.Direction_Count),
			vp.rnd.Intn(vp.W), vp.rnd.Intn(vp.H),
			vp.W, vp.H,
		)
	}

	return vp
}

func (vp *Viewport2d) drawObj() {
	vp.background.DrawTo(vp.context2d)
	for _, v := range vp.ballTeams {
		v.DrawTo(vp.context2d)
	}
	for _, v := range vp.cloudObjs {
		v.DrawTo(vp.context2d)
	}
}

func (app *WasmClient) drawCanvas(this js.Value, args []js.Value) interface{} {
	defer func() {
		js.Global().Call("requestAnimationFrame", js.FuncOf(app.drawCanvas))
	}()
	dispCount := app.DispInterDur.GetCount()
	_ = dispCount
	act := app.DispInterDur.BeginAct()
	defer act.End()

	app.vp.drawObj()

	return nil
}
