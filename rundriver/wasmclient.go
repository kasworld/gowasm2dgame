// +build ignore

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

package main

import (
	"context"
	"fmt"
	"sync"
	"syscall/js"
	"time"

	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_connwasm"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_idcmd"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_json"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_obj"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_packet"
)

var done chan struct{}

func main() {
	InitApp()
	<-done
}

type App struct {
	wsc      *w2d_connwasm.Connection
	lasttime time.Time
	pid      uint32
}

func InitApp() {
	dst := "ws://localhost:8080/ws"
	app := App{}
	app.wsc = w2d_connwasm.New(dst, w2d_json.MarshalBodyFn, handleRecvPacket, handleSentPacket)

	var err error
	ctx := context.Background()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		err = app.wsc.Connect(ctx, &wg)
	}()
	wg.Wait()
	if err != nil {
		fmt.Printf("w2d_connwasm.Connect err %v\n", err)
		return
	}

	fmt.Printf("%v %v", dst, err)
	js.Global().Call("requestAnimationFrame", js.FuncOf(app.jsFrame))
	app.displayFrame()
}

func (app *App) jsFrame(js.Value, []js.Value) interface{} {
	app.displayFrame()
	js.Global().Call("requestAnimationFrame", js.FuncOf(app.jsFrame))
	return nil
}

func (app *App) displayFrame() {
	thistime := time.Now()
	if app.lasttime.Second() == thistime.Second() {
		return
	}
	app.lasttime = thistime
	fmt.Println(thistime)
	err := app.wsc.EnqueueSendPacket(app.makePacket())
	if err != nil {
		done <- struct{}{}
	}
}

func (app *App) makePacket() w2d_packet.Packet {
	body := w2d_obj.ReqHeartbeat_data{}
	hd := w2d_packet.Header{
		Cmd:      uint16(w2d_idcmd.Heartbeat),
		ID:       app.pid,
		FlowType: w2d_packet.Request,
	}
	app.pid++

	return w2d_packet.Packet{
		Header: hd,
		Body:   body,
	}
}

func handleRecvPacket(header w2d_packet.Header, body []byte) error {
	robj, err := w2d_json.UnmarshalPacket(header, body)
	fmt.Println(header, robj, err)
	return err
}

func handleSentPacket(header w2d_packet.Header) error {
	return nil
}
