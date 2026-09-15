package main

import (
	"wall/wl"
	"wall/wlr"
)

type CursorMode int

const (
	CursorPassthrough CursorMode = iota
	CursorModeMove
	CursorModeResize
)

type Toplevel struct {
	link				wl.List
	server			*Server
	xdgToplevel		wlr.XDGToplevel
	sceneTree			wlr.SceneTree
	mapped			wl.Listener
	unmap			wl.Listener
	commit			wl.Listener
	destroy			wl.Listener
	requestMove		wl.Listener
	requestResize		wl.Listener
	requestMaximize	wl.Listener
	requestFullscreen	wl.Listener
}

type Server struct {
	display				wl.Display
	backend				wlr.Backend
	renderer				wlr.Renderer
	allocator				wlr.Allocator
	scene				wlr.Scene
	sceneLayout			wlr.SceneOutputLayout

	xdgShell				wlr.XDGShell
	newXdgToplevel		wl.Listener
	newXdgPopup			wl.Listener
	// could this be a shim to containers/list or a dynamic array when possible?
	toplevels				wl.List 

	cursor				wlr.Cursor
	cursorManager			wlr.XCursorManager
	cursorMotion			wl.Listener
	cursorMotionAbsolute	wl.Listener
	cursorButton			wl.Listener
	cursorAxis			wl.Listener
	cursorFrame			wl.Listener

	seat					wlr.Seat
	newInput				wl.Listener
	requestCursor			wl.Listener
	pointerFocusChange	wl.Listener
	requestSetSelection		wl.Listener
	keyboards			wl.List
	cursorMode			CursorMode
	grabbedToplevel		*Toplevel
	grabX, grabY			float64
	grabGeobox			wlr.Box
	resizeEdges			uint32

	outputLayout			wlr.OutputLayout
	// ditto shim idea
	outputs				wl.List
	newOutput			wl.Listener			
}

func main() {
}
