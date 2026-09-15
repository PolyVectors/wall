package wl

/*
import (
	"unsafe"
)
*/

// #cgo pkg-config: wayland-server
// #cgo CFLAGS: -D_GNU_SOURCE -DWLR_USE_UNSTABLE
// #include <wayland-server-core.h>
import "C"

type Display struct {
	p *C.struct_wl_display
}

type Listener struct {
	p *C.struct_wl_listener
}
