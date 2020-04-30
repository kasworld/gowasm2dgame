// Code generated by "genenum -typename=EffectType -packagename=effecttype -basedir=enum -vectortype=int"

package effecttype_vector

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
)

type EffectTypeVector [effecttype.EffectType_Count]int

func (es *EffectTypeVector) String() string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "EffectTypeVector[")
	for i, v := range es {
		fmt.Fprintf(&buf,
			"%v:%v ",
			effecttype.EffectType(i), v)
	}
	buf.WriteString("]")
	return buf.String()
}
func (es *EffectTypeVector) Dec(e effecttype.EffectType) {
	es[e] -= 1
}
func (es *EffectTypeVector) Inc(e effecttype.EffectType) {
	es[e] += 1
}
func (es *EffectTypeVector) Add(e effecttype.EffectType, v int) {
	es[e] += v
}
func (es *EffectTypeVector) SetIfGt(e effecttype.EffectType, v int) {
	if es[e] < v {
		es[e] = v
	}
}
func (es *EffectTypeVector) Get(e effecttype.EffectType) int {
	return es[e]
}

// Iter return true if iter stop, return false if iter all
// fn return true to stop iter
func (es EffectTypeVector) Iter(fn func(i int, v int) bool) bool {
	for i := 0; i < effecttype.EffectType_Count; i++ {
		if fn(i, es[i]) {
			return true
		}
	}
	return false
}

// VectorAdd add element to element
func (es EffectTypeVector) VectorAdd(arg EffectTypeVector) EffectTypeVector {
	var rtn EffectTypeVector
	for i := 0; i < effecttype.EffectType_Count; i++ {
		rtn[i] = es[i] + arg[i]
	}
	return rtn
}

// VectorSub sub element to element
func (es EffectTypeVector) VectorSub(arg EffectTypeVector) EffectTypeVector {
	var rtn EffectTypeVector
	for i := 0; i < effecttype.EffectType_Count; i++ {
		rtn[i] = es[i] - arg[i]
	}
	return rtn
}

func (es *EffectTypeVector) ToWeb(w http.ResponseWriter, r *http.Request) error {
	tplIndex, err := template.New("index").Funcs(IndexFn).Parse(`
		<html>
		<head>
		<title>EffectType statistics</title>
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
	return effecttype.EffectType(i).String()
}

var IndexFn = template.FuncMap{
	"EffectTypeIndex": Index,
}

const (
	HTML_tableheader = `<tr>
		<th>Name</th>
		<th>Value</th>
		</tr>`
	HTML_row = `<tr>
		<td>{{EffectTypeIndex $i}}</td>
		<td>{{$v}}</td>
		</tr>
		`
)
