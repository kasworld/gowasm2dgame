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
	"flag"
	"fmt"
	"time"

	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_conntcp"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_connwsgorilla"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_error"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_idcmd"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_json"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_packet"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_pid2rspfn"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_statapierror"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_statcallapi"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_statnoti"
)

// service const
const (
	// for client
	readTimeoutSec  = 6 * time.Second
	writeTimeoutSec = 3 * time.Second
)

var gMarshalBodyFn func(body interface{}, oldBuffToAppend []byte) ([]byte, byte, error)
var gUnmarshalPacket func(h w2d_packet.Header, bodyData []byte) (interface{}, error)

func main() {
	addr := flag.String("addr", "localhost:8080", "server addr")
	flag.Parse()
	fmt.Printf("addr %v \n", *addr)

	gMarshalBodyFn = w2d_json.MarshalBodyFn
	gUnmarshalPacket = w2d_json.UnmarshalPacket

	app := NewApp(*addr)
	app.Run()
}

type App struct {
	addr string

	c2scWS            *w2d_connwsgorilla.Connection
	c2scTCP           *w2d_conntcp.Connection
	EnqueueSendPacket func(pk w2d_packet.Packet) error

	sendRecvStop func()
	apistat      *w2d_statcallapi.StatCallAPI
	pid2statobj  *w2d_statcallapi.PacketID2StatObj
	notistat     *w2d_statnoti.StatNotification
	errstat      *w2d_statapierror.StatAPIError
	pid2recv     *w2d_pid2rspfn.PID2RspFn
}

func NewApp(addr string) *App {
	app := &App{
		addr:        addr,
		apistat:     w2d_statcallapi.New(),
		pid2statobj: w2d_statcallapi.NewPacketID2StatObj(),
		notistat:    w2d_statnoti.New(),
		errstat:     w2d_statapierror.New(),
		pid2recv:    w2d_pid2rspfn.New(),
	}
	app.sendRecvStop = func() {
		fmt.Printf("Too early sendRecvStop call\n")
	}
	app.EnqueueSendPacket = func(pk w2d_packet.Packet) error {
		fmt.Printf("Too early EnqueueSendPacket call\n")
		return nil
	}
	return app
}

func (app *App) Run() {
	ctx, stopFn := context.WithCancel(context.Background())
	app.sendRecvStop = stopFn
	defer app.sendRecvStop()

	go app.connectWS(ctx)

	time.Sleep(time.Second)
	app.sendTestPacket()
	app.sendTestPacket()

	for {
		select {
		case <-ctx.Done():
			return
		}
	}
}

func (app *App) sendTestPacket() error {
	return nil
	// return app.ReqWithRspFn(
	// 	w2d_idcmd.Heartbeat,
	// 	w2d_obj.ReqHeartbeat_data{time.Now()},
	// 	func(hd w2d_packet.Header, rbody interface{}) error {
	// 		robj, err := gUnmarshalPacket(hd, rbody.([]byte))
	// 		if err != nil {
	// 			return fmt.Errorf("Packet type miss match %v", rbody)
	// 		}
	// 		recvBody, ok := robj.(*w2d_obj.RspHeartbeat_data)
	// 		if !ok {
	// 			return fmt.Errorf("Packet type miss match %v", robj)
	// 		}

	// 		fmt.Printf("ping %v\n", time.Now().Sub(recvBody.Now))
	// 		return nil
	// 	},
	// )
}

func (app *App) connectWS(ctx context.Context) {
	app.c2scWS = w2d_connwsgorilla.New(
		readTimeoutSec, writeTimeoutSec,
		gMarshalBodyFn,
		app.handleRecvPacket,
		app.handleSentPacket,
	)
	if err := app.c2scWS.ConnectTo(app.addr); err != nil {
		fmt.Printf("%v\n", err)
		app.sendRecvStop()
		return
	}
	app.EnqueueSendPacket = app.c2scWS.EnqueueSendPacket
	app.c2scWS.Run(ctx)
}

func (app *App) handleSentPacket(header w2d_packet.Header) error {
	if err := app.apistat.AfterSendReq(header); err != nil {
		return err
	}
	return nil
}

func (app *App) handleRecvPacket(header w2d_packet.Header, body []byte) error {
	switch header.FlowType {
	default:
		return fmt.Errorf("Invalid packet type %v %v", header, body)
	case w2d_packet.Notification:
		// noti stat
		app.notistat.Add(header)
		//process noti here
		// robj, err := w2d_json.UnmarshalPacket(header, body)

	case w2d_packet.Response:
		// error stat
		app.errstat.Inc(w2d_idcmd.CommandID(header.Cmd), header.ErrorCode)
		// api stat
		if err := app.apistat.AfterRecvRsp(header); err != nil {
			fmt.Printf("%v %v\n", app, err)
			return err
		}
		psobj := app.pid2statobj.Get(header.ID)
		if psobj == nil {
			return fmt.Errorf("no statobj for %v", header.ID)
		}
		psobj.CallServerEnd(header.ErrorCode == w2d_error.None)
		app.pid2statobj.Del(header.ID)

		// process response
		if err := app.pid2recv.HandleRsp(header, body); err != nil {
			return err
		}

		// send new test packet
		go app.sendTestPacket()
	}
	return nil
}

func (app *App) ReqWithRspFn(cmd w2d_idcmd.CommandID, body interface{},
	fn w2d_pid2rspfn.HandleRspFn) error {

	pid := app.pid2recv.NewPID(fn)
	spk := w2d_packet.Packet{
		Header: w2d_packet.Header{
			Cmd:      uint16(cmd),
			ID:       pid,
			FlowType: w2d_packet.Request,
		},
		Body: body,
	}

	// add api stat
	psobj, err := app.apistat.BeforeSendReq(spk.Header)
	if err != nil {
		return nil
	}
	app.pid2statobj.Add(spk.Header.ID, psobj)

	if err := app.EnqueueSendPacket(spk); err != nil {
		fmt.Printf("End %v %v %v\n", app, spk, err)
		app.sendRecvStop()
		return fmt.Errorf("Send fail %v %v", app, err)
	}
	return nil
}