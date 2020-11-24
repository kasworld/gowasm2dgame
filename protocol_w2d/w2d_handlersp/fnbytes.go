// Code generated by "genprotocol.exe -ver=56a40c9afaeea01bb7ff0ceed1dabb8a62deedd1dfa2e5c804d9e37c44d134ca -basedir=protocol_w2d -prefix=w2d -statstype=int"

package w2d_handlersp

import (
	"fmt"

	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_idcmd"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_json"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_obj"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_packet"
)

// bytes base demux fn map

var DemuxRsp2BytesFnMap = [...]func(me interface{}, hd w2d_packet.Header, rbody []byte) error{
	w2d_idcmd.Invalid:   bytesRecvRspFn_Invalid,   // Invalid not used, make empty packet error
	w2d_idcmd.Login:     bytesRecvRspFn_Login,     // Login make session with nickname and enter stage
	w2d_idcmd.Heartbeat: bytesRecvRspFn_Heartbeat, // Heartbeat prevent connection timeout
	w2d_idcmd.Chat:      bytesRecvRspFn_Chat,      // Chat chat to stage
	w2d_idcmd.Act:       bytesRecvRspFn_Act,       // Act send user action

}

// Invalid not used, make empty packet error
func bytesRecvRspFn_Invalid(me interface{}, hd w2d_packet.Header, rbody []byte) error {
	robj, err := w2d_json.UnmarshalPacket(hd, rbody)
	if err != nil {
		return fmt.Errorf("Packet type miss match %v", rbody)
	}
	recved, ok := robj.(*w2d_obj.RspInvalid_data)
	if !ok {
		return fmt.Errorf("packet mismatch %v", robj)
	}
	return fmt.Errorf("Not implemented %v", recved)
}

// Login make session with nickname and enter stage
func bytesRecvRspFn_Login(me interface{}, hd w2d_packet.Header, rbody []byte) error {
	robj, err := w2d_json.UnmarshalPacket(hd, rbody)
	if err != nil {
		return fmt.Errorf("Packet type miss match %v", rbody)
	}
	recved, ok := robj.(*w2d_obj.RspLogin_data)
	if !ok {
		return fmt.Errorf("packet mismatch %v", robj)
	}
	return fmt.Errorf("Not implemented %v", recved)
}

// Heartbeat prevent connection timeout
func bytesRecvRspFn_Heartbeat(me interface{}, hd w2d_packet.Header, rbody []byte) error {
	robj, err := w2d_json.UnmarshalPacket(hd, rbody)
	if err != nil {
		return fmt.Errorf("Packet type miss match %v", rbody)
	}
	recved, ok := robj.(*w2d_obj.RspHeartbeat_data)
	if !ok {
		return fmt.Errorf("packet mismatch %v", robj)
	}
	return fmt.Errorf("Not implemented %v", recved)
}

// Chat chat to stage
func bytesRecvRspFn_Chat(me interface{}, hd w2d_packet.Header, rbody []byte) error {
	robj, err := w2d_json.UnmarshalPacket(hd, rbody)
	if err != nil {
		return fmt.Errorf("Packet type miss match %v", rbody)
	}
	recved, ok := robj.(*w2d_obj.RspChat_data)
	if !ok {
		return fmt.Errorf("packet mismatch %v", robj)
	}
	return fmt.Errorf("Not implemented %v", recved)
}

// Act send user action
func bytesRecvRspFn_Act(me interface{}, hd w2d_packet.Header, rbody []byte) error {
	robj, err := w2d_json.UnmarshalPacket(hd, rbody)
	if err != nil {
		return fmt.Errorf("Packet type miss match %v", rbody)
	}
	recved, ok := robj.(*w2d_obj.RspAct_data)
	if !ok {
		return fmt.Errorf("packet mismatch %v", robj)
	}
	return fmt.Errorf("Not implemented %v", recved)
}