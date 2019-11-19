// Copyright 2015,2016,2017,2018,2019 SeukWon Kang (kasworld@gmail.com)
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import (
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_error"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_idcmd"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_obj"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_packet"
)

///////////////////////////////////////////////////////////////

var DemuxReq2BytesAPIFnMap = [...]func(
	me interface{}, hd w2d_packet.Header, rbody []byte) (
	w2d_packet.Header, interface{}, error){
	w2d_idcmd.Invalid:   bytesAPIFn_ReqInvalid,
	w2d_idcmd.MakeTeam:  bytesAPIFn_ReqMakeTeam,
	w2d_idcmd.Act:       bytesAPIFn_ReqAct,
	w2d_idcmd.State:     bytesAPIFn_ReqState,
	w2d_idcmd.Heartbeat: bytesAPIFn_ReqHeartbeat,
} // DemuxReq2BytesAPIFnMap

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
		ErrorCode: w2d_error.None,
	}
	sendBody := &w2d_obj.RspInvalid_data{}
	return sendHeader, sendBody, nil
}

func bytesAPIFn_ReqMakeTeam(
	me interface{}, hd w2d_packet.Header, rbody []byte) (
	w2d_packet.Header, interface{}, error) {
	// robj, err := w2d_json.UnmarshalPacket(hd, rbody)
	// if err != nil {
	// 	return hd, nil, fmt.Errorf("Packet type miss match %v", rbody)
	// }
	// recvBody, ok := robj.(*w2d_obj.ReqMakeTeam_data)
	// if !ok {
	// 	return hd, nil, fmt.Errorf("Packet type miss match %v", robj)
	// }
	// _ = recvBody

	sendHeader := w2d_packet.Header{
		ErrorCode: w2d_error.None,
	}
	sendBody := &w2d_obj.RspMakeTeam_data{}
	return sendHeader, sendBody, nil
}

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
		ErrorCode: w2d_error.None,
	}
	sendBody := &w2d_obj.RspAct_data{}
	return sendHeader, sendBody, nil
}

func bytesAPIFn_ReqState(
	me interface{}, hd w2d_packet.Header, rbody []byte) (
	w2d_packet.Header, interface{}, error) {
	// robj, err := w2d_json.UnmarshalPacket(hd, rbody)
	// if err != nil {
	// 	return hd, nil, fmt.Errorf("Packet type miss match %v", rbody)
	// }
	// recvBody, ok := robj.(*w2d_obj.ReqState_data)
	// if !ok {
	// 	return hd, nil, fmt.Errorf("Packet type miss match %v", robj)
	// }
	// _ = recvBody

	sendHeader := w2d_packet.Header{
		ErrorCode: w2d_error.None,
	}
	sendBody := &w2d_obj.RspState_data{}
	return sendHeader, sendBody, nil
}

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
		ErrorCode: w2d_error.None,
	}
	sendBody := &w2d_obj.RspHeartbeat_data{}
	return sendHeader, sendBody, nil
}
