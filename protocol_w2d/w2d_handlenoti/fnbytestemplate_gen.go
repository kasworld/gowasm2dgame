// Code generated by "genprotocol -ver=ed65a653bd268dc21902d5d07939f7bfc1ba6b98026a426c30526d9f59ba8d12 -basedir=. -prefix=w2d -statstype=int"

package w2d_handlenoti

/* bytes base demux fn map template

var DemuxNoti2ByteFnMap = [...]func(me interface{}, hd w2d_packet.Header, rbody []byte) error {
w2d_idnoti.Invalid : bytesRecvNotiFn_Invalid,
w2d_idnoti.StageInfo : bytesRecvNotiFn_StageInfo,
w2d_idnoti.StatsInfo : bytesRecvNotiFn_StatsInfo,
w2d_idnoti.StageChat : bytesRecvNotiFn_StageChat,

}

	func bytesRecvNotiFn_Invalid(me interface{}, hd w2d_packet.Header, rbody []byte) error {
		robj, err := w2d_json.UnmarshalPacket(hd, rbody)
		if err != nil {
			return fmt.Errorf("Packet type miss match %v", rbody)
		}
		recved , ok := robj.(*w2d_obj.NotiInvalid_data)
		if !ok {
			return fmt.Errorf("packet mismatch %v", robj )
		}
		return fmt.Errorf("Not implemented %v", recved)
	}

	func bytesRecvNotiFn_StageInfo(me interface{}, hd w2d_packet.Header, rbody []byte) error {
		robj, err := w2d_json.UnmarshalPacket(hd, rbody)
		if err != nil {
			return fmt.Errorf("Packet type miss match %v", rbody)
		}
		recved , ok := robj.(*w2d_obj.NotiStageInfo_data)
		if !ok {
			return fmt.Errorf("packet mismatch %v", robj )
		}
		return fmt.Errorf("Not implemented %v", recved)
	}

	func bytesRecvNotiFn_StatsInfo(me interface{}, hd w2d_packet.Header, rbody []byte) error {
		robj, err := w2d_json.UnmarshalPacket(hd, rbody)
		if err != nil {
			return fmt.Errorf("Packet type miss match %v", rbody)
		}
		recved , ok := robj.(*w2d_obj.NotiStatsInfo_data)
		if !ok {
			return fmt.Errorf("packet mismatch %v", robj )
		}
		return fmt.Errorf("Not implemented %v", recved)
	}

	func bytesRecvNotiFn_StageChat(me interface{}, hd w2d_packet.Header, rbody []byte) error {
		robj, err := w2d_json.UnmarshalPacket(hd, rbody)
		if err != nil {
			return fmt.Errorf("Packet type miss match %v", rbody)
		}
		recved , ok := robj.(*w2d_obj.NotiStageChat_data)
		if !ok {
			return fmt.Errorf("packet mismatch %v", robj )
		}
		return fmt.Errorf("Not implemented %v", recved)
	}

*/
