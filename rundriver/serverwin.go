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

package main

import (
	"flag"
	"fmt"

	"github.com/kasworld/argdefault"
	"github.com/kasworld/configutil"
	"github.com/kasworld/go-profile"
	"github.com/kasworld/gowasm2dgame/config/dataversion"
	"github.com/kasworld/gowasm2dgame/config/serverconfig"
	"github.com/kasworld/gowasm2dgame/game/server"
	"github.com/kasworld/gowasm2dgame/lib/w2dlog"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_version"
	"github.com/kasworld/log/logflags"
	"github.com/kasworld/signalhandlewin"
	"github.com/kasworld/version"
)

var Ver = ""

func init() {
	version.Set(Ver)
}

func printVersion() {
	fmt.Println("gowasm2dgame")
	fmt.Println("Build     ", version.GetVersion())
	fmt.Println("Data      ", dataversion.DataVersion)
	fmt.Println("Protocol  ", w2d_version.ProtocolVersion)
	fmt.Println()
}

func main() {
	printVersion()

	configurl := flag.String("i", "", "server config file or url")
	signalhandlewin.AddArgs()
	profile.AddArgs()

	ads := argdefault.New(&serverconfig.Config{})
	ads.RegisterFlag()
	flag.Parse()
	config := &serverconfig.Config{
		LogLevel:      w2dlog.LL_All,
		SplitLogLevel: 0,
	}
	ads.SetDefaultToNonZeroField(config)
	if *configurl != "" {
		if err := configutil.LoadIni(*configurl, &config); err != nil {
			w2dlog.Fatal("%v", err)
		}
	}
	ads.ApplyFlagTo(config)
	if *configurl == "" {
		configutil.SaveIni("w2dserver.ini", &config)
	}
	if profile.IsCpu() {
		fn := profile.StartCPUProfile()
		defer fn()
	}

	l, err := w2dlog.NewWithDstDir(
		"",
		config.MakeLogDir(),
		logflags.DefaultValue(false).BitClear(logflags.LF_functionname),
		config.LogLevel,
		config.SplitLogLevel,
	)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	w2dlog.GlobalLogger = l

	svr := server.New(*config)
	if err := signalhandlewin.StartByArgs(svr); err != nil {
		w2dlog.Error("%v", err)
	}

	if profile.IsMem() {
		profile.WriteHeapProfile()
	}
}
