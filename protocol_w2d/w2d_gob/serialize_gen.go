// Code generated by "genprotocol -ver=8cb09769c07a0cf3e7042afbf364a4eff2c960eafe6d9a8ccbb46041a984713a -basedir=. -prefix=w2d -statstype=int"

package w2d_gob

import (
	"bytes"
	"encoding/gob"
	"fmt"

	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_idcmd"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_idnoti"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_obj"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_packet"
)

// marshal body and append to oldBufferToAppend
// and return newbuffer, body type, error
func MarshalBodyFn(body interface{}, oldBuffToAppend []byte) ([]byte, byte, error) {
	network := bytes.NewBuffer(oldBuffToAppend)
	enc := gob.NewEncoder(network)
	err := enc.Encode(body)
	return network.Bytes(), 0, err
}

func UnmarshalPacket(h w2d_packet.Header, bodyData []byte) (interface{}, error) {
	switch h.FlowType {
	case w2d_packet.Request:
		if int(h.Cmd) >= len(ReqUnmarshalMap) {
			return nil, fmt.Errorf("unknown request command: %v %v",
				h.FlowType, w2d_idcmd.CommandID(h.Cmd))
		}
		return ReqUnmarshalMap[h.Cmd](h, bodyData)

	case w2d_packet.Response:
		if int(h.Cmd) >= len(RspUnmarshalMap) {
			return nil, fmt.Errorf("unknown response command: %v %v",
				h.FlowType, w2d_idcmd.CommandID(h.Cmd))
		}
		return RspUnmarshalMap[h.Cmd](h, bodyData)

	case w2d_packet.Notification:
		if int(h.Cmd) >= len(NotiUnmarshalMap) {
			return nil, fmt.Errorf("unknown notification command: %v %v",
				h.FlowType, w2d_idcmd.CommandID(h.Cmd))
		}
		return NotiUnmarshalMap[h.Cmd](h, bodyData)
	}
	return nil, fmt.Errorf("unknown packet FlowType %v", h.FlowType)
}

var ReqUnmarshalMap = [...]func(h w2d_packet.Header, bodyData []byte) (interface{}, error){
	w2d_idcmd.Invalid:     unmarshal_ReqInvalid,
	w2d_idcmd.EnterStage:  unmarshal_ReqEnterStage,
	w2d_idcmd.ChatToStage: unmarshal_ReqChatToStage,
	w2d_idcmd.Heartbeat:   unmarshal_ReqHeartbeat,
}

var RspUnmarshalMap = [...]func(h w2d_packet.Header, bodyData []byte) (interface{}, error){
	w2d_idcmd.Invalid:     unmarshal_RspInvalid,
	w2d_idcmd.EnterStage:  unmarshal_RspEnterStage,
	w2d_idcmd.ChatToStage: unmarshal_RspChatToStage,
	w2d_idcmd.Heartbeat:   unmarshal_RspHeartbeat,
}

var NotiUnmarshalMap = [...]func(h w2d_packet.Header, bodyData []byte) (interface{}, error){
	w2d_idnoti.Invalid:        unmarshal_NotiInvalid,
	w2d_idnoti.StageInfo:      unmarshal_NotiStageInfo,
	w2d_idnoti.StatsInfo:      unmarshal_NotiStatsInfo,
	w2d_idnoti.StageBroadCast: unmarshal_NotiStageBroadCast,
}

func unmarshal_ReqInvalid(h w2d_packet.Header, bodyData []byte) (interface{}, error) {
	var args w2d_obj.ReqInvalid_data
	network := bytes.NewBuffer(bodyData)
	dec := gob.NewDecoder(network)
	err := dec.Decode(&args)
	return &args, err
}

func unmarshal_RspInvalid(h w2d_packet.Header, bodyData []byte) (interface{}, error) {
	var args w2d_obj.RspInvalid_data
	network := bytes.NewBuffer(bodyData)
	dec := gob.NewDecoder(network)
	err := dec.Decode(&args)
	return &args, err
}

func unmarshal_ReqEnterStage(h w2d_packet.Header, bodyData []byte) (interface{}, error) {
	var args w2d_obj.ReqEnterStage_data
	network := bytes.NewBuffer(bodyData)
	dec := gob.NewDecoder(network)
	err := dec.Decode(&args)
	return &args, err
}

func unmarshal_RspEnterStage(h w2d_packet.Header, bodyData []byte) (interface{}, error) {
	var args w2d_obj.RspEnterStage_data
	network := bytes.NewBuffer(bodyData)
	dec := gob.NewDecoder(network)
	err := dec.Decode(&args)
	return &args, err
}

func unmarshal_ReqChatToStage(h w2d_packet.Header, bodyData []byte) (interface{}, error) {
	var args w2d_obj.ReqChatToStage_data
	network := bytes.NewBuffer(bodyData)
	dec := gob.NewDecoder(network)
	err := dec.Decode(&args)
	return &args, err
}

func unmarshal_RspChatToStage(h w2d_packet.Header, bodyData []byte) (interface{}, error) {
	var args w2d_obj.RspChatToStage_data
	network := bytes.NewBuffer(bodyData)
	dec := gob.NewDecoder(network)
	err := dec.Decode(&args)
	return &args, err
}

func unmarshal_ReqHeartbeat(h w2d_packet.Header, bodyData []byte) (interface{}, error) {
	var args w2d_obj.ReqHeartbeat_data
	network := bytes.NewBuffer(bodyData)
	dec := gob.NewDecoder(network)
	err := dec.Decode(&args)
	return &args, err
}

func unmarshal_RspHeartbeat(h w2d_packet.Header, bodyData []byte) (interface{}, error) {
	var args w2d_obj.RspHeartbeat_data
	network := bytes.NewBuffer(bodyData)
	dec := gob.NewDecoder(network)
	err := dec.Decode(&args)
	return &args, err
}

func unmarshal_NotiInvalid(h w2d_packet.Header, bodyData []byte) (interface{}, error) {
	var args w2d_obj.NotiInvalid_data
	network := bytes.NewBuffer(bodyData)
	dec := gob.NewDecoder(network)
	err := dec.Decode(&args)
	return &args, err
}

func unmarshal_NotiStageInfo(h w2d_packet.Header, bodyData []byte) (interface{}, error) {
	var args w2d_obj.NotiStageInfo_data
	network := bytes.NewBuffer(bodyData)
	dec := gob.NewDecoder(network)
	err := dec.Decode(&args)
	return &args, err
}

func unmarshal_NotiStatsInfo(h w2d_packet.Header, bodyData []byte) (interface{}, error) {
	var args w2d_obj.NotiStatsInfo_data
	network := bytes.NewBuffer(bodyData)
	dec := gob.NewDecoder(network)
	err := dec.Decode(&args)
	return &args, err
}

func unmarshal_NotiStageBroadCast(h w2d_packet.Header, bodyData []byte) (interface{}, error) {
	var args w2d_obj.NotiStageBroadCast_data
	network := bytes.NewBuffer(bodyData)
	dec := gob.NewDecoder(network)
	err := dec.Decode(&args)
	return &args, err
}
