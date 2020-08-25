// Code generated by "genprotocol.exe -ver=8ce3a4010a59de778695c59389c0fd9a3938197ee346a449f005b536d94d0e60 -basedir=protocol_w2d -prefix=w2d -statstype=int"

package w2d_statcallapi

import (
	"fmt"
	"html/template"
	"net/http"
	"sync"
	"time"

	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_idcmd"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_packet"
)

func (cps *StatCallAPI) String() string {
	return fmt.Sprintf("StatCallAPI[%v]",
		len(cps))
}

type StatCallAPI [w2d_idcmd.CommandID_Count]StatRow

func New() *StatCallAPI {
	cps := new(StatCallAPI)
	for i := 0; i < w2d_idcmd.CommandID_Count; i++ {
		cps[i].Name = w2d_idcmd.CommandID(i).String()
	}
	return cps
}
func (cps *StatCallAPI) BeforeSendReq(header w2d_packet.Header) (*statObj, error) {
	if int(header.Cmd) >= w2d_idcmd.CommandID_Count {
		return nil, fmt.Errorf("CommandID out of range %v %v",
			header, w2d_idcmd.CommandID_Count)
	}
	return cps[header.Cmd].open(), nil
}
func (cps *StatCallAPI) AfterSendReq(header w2d_packet.Header) error {
	if int(header.Cmd) >= w2d_idcmd.CommandID_Count {
		return fmt.Errorf("CommandID out of range %v %v", header, w2d_idcmd.CommandID_Count)
	}
	n := int(header.BodyLen()) + w2d_packet.HeaderLen
	cps[header.Cmd].addTx(n)
	return nil
}
func (cps *StatCallAPI) AfterRecvRsp(header w2d_packet.Header) error {
	if int(header.Cmd) >= w2d_idcmd.CommandID_Count {
		return fmt.Errorf("CommandID out of range %v %v", header, w2d_idcmd.CommandID_Count)
	}
	n := int(header.BodyLen()) + w2d_packet.HeaderLen
	cps[header.Cmd].addRx(n)
	return nil
}
func (ws *StatCallAPI) ToWeb(w http.ResponseWriter, r *http.Request) error {
	tplIndex, err := template.New("index").Parse(`
	<html><head><title>Call API Stat Info</title></head><body>
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
type statObj struct {
	StartTime time.Time
	StatRef   *StatRow
}

func (so *statObj) CallServerEnd(success bool) {
	so.StatRef.close(success, so.StartTime)
}

////////////////////////////////////////////////////////////////////////////////
type PacketID2StatObj struct {
	mutex sync.RWMutex
	stats map[uint32]*statObj
}

func NewPacketID2StatObj() *PacketID2StatObj {
	return &PacketID2StatObj{
		stats: make(map[uint32]*statObj),
	}
}
func (som *PacketID2StatObj) Add(pkid uint32, so *statObj) error {
	som.mutex.Lock()
	defer som.mutex.Unlock()
	if _, exist := som.stats[pkid]; exist {
		return fmt.Errorf("pkid exist %v", pkid)
	}
	som.stats[pkid] = so
	return nil
}
func (som *PacketID2StatObj) Del(pkid uint32) *statObj {
	som.mutex.Lock()
	defer som.mutex.Unlock()
	so := som.stats[pkid]
	delete(som.stats, pkid)
	return so
}
func (som *PacketID2StatObj) Get(pkid uint32) *statObj {
	som.mutex.RLock()
	defer som.mutex.RUnlock()
	return som.stats[pkid]
}

////////////////////////////////////////////////////////////////////////////////
const (
	HTML_tableheader = `<tr>
	<th>Name</th>
	<th>Start</th>
	<th>End</th>
	<th>Success</th>
	<th>Running</th>
	<th>Fail</th>
	<th>Avg ms</th>
	<th>TxAvg Byte</th>
	<th>RxAvg Byte</th>
	</tr>`
	HTML_row = `<tr>
	<td>{{$v.Name}}</td>
	<td>{{$v.StartCount}}</td>
	<td>{{$v.EndCount}}</td>
	<td>{{$v.SuccessCount}}</td>
	<td>{{$v.RunCount}}</td>
	<td>{{$v.FailCount}}</td>
	<td>{{printf "%13.6f" $v.Avgms }}</td>
	<td>{{printf "%10.3f" $v.AvgTx }}</td>
	<td>{{printf "%10.3f" $v.AvgRx }}</td>
	</tr>
	`
)

type StatRow struct {
	mutex        sync.Mutex
	Name         string
	TxCount      int
	TxByte       int
	RxCount      int
	RxByte       int
	StartCount   int
	EndCount     int
	SuccessCount int
	Sum          time.Duration
}

func (sr *StatRow) open() *statObj {
	sr.mutex.Lock()
	defer sr.mutex.Unlock()
	sr.StartCount++
	return &statObj{
		StartTime: time.Now(),
		StatRef:   sr,
	}
}
func (sr *StatRow) close(success bool, startTime time.Time) {
	sr.mutex.Lock()
	defer sr.mutex.Unlock()
	sr.EndCount++
	if success {
		sr.SuccessCount++
		sr.Sum += time.Now().Sub(startTime)
	}
}
func (sr *StatRow) addTx(n int) {
	sr.mutex.Lock()
	defer sr.mutex.Unlock()
	sr.TxCount++
	sr.TxByte += n
}
func (sr *StatRow) addRx(n int) {
	sr.mutex.Lock()
	defer sr.mutex.Unlock()
	sr.RxCount++
	sr.RxByte += n
}
func (sr *StatRow) RunCount() int {
	return sr.StartCount - sr.EndCount
}
func (sr *StatRow) FailCount() int {
	return sr.EndCount - sr.SuccessCount
}
func (sr *StatRow) Avgms() float64 {
	if sr.EndCount != 0 {
		return float64(sr.Sum) / float64(sr.EndCount*1000000)
	}
	return 0.0
}
func (sr *StatRow) AvgRx() float64 {
	if sr.EndCount != 0 {
		return float64(sr.RxByte) / float64(sr.RxCount)
	}
	return 0.0
}
func (sr *StatRow) AvgTx() float64 {
	if sr.EndCount != 0 {
		return float64(sr.TxByte) / float64(sr.TxCount)
	}
	return 0.0
}
