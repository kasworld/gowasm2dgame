// Copyright 2015,2016,2017,2018,2019,2020 SeukWon Kang (kasworld@gmail.com)

package imagecanvas

import (
	"syscall/js"

	"github.com/kasworld/gowasmlib/jslog"
)

type ImageCanvas struct {
	Img js.Value
	W   int
	H   int
	Cnv js.Value
	Ctx js.Value
}

func NewByID(srcImageID string) *ImageCanvas {
	img := js.Global().Get("document").Call("getElementById", srcImageID)
	if !img.Truthy() {
		jslog.Errorf("fail to get %v", srcImageID)
		return nil
	}
	srcw := img.Get("naturalWidth").Int()
	srch := img.Get("naturalHeight").Int()

	cnv := js.Global().Get("document").Call("createElement", "CANVAS")
	ctx := cnv.Call("getContext", "2d")
	if !ctx.Truthy() {
		jslog.Errorf("fail to get context", srcImageID)
		return nil
	}
	// ctx.Set("imageSmoothingEnabled", false)
	// cnv.Set("width", srcw)
	// cnv.Set("height", srch)
	// ctx.Call("clearRect", 0, 0, srcw, srch)
	// ctx.Call("drawImage", img, 0, 0)

	return &ImageCanvas{
		Img: img,
		W:   srcw,
		H:   srch,
		Cnv: cnv,
		Ctx: ctx,
	}
}
