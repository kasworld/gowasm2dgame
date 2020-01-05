// Copyright 2015,2016,2017,2018,2019,2020 SeukWon Kang (kasworld@gmail.com)
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
	"fmt"

	"github.com/kasworld/gowasm2dgame/game/conndata"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_error"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_gob"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_idnoti"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_obj"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_packet"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_serveconnbyte"
)

///////////////////////////////////////////////////////////////

func (svr *Server) bytesAPIFn_ReqInvalid(
	me interface{}, hd w2d_packet.Header, rbody []byte) (
	w2d_packet.Header, interface{}, error) {
	sendHeader := w2d_packet.Header{}
	return sendHeader, nil, fmt.Errorf("invalid packet")
}

func (svr *Server) bytesAPIFn_ReqEnterStage(
	me interface{}, hd w2d_packet.Header, rbody []byte) (
	w2d_packet.Header, interface{}, error) {
	robj, err := w2d_gob.UnmarshalPacket(hd, rbody)
	if err != nil {
		return hd, nil, fmt.Errorf("Packet type miss match %v", rbody)
	}
	recvBody, ok := robj.(*w2d_obj.ReqEnterStage_data)
	if !ok {
		return hd, nil, fmt.Errorf("Packet type miss match %v", robj)
	}
	_ = recvBody

	sendHeader := w2d_packet.Header{
		ErrorCode: w2d_error.None,
	}
	sendBody := &w2d_obj.RspEnterStage_data{}
	return sendHeader, sendBody, nil
}

func (svr *Server) bytesAPIFn_ReqChatToStage(
	me interface{}, hd w2d_packet.Header, rbody []byte) (
	w2d_packet.Header, interface{}, error) {
	robj, err := w2d_gob.UnmarshalPacket(hd, rbody)
	if err != nil {
		return hd, nil, fmt.Errorf("Packet type miss match %v", rbody)
	}
	recvBody, ok := robj.(*w2d_obj.ReqChatToStage_data)
	if !ok {
		return hd, nil, fmt.Errorf("Packet type miss match %v", robj)
	}
	_ = recvBody

	conn, ok := me.(*w2d_serveconnbyte.ServeConnByte)
	if !ok {
		return hd, nil, fmt.Errorf("Packet type miss match %v", me)
	}
	connData, ok := conn.GetConnData().(*conndata.ConnData)
	if !ok {
		return hd, nil, fmt.Errorf("Packet type miss match %v", conn.GetConnData())
	}
	stg := svr.stageManager.GetByUUID(connData.StageID)
	connList := stg.Conns.GetList()
	noti := &w2d_obj.NotiStageChat_data{
		SenderNick: connData.UUID,
		Chat:       recvBody.Chat,
	}
	for _, v := range connList {
		v.SendNotiPacket(w2d_idnoti.StageChat,
			noti,
		)
	}

	sendHeader := w2d_packet.Header{
		ErrorCode: w2d_error.None,
	}
	sendBody := &w2d_obj.RspChatToStage_data{}
	return sendHeader, sendBody, nil
}

func (svr *Server) bytesAPIFn_ReqHeartbeat(
	me interface{}, hd w2d_packet.Header, rbody []byte) (
	w2d_packet.Header, interface{}, error) {
	robj, err := w2d_gob.UnmarshalPacket(hd, rbody)
	if err != nil {
		return hd, nil, fmt.Errorf("Packet type miss match %v", rbody)
	}
	recvBody, ok := robj.(*w2d_obj.ReqHeartbeat_data)
	if !ok {
		return hd, nil, fmt.Errorf("Packet type miss match %v", robj)
	}
	_ = recvBody

	sendHeader := w2d_packet.Header{
		ErrorCode: w2d_error.None,
	}
	sendBody := &w2d_obj.RspHeartbeat_data{
		Tick: recvBody.Tick,
	}
	return sendHeader, sendBody, nil
}
