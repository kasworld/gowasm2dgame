// Code generated by "genprotocol -ver=ed65a653bd268dc21902d5d07939f7bfc1ba6b98026a426c30526d9f59ba8d12 -basedir=. -prefix=w2d -statstype=int"

package w2d_statserveapi

import (
	"fmt"
	"net/http"
	"sync"
	"text/template"
	"time"

	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_idcmd"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_packet"
)

func (ps *StatServeAPI) String() string {
	return fmt.Sprintf("StatServeAPI[%v]", len(ps))
}

type StatServeAPI [w2d_idcmd.CommandID_Count]StatRow

func New() *StatServeAPI {
	ps := new(StatServeAPI)
	for i := 0; i < w2d_idcmd.CommandID_Count; i++ {
		ps[i].Name = w2d_idcmd.CommandID(i).String()
	}
	return ps
}
func (ps *StatServeAPI) AfterRecvReqHeader(header w2d_packet.Header) (*StatObj, error) {
	if int(header.Cmd) >= w2d_idcmd.CommandID_Count {
		return nil, fmt.Errorf("CommandID out of range %v %v", header, w2d_idcmd.CommandID_Count)
	}
	return ps[header.Cmd].open(header), nil
}
func (ws *StatServeAPI) ToWeb(w http.ResponseWriter, r *http.Request) error {
	tplIndex, err := template.New("index").Parse(`
	<html><head><title>Serve API stat Info</title></head><body>
	<table border=1 style="border-collapse:collapse;">` +
		HTML_tableheader +
		`{{range $i, $v := .}}` +
		HTML_row +
		`{{end}}` +
		HTML_tableheader +
		`</table><br/>
	</body></html>`)
	if err != nil {
		return err
	}
	if err := tplIndex.Execute(w, ws); err != nil {
		return err
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
type StatObj struct {
	RecvTime    time.Time
	APICallTime time.Time
	StatRef     *StatRow
}

func (sm *StatObj) BeforeAPICall() {
	sm.APICallTime = time.Now().UTC()
	sm.StatRef.afterAuth()
}
func (sm *StatObj) AfterAPICall() {
	sm.StatRef.apiEnd(time.Now().UTC().Sub(sm.APICallTime))
}
func (sm *StatObj) AfterSendRsp(hd w2d_packet.Header) {
	sm.StatRef.afterSend(time.Now().UTC().Sub(sm.RecvTime), hd)
}

////////////////////////////////////////////////////////////////////////////////
type PacketID2StatObj struct {
	mutex sync.RWMutex
	stats map[uint32]*StatObj
}

func NewPacketID2StatObj() *PacketID2StatObj {
	return &PacketID2StatObj{
		stats: make(map[uint32]*StatObj),
	}
}
func (som *PacketID2StatObj) Add(pkid uint32, so *StatObj) error {
	som.mutex.Lock()
	defer som.mutex.Unlock()
	if _, exist := som.stats[pkid]; exist {
		return fmt.Errorf("pkid exist %v", pkid)
	}
	som.stats[pkid] = so
	return nil
}
func (som *PacketID2StatObj) Del(pkid uint32) *StatObj {
	som.mutex.Lock()
	defer som.mutex.Unlock()
	so := som.stats[pkid]
	delete(som.stats, pkid)
	return so
}
func (som *PacketID2StatObj) Get(pkid uint32) *StatObj {
	som.mutex.RLock()
	defer som.mutex.RUnlock()
	return som.stats[pkid]
}

////////////////////////////////////////////////////////////////////////////////
const (
	HTML_tableheader = `<tr>
	<th>Name</th>
	<th>Recv Count</th>
	<th>Auth Count</th>
	<th>APIEnd Count</th>
	<th>Send Count</th>
	<th>Run Count</th>
	<th>Fail Count</th>
	<th>RecvSend Avg ms</th>
	<th>API Avg ms</th>
	<th>Rx Avg Byte</th>
	<th>Rx Max Byte</th>
	<th>Tx Avg Byte</th>
	<th>Tx Max Byte</th>
	</tr>`
	HTML_row = `<tr>
	<td>{{$v.Name}}</td>
	<td>{{$v.RecvCount}}</td>
	<td>{{$v.AuthCount}}</td>
	<td>{{$v.APIEndCount}}</td>
	<td>{{$v.SendCount}}</td>
	<td>{{$v.RunCount}}</td>
	<td>{{$v.FailCount}}</td>
	<td>{{printf "%13.6f" $v.RSAvgms }}</td>
	<td>{{printf "%13.6f" $v.APIAvgms }}</td>
	<td>{{printf "%10.3f" $v.AvgRxByte }}</td>
	<td>{{$v.MaxRecvBytes }}</td>
	<td>{{printf "%10.3f" $v.AvgTxByte }}</td>
	<td>{{$v.MaxSendBytes }}</td>
	</tr>
	`
)

type StatRow struct {
	mutex          sync.Mutex
	Name           string
	RecvCount      int
	MaxRecvBytes   int
	RecvBytes      int
	SendCount      int
	MaxSendBytes   int
	SendBytes      int
	RecvSendDurSum time.Duration
	AuthCount      int
	APIEndCount    int
	APIDurSum      time.Duration
}

func (sr *StatRow) open(hd w2d_packet.Header) *StatObj {
	sr.mutex.Lock()
	defer sr.mutex.Unlock()
	sr.RecvCount++
	rxbyte := int(hd.BodyLen()) + w2d_packet.HeaderLen
	sr.RecvBytes += rxbyte
	if sr.MaxRecvBytes < rxbyte {
		sr.MaxRecvBytes = rxbyte
	}
	rtn := &StatObj{
		RecvTime: time.Now().UTC(),
		StatRef:  sr,
	}
	return rtn
}
func (sr *StatRow) afterAuth() {
	sr.mutex.Lock()
	defer sr.mutex.Unlock()
	sr.AuthCount++
}
func (sr *StatRow) apiEnd(diffDur time.Duration) {
	sr.mutex.Lock()
	defer sr.mutex.Unlock()
	sr.APIEndCount++
	sr.APIDurSum += diffDur
}
func (sr *StatRow) afterSend(diffDur time.Duration, hd w2d_packet.Header) {
	sr.mutex.Lock()
	defer sr.mutex.Unlock()
	sr.SendCount++
	txbyte := int(hd.BodyLen()) + w2d_packet.HeaderLen
	sr.SendBytes += txbyte
	if sr.MaxSendBytes < txbyte {
		sr.MaxSendBytes = txbyte
	}
	sr.RecvSendDurSum += diffDur
}

////////////////////////////////////////////////////////////////////////////////
func (sr *StatRow) RunCount() int {
	return sr.AuthCount - sr.APIEndCount
}
func (sr *StatRow) FailCount() int {
	return sr.APIEndCount - sr.SendCount
}
func (sr *StatRow) RSAvgms() float64 {
	if sr.SendCount == 0 {
		return 0
	}
	return float64(sr.RecvSendDurSum) / float64(sr.SendCount*1000000)
}
func (sr *StatRow) APIAvgms() float64 {
	if sr.APIEndCount == 0 {
		return 0
	}
	return float64(sr.APIDurSum) / float64(sr.APIEndCount*1000000)
}
func (sr *StatRow) AvgRxByte() float64 {
	if sr.RecvCount == 0 {
		return 0
	}
	return float64(sr.RecvBytes) / float64(sr.RecvCount)
}
func (sr *StatRow) AvgTxByte() float64 {
	if sr.SendCount == 0 {
		return 0
	}
	return float64(sr.SendBytes) / float64(sr.SendCount)
}