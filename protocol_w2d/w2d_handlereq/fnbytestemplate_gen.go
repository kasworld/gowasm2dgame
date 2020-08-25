// Code generated by "genprotocol -ver=8ce3a4010a59de778695c59389c0fd9a3938197ee346a449f005b536d94d0e60 -basedir=protocol_w2d -prefix=w2d -statstype=int"

package w2d_handlereq

/* bytes base fn map api template , unmarshal in api
	var DemuxReq2BytesAPIFnMap = [...]func(
		me interface{}, hd w2d_packet.Header, rbody []byte) (
		w2d_packet.Header, interface{}, error){
	w2d_idcmd.Invalid: bytesAPIFn_ReqInvalid,// Invalid not used, make empty packet error
w2d_idcmd.Login: bytesAPIFn_ReqLogin,// Login make session with nickname and enter stage
w2d_idcmd.Heartbeat: bytesAPIFn_ReqHeartbeat,// Heartbeat prevent connection timeout
w2d_idcmd.Chat: bytesAPIFn_ReqChat,// Chat chat to stage
w2d_idcmd.Act: bytesAPIFn_ReqAct,// Act send user action

}   // DemuxReq2BytesAPIFnMap

	// Invalid not used, make empty packet error
	func bytesAPIFn_ReqInvalid(
		me interface{}, hd w2d_packet.Header, rbody []byte) (
		w2d_packet.Header, interface{}, error) {
		// robj, err := w2d_json.UnmarshalPacket(hd, rbody)
		// if err != nil {
		// 	return hd, nil, fmt.Errorf("Packet type miss match %v", rbody)
		// }
		// recvBody, ok := robj.(*w2d_obj.ReqInvalid_data)
		// if !ok {
		// 	return hd, nil, fmt.Errorf("Packet type miss match %v", robj)
		// }
		// _ = recvBody

		sendHeader := w2d_packet.Header{
			ErrorCode : w2d_error.None,
		}
		sendBody := &w2d_obj.RspInvalid_data{
		}
		return sendHeader, sendBody, nil
	}

	// Login make session with nickname and enter stage
	func bytesAPIFn_ReqLogin(
		me interface{}, hd w2d_packet.Header, rbody []byte) (
		w2d_packet.Header, interface{}, error) {
		// robj, err := w2d_json.UnmarshalPacket(hd, rbody)
		// if err != nil {
		// 	return hd, nil, fmt.Errorf("Packet type miss match %v", rbody)
		// }
		// recvBody, ok := robj.(*w2d_obj.ReqLogin_data)
		// if !ok {
		// 	return hd, nil, fmt.Errorf("Packet type miss match %v", robj)
		// }
		// _ = recvBody

		sendHeader := w2d_packet.Header{
			ErrorCode : w2d_error.None,
		}
		sendBody := &w2d_obj.RspLogin_data{
		}
		return sendHeader, sendBody, nil
	}

	// Heartbeat prevent connection timeout
	func bytesAPIFn_ReqHeartbeat(
		me interface{}, hd w2d_packet.Header, rbody []byte) (
		w2d_packet.Header, interface{}, error) {
		// robj, err := w2d_json.UnmarshalPacket(hd, rbody)
		// if err != nil {
		// 	return hd, nil, fmt.Errorf("Packet type miss match %v", rbody)
		// }
		// recvBody, ok := robj.(*w2d_obj.ReqHeartbeat_data)
		// if !ok {
		// 	return hd, nil, fmt.Errorf("Packet type miss match %v", robj)
		// }
		// _ = recvBody

		sendHeader := w2d_packet.Header{
			ErrorCode : w2d_error.None,
		}
		sendBody := &w2d_obj.RspHeartbeat_data{
		}
		return sendHeader, sendBody, nil
	}

	// Chat chat to stage
	func bytesAPIFn_ReqChat(
		me interface{}, hd w2d_packet.Header, rbody []byte) (
		w2d_packet.Header, interface{}, error) {
		// robj, err := w2d_json.UnmarshalPacket(hd, rbody)
		// if err != nil {
		// 	return hd, nil, fmt.Errorf("Packet type miss match %v", rbody)
		// }
		// recvBody, ok := robj.(*w2d_obj.ReqChat_data)
		// if !ok {
		// 	return hd, nil, fmt.Errorf("Packet type miss match %v", robj)
		// }
		// _ = recvBody

		sendHeader := w2d_packet.Header{
			ErrorCode : w2d_error.None,
		}
		sendBody := &w2d_obj.RspChat_data{
		}
		return sendHeader, sendBody, nil
	}

	// Act send user action
	func bytesAPIFn_ReqAct(
		me interface{}, hd w2d_packet.Header, rbody []byte) (
		w2d_packet.Header, interface{}, error) {
		// robj, err := w2d_json.UnmarshalPacket(hd, rbody)
		// if err != nil {
		// 	return hd, nil, fmt.Errorf("Packet type miss match %v", rbody)
		// }
		// recvBody, ok := robj.(*w2d_obj.ReqAct_data)
		// if !ok {
		// 	return hd, nil, fmt.Errorf("Packet type miss match %v", robj)
		// }
		// _ = recvBody

		sendHeader := w2d_packet.Header{
			ErrorCode : w2d_error.None,
		}
		sendBody := &w2d_obj.RspAct_data{
		}
		return sendHeader, sendBody, nil
	}

*/
