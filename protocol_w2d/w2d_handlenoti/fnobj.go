// Code generated by "genprotocol.exe -ver=56a40c9afaeea01bb7ff0ceed1dabb8a62deedd1dfa2e5c804d9e37c44d134ca -basedir=protocol_w2d -prefix=w2d -statstype=int"

package w2d_handlenoti

import (
	"fmt"

	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_idnoti"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_obj"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_packet"
)

// obj base demux fn map

var DemuxNoti2ObjFnMap = [...]func(me interface{}, hd w2d_packet.Header, body interface{}) error{
	w2d_idnoti.Invalid:   objRecvNotiFn_Invalid,   // Invalid
	w2d_idnoti.StageInfo: objRecvNotiFn_StageInfo, // StageInfo // game stage info to display
	w2d_idnoti.StatsInfo: objRecvNotiFn_StatsInfo, // StatsInfo // game stats info
	w2d_idnoti.StageChat: objRecvNotiFn_StageChat, // StageChat

}

// Invalid
func objRecvNotiFn_Invalid(me interface{}, hd w2d_packet.Header, body interface{}) error {
	robj, ok := body.(*w2d_obj.NotiInvalid_data)
	if !ok {
		return fmt.Errorf("packet mismatch %v", body)
	}
	return fmt.Errorf("Not implemented %v", robj)
}

// StageInfo // game stage info to display
func objRecvNotiFn_StageInfo(me interface{}, hd w2d_packet.Header, body interface{}) error {
	robj, ok := body.(*w2d_obj.NotiStageInfo_data)
	if !ok {
		return fmt.Errorf("packet mismatch %v", body)
	}
	return fmt.Errorf("Not implemented %v", robj)
}

// StatsInfo // game stats info
func objRecvNotiFn_StatsInfo(me interface{}, hd w2d_packet.Header, body interface{}) error {
	robj, ok := body.(*w2d_obj.NotiStatsInfo_data)
	if !ok {
		return fmt.Errorf("packet mismatch %v", body)
	}
	return fmt.Errorf("Not implemented %v", robj)
}

// StageChat
func objRecvNotiFn_StageChat(me interface{}, hd w2d_packet.Header, body interface{}) error {
	robj, ok := body.(*w2d_obj.NotiStageChat_data)
	if !ok {
		return fmt.Errorf("packet mismatch %v", body)
	}
	return fmt.Errorf("Not implemented %v", robj)
}
