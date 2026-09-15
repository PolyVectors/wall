package wlr

// #cgo pkg-config: wlroots-0.20
// #cgo CFLAGS: -D_GNU_SOURCE -DWLR_USE_UNSTABLE
// #include <wlr/types/wlr_xcursor_manager.h>
import "C"

type XCursorManager struct {
	p *C.struct_wlr_xcursor_manager
}
