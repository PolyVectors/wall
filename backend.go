package wall

// #cgo pkg-config: wlroots-0.20
// #cgo CFLAGS: -D_GNU_SOURCE -DWLR_USE_UNSTABLE
// #include <wlr/backend.h>
// #include <wlr/backend/session.h>
import "C"

type Session struct {
	p *C.struct_wlr_session
}

type Backend struct {
	p *C.struct_wlr_backend
}

func (e EventLoop) BackendAutocreate(session *Session) Backend {
	return Backend{p: C.wlr_backend_autocreate(e.p, nil)}
}
