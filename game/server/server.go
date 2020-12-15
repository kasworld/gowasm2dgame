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

package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/kasworld/actpersec"
	"github.com/kasworld/g2rand"
	"github.com/kasworld/gowasm2dgame/config/gameconst"
	"github.com/kasworld/gowasm2dgame/config/serverconfig"
	"github.com/kasworld/gowasm2dgame/game/stage"
	"github.com/kasworld/gowasm2dgame/game/stagemanager"
	"github.com/kasworld/gowasm2dgame/lib/sessionmanager"
	"github.com/kasworld/gowasm2dgame/lib/w2dlog"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_connbytemanager"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_idcmd"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_packet"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_statapierror"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_statnoti"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_statserveapi"
	"github.com/kasworld/log/logflags"
	"github.com/kasworld/prettystring"
	"github.com/kasworld/weblib/retrylistenandserve"
)

type Server struct {
	rnd       *g2rand.G2Rand  `prettystring:"hide"`
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

	connManager    *w2d_connbytemanager.Manager
	sessionManager *sessionmanager.SessionManager

	stageManager *stagemanager.Manager
}

func New(config serverconfig.Config) *Server {
	fmt.Printf("%v\n", config.StringForm())

	if config.BaseLogDir != "" {
		log, err := w2dlog.NewWithDstDir(
			"",
			config.MakeLogDir(),
			logflags.DefaultValue(false).BitClear(logflags.LF_functionname),
			config.LogLevel,
			config.SplitLogLevel,
		)
		if err == nil {
			w2dlog.GlobalLogger = log
		} else {
			fmt.Printf("%v\n", err)
			w2dlog.GlobalLogger.SetFlags(
				w2dlog.GlobalLogger.GetFlags().BitClear(logflags.LF_functionname))
			w2dlog.GlobalLogger.SetLevel(
				config.LogLevel)
		}
	} else {
		w2dlog.GlobalLogger.SetFlags(
			w2dlog.GlobalLogger.GetFlags().BitClear(logflags.LF_functionname))
		w2dlog.GlobalLogger.SetLevel(
			config.LogLevel)
	}
	svr := &Server{
		config: config,
		log:    w2dlog.GlobalLogger,
		rnd:    g2rand.New(),

		SendStat: actpersec.New(),
		RecvStat: actpersec.New(),

		apiStat:        w2d_statserveapi.New(),
		notiStat:       w2d_statnoti.New(),
		errorStat:      w2d_statapierror.New(),
		connManager:    w2d_connbytemanager.New(),
		sessionManager: sessionmanager.New("", 100, w2dlog.GlobalLogger),

		stageManager: stagemanager.New(w2dlog.GlobalLogger),
	}
	svr.sendRecvStop = func() {
		fmt.Printf("Too early sendRecvStop call\n")
	}
	return svr
}

// called from signal handler
func (svr *Server) GetServiceLockFilename() string {
	return svr.config.MakePIDFileFullpath()
}

// called from signal handler
// return implement signalhandle.LoggerI
func (svr *Server) GetLogger() interface{} {
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
func (svr *Server) ServiceMain(mainctx context.Context) {
	fmt.Println(prettystring.PrettyString(svr.config, 4))
	svr.startTime = time.Now()

	ctx, stopFn := context.WithCancel(mainctx)
	svr.sendRecvStop = stopFn
	defer svr.sendRecvStop()

	svr.initAdminWeb()
	svr.initServiceWeb(ctx)

	fmt.Printf("WebAdmin  : %v:%v id:%v pass:%v\n",
		svr.config.ServiceHostBase, svr.config.AdminPort, svr.config.WebAdminID, svr.config.WebAdminPass)
	fmt.Printf("WebClient : %v:%v/\n", svr.config.ServiceHostBase, svr.config.ServicePort)
	fmt.Printf("WebClientGL : %v:%v//gl.html\n", svr.config.ServiceHostBase, svr.config.ServicePort)

	go retrylistenandserve.RetryListenAndServe(svr.adminWeb, svr.log, "serveAdminWeb")
	go retrylistenandserve.RetryListenAndServe(svr.clientWeb, svr.log, "serveServiceWeb")

	for i := 0; i < gameconst.StagePerServer; i++ {
		stg := stage.New(svr.log, svr.config, svr.rnd.Int63())
		svr.stageManager.Add(stg)
		go stg.Run(ctx)
	}

	timerInfoTk := time.NewTicker(1 * time.Second)
	defer timerInfoTk.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-timerInfoTk.C:
			svr.SendStat.UpdateLap()
			svr.RecvStat.UpdateLap()
		}
	}
}
