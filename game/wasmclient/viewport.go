// Copyright 2015,2016,2017,2018,2019,2020,2021 SeukWon Kang (kasworld@gmail.com)
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

	"github.com/kasworld/gowasm2dgame/config/gameconst"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_obj"
	"github.com/kasworld/htmlcolors"
)

type Viewport struct {
	rnd *rand.Rand

	Canvas    js.Value
	context2d js.Value

	W float64
	H float64

	stageInfo *w2d_obj.NotiStageInfo_data
}

func NewViewport() *Viewport {
	vp := &Viewport{
		rnd: rand.New(rand.NewSource(time.Now().UnixNano())),
		W:   gameconst.StageW,
		H:   gameconst.StageH,
	}

	vp.Canvas, vp.context2d = getCnv2dCtx("viewport2DCanvas")
	vp.Canvas.Set("width", vp.W)
	vp.Canvas.Set("height", vp.H)
	return vp
}

func (vp *Viewport) Focus() {
	vp.Canvas.Call("focus")
}
func (vp *Viewport) AddEventListener(evt string, fn func(this js.Value, args []js.Value) interface{}) {
	vp.Canvas.Call("addEventListener", evt, js.FuncOf(fn))
}
func (vp *Viewport) Resized() {
	vp.Canvas.Call("setAttribute", "width", vp.W)
	vp.Canvas.Call("setAttribute", "height", vp.H)
}

func (vp *Viewport) DrawTitle() {
	win := js.Global().Get("window")
	winW := win.Get("innerWidth").Int()
	winH := win.Get("innerHeight").Int()

	msgList := []string{
		"Go 2D game",
	}

	cellW := winW / len(msgList[0])
	cellH := winH / len(msgList)
	if cellW > cellH {
		cellW = cellH
	} else {
		cellH = cellW
	}

	cnvW := cellW * len(msgList[0])
	cnvH := cellH * len(msgList)
	vp.Canvas.Call("setAttribute", "width", cnvW)
	vp.Canvas.Call("setAttribute", "height", cnvH)

	vp.context2d.Set("fillStyle", "gray")
	vp.context2d.Call("fillRect", 0, 0, cnvW, cnvH)

	fontH := cellH
	vp.context2d.Set("font", fmt.Sprintf("%dpx sans-serif", fontH))
	posx := cellW
	posy := cellH - cellH/4
	co := htmlcolors.Color24List[int(time.Now().UnixNano())%len(htmlcolors.Color24List)]
	vp.context2d.Set("fillStyle", co.ToHTMLColorString())
	for _, v := range msgList {
		vp.context2d.Call("fillText", v, posx, posy)
		posy += cellH
	}
}
