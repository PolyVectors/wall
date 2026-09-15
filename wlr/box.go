package wlr

// #cgo pkg-config: wlroots-0.20
// #cgo CFLAGS: -D_GNU_SOURCE -DWLR_USE_UNSTABLE
// #include <wlr/util/box.h>
import "C"

// can be a wrapper around Box2i via concorde
type Box struct {
	p *C.struct_wlr_box
}
