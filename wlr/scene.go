package wlr

// #cgo pkg-config: wlroots-0.20
// #cgo CFLAGS: -D_GNU_SOURCE -DWLR_USE_UNSTABLE
// #include <wlr/types/wlr_scene.h>
import "C"

type Scene struct {
	p *C.struct_wlr_scene
}

type SceneOutputLayout struct {
	p *C.struct_wlr_scene_output_layout
}

type SceneTree struct {
	p *C.struct_wlr_scene_tree
}
