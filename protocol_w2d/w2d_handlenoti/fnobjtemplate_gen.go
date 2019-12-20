// Code generated by "genprotocol -ver=8cb09769c07a0cf3e7042afbf364a4eff2c960eafe6d9a8ccbb46041a984713a -basedir=. -prefix=w2d -statstype=int"

package w2d_handlenoti

/* obj base demux fn map template

var DemuxNoti2ObjFnMap = [...]func(me interface{}, hd w2d_packet.Header, body interface{}) error {
w2d_idnoti.Invalid : objRecvNotiFn_Invalid,
w2d_idnoti.StageInfo : objRecvNotiFn_StageInfo,
w2d_idnoti.StatsInfo : objRecvNotiFn_StatsInfo,
w2d_idnoti.StageBroadCast : objRecvNotiFn_StageBroadCast,

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

	func objRecvNotiFn_StageBroadCast(me interface{}, hd w2d_packet.Header, body interface{}) error {
		robj , ok := body.(*w2d_obj.NotiStageBroadCast_data)
		if !ok {
			return fmt.Errorf("packet mismatch %v", body )
		}
		return fmt.Errorf("Not implemented %v", robj)
	}

*/
