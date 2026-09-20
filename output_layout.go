package wall

// #cgo pkg-config: wlroots-0.20
// #cgo CFLAGS: -D_GNU_SOURCE -DWLR_USE_UNSTABLE
// #include <wlr/types/wlr_output_layout.h>
import "C"

type OutputLayout struct {
	p *C.struct_wlr_box
}
