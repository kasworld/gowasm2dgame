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
	"context"
	"fmt"
	"syscall/js"
	"time"

	"github.com/kasworld/actjitter"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_connwasm"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_pid2rspfn"
	"github.com/kasworld/intervalduration"
)

type WasmClient struct {
	DoClose              func()
	pid2recv             *w2d_pid2rspfn.PID2RspFn
	wsConn               *w2d_connwasm.Connection
	ServerJitter         *actjitter.ActJitter
	ClientJitter         *actjitter.ActJitter
	PingDur              time.Duration
	ServerClientTimeDiff time.Duration
	DispInterDur         *intervalduration.IntervalDuration

	vp *Viewport2d
}

func InitApp() {
	// dst := "ws://localhost:8080/ws"
	app := &WasmClient{
		DoClose:      func() { fmt.Println("Too early DoClose call") },
		pid2recv:     w2d_pid2rspfn.New(),
		ServerJitter: actjitter.New("Server"),
		ClientJitter: actjitter.New("Client"),
		DispInterDur: intervalduration.New("Display"),
	}
	gSprites = LoadSprites()
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
			div := js.Global().Get("document").Call("getElementById", "sysmsg")
			div.Set("innerHTML", fmt.Sprintf("%v", app.DispInterDur))
		}
	}
}
