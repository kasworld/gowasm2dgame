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
	"bytes"
	"context"
	"fmt"
	"syscall/js"
	"time"

	"github.com/kasworld/actjitter"
	"github.com/kasworld/gowasm2dgame/enums/acttype"
	"github.com/kasworld/gowasm2dgame/enums/teamtype"
	"github.com/kasworld/gowasm2dgame/game/gameconst"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_connwasm"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_obj"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_pid2rspfn"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_version"
	"github.com/kasworld/intervalduration"
)

type WasmClient struct {
	DoClose              func()
	pid2recv             *w2d_pid2rspfn.PID2RspFn
	wsConn               *w2d_connwasm.Connection
	ServerJitter         *actjitter.ActJitter
	ClientJitter         *actjitter.ActJitter
	PingDur              int64
	ServerClientTictDiff int64
	DispInterDur         *intervalduration.IntervalDuration

	vp        *Viewport2d
	statsInfo *w2d_obj.NotiStatsInfo_data
}

func InitApp() {
	// dst := "ws://localhost:8080/ws"
	app := &WasmClient{
		DoClose:      func() { fmt.Println("Too early DoClose call") },
		pid2recv:     w2d_pid2rspfn.New(),
		ServerJitter: actjitter.New("Server"),
		ClientJitter: actjitter.New("Client"),
	}
	gSprites = LoadSprites()
	app.DispInterDur = intervalduration.New("Display")
	app.vp = NewViewport2d()
	go app.run()
}

func (app *WasmClient) run() {
	ctx, closeCtx := context.WithCancel(context.Background())
	app.DoClose = closeCtx
	defer app.DoClose()
	if err := app.NetInit(ctx); err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	defer app.Cleanup()

	app.updataClientInfoHTML()
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
			app.updateSysmsg()
		}
	}
}

func (app *WasmClient) updateSysmsg() {
	var buf bytes.Buffer
	fmt.Fprintf(&buf,
		"%v<br/>Ping %v<br/>ServerClientTickDiff %v<br/>",
		app.DispInterDur, app.PingDur, app.ServerClientTictDiff,
	)

	buf.WriteString(`<table border=1 style="border-collapse:collapse;">
	<tr><th>act\team</th>`)
	for teami := 0; teami < teamtype.TeamType_Count; teami++ {
		fmt.Fprintf(&buf, "<th>%v</th>", teamtype.TeamType(teami))
	}
	buf.WriteString(`</tr>`)

	for acti := 0; acti < acttype.ActType_Count; acti++ {
		fmt.Fprintf(&buf, "<tr><td>%v</td>", acttype.ActType(acti))
		for teami := 0; teami < teamtype.TeamType_Count; teami++ {
			fmt.Fprintf(&buf, "<td>%v</td>", app.statsInfo.ActStats[teami][acti])
		}
		buf.WriteString(`</tr>`)
	}
	buf.WriteString(`</table>`)

	div := js.Global().Get("document").Call("getElementById", "sysmsg")
	div.Set("innerHTML", buf.String())
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

func (app *WasmClient) updataClientInfoHTML() {
	msgCopyright := `</hr>Copyright 2019 SeukWon Kang 
		<a href="https://github.com/kasworld/gowasm2dgame" target="_blank">gowasm2dgame</a>`

	var buf bytes.Buffer
	fmt.Fprintf(&buf, "gowasm2dgame webclient<br/>")
	fmt.Fprintf(&buf, "Protocol %v<br/>", w2d_version.ProtocolVersion)
	fmt.Fprintf(&buf, "Data %v<br/>", gameconst.DataVersion)
	fmt.Fprintf(&buf, "%v<br/>", msgCopyright)
	div := js.Global().Get("document").Call("getElementById", "serviceinfo")
	div.Set("innerHTML", buf.String())
}
