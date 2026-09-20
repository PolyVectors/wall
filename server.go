package wall

/*
import (
	"unsafe"
)
*/

// #cgo pkg-config: wlroots-0.20 wayland-server
// #cgo CFLAGS: -D_GNU_SOURCE -DWLR_USE_UNSTABLE
// #include <wayland-server-core.h>
// #include <wlr/backend.h>
import "C"

type EventLoop struct {
	p *C.struct_wl_event_loop
}

type Display struct {
	p *C.struct_wl_display
}

func NewDisplay() Display {
	return Display{p: C.wl_display_create()}
}

func (d Display) GetEventLoop() EventLoop {
	return EventLoop{p: C.wl_display_get_event_loop(d.p)}
}
