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
	"math"
	"syscall/js"

	"github.com/kasworld/gowasm2dgame/enum/teamtype"

	"github.com/kasworld/gowasm2dgame/config/gameconst"
	"github.com/kasworld/gowasm2dgame/enum/gameobjtype"
	"github.com/kasworld/gowasm2dgame/protocol_w2d/w2d_obj"
)

func (vp *Viewport) ThreeJsNew(name string, args ...interface{}) js.Value {
	return vp.threejs.Get(name).New(args...)
}

func (vp *Viewport) initGrid() {
	outerStageSize := gameconst.StageSize + gameobjtype.MaxRadius*2
	innerStageSize := gameconst.StageSize

	helper := vp.ThreeJsNew("GridHelper", outerStageSize, 100, 0x0000ff, 0x404040)
	helper.Get("position").Set("x", innerStageSize/2)
	helper.Get("position").Set("y", -gameobjtype.MaxRadius)
	helper.Get("position").Set("z", innerStageSize/2)
	vp.scene.Call("add", helper)

	helper = vp.ThreeJsNew("GridHelper", outerStageSize, 100, 0xffff00, 0x404040)
	helper.Get("position").Set("x", innerStageSize/2)
	helper.Get("position").Set("y", gameconst.StageSize+gameobjtype.MaxRadius)
	helper.Get("position").Set("z", innerStageSize/2)
	vp.scene.Call("add", helper)

	helper = vp.ThreeJsNew("GridHelper", outerStageSize, 100, 0xff0000, 0x404040)
	helper.Get("rotation").Set("z", math.Pi/2)
	helper.Get("position").Set("x", -gameobjtype.MaxRadius)
	helper.Get("position").Set("y", innerStageSize/2)
	helper.Get("position").Set("z", innerStageSize/2)
	vp.scene.Call("add", helper)

	helper = vp.ThreeJsNew("GridHelper", outerStageSize, 100, 0x00ffff, 0x404040)
	helper.Get("rotation").Set("z", math.Pi/2)
	helper.Get("position").Set("x", gameconst.StageSize+gameobjtype.MaxRadius)
	helper.Get("position").Set("y", innerStageSize/2)
	helper.Get("position").Set("z", innerStageSize/2)
	vp.scene.Call("add", helper)

	helper = vp.ThreeJsNew("GridHelper", outerStageSize, 100, 0x00ff00, 0x404040)
	helper.Get("rotation").Set("x", math.Pi/2)
	helper.Get("position").Set("x", innerStageSize/2)
	helper.Get("position").Set("y", innerStageSize/2)
	helper.Get("position").Set("z", -gameobjtype.MaxRadius)
	vp.scene.Call("add", helper)

	helper = vp.ThreeJsNew("GridHelper", outerStageSize, 100, 0xff00ff, 0x404040)
	helper.Get("rotation").Set("x", math.Pi/2)
	helper.Get("position").Set("x", innerStageSize/2)
	helper.Get("position").Set("y", innerStageSize/2)
	helper.Get("position").Set("z", gameconst.StageSize+gameobjtype.MaxRadius)
	vp.scene.Call("add", helper)

	box3 := vp.ThreeJsNew("Box3",
		vp.ThreeJsNew("Vector3", 0, 0, 0),
		vp.ThreeJsNew("Vector3", innerStageSize, innerStageSize, innerStageSize),
	)
	helper = vp.ThreeJsNew("Box3Helper", box3, 0xffffff)
	vp.scene.Call("add", helper)

	// axisHelper := vp.ThreeJsNew("AxesHelper", gameconst.StageSize)
	// vp.scene.Call("add", axisHelper)
}
func (vp *Viewport) initLight() {
	vp.light = vp.ThreeJsNew("PointLight", 0x808080, 1)
	vp.scene.Call("add", vp.light)
	// vp.light.Get("position").Set("x", vt[0])
	// vp.light.Get("position").Set("y", vt[1])
	// vp.light.Get("position").Set("z", vt[2])
}

func (vp *Viewport) getGeometry(gotype gameobjtype.GameObjType) js.Value {
	geo, exist := vp.geometryCache[gotype]
	if !exist {
		radius := gameobjtype.Attrib[gotype].Radius
		switch gotype {
		default:
			geo = vp.ThreeJsNew("SphereGeometry", radius, 32, 16)
		case gameobjtype.Ball:
			geo = vp.ThreeJsNew("TorusGeometry", radius, radius/2, 16, 64)
		case gameobjtype.Shield:
			geo = vp.ThreeJsNew("IcosahedronGeometry", radius)
		case gameobjtype.HommingShield:
			geo = vp.ThreeJsNew("OctahedronGeometry", radius)
			// geo = vp.ThreeJsNew("IcosahedronGeometry", radius)
		case gameobjtype.Bullet:
			geo = vp.ThreeJsNew("DodecahedronGeometry", radius)
		case gameobjtype.HommingBullet:
			geo = vp.ThreeJsNew("TetrahedronGeometry", radius)
			//geo = vp.ThreeJsNew("OctahedronGeometry", radius)
		case gameobjtype.SuperBullet:
			geo = vp.ThreeJsNew("ConeGeometry", radius, radius*2, 16)
			// geo = vp.ThreeJsNew("TetrahedronGeometry", radius)
		}
		vp.geometryCache[gotype] = geo
	}
	return geo
}

func (vp *Viewport) getMaterial(co uint32) js.Value {
	mat, exist := vp.materialCache[co]
	if !exist {
		mat = vp.ThreeJsNew("MeshStandardMaterial")
		// material.Set("color", vp.ToThColor(htmlcolors.Gray))
		mat.Set("emissive", vp.ThreeJsNew("Color", co))
		mat.Set("shininess", 30)
		vp.materialCache[co] = mat
	}
	return mat
}

func (vp *Viewport) add2Scene(o *w2d_obj.GameObj, co uint32) js.Value {
	if jso, exist := vp.jsSceneObjs[o.UUID]; exist {
		jso.Get("position").Set("x", o.PosVt[0])
		jso.Get("position").Set("y", o.PosVt[1])
		jso.Get("position").Set("z", 0)
		// jso.Get("rotation").Set("x", o.RotVt[0])
		// jso.Get("rotation").Set("y", o.RotVt[1])
		// jso.Get("rotation").Set("z", o.RotVt[2])
		return jso
	}
	geometry := vp.getGeometry(o.GOType)
	material := vp.getMaterial(co)
	jso := vp.ThreeJsNew("Mesh", geometry, material)
	jso.Get("position").Set("x", o.PosVt[0])
	jso.Get("position").Set("y", o.PosVt[1])
	jso.Get("position").Set("z", 0)
	// jso.Get("rotation").Set("x", o.RotVt[0])
	// jso.Get("rotation").Set("y", o.RotVt[1])
	// jso.Get("rotation").Set("z", o.RotVt[2])
	vp.scene.Call("add", jso)
	vp.jsSceneObjs[o.UUID] = jso
	return jso
}

func (vp *Viewport) processRecvStageInfo(stageInfo *w2d_obj.NotiStageInfo_data) {
	setCamera := false
	addUUID := make(map[string]bool)
	for _, tm := range stageInfo.Teams {
		if tm == nil {
			continue
		}
		if !setCamera {
			setCamera = true

			vp.camera.Get("position").Set("x", gameconst.StageSize/2)
			vp.camera.Get("position").Set("y", gameconst.StageSize/2)
			vp.camera.Get("position").Set("z", gameconst.StageSize)

			vp.camera.Call("lookAt",
				vp.ThreeJsNew("Vector3",
					gameconst.StageSize/2, gameconst.StageSize/2, 0,
				),
			)
			vp.camera.Call("updateProjectionMatrix")
		}
		v := tm.Ball
		vp.add2Scene(v,
			uint32(teamtype.Attrib[tm.TeamType].Color24))
		addUUID[v.UUID] = true
		for _, v := range tm.Objs {
			if v == nil {
				continue
			}
			vp.add2Scene(v,
				uint32(teamtype.Attrib[tm.TeamType].Color24))
			addUUID[v.UUID] = true
		}
	}
	for id, jso := range vp.jsSceneObjs {
		if !addUUID[id] {
			vp.scene.Call("remove", jso)
			delete(vp.jsSceneObjs, id)
		}
	}
}
