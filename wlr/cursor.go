package wlr

// #cgo pkg-config: wlroots-0.20
// #cgo CFLAGS: -D_GNU_SOURCE -DWLR_USE_UNSTABLE
// #include <wlr/types/wlr_cursor.h>
import "C"

type Cursor struct {
	p *C.struct_wlr_cursor
}
