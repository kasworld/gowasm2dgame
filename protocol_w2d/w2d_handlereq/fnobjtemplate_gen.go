// Code generated by "genprotocol -ver=ed65a653bd268dc21902d5d07939f7bfc1ba6b98026a426c30526d9f59ba8d12 -basedir=. -prefix=w2d -statstype=int"

package w2d_handlereq

/* obj base demux fn map template
	var DemuxReq2ObjAPIFnMap = [...]func(
		me interface{}, hd w2d_packet.Header, robj interface{}) (
		w2d_packet.Header, interface{}, error){
	w2d_idcmd.Invalid: Req2ObjAPI_Invalid,
w2d_idcmd.EnterStage: Req2ObjAPI_EnterStage,
w2d_idcmd.ChatToStage: Req2ObjAPI_ChatToStage,
w2d_idcmd.Heartbeat: Req2ObjAPI_Heartbeat,

}   // DemuxReq2ObjAPIFnMap

	func Req2ObjAPI_Invalid(
		me interface{}, hd w2d_packet.Header, robj interface{}) (
		w2d_packet.Header, interface{},  error) {
		req, ok := robj.(*w2d_obj.ReqInvalid_data)
		if !ok {
			return hd, nil, fmt.Errorf("Packet type miss match %v", robj)
		}
		rhd, rsp, err := objAPIFn_ReqInvalid(me, hd, req)
		return rhd, rsp, err
	}
	func objAPIFn_ReqInvalid(
		me interface{}, hd w2d_packet.Header, robj *w2d_obj.ReqInvalid_data) (
		w2d_packet.Header, *w2d_obj.RspInvalid_data, error) {
		sendHeader := w2d_packet.Header{
			ErrorCode : w2d_error.None,
		}
		sendBody := &w2d_obj.RspInvalid_data{
		}
		return sendHeader, sendBody, nil
	}

	func Req2ObjAPI_EnterStage(
		me interface{}, hd w2d_packet.Header, robj interface{}) (
		w2d_packet.Header, interface{},  error) {
		req, ok := robj.(*w2d_obj.ReqEnterStage_data)
		if !ok {
			return hd, nil, fmt.Errorf("Packet type miss match %v", robj)
		}
		rhd, rsp, err := objAPIFn_ReqEnterStage(me, hd, req)
		return rhd, rsp, err
	}
	func objAPIFn_ReqEnterStage(
		me interface{}, hd w2d_packet.Header, robj *w2d_obj.ReqEnterStage_data) (
		w2d_packet.Header, *w2d_obj.RspEnterStage_data, error) {
		sendHeader := w2d_packet.Header{
			ErrorCode : w2d_error.None,
		}
		sendBody := &w2d_obj.RspEnterStage_data{
		}
		return sendHeader, sendBody, nil
	}

	func Req2ObjAPI_ChatToStage(
		me interface{}, hd w2d_packet.Header, robj interface{}) (
		w2d_packet.Header, interface{},  error) {
		req, ok := robj.(*w2d_obj.ReqChatToStage_data)
		if !ok {
			return hd, nil, fmt.Errorf("Packet type miss match %v", robj)
		}
		rhd, rsp, err := objAPIFn_ReqChatToStage(me, hd, req)
		return rhd, rsp, err
	}
	func objAPIFn_ReqChatToStage(
		me interface{}, hd w2d_packet.Header, robj *w2d_obj.ReqChatToStage_data) (
		w2d_packet.Header, *w2d_obj.RspChatToStage_data, error) {
		sendHeader := w2d_packet.Header{
			ErrorCode : w2d_error.None,
		}
		sendBody := &w2d_obj.RspChatToStage_data{
		}
		return sendHeader, sendBody, nil
	}

	func Req2ObjAPI_Heartbeat(
		me interface{}, hd w2d_packet.Header, robj interface{}) (
		w2d_packet.Header, interface{},  error) {
		req, ok := robj.(*w2d_obj.ReqHeartbeat_data)
		if !ok {
			return hd, nil, fmt.Errorf("Packet type miss match %v", robj)
		}
		rhd, rsp, err := objAPIFn_ReqHeartbeat(me, hd, req)
		return rhd, rsp, err
	}
	func objAPIFn_ReqHeartbeat(
		me interface{}, hd w2d_packet.Header, robj *w2d_obj.ReqHeartbeat_data) (
		w2d_packet.Header, *w2d_obj.RspHeartbeat_data, error) {
		sendHeader := w2d_packet.Header{
			ErrorCode : w2d_error.None,
		}
		sendBody := &w2d_obj.RspHeartbeat_data{
		}
		return sendHeader, sendBody, nil
	}

*/
