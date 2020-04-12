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

package wasmclient

import (
	"bytes"
	"fmt"
	"syscall/js"

	"github.com/kasworld/gowasm2dgame/enums/acttype"
	"github.com/kasworld/gowasm2dgame/enums/teamtype"
	"github.com/kasworld/gowasm2dgame/game/gameconst"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_version"
)

func (app *WasmClient) updataServiceInfo() {
	msgCopyright := `</hr>Copyright 2019 SeukWon Kang 
		<a href="https://github.com/kasworld/gowasm2dgame" target="_blank">gowasm2dgame</a>`

	var buf bytes.Buffer
	fmt.Fprintf(&buf, "gowasm2dgame webclient<br/>")
	fmt.Fprintf(&buf, "Protocol %v<br/>", w2d_version.ProtocolVersion)
	fmt.Fprintf(&buf, "Data %v<br/>", gameconst.DataVersion)
	fmt.Fprintf(&buf, "%v<br/>", msgCopyright)
	div := js.Global().Get("document").Call("getElementById", "serviceinfo")
	div.Set("innerHTML", buf.String())
}

func (app *WasmClient) updateDebugInfo() {
	var buf bytes.Buffer
	fmt.Fprintf(&buf,
		"%v<br/>Ping %v<br/>ServerClientTickDiff %v<br/>",
		app.DispInterDur, app.PingDur, app.ServerClientTictDiff,
	)

	div := js.Global().Get("document").Call("getElementById", "debuginfo")
	div.Set("innerHTML", buf.String())
}

func (app *WasmClient) updateSysmsg() {
	app.systemMessage = app.systemMessage.GetLastN(100)
	div := js.Global().Get("document").Call("getElementById", "sysmsg")
	div.Set("innerHTML", app.systemMessage.ToHtmlStringRev())
}

func (app *WasmClient) updateTeamStatsInfo() {
	stats := app.statsInfo
	if stats == nil {
		return
	}
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "Stage %v<br/>", stats.UUID)

	buf.WriteString(`<table border=1 style="border-collapse:collapse;">`)

	buf.WriteString(`<tr><th>act\team</th>`)
	for teami := 0; teami < teamtype.TeamType_Count; teami++ {
		fmt.Fprintf(&buf, "<th>%v</th>", teamtype.TeamType(teami))
	}
	buf.WriteString(`</tr>`)

	buf.WriteString(`<tr><td>UUID</td>`)
	for _, v := range stats.Stats {
		fmt.Fprintf(&buf, "<td>%v</td>", v.UUID)
	}
	buf.WriteString(`</tr>`)

	buf.WriteString(`<tr><td>Alive</td>`)
	for _, v := range stats.Stats {
		fmt.Fprintf(&buf, "<td>%v</td>", v.Alive)
	}
	buf.WriteString(`</tr>`)

	for acti := 0; acti < acttype.ActType_Count; acti++ {
		fmt.Fprintf(&buf, "<tr><td>%v</td>", acttype.ActType(acti))
		for teami := 0; teami < teamtype.TeamType_Count; teami++ {
			fmt.Fprintf(&buf, "<td>%v</td>", stats.Stats[teami].ActStats[acti])
		}
		buf.WriteString(`</tr>`)
	}
	buf.WriteString(`</table>`)
	div := js.Global().Get("document").Call("getElementById", "teamstatsinfo")
	div.Set("innerHTML", buf.String())
}
