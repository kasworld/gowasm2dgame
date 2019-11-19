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

package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/kasworld/gowasm2dgame/game/serverconfig"

	"github.com/gorilla/websocket"
	"github.com/kasworld/gowasm2dgame/game/gameconst"
	"github.com/kasworld/gowasm2dgame/lib/w2dlog"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_authorize"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_error"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_idcmd"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_json"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_obj"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_packet"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_serveconnbyte"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_statapierror"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_statnoti"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_statserveapi"
	"github.com/kasworld/signalhandle"
)

type Server struct {
	config                 serverconfig.Config
	sendRecvStop           func()
	apiStat                *w2d_statserveapi.StatServeAPI
	notiStat               *w2d_statnoti.StatNotification
	errStat                *w2d_statapierror.StatAPIError
	marshalBodyFn          func(body interface{}, oldBuffToAppend []byte) ([]byte, byte, error)
	unmarshalPacketFn      func(h w2d_packet.Header, bodyData []byte) (interface{}, error)
	DemuxReq2BytesAPIFnMap [w2d_idcmd.CommandID_Count]func(
		me interface{}, hd w2d_packet.Header, rbody []byte) (
		w2d_packet.Header, interface{}, error)
}

func New(config serverconfig.Config) *Server {
	svr := &Server{
		config:   config,
		apiStat:  w2d_statserveapi.New(),
		notiStat: w2d_statnoti.New(),
		errStat:  w2d_statapierror.New(),
	}
	svr.sendRecvStop = func() {
		fmt.Printf("Too early sendRecvStop call\n")
	}
	svr.marshalBodyFn = w2d_json.MarshalBodyFn
	svr.unmarshalPacketFn = w2d_json.UnmarshalPacket
	return svr
}

// called from signal handler
func (svr *Server) GetServiceLockFilename() string {
	return svr.config.MakePIDFileFullpath()
}

// called from signal handler
func (svr *Server) GetLogger() signalhandle.LoggerI {
	return w2dlog.GlobalLogger
}

// called from signal handler
func (svr *Server) ServiceInit() error {
	return nil
}

// called from signal handler
func (svr *Server) ServiceCleanup() {
}

// called from signal handler
func (svr *Server) ServiceMain(ctx context.Context) {
	ctx, stopFn := context.WithCancel(context.Background())
	svr.sendRecvStop = stopFn
	defer svr.sendRecvStop()

	go svr.serveHTTP(ctx, svr.config.ServicePort, svr.config.ClientDataFolder)
	timerInfoTk := time.NewTicker(1 * time.Second)
	defer timerInfoTk.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-timerInfoTk.C:

		}
	}
}

func (svr *Server) serveHTTP(ctx context.Context, port string, folder string) {
	webMux := http.NewServeMux()
	webMux.Handle("/",
		http.FileServer(http.Dir(folder)),
	)
	webMux.HandleFunc("/svr", func(w http.ResponseWriter, r *http.Request) {
		svr.serveWebSocketClient(ctx, w, r)
	})
	if err := http.ListenAndServe(port, webMux); err != nil {
		fmt.Println(err.Error())
	}
}

func CheckOrigin(r *http.Request) bool {
	return true
}

func (svr *Server) serveWebSocketClient(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		CheckOrigin: CheckOrigin,
	}
	wsConn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("upgrade %v\n", err)
		return
	}
	c2sc := w2d_serveconnbyte.NewWithStats(
		nil,
		gameconst.SendBufferSize,
		w2d_authorize.NewAllSet(),
		svr.apiStat,
		svr.notiStat,
		svr.errStat,
		DemuxReq2BytesAPIFnMap)
	c2sc.StartServeWS(ctx, wsConn,
		gameconst.ReadTimeoutSec, gameconst.WriteTimeoutSec, svr.marshalBodyFn)
	wsConn.Close()
}

///////////////////////////////////////////////////////////////

var DemuxReq2BytesAPIFnMap = [...]func(
	me interface{}, hd w2d_packet.Header, rbody []byte) (
	w2d_packet.Header, interface{}, error){
	w2d_idcmd.Invalid:   bytesAPIFn_ReqInvalid,
	w2d_idcmd.MakeTeam:  bytesAPIFn_ReqMakeTeam,
	w2d_idcmd.Act:       bytesAPIFn_ReqAct,
	w2d_idcmd.State:     bytesAPIFn_ReqState,
	w2d_idcmd.Heartbeat: bytesAPIFn_ReqHeartbeat,
} // DemuxReq2BytesAPIFnMap

func bytesAPIFn_ReqInvalid(
	me interface{}, hd w2d_packet.Header, rbody []byte) (
	w2d_packet.Header, interface{}, error) {
	// robj, err := w2d_json.UnmarshalPacket(hd, rbody)
	// if err != nil {
	// 	return hd, nil, fmt.Errorf("Packet type miss match %v", rbody)
	// }
	// recvBody, ok := robj.(*w2d_obj.ReqInvalid_data)
	// if !ok {
	// 	return hd, nil, fmt.Errorf("Packet type miss match %v", robj)
	// }
	// _ = recvBody

	sendHeader := w2d_packet.Header{
		ErrorCode: w2d_error.None,
	}
	sendBody := &w2d_obj.RspInvalid_data{}
	return sendHeader, sendBody, nil
}

func bytesAPIFn_ReqMakeTeam(
	me interface{}, hd w2d_packet.Header, rbody []byte) (
	w2d_packet.Header, interface{}, error) {
	// robj, err := w2d_json.UnmarshalPacket(hd, rbody)
	// if err != nil {
	// 	return hd, nil, fmt.Errorf("Packet type miss match %v", rbody)
	// }
	// recvBody, ok := robj.(*w2d_obj.ReqMakeTeam_data)
	// if !ok {
	// 	return hd, nil, fmt.Errorf("Packet type miss match %v", robj)
	// }
	// _ = recvBody

	sendHeader := w2d_packet.Header{
		ErrorCode: w2d_error.None,
	}
	sendBody := &w2d_obj.RspMakeTeam_data{}
	return sendHeader, sendBody, nil
}

func bytesAPIFn_ReqAct(
	me interface{}, hd w2d_packet.Header, rbody []byte) (
	w2d_packet.Header, interface{}, error) {
	// robj, err := w2d_json.UnmarshalPacket(hd, rbody)
	// if err != nil {
	// 	return hd, nil, fmt.Errorf("Packet type miss match %v", rbody)
	// }
	// recvBody, ok := robj.(*w2d_obj.ReqAct_data)
	// if !ok {
	// 	return hd, nil, fmt.Errorf("Packet type miss match %v", robj)
	// }
	// _ = recvBody

	sendHeader := w2d_packet.Header{
		ErrorCode: w2d_error.None,
	}
	sendBody := &w2d_obj.RspAct_data{}
	return sendHeader, sendBody, nil
}

func bytesAPIFn_ReqState(
	me interface{}, hd w2d_packet.Header, rbody []byte) (
	w2d_packet.Header, interface{}, error) {
	// robj, err := w2d_json.UnmarshalPacket(hd, rbody)
	// if err != nil {
	// 	return hd, nil, fmt.Errorf("Packet type miss match %v", rbody)
	// }
	// recvBody, ok := robj.(*w2d_obj.ReqState_data)
	// if !ok {
	// 	return hd, nil, fmt.Errorf("Packet type miss match %v", robj)
	// }
	// _ = recvBody

	sendHeader := w2d_packet.Header{
		ErrorCode: w2d_error.None,
	}
	sendBody := &w2d_obj.RspState_data{}
	return sendHeader, sendBody, nil
}

func bytesAPIFn_ReqHeartbeat(
	me interface{}, hd w2d_packet.Header, rbody []byte) (
	w2d_packet.Header, interface{}, error) {
	// robj, err := w2d_json.UnmarshalPacket(hd, rbody)
	// if err != nil {
	// 	return hd, nil, fmt.Errorf("Packet type miss match %v", rbody)
	// }
	// recvBody, ok := robj.(*w2d_obj.ReqHeartbeat_data)
	// if !ok {
	// 	return hd, nil, fmt.Errorf("Packet type miss match %v", robj)
	// }
	// _ = recvBody

	sendHeader := w2d_packet.Header{
		ErrorCode: w2d_error.None,
	}
	sendBody := &w2d_obj.RspHeartbeat_data{}
	return sendHeader, sendBody, nil
}