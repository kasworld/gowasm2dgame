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

package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/kasworld/gowasm2dgame/config/authdata"
	"github.com/kasworld/gowasm2dgame/config/gameconst"
	"github.com/kasworld/gowasm2dgame/lib/conndata"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_gob"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_serveconnbyte"
	"github.com/kasworld/uuidstr"
)

func (svr *Server) initServiceWeb(ctx context.Context) {
	webMux := http.NewServeMux()
	webMux.Handle("/",
		http.FileServer(http.Dir(svr.config.ClientDataFolder)),
	)
	webMux.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		svr.serveWebSocketClient(ctx, w, r)
	})
	svr.clientWeb = &http.Server{
		Handler: webMux,
		Addr:    fmt.Sprintf(":%v", svr.config.ServicePort),
	}
	svr.marshalBodyFn = w2d_gob.MarshalBodyFn
	svr.unmarshalPacketFn = w2d_gob.UnmarshalPacket
	svr.setFnMap()
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

	connData := &conndata.ConnData{
		UUID:       uuidstr.New(),
		RemoteAddr: r.RemoteAddr,
		// Session set at login
	}
	c2sc := w2d_serveconnbyte.NewWithStats(
		connData,
		gameconst.SendBufferSize,
		authdata.NewPreLoginAuthorCmdIDList(),
		svr.SendStat, svr.RecvStat,
		svr.apiStat,
		svr.notiStat,
		svr.errorStat,
		svr.DemuxReq2BytesAPIFnMap)

	// add to conn manager
	svr.connManager.Add(connData.UUID, c2sc)

	// start client service
	c2sc.StartServeWS(ctx, wsConn,
		gameconst.ReadTimeoutSec, gameconst.WriteTimeoutSec, svr.marshalBodyFn)
	wsConn.Close()
	// connection cleanup here

	// del from conn manager
	svr.connManager.Del(connData.UUID)
	// stage may be changed
	if ss := connData.Session; ss != nil {
		stg := svr.stageManager.GetByUUID(ss.StageID)
		if stg != nil {
			stg.GetConnManager().Del(connData.UUID)
		}
	}
}
