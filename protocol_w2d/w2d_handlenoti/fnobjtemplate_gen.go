// Code generated by "genprotocol -ver=eb884c961074aeaf0b613a0d0567567c029f9b9d5b9a686f9b1b7ade5f686087 -basedir=. -prefix=w2d -statstype=int"

package w2d_handlenoti

/* obj base demux fn map template

var DemuxNoti2ObjFnMap = [...]func(me interface{}, hd w2d_packet.Header, body interface{}) error {
w2d_idnoti.Invalid : objRecvNotiFn_Invalid,
w2d_idnoti.StageInfo : objRecvNotiFn_StageInfo,
w2d_idnoti.StatsInfo : objRecvNotiFn_StatsInfo,
w2d_idnoti.StageChat : objRecvNotiFn_StageChat,

}

	func objRecvNotiFn_Invalid(me interface{}, hd w2d_packet.Header, body interface{}) error {
		robj , ok := body.(*w2d_obj.NotiInvalid_data)
		if !ok {
			return fmt.Errorf("packet mismatch %v", body )
		}
		return fmt.Errorf("Not implemented %v", robj)
	}

	func objRecvNotiFn_StageInfo(me interface{}, hd w2d_packet.Header, body interface{}) error {
		robj , ok := body.(*w2d_obj.NotiStageInfo_data)
		if !ok {
			return fmt.Errorf("packet mismatch %v", body )
		}
		return fmt.Errorf("Not implemented %v", robj)
	}

	func objRecvNotiFn_StatsInfo(me interface{}, hd w2d_packet.Header, body interface{}) error {
		robj , ok := body.(*w2d_obj.NotiStatsInfo_data)
		if !ok {
			return fmt.Errorf("packet mismatch %v", body )
		}
		return fmt.Errorf("Not implemented %v", robj)
	}

	func objRecvNotiFn_StageChat(me interface{}, hd w2d_packet.Header, body interface{}) error {
		robj , ok := body.(*w2d_obj.NotiStageChat_data)
		if !ok {
			return fmt.Errorf("packet mismatch %v", body )
		}
		return fmt.Errorf("Not implemented %v", robj)
	}

*/
