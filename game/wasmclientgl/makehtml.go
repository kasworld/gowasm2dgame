// Copyright 2015,2016,2017,2018,2019,2020,2021 SeukWon Kang (kasworld@gmail.com)
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package wasmclientgl

import (
	"bytes"
	"fmt"

	"github.com/kasworld/gowasm2dgame/config/dataversion"
	"github.com/kasworld/gowasm2dgame/enum/acttype"
	"github.com/kasworld/gowasm2dgame/enum/teamtype"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_version"
)

func (app *WasmClient) makeButtons() string {
	var buf bytes.Buffer
	gameOptions.MakeButtonToolTipTop(&buf)
	return buf.String()
}

func (app *WasmClient) DisplayTextInfo() {
	app.updateLeftInfo()
	app.updateRightInfo()
	app.updateCenterInfo()
}

func (app *WasmClient) makeServiceInfo() string {
	msgCopyright := `</hr>Copyright 2019,2020 SeukWon Kang 
		<a href="https://github.com/kasworld/gowasm2dgame" target="_blank">gowasm2dgame</a>`

	var buf bytes.Buffer
	fmt.Fprintf(&buf, "gowasm2dgame webclient<br/>")
	fmt.Fprintf(&buf, "Protocol %v<br/>", w2d_version.ProtocolVersion)
	fmt.Fprintf(&buf, "Data %v<br/>", dataversion.DataVersion)
	fmt.Fprintf(&buf, "%v<br/>", msgCopyright)
	return buf.String()
}

func (app *WasmClient) makeDebugInfo() string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf,
		"%v<br/>Ping %v<br/>ServerClientTickDiff %v<br/>",
		app.DispInterDur, app.PingDur, app.ServerClientTictDiff,
	)
	return buf.String()
}

func (app *WasmClient) makeTeamStatsInfo() string {
	stats := app.statsInfo
	if stats == nil {
		return ""
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
	return buf.String()
}
