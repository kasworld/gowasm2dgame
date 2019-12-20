// Code generated by "genprotocol -ver=8cb09769c07a0cf3e7042afbf364a4eff2c960eafe6d9a8ccbb46041a984713a -basedir=. -prefix=w2d -statstype=int"

package w2d_idcmd

import "fmt"

type CommandID uint16 // use in packet header, DO NOT CHANGE
const (
	Invalid     CommandID = iota //
	EnterStage                   //
	ChatToStage                  //
	Heartbeat                    // // sayHello?

	CommandID_Count int = iota
)

var _CommandID2string = [CommandID_Count]string{
	Invalid:     "Invalid",
	EnterStage:  "EnterStage",
	ChatToStage: "ChatToStage",
	Heartbeat:   "Heartbeat",
}

func (e CommandID) String() string {
	if e >= 0 && e < CommandID(CommandID_Count) {
		return _CommandID2string[e]
	}
	return fmt.Sprintf("CommandID%d", uint16(e))
}

var _string2CommandID = map[string]CommandID{
	"Invalid":     Invalid,
	"EnterStage":  EnterStage,
	"ChatToStage": ChatToStage,
	"Heartbeat":   Heartbeat,
}

func String2CommandID(s string) (CommandID, bool) {
	v, b := _string2CommandID[s]
	return v, b
}
