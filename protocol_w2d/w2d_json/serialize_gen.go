// Code generated by "genprotocol -ver=311c9c290570c203090ea3d20ebbf006c810eb958a7a96aef942015fbfd89d2f -basedir=. -prefix=w2d -statstype=int"

package w2d_json

import (
	"encoding/json"
	"fmt"

	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_idcmd"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_idnoti"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_obj"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_packet"
)

// marshal body and append to oldBufferToAppend
// and return newbuffer, body type, error
func MarshalBodyFn(body interface{}, oldBuffToAppend []byte) ([]byte, byte, error) {
	var newBuffer []byte
	mdata, err := json.Marshal(body)
	if err == nil {
		newBuffer = append(oldBuffToAppend, mdata...)
	}
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
	w2d_idcmd.Invalid:     unmarshal_ReqInvalid,
	w2d_idcmd.ListStage:   unmarshal_ReqListStage,
	w2d_idcmd.EnterStage:  unmarshal_ReqEnterStage,
	w2d_idcmd.LeaveStage:  unmarshal_ReqLeaveStage,
	w2d_idcmd.ChatToStage: unmarshal_ReqChatToStage,
	w2d_idcmd.MakeTeam:    unmarshal_ReqMakeTeam,
	w2d_idcmd.Act:         unmarshal_ReqAct,
	w2d_idcmd.Heartbeat:   unmarshal_ReqHeartbeat,
}

var RspUnmarshalMap = [...]func(h w2d_packet.Header, bodyData []byte) (interface{}, error){
	w2d_idcmd.Invalid:     unmarshal_RspInvalid,
	w2d_idcmd.ListStage:   unmarshal_RspListStage,
	w2d_idcmd.EnterStage:  unmarshal_RspEnterStage,
	w2d_idcmd.LeaveStage:  unmarshal_RspLeaveStage,
	w2d_idcmd.ChatToStage: unmarshal_RspChatToStage,
	w2d_idcmd.MakeTeam:    unmarshal_RspMakeTeam,
	w2d_idcmd.Act:         unmarshal_RspAct,
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
	if err := json.Unmarshal(bodyData, &args); err != nil {
		return nil, err
	}
	return &args, nil
}

func unmarshal_RspInvalid(h w2d_packet.Header, bodyData []byte) (interface{}, error) {
	var args w2d_obj.RspInvalid_data
	if err := json.Unmarshal(bodyData, &args); err != nil {
		return nil, err
	}
	return &args, nil
}

func unmarshal_ReqListStage(h w2d_packet.Header, bodyData []byte) (interface{}, error) {
	var args w2d_obj.ReqListStage_data
	if err := json.Unmarshal(bodyData, &args); err != nil {
		return nil, err
	}
	return &args, nil
}

func unmarshal_RspListStage(h w2d_packet.Header, bodyData []byte) (interface{}, error) {
	var args w2d_obj.RspListStage_data
	if err := json.Unmarshal(bodyData, &args); err != nil {
		return nil, err
	}
	return &args, nil
}

func unmarshal_ReqEnterStage(h w2d_packet.Header, bodyData []byte) (interface{}, error) {
	var args w2d_obj.ReqEnterStage_data
	if err := json.Unmarshal(bodyData, &args); err != nil {
		return nil, err
	}
	return &args, nil
}

func unmarshal_RspEnterStage(h w2d_packet.Header, bodyData []byte) (interface{}, error) {
	var args w2d_obj.RspEnterStage_data
	if err := json.Unmarshal(bodyData, &args); err != nil {
		return nil, err
	}
	return &args, nil
}

func unmarshal_ReqLeaveStage(h w2d_packet.Header, bodyData []byte) (interface{}, error) {
	var args w2d_obj.ReqLeaveStage_data
	if err := json.Unmarshal(bodyData, &args); err != nil {
		return nil, err
	}
	return &args, nil
}

func unmarshal_RspLeaveStage(h w2d_packet.Header, bodyData []byte) (interface{}, error) {
	var args w2d_obj.RspLeaveStage_data
	if err := json.Unmarshal(bodyData, &args); err != nil {
		return nil, err
	}
	return &args, nil
}

func unmarshal_ReqChatToStage(h w2d_packet.Header, bodyData []byte) (interface{}, error) {
	var args w2d_obj.ReqChatToStage_data
	if err := json.Unmarshal(bodyData, &args); err != nil {
		return nil, err
	}
	return &args, nil
}

func unmarshal_RspChatToStage(h w2d_packet.Header, bodyData []byte) (interface{}, error) {
	var args w2d_obj.RspChatToStage_data
	if err := json.Unmarshal(bodyData, &args); err != nil {
		return nil, err
	}
	return &args, nil
}

func unmarshal_ReqMakeTeam(h w2d_packet.Header, bodyData []byte) (interface{}, error) {
	var args w2d_obj.ReqMakeTeam_data
	if err := json.Unmarshal(bodyData, &args); err != nil {
		return nil, err
	}
	return &args, nil
}

func unmarshal_RspMakeTeam(h w2d_packet.Header, bodyData []byte) (interface{}, error) {
	var args w2d_obj.RspMakeTeam_data
	if err := json.Unmarshal(bodyData, &args); err != nil {
		return nil, err
	}
	return &args, nil
}

func unmarshal_ReqAct(h w2d_packet.Header, bodyData []byte) (interface{}, error) {
	var args w2d_obj.ReqAct_data
	if err := json.Unmarshal(bodyData, &args); err != nil {
		return nil, err
	}
	return &args, nil
}

func unmarshal_RspAct(h w2d_packet.Header, bodyData []byte) (interface{}, error) {
	var args w2d_obj.RspAct_data
	if err := json.Unmarshal(bodyData, &args); err != nil {
		return nil, err
	}
	return &args, nil
}

func unmarshal_ReqHeartbeat(h w2d_packet.Header, bodyData []byte) (interface{}, error) {
	var args w2d_obj.ReqHeartbeat_data
	if err := json.Unmarshal(bodyData, &args); err != nil {
		return nil, err
	}
	return &args, nil
}

func unmarshal_RspHeartbeat(h w2d_packet.Header, bodyData []byte) (interface{}, error) {
	var args w2d_obj.RspHeartbeat_data
	if err := json.Unmarshal(bodyData, &args); err != nil {
		return nil, err
	}
	return &args, nil
}

func unmarshal_NotiInvalid(h w2d_packet.Header, bodyData []byte) (interface{}, error) {
	var args w2d_obj.NotiInvalid_data
	if err := json.Unmarshal(bodyData, &args); err != nil {
		return nil, err
	}
	return &args, nil
}

func unmarshal_NotiStageInfo(h w2d_packet.Header, bodyData []byte) (interface{}, error) {
	var args w2d_obj.NotiStageInfo_data
	if err := json.Unmarshal(bodyData, &args); err != nil {
		return nil, err
	}
	return &args, nil
}

func unmarshal_NotiStatsInfo(h w2d_packet.Header, bodyData []byte) (interface{}, error) {
	var args w2d_obj.NotiStatsInfo_data
	if err := json.Unmarshal(bodyData, &args); err != nil {
		return nil, err
	}
	return &args, nil
}

func unmarshal_NotiStageBroadCast(h w2d_packet.Header, bodyData []byte) (interface{}, error) {
	var args w2d_obj.NotiStageBroadCast_data
	if err := json.Unmarshal(bodyData, &args); err != nil {
		return nil, err
	}
	return &args, nil
}
