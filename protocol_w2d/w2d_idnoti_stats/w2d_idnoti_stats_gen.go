// Code generated by "genprotocol -ver=311c9c290570c203090ea3d20ebbf006c810eb958a7a96aef942015fbfd89d2f -basedir=. -prefix=w2d -statstype=int"

package w2d_idnoti_stats

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"

	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_idnoti"
)

type NotiIDStat [w2d_idnoti.NotiID_Count]int

func (es *NotiIDStat) String() string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "NotiIDStats[")
	for i, v := range es {
		fmt.Fprintf(&buf,
			"%v:%v ",
			w2d_idnoti.NotiID(i), v)
	}
	buf.WriteString("]")
	return buf.String()
}
func (es *NotiIDStat) Inc(e w2d_idnoti.NotiID) {
	es[e] += 1
}
func (es *NotiIDStat) Add(e w2d_idnoti.NotiID, v int) {
	es[e] += v
}
func (es *NotiIDStat) SetIfGt(e w2d_idnoti.NotiID, v int) {
	if es[e] < v {
		es[e] = v
	}
}
func (es *NotiIDStat) Get(e w2d_idnoti.NotiID) int {
	return es[e]
}

func (es *NotiIDStat) ToWeb(w http.ResponseWriter, r *http.Request) error {
	tplIndex, err := template.New("index").Funcs(IndexFn).Parse(`
		<html>
		<head>
		<title>NotiID statistics</title>
		</head>
		<body>
		<table border=1 style="border-collapse:collapse;">` +
		HTML_tableheader +
		`{{range $i, $v := .}}` +
		HTML_row +
		`{{end}}` +
		HTML_tableheader +
		`</table>
	
		<br/>
		</body>
		</html>
		`)
	if err != nil {
		return err
	}
	if err := tplIndex.Execute(w, es); err != nil {
		return err
	}
	return nil
}

func Index(i int) string {
	return w2d_idnoti.NotiID(i).String()
}

var IndexFn = template.FuncMap{
	"NotiIDIndex": Index,
}

const (
	HTML_tableheader = `<tr>
		<th>Name</th>
		<th>Value</th>
		</tr>`
	HTML_row = `<tr>
		<td>{{NotiIDIndex $i}}</td>
		<td>{{$v}}</td>
		</tr>
		`
)
