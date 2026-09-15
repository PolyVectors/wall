package wlr

// #cgo pkg-config: wlroots-0.20
// #cgo CFLAGS: -D_GNU_SOURCE -DWLR_USE_UNSTABLE
// #include <wlr/backend.h>
import "C"

type Backend struct {
	p *C.struct_wlr_backend
}
