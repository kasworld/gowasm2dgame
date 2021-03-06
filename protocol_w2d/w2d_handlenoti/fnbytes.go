// Code generated by "genprotocol.exe -ver=56a40c9afaeea01bb7ff0ceed1dabb8a62deedd1dfa2e5c804d9e37c44d134ca -basedir=protocol_w2d -prefix=w2d -statstype=int"

package w2d_handlenoti

import (
	"fmt"

	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_idnoti"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_json"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_obj"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_packet"
)

// bytes base demux fn map

var DemuxNoti2ByteFnMap = [...]func(me interface{}, hd w2d_packet.Header, rbody []byte) error{
	w2d_idnoti.Invalid:   bytesRecvNotiFn_Invalid,   // Invalid
	w2d_idnoti.StageInfo: bytesRecvNotiFn_StageInfo, // StageInfo // game stage info to display
	w2d_idnoti.StatsInfo: bytesRecvNotiFn_StatsInfo, // StatsInfo // game stats info
	w2d_idnoti.StageChat: bytesRecvNotiFn_StageChat, // StageChat

}

// Invalid
func bytesRecvNotiFn_Invalid(me interface{}, hd w2d_packet.Header, rbody []byte) error {
	robj, err := w2d_json.UnmarshalPacket(hd, rbody)
	if err != nil {
		return fmt.Errorf("Packet type miss match %v", rbody)
	}
	recved, ok := robj.(*w2d_obj.NotiInvalid_data)
	if !ok {
		return fmt.Errorf("packet mismatch %v", robj)
	}
	return fmt.Errorf("Not implemented %v", recved)
}

// StageInfo // game stage info to display
func bytesRecvNotiFn_StageInfo(me interface{}, hd w2d_packet.Header, rbody []byte) error {
	robj, err := w2d_json.UnmarshalPacket(hd, rbody)
	if err != nil {
		return fmt.Errorf("Packet type miss match %v", rbody)
	}
	recved, ok := robj.(*w2d_obj.NotiStageInfo_data)
	if !ok {
		return fmt.Errorf("packet mismatch %v", robj)
	}
	return fmt.Errorf("Not implemented %v", recved)
}

// StatsInfo // game stats info
func bytesRecvNotiFn_StatsInfo(me interface{}, hd w2d_packet.Header, rbody []byte) error {
	robj, err := w2d_json.UnmarshalPacket(hd, rbody)
	if err != nil {
		return fmt.Errorf("Packet type miss match %v", rbody)
	}
	recved, ok := robj.(*w2d_obj.NotiStatsInfo_data)
	if !ok {
		return fmt.Errorf("packet mismatch %v", robj)
	}
	return fmt.Errorf("Not implemented %v", recved)
}

// StageChat
func bytesRecvNotiFn_StageChat(me interface{}, hd w2d_packet.Header, rbody []byte) error {
	robj, err := w2d_json.UnmarshalPacket(hd, rbody)
	if err != nil {
		return fmt.Errorf("Packet type miss match %v", rbody)
	}
	recved, ok := robj.(*w2d_obj.NotiStageChat_data)
	if !ok {
		return fmt.Errorf("packet mismatch %v", robj)
	}
	return fmt.Errorf("Not implemented %v", recved)
}
