package wall

/*
import (
	"unsafe"
)
*/

// #cgo pkg-config: wayland-server
// #cgo CFLAGS: -D_GNU_SOURCE -DWLR_USE_UNSTABLE
// #include <wayland-util.h>
import "C"

type List struct {
	p *C.struct_wl_list
}
