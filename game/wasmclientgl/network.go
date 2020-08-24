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

package wasmclientgl

import (
	"context"
	"fmt"
	"net/url"
	"sync"
	"syscall/js"
	"time"

	"github.com/kasworld/gowasm2dgame/lib/clientcookie"
	"github.com/kasworld/gowasm2dgame/lib/jsobj"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_connwasm"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_gob"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_idcmd"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_obj"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_packet"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_pid2rspfn"
	"github.com/kasworld/gowasmlib/jslog"
	"github.com/kasworld/gowasmlib/wasmcookie"
)

func getConnURL() string {
	loc := js.Global().Get("window").Get("location").Get("href")
	u, err := url.Parse(loc.String())
	if err != nil {
		fmt.Printf("%v\n", err)
		return ""
	}
	u.Path = "ws"
	u.Scheme = "ws"
	return u.String()
}

func (app *WasmClient) NetInit(ctx context.Context) (*w2d_obj.RspLogin_data, error) {
	app.wsConn = w2d_connwasm.New(
		getConnURL(),
		w2d_gob.MarshalBodyFn,
		app.handleRecvPacket,
		app.handleSentPacket)

	fmt.Println(getConnURL())

	var wg sync.WaitGroup

	// connect
	wg.Add(1)
	go func() {
		err := app.wsConn.Connect(ctx, &wg)
		if err != nil {
			jslog.Errorf("wsConn.Connect err %v", err)
			app.DoClose()
		}
	}()
	authkey := clientcookie.GetQuery().Get("authkey")
	nick := jsobj.GetTextValueFromInputText("nickname")
	ck := wasmcookie.GetMap()
	sessionkey := ck[clientcookie.SessionKeyName()]
	wg.Wait()
	jslog.Info("connected")

	// login
	var rtn *w2d_obj.RspLogin_data
	wg.Add(1)
	app.ReqWithRspFn(
		w2d_idcmd.Login,
		&w2d_obj.ReqLogin_data{
			SessionKey: sessionkey,
			NickName:   nick,
			AuthKey:    authkey,
		},
		func(hd w2d_packet.Header, rsp interface{}) error {
			rtn = rsp.(*w2d_obj.RspLogin_data)
			wg.Done()
			return nil
		},
	)
	wg.Wait()
	jslog.Info("logined")

	return rtn, nil
}

func (app *WasmClient) Cleanup() {
	app.wsConn.SendRecvStop()
}

func (app *WasmClient) handleSentPacket(header w2d_packet.Header) error {
	return nil
}

func (app *WasmClient) handleRecvPacket(header w2d_packet.Header, body []byte) error {
	robj, err := w2d_gob.UnmarshalPacket(header, body)
	if err != nil {
		return err
	}
	switch header.FlowType {
	default:
		return fmt.Errorf("Invalid packet type %v %v", header, robj)
	case w2d_packet.Response:
		if err := app.pid2recv.HandleRsp(header, robj); err != nil {
			return err
		}
	case w2d_packet.Notification:
		fn := DemuxNoti2ObjFnMap[header.Cmd]
		if err := fn(app, header, robj); err != nil {
			return err
		}
	}
	return nil
}

func (app *WasmClient) ReqWithRspFn(cmd w2d_idcmd.CommandID, body interface{},
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
	if err := app.wsConn.EnqueueSendPacket(spk); err != nil {
		app.wsConn.SendRecvStop()
		return fmt.Errorf("Send fail %s %v:%v %v", app, cmd, pid, err)
	}
	return nil
}

func (app *WasmClient) reqHeartbeat() error {
	return app.ReqWithRspFn(
		w2d_idcmd.Heartbeat,
		&w2d_obj.ReqHeartbeat_data{
			Tick: time.Now().UnixNano(),
		},
		func(hd w2d_packet.Header, rsp interface{}) error {
			rpk := rsp.(*w2d_obj.RspHeartbeat_data)
			pingDur := time.Now().UnixNano() - rpk.Tick
			app.PingDur = (app.PingDur + pingDur) / 2
			return nil
		},
	)
}

func (app *WasmClient) ReqWithRspFnWithAuth(cmd w2d_idcmd.CommandID, body interface{},
	fn w2d_pid2rspfn.HandleRspFn) error {
	if !app.CanUseCmd(cmd) {
		return fmt.Errorf("Cmd not allowed %v", cmd)
	}
	return app.ReqWithRspFn(cmd, body, fn)
}

func (app *WasmClient) CanUseCmd(cmd w2d_idcmd.CommandID) bool {
	if app.loginData == nil {
		return false
	}
	return app.loginData.CmdList[cmd]
}

func (app *WasmClient) sendPacket(cmd w2d_idcmd.CommandID, arg interface{}) {
	app.ReqWithRspFnWithAuth(
		cmd, arg,
		func(hd w2d_packet.Header, rsp interface{}) error {
			return nil
		},
	)
}
