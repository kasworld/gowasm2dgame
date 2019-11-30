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
	"math/rand"
	"net/http"
	"time"

	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_gob"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_idnoti"

	"github.com/kasworld/prettystring"

	"github.com/kasworld/actpersec"
	"github.com/kasworld/gowasm2dgame/game/connectionmanager"
	"github.com/kasworld/gowasm2dgame/game/serverconfig"
	"github.com/kasworld/gowasm2dgame/game/stage"
	"github.com/kasworld/gowasm2dgame/lib/w2dlog"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_idcmd"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_packet"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_statapierror"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_statnoti"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_statserveapi"
	"github.com/kasworld/signalhandle"
	"github.com/kasworld/weblib/retrylistenandserve"
)

type Server struct {
	rnd       *rand.Rand      `prettystring:"hide"`
	log       *w2dlog.LogBase `prettystring:"hide"`
	config    serverconfig.Config
	adminWeb  *http.Server `prettystring:"simple"`
	clientWeb *http.Server `prettystring:"simple"`
	startTime time.Time    `prettystring:"simple"`

	sendRecvStop func()
	SendStat     *actpersec.ActPerSec `prettystring:"simple"`
	RecvStat     *actpersec.ActPerSec `prettystring:"simple"`

	apiStat   *w2d_statserveapi.StatServeAPI
	notiStat  *w2d_statnoti.StatNotification
	errorStat *w2d_statapierror.StatAPIError

	marshalBodyFn          func(body interface{}, oldBuffToAppend []byte) ([]byte, byte, error)
	unmarshalPacketFn      func(h w2d_packet.Header, bodyData []byte) (interface{}, error)
	DemuxReq2BytesAPIFnMap [w2d_idcmd.CommandID_Count]func(
		me interface{}, hd w2d_packet.Header, rbody []byte) (
		w2d_packet.Header, interface{}, error)

	connManager *connectionmanager.Manager
	stage       *stage.Stage
}

func New(config serverconfig.Config) *Server {
	svr := &Server{
		config: config,
		log:    w2dlog.GlobalLogger,
		rnd:    rand.New(rand.NewSource(time.Now().UnixNano())),

		SendStat: actpersec.New(),
		RecvStat: actpersec.New(),

		apiStat:     w2d_statserveapi.New(),
		notiStat:    w2d_statnoti.New(),
		errorStat:   w2d_statapierror.New(),
		connManager: connectionmanager.New(),
	}
	svr.sendRecvStop = func() {
		fmt.Printf("Too early sendRecvStop call\n")
	}
	svr.marshalBodyFn = w2d_gob.MarshalBodyFn
	svr.unmarshalPacketFn = w2d_gob.UnmarshalPacket
	svr.DemuxReq2BytesAPIFnMap = [...]func(
		me interface{}, hd w2d_packet.Header, rbody []byte) (
		w2d_packet.Header, interface{}, error){
		w2d_idcmd.Invalid:   svr.bytesAPIFn_ReqInvalid,
		w2d_idcmd.MakeTeam:  svr.bytesAPIFn_ReqMakeTeam,
		w2d_idcmd.Act:       svr.bytesAPIFn_ReqAct,
		w2d_idcmd.Heartbeat: svr.bytesAPIFn_ReqHeartbeat,
	} // DemuxReq2BytesAPIFnMap
	svr.stage = stage.New(svr.log)
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
	fmt.Println(prettystring.PrettyString(svr.config, 4))
	svr.startTime = time.Now()

	ctx, stopFn := context.WithCancel(ctx)
	svr.sendRecvStop = stopFn
	defer svr.sendRecvStop()

	svr.initAdminWeb()
	svr.initServiceWeb(ctx)

	fmt.Printf("open admin web\nhttp://localhost%v/\n", svr.config.AdminPort)
	fmt.Printf("open client web\nhttp://localhost%v/\n", svr.config.ServicePort)

	go retrylistenandserve.RetryListenAndServe(svr.adminWeb, svr.log, "serveAdminWeb")
	go retrylistenandserve.RetryListenAndServe(svr.clientWeb, svr.log, "serveServiceWeb")

	timerInfoTk := time.NewTicker(1 * time.Second)
	defer timerInfoTk.Stop()

	turnDur := time.Duration(float64(time.Second) / svr.config.ActTurnPerSec)
	timerTurnTk := time.NewTicker(turnDur)
	defer timerTurnTk.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-timerInfoTk.C:
			svr.SendStat.UpdateLap()
			svr.RecvStat.UpdateLap()
			si := svr.stage.ToStatsInfo()
			conlist := svr.connManager.GetList()
			for _, v := range conlist {
				v.SendNotiPacket(w2d_idnoti.StatsInfo,
					si,
				)
			}

		case <-timerTurnTk.C:
			svr.stage.Turn()
			si := svr.stage.ToStageInfo()
			conlist := svr.connManager.GetList()
			for _, v := range conlist {
				v.SendNotiPacket(w2d_idnoti.StageInfo,
					si,
				)
			}
		}
	}
}
