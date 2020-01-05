// Code generated by "genprotocol -ver=ed65a653bd268dc21902d5d07939f7bfc1ba6b98026a426c30526d9f59ba8d12 -basedir=. -prefix=w2d -statstype=int"

package w2d_idnoti

import "fmt"

type NotiID uint16 // use in packet header, DO NOT CHANGE
const (
	Invalid   NotiID = iota //
	StageInfo               // // game stage info to display
	StatsInfo               // // game stats info
	StageChat               //

	NotiID_Count int = iota
)

var _NotiID2string = [NotiID_Count]string{
	Invalid:   "Invalid",
	StageInfo: "StageInfo",
	StatsInfo: "StatsInfo",
	StageChat: "StageChat",
}

func (e NotiID) String() string {
	if e >= 0 && e < NotiID(NotiID_Count) {
		return _NotiID2string[e]
	}
	return fmt.Sprintf("NotiID%d", uint16(e))
}

var _string2NotiID = map[string]NotiID{
	"Invalid":   Invalid,
	"StageInfo": StageInfo,
	"StatsInfo": StatsInfo,
	"StageChat": StageChat,
}

func String2NotiID(s string) (NotiID, bool) {
	v, b := _string2NotiID[s]
	return v, b
}