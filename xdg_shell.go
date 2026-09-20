package wall

// #cgo pkg-config: wlroots-0.20
// #cgo CFLAGS: -D_GNU_SOURCE -DWLR_USE_UNSTABLE
// #include <wlr/types/wlr_xdg_shell.h>
import "C"

type XDGShell struct {
	p *C.struct_wlr_xdg_shell
}

type XDGToplevel struct {
	p *C.struct_wlr_xdg_toplevel
}
