package wlr

// #cgo pkg-config: wlroots-0.20
// #cgo CFLAGS: -D_GNU_SOURCE -DWLR_USE_UNSTABLE
// #include <wlr/render/wlr_renderer.h>
import "C"

type Renderer struct {
	p *C.struct_wlr_renderer
}
