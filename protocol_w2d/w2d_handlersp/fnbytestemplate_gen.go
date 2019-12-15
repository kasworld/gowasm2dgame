// Code generated by "genprotocol -ver=311c9c290570c203090ea3d20ebbf006c810eb958a7a96aef942015fbfd89d2f -basedir=. -prefix=w2d -statstype=int"

package w2d_handlersp

/* bytes base demux fn map template

var DemuxRsp2BytesFnMap = [...]func(me interface{}, hd w2d_packet.Header, rbody []byte) error {
w2d_idcmd.Invalid : bytesRecvRspFn_Invalid,
w2d_idcmd.ListStage : bytesRecvRspFn_ListStage,
w2d_idcmd.EnterStage : bytesRecvRspFn_EnterStage,
w2d_idcmd.LeaveStage : bytesRecvRspFn_LeaveStage,
w2d_idcmd.ChatToStage : bytesRecvRspFn_ChatToStage,
w2d_idcmd.MakeTeam : bytesRecvRspFn_MakeTeam,
w2d_idcmd.Act : bytesRecvRspFn_Act,
w2d_idcmd.Heartbeat : bytesRecvRspFn_Heartbeat,

}

	func bytesRecvRspFn_Invalid(me interface{}, hd w2d_packet.Header, rbody []byte) error {
		robj, err := w2d_json.UnmarshalPacket(hd, rbody)
		if err != nil {
			return  fmt.Errorf("Packet type miss match %v", rbody)
		}
		recved , ok := robj.(*w2d_obj.RspInvalid_data)
		if !ok {
			return fmt.Errorf("packet mismatch %v", robj )
		}
		return fmt.Errorf("Not implemented %v", recved)
	}

	func bytesRecvRspFn_ListStage(me interface{}, hd w2d_packet.Header, rbody []byte) error {
		robj, err := w2d_json.UnmarshalPacket(hd, rbody)
		if err != nil {
			return  fmt.Errorf("Packet type miss match %v", rbody)
		}
		recved , ok := robj.(*w2d_obj.RspListStage_data)
		if !ok {
			return fmt.Errorf("packet mismatch %v", robj )
		}
		return fmt.Errorf("Not implemented %v", recved)
	}

	func bytesRecvRspFn_EnterStage(me interface{}, hd w2d_packet.Header, rbody []byte) error {
		robj, err := w2d_json.UnmarshalPacket(hd, rbody)
		if err != nil {
			return  fmt.Errorf("Packet type miss match %v", rbody)
		}
		recved , ok := robj.(*w2d_obj.RspEnterStage_data)
		if !ok {
			return fmt.Errorf("packet mismatch %v", robj )
		}
		return fmt.Errorf("Not implemented %v", recved)
	}

	func bytesRecvRspFn_LeaveStage(me interface{}, hd w2d_packet.Header, rbody []byte) error {
		robj, err := w2d_json.UnmarshalPacket(hd, rbody)
		if err != nil {
			return  fmt.Errorf("Packet type miss match %v", rbody)
		}
		recved , ok := robj.(*w2d_obj.RspLeaveStage_data)
		if !ok {
			return fmt.Errorf("packet mismatch %v", robj )
		}
		return fmt.Errorf("Not implemented %v", recved)
	}

	func bytesRecvRspFn_ChatToStage(me interface{}, hd w2d_packet.Header, rbody []byte) error {
		robj, err := w2d_json.UnmarshalPacket(hd, rbody)
		if err != nil {
			return  fmt.Errorf("Packet type miss match %v", rbody)
		}
		recved , ok := robj.(*w2d_obj.RspChatToStage_data)
		if !ok {
			return fmt.Errorf("packet mismatch %v", robj )
		}
		return fmt.Errorf("Not implemented %v", recved)
	}

	func bytesRecvRspFn_MakeTeam(me interface{}, hd w2d_packet.Header, rbody []byte) error {
		robj, err := w2d_json.UnmarshalPacket(hd, rbody)
		if err != nil {
			return  fmt.Errorf("Packet type miss match %v", rbody)
		}
		recved , ok := robj.(*w2d_obj.RspMakeTeam_data)
		if !ok {
			return fmt.Errorf("packet mismatch %v", robj )
		}
		return fmt.Errorf("Not implemented %v", recved)
	}

	func bytesRecvRspFn_Act(me interface{}, hd w2d_packet.Header, rbody []byte) error {
		robj, err := w2d_json.UnmarshalPacket(hd, rbody)
		if err != nil {
			return  fmt.Errorf("Packet type miss match %v", rbody)
		}
		recved , ok := robj.(*w2d_obj.RspAct_data)
		if !ok {
			return fmt.Errorf("packet mismatch %v", robj )
		}
		return fmt.Errorf("Not implemented %v", recved)
	}

	func bytesRecvRspFn_Heartbeat(me interface{}, hd w2d_packet.Header, rbody []byte) error {
		robj, err := w2d_json.UnmarshalPacket(hd, rbody)
		if err != nil {
			return  fmt.Errorf("Packet type miss match %v", rbody)
		}
		recved , ok := robj.(*w2d_obj.RspHeartbeat_data)
		if !ok {
			return fmt.Errorf("packet mismatch %v", robj )
		}
		return fmt.Errorf("Not implemented %v", recved)
	}

*/
