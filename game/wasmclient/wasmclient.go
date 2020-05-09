// Copyright 2015,2016,2017,2018,2019,2020 SeukWon Kang (kasworld@gmail.com)
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
	"context"
	"fmt"
	"syscall/js"
	"time"

	"github.com/kasworld/actjitter"
	"github.com/kasworld/gowasm2dgame/config/gameconst"
	"github.com/kasworld/gowasm2dgame/lib/clientcookie"
	"github.com/kasworld/gowasm2dgame/lib/jskeypressmap"
	"github.com/kasworld/gowasm2dgame/lib/jsobj"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_connwasm"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_obj"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_pid2rspfn"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_version"
	"github.com/kasworld/gowasmlib/jslog"
	"github.com/kasworld/gowasmlib/textncount"
	"github.com/kasworld/intervalduration"
)

type WasmClient struct {
	DoClose  func()
	pid2recv *w2d_pid2rspfn.PID2RspFn
	wsConn   *w2d_connwasm.Connection

	ServerJitter         *actjitter.ActJitter
	ClientJitter         *actjitter.ActJitter
	PingDur              int64
	ServerClientTictDiff int64

	DispInterDur  *intervalduration.IntervalDuration
	systemMessage textncount.TextNCountList

	KeyboardPressedMap *jskeypressmap.KeyPressMap
	vp                 *Viewport2d

	loginData *w2d_obj.RspLogin_data
	statsInfo *w2d_obj.NotiStatsInfo_data
}

func InitApp() {
	app := &WasmClient{
		DoClose:            func() { fmt.Println("Too early DoClose call") },
		pid2recv:           w2d_pid2rspfn.New(),
		ServerJitter:       actjitter.New("Server"),
		ClientJitter:       actjitter.New("Client"),
		DispInterDur:       intervalduration.New(""),
		KeyboardPressedMap: jskeypressmap.New(),
		systemMessage:      make(textncount.TextNCountList, 0),
	}
	gSprites = LoadSprites()
	app.vp = NewViewport2d()

	jsdoc := js.Global().Get("document")
	jsobj.Hide(jsdoc.Call("getElementById", "loadmsg"))
	jsdoc.Call("getElementById", "leftinfo").Set("style",
		"color: white; position: fixed; top: 0; left: 0; overflow: hidden;")
	jsdoc.Call("getElementById", "rightinfo").Set("style",
		"color: white; position: fixed; top: 0; right: 0; overflow: hidden; text-align: right;")
	jsdoc.Call("getElementById", "centerinfo").Set("style",
		"color: white; position: fixed; top: 0%; left: 25%; overflow: hidden;")

	js.Global().Set("clearNickname", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		go clientcookie.ClearSession()
		return nil
	}))
	clientcookie.InitNickname()
	go app.enterStage()
}

func (app *WasmClient) enterStage() {
	ctx, closeCtx := context.WithCancel(context.Background())
	app.DoClose = closeCtx
	defer app.DoClose()

	var err error
	if app.loginData, err = app.NetInit(ctx); err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	defer app.Cleanup()

	if gameconst.DataVersion != app.loginData.DataVersion {
		jslog.Errorf("DataVersion mismatch client %v server %v",
			gameconst.DataVersion, app.loginData.DataVersion)
	}
	if w2d_version.ProtocolVersion != app.loginData.ProtocolVersion {
		jslog.Errorf("ProtocolVersion mismatch client %v server %v",
			w2d_version.ProtocolVersion, app.loginData.ProtocolVersion)
	}
	clientcookie.SetSession(app.loginData.SessionKey, app.loginData.NickName)
	jsdoc := js.Global().Get("document")
	jsobj.Hide(jsdoc.Call("getElementById", "titleform"))
	jsobj.Show(jsdoc.Call("getElementById", "cmdrow"))

	app.updataServiceInfo()
	js.Global().Call("requestAnimationFrame", js.FuncOf(app.drawCanvas))

	timerPingTk := time.NewTicker(time.Second)
	defer timerPingTk.Stop()
loop:
	for {
		select {
		case <-ctx.Done():
			break loop

		case <-timerPingTk.C:
			go app.reqHeartbeat()
			app.updateDebugInfo()
			app.updateTeamStatsInfo()
			app.updateSysmsg()
		}
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

	now := app.GetEstServerTick()
	app.vp.draw(now)

	return nil
}

func (app *WasmClient) GetEstServerTick() int64 {
	return time.Now().UnixNano() + app.ServerClientTictDiff
}
