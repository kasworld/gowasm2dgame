// Copyright 2015,2016,2017,2018,2019,2020 SeukWon Kang (kasworld@gmail.com)

package serverconfig

import (
	"fmt"
	"path/filepath"

	"github.com/kasworld/gowasm2dgame/lib/w2dlog"
	"github.com/kasworld/prettystring"
)

type Config struct {
	LogLevel              w2dlog.LL_Type `default:"31" argname:""`
	SplitLogLevel         w2dlog.LL_Type `default:"0" argname:""`
	BaseLogDir            string         `default:"/tmp/"  argname:""`
	ServerDataFolder      string         `default:"./serverdata" argname:""`
	ClientDataFolder      string         `default:"./clientdata" argname:""`
	ServicePort           string         `default:"24101"  argname:""`
	AdminPort             string         `default:"24201"  argname:""`
	WebAdminID            string         `default:"root" argname:""`
	WebAdminPass          string         `default:"password" argname:"" prettystring:"hidevalue"`
	ServiceHostBase       string         `default:"http://localhost" argname:""`
	ConcurrentConnections int            `default:"1000" argname:""`
	ActTurnPerSec         float64        `default:"30.0" argname:""`
}

func (config Config) MakeLogDir() string {
	rstr := filepath.Join(
		config.BaseLogDir,
		"w2dserver.logfiles",
	)
	rtn, err := filepath.Abs(rstr)
	if err != nil {
		fmt.Println(rstr, rtn, err.Error())
		return rstr
	}
	return rtn
}

func (config Config) MakePIDFileFullpath() string {
	rstr := filepath.Join(
		config.BaseLogDir,
		"w2dserver.pid",
	)
	rtn, err := filepath.Abs(rstr)
	if err != nil {
		fmt.Println(rstr, rtn, err.Error())
		return rstr
	}
	return rtn
}

func (config Config) MakeOutfileFullpath() string {
	rstr := "w2dserver.out"
	rtn, err := filepath.Abs(rstr)
	if err != nil {
		fmt.Println(rstr, rtn, err.Error())
		return rstr
	}
	return rtn
}

func (config Config) StringForm() string {
	return prettystring.PrettyString(config, 4)
}
