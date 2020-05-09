// Code generated by "genprotocol -ver=eb884c961074aeaf0b613a0d0567567c029f9b9d5b9a686f9b1b7ade5f686087 -basedir=. -prefix=w2d -statstype=int"

package w2d_msgp

import (
	"fmt"

	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_idcmd"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_idnoti"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_obj"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_packet"
	"github.com/tinylib/msgp/msgp"
)

// MarshalBodyFn marshal body and append to oldBufferToAppend
// and return newbuffer, body type, error
func MarshalBodyFn(body interface{}, oldBuffToAppend []byte) ([]byte, byte, error) {
	newBuffer, err := body.(msgp.Marshaler).MarshalMsg(oldBuffToAppend)
	return newBuffer, 0, err
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
	w2d_idcmd.Invalid:    unmarshal_ReqInvalid,
	w2d_idcmd.Login:      unmarshal_ReqLogin,
	w2d_idcmd.Heartbeat:  unmarshal_ReqHeartbeat,
	w2d_idcmd.Chat:       unmarshal_ReqChat,
	w2d_idcmd.EnterStage: unmarshal_ReqEnterStage,
}

var RspUnmarshalMap = [...]func(h w2d_packet.Header, bodyData []byte) (interface{}, error){
	w2d_idcmd.Invalid:    unmarshal_RspInvalid,
	w2d_idcmd.Login:      unmarshal_RspLogin,
	w2d_idcmd.Heartbeat:  unmarshal_RspHeartbeat,
	w2d_idcmd.Chat:       unmarshal_RspChat,
	w2d_idcmd.EnterStage: unmarshal_RspEnterStage,
}

var NotiUnmarshalMap = [...]func(h w2d_packet.Header, bodyData []byte) (interface{}, error){
	w2d_idnoti.Invalid:   unmarshal_NotiInvalid,
	w2d_idnoti.StageInfo: unmarshal_NotiStageInfo,
	w2d_idnoti.StatsInfo: unmarshal_NotiStatsInfo,
	w2d_idnoti.StageChat: unmarshal_NotiStageChat,
}

func unmarshal_ReqInvalid(h w2d_packet.Header, bodyData []byte) (interface{}, error) {
	var args w2d_obj.ReqInvalid_data
	if _, err := args.UnmarshalMsg(bodyData); err != nil {
		return nil, err
	}
	return &args, nil
}

func unmarshal_RspInvalid(h w2d_packet.Header, bodyData []byte) (interface{}, error) {
	var args w2d_obj.RspInvalid_data
	if _, err := args.UnmarshalMsg(bodyData); err != nil {
		return nil, err
	}
	return &args, nil
}

func unmarshal_ReqLogin(h w2d_packet.Header, bodyData []byte) (interface{}, error) {
	var args w2d_obj.ReqLogin_data
	if _, err := args.UnmarshalMsg(bodyData); err != nil {
		return nil, err
	}
	return &args, nil
}

func unmarshal_RspLogin(h w2d_packet.Header, bodyData []byte) (interface{}, error) {
	var args w2d_obj.RspLogin_data
	if _, err := args.UnmarshalMsg(bodyData); err != nil {
		return nil, err
	}
	return &args, nil
}

func unmarshal_ReqHeartbeat(h w2d_packet.Header, bodyData []byte) (interface{}, error) {
	var args w2d_obj.ReqHeartbeat_data
	if _, err := args.UnmarshalMsg(bodyData); err != nil {
		return nil, err
	}
	return &args, nil
}

func unmarshal_RspHeartbeat(h w2d_packet.Header, bodyData []byte) (interface{}, error) {
	var args w2d_obj.RspHeartbeat_data
	if _, err := args.UnmarshalMsg(bodyData); err != nil {
		return nil, err
	}
	return &args, nil
}

func unmarshal_ReqChat(h w2d_packet.Header, bodyData []byte) (interface{}, error) {
	var args w2d_obj.ReqChat_data
	if _, err := args.UnmarshalMsg(bodyData); err != nil {
		return nil, err
	}
	return &args, nil
}

func unmarshal_RspChat(h w2d_packet.Header, bodyData []byte) (interface{}, error) {
	var args w2d_obj.RspChat_data
	if _, err := args.UnmarshalMsg(bodyData); err != nil {
		return nil, err
	}
	return &args, nil
}

func unmarshal_ReqEnterStage(h w2d_packet.Header, bodyData []byte) (interface{}, error) {
	var args w2d_obj.ReqEnterStage_data
	if _, err := args.UnmarshalMsg(bodyData); err != nil {
		return nil, err
	}
	return &args, nil
}

func unmarshal_RspEnterStage(h w2d_packet.Header, bodyData []byte) (interface{}, error) {
	var args w2d_obj.RspEnterStage_data
	if _, err := args.UnmarshalMsg(bodyData); err != nil {
		return nil, err
	}
	return &args, nil
}

func unmarshal_NotiInvalid(h w2d_packet.Header, bodyData []byte) (interface{}, error) {
	var args w2d_obj.NotiInvalid_data
	if _, err := args.UnmarshalMsg(bodyData); err != nil {
		return nil, err
	}
	return &args, nil
}

func unmarshal_NotiStageInfo(h w2d_packet.Header, bodyData []byte) (interface{}, error) {
	var args w2d_obj.NotiStageInfo_data
	if _, err := args.UnmarshalMsg(bodyData); err != nil {
		return nil, err
	}
	return &args, nil
}

func unmarshal_NotiStatsInfo(h w2d_packet.Header, bodyData []byte) (interface{}, error) {
	var args w2d_obj.NotiStatsInfo_data
	if _, err := args.UnmarshalMsg(bodyData); err != nil {
		return nil, err
	}
	return &args, nil
}

func unmarshal_NotiStageChat(h w2d_packet.Header, bodyData []byte) (interface{}, error) {
	var args w2d_obj.NotiStageChat_data
	if _, err := args.UnmarshalMsg(bodyData); err != nil {
		return nil, err
	}
	return &args, nil
}
