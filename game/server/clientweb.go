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

	"github.com/gorilla/websocket"
	"github.com/kasworld/gowasm2dgame/game/gameconst"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_authorize"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_idcmd"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_packet"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_serveconnbyte"
)

func (svr *Server) initServiceWeb(ctx context.Context) {
	webMux := http.NewServeMux()
	webMux.Handle("/",
		http.FileServer(http.Dir(svr.config.ClientDataFolder)),
	)
	// webMux.HandleFunc("/ServiceInfo", svr.json_ServiceInfo)
	webMux.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		svr.serveWebSocketClient(ctx, w, r)
	})
	svr.clientWeb = &http.Server{
		Handler: webMux,
		Addr:    svr.config.ServicePort,
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
	DemuxReq2BytesAPIFnMap := [...]func(
		me interface{}, hd w2d_packet.Header, rbody []byte) (
		w2d_packet.Header, interface{}, error){
		w2d_idcmd.Invalid:   svr.bytesAPIFn_ReqInvalid,
		w2d_idcmd.MakeTeam:  svr.bytesAPIFn_ReqMakeTeam,
		w2d_idcmd.Act:       svr.bytesAPIFn_ReqAct,
		w2d_idcmd.Heartbeat: svr.bytesAPIFn_ReqHeartbeat,
	} // DemuxReq2BytesAPIFnMap

	c2sc := w2d_serveconnbyte.NewWithStats(
		nil,
		gameconst.SendBufferSize,
		w2d_authorize.NewAllSet(),
		svr.apiStat,
		svr.notiStat,
		svr.errorStat,
		DemuxReq2BytesAPIFnMap)
	c2sc.StartServeWS(ctx, wsConn,
		gameconst.ReadTimeoutSec, gameconst.WriteTimeoutSec, svr.marshalBodyFn)
	wsConn.Close()

	// connection cleanup here
}
