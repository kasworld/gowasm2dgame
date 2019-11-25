// Copyright 2015,2016,2017,2018,2019 SeukWon Kang (kasworld@gmail.com)
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
	"fmt"
	"math"
	"syscall/js"
)

func calcCircularPos(cx, cy int, angle, r int) (int, int) {
	rad := float64(angle) * math.Pi / 180
	dstx := float64(cx) + float64(r)*math.Cos(rad)
	dsty := float64(cy) + float64(r)*math.Sin(rad)
	return int(dstx), int(dsty)
}

func getImgWH(srcImageID string) (js.Value, int, int) {
	img := js.Global().Get("document").Call("getElementById", srcImageID)
	if !img.Truthy() {
		fmt.Printf("fail to get %v", srcImageID)
		return js.Null(), 0, 0
	}
	srcw := img.Get("naturalWidth").Int()
	srch := img.Get("naturalHeight").Int()
	return img, srcw, srch
}

func getCnv2dCtx(dstCanvasID string) (js.Value, js.Value) {
	dstcnv := js.Global().Get("document").Call("getElementById", dstCanvasID)
	if !dstcnv.Truthy() {
		fmt.Printf("fail to get canvas\n")
		return js.Null(), js.Null()
	}
	dstctx := dstcnv.Call("getContext", "2d")
	if !dstctx.Truthy() {
		fmt.Printf("fail to get context\n")
		return js.Null(), js.Null()
	}
	dstctx.Set("imageSmoothingEnabled", false)
	return dstcnv, dstctx
}
