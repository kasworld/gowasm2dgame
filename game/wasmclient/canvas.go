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
	"github.com/kasworld/gowasm2dgame/enums/gameobjtype"
)

type Viewport2d struct {
	Canvas    js.Value
	context2d js.Value
	rnd       *rand.Rand

	W int
	H int

	background *BGObj
	clouds     []*Cloud
	ball       []*Ball

	explodeetc  *Sprite
	explodeball *Sprite
	spawn       *Sprite
}

func NewViewport2d() *Viewport2d {
	vp := &Viewport2d{
		rnd: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
	vp.W = 1000
	vp.H = 1000

	vp.Canvas, vp.context2d = getCnv2dCtx("viewport2DCanvas")
	if !vp.Canvas.Truthy() {
		fmt.Printf("fail to get canvas viewport2DCanvas\n")
	}
	if !vp.context2d.Truthy() {
		fmt.Printf("fail to get context2d\n")
	}
	vp.Canvas.Set("width", vp.W)
	vp.Canvas.Set("height", vp.H)

	vp.background = NewBG()

	cloudsp := LoadSpriteXYN("clouds", "cloudStore", 1, 4)
	vp.clouds = make([]*Cloud, 10)
	for i := range vp.clouds {
		vp.clouds[i] = NewCloud(cloudsp, i,
			direction.Direction_Type(i%direction.Direction_Count),
			vp.rnd.Intn(vp.W), vp.rnd.Intn(vp.H),
			vp.W, vp.H,
		)
	}

	team := [...]struct {
		name  string
		index int
		value int
	}{
		{"red", 0, 255},
		{"blue", 1, 255},
		{"green", 2, 255},
	}
	vp.ball = make([]*Ball, len(team))
	for i := range vp.ball {
		ballSprite := vp.loadBallSprite(
			team[i].name,
			team[i].index,
			team[i].value)
		vp.ball[i] = NewBall(
			ballSprite,
			direction.Direction_Type(i%direction.Direction_Count),
			vp.rnd.Intn(vp.W), vp.rnd.Intn(vp.H),
			vp.W, vp.H,
		)
	}

	vp.spawn = LoadSpriteXYN("spawn", "spawnStore", 1, 6)
	vp.explodeetc = LoadSpriteXYN("explodeetc", "explodeetcStore", 1, 8)
	vp.explodeball = LoadSpriteXYN("explodeball", "explodeballStore", 8, 1)

	/*
		vp.grayball = vp.LoadImage("grayball") // color change
		vp.spiral = vp.LoadImage("spiral")     // color change, rotate (0, 360, 10)
		// ('bounceball', "grayball.png", None),
		// ('bullet', "grayball.png", None),
		// ('hommingbullet', "spiral.png", (0, 360, 10)),
		// ('superbullet', "spiral.png", (0, 360, 10)),
		// ('circularbullet', "grayball.png", None),
		// ('shield', "grayball.png", None),
		// ('supershield', "spiral.png", (0, 360, 10))

		vp.clouds = vp.LoadImage("clouds")           // slicearg=(1, 4
		vp.explodeetc = vp.LoadImage("explodeetc")   // slicearg=(1, 8, spriteexplosioneffect
		vp.explodeball = vp.LoadImage("explodeball") // slicearg=(8, 1 ballexplosioneffect
		vp.spawn = vp.LoadImage("spawn")             // slicearg=(1, 6, reverse spawneffect
	*/
	return vp
}

func (vp *Viewport2d) loadBallSprite(
	teamname string,
	filterindex int, filtervalue int,
) [gameobjtype.GameObjType_Count]*Sprite {
	var rtn [gameobjtype.GameObjType_Count]*Sprite
	rtn[gameobjtype.Ball] = LoadSpriteXYNResize(
		"grayball", teamname+"_ball",
		1, 1,
		gameobjtype.Attrib[gameobjtype.Ball].Size,
		gameobjtype.Attrib[gameobjtype.Ball].Size,
	)

	rtn[gameobjtype.Shield] = LoadSpriteXYNResize(
		"grayball", teamname+"_shield",
		1, 1,
		gameobjtype.Attrib[gameobjtype.Shield].Size,
		gameobjtype.Attrib[gameobjtype.Shield].Size,
	)

	rtn[gameobjtype.SuperShield] = LoadSpriteRotateResize(
		"spiral", teamname+"_supershield",
		0, 360, 10,
		gameobjtype.Attrib[gameobjtype.SuperShield].Size,
		gameobjtype.Attrib[gameobjtype.SuperShield].Size,
	)
	rtn[gameobjtype.HommingShield] = LoadSpriteRotateResize(
		"spiral", teamname+"_hommingshield",
		0, 360, 10,
		gameobjtype.Attrib[gameobjtype.HommingShield].Size,
		gameobjtype.Attrib[gameobjtype.HommingShield].Size,
	)
	rtn[gameobjtype.Bullet] = LoadSpriteXYNResize(
		"grayball", teamname+"_bullet",
		1, 1,
		gameobjtype.Attrib[gameobjtype.Bullet].Size,
		gameobjtype.Attrib[gameobjtype.Bullet].Size,
	)
	rtn[gameobjtype.SuperBullet] = LoadSpriteRotateResize(
		"spiral", teamname+"_superbullet",
		0, 360, 10,
		gameobjtype.Attrib[gameobjtype.SuperBullet].Size,
		gameobjtype.Attrib[gameobjtype.SuperBullet].Size,
	)
	rtn[gameobjtype.HommingBullet] = LoadSpriteRotateResize(
		"spiral", teamname+"_hommingbullet",
		0, 360, 10,
		gameobjtype.Attrib[gameobjtype.HommingBullet].Size,
		gameobjtype.Attrib[gameobjtype.HommingBullet].Size,
	)
	for i := range rtn {
		rtn[i].Filter(filterindex, filtervalue)
	}
	return rtn
}

func (vp *Viewport2d) drawBG() {
	vp.background.DrawTo(vp.context2d)
}

func (vp *Viewport2d) drawObj() {
	for _, v := range vp.ball {
		v.DrawTo(vp.context2d)
	}
	for _, v := range vp.clouds {
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

	app.vp.drawBG()
	app.vp.drawObj()

	return nil
}
