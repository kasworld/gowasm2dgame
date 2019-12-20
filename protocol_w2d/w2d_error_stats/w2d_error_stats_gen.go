// Code generated by "genprotocol -ver=8cb09769c07a0cf3e7042afbf364a4eff2c960eafe6d9a8ccbb46041a984713a -basedir=. -prefix=w2d -statstype=int"

package w2d_error_stats

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"

	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_error"
)

type ErrorCodeStat [w2d_error.ErrorCode_Count]int

func (es *ErrorCodeStat) String() string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "ErrorCodeStats[")
	for i, v := range es {
		fmt.Fprintf(&buf,
			"%v:%v ",
			w2d_error.ErrorCode(i), v)
	}
	buf.WriteString("]")
	return buf.String()
}
func (es *ErrorCodeStat) Inc(e w2d_error.ErrorCode) {
	es[e] += 1
}
func (es *ErrorCodeStat) Add(e w2d_error.ErrorCode, v int) {
	es[e] += v
}
func (es *ErrorCodeStat) SetIfGt(e w2d_error.ErrorCode, v int) {
	if es[e] < v {
		es[e] = v
	}
}
func (es *ErrorCodeStat) Get(e w2d_error.ErrorCode) int {
	return es[e]
}

func (es *ErrorCodeStat) ToWeb(w http.ResponseWriter, r *http.Request) error {
	tplIndex, err := template.New("index").Funcs(IndexFn).Parse(`
		<html>
		<head>
		<title>ErrorCode statistics</title>
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
	return w2d_error.ErrorCode(i).String()
}

var IndexFn = template.FuncMap{
	"ErrorCodeIndex": Index,
}

const (
	HTML_tableheader = `<tr>
		<th>Name</th>
		<th>Value</th>
		</tr>`
	HTML_row = `<tr>
		<td>{{ErrorCodeIndex $i}}</td>
		<td>{{$v}}</td>
		</tr>
		`
)
