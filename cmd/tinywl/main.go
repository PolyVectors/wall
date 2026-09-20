package main

import (
	"fmt"

	"wall"
)

type CursorMode int

const (
	CursorPassthrough CursorMode = iota
	CursorModeMove
	CursorModeResize
)

type Toplevel struct {
	link              wall.List
	server            *Server
	xdgToplevel       wall.XDGToplevel
	sceneTree         wall.SceneTree
	mapped            wall.Listener
	unmap             wall.Listener
	commit            wall.Listener
	destroy           wall.Listener
	requestMove       wall.Listener
	requestResize     wall.Listener
	requestMaximize   wall.Listener
	requestFullscreen wall.Listener
}

type Server struct {
	display     wall.Display
	backend     wall.Backend
	renderer    wall.Renderer
	allocator   wall.Allocator
	scene       wall.Scene
	sceneLayout wall.SceneOutputLayout

	xdgShell       wall.XDGShell
	newXdgToplevel wall.Listener
	newXdgPopup    wall.Listener
	// could this be a shim to containers/list or a dynamic array when possible?
	toplevels wall.List

	cursor               wall.Cursor
	cursorManager        wall.XCursorManager
	cursorMotion         wall.Listener
	cursorMotionAbsolute wall.Listener
	cursorButton         wall.Listener
	cursorAxis           wall.Listener
	cursorFrame          wall.Listener

	seat                wall.Seat
	newInput            wall.Listener
	requestCursor       wall.Listener
	pointerFocusChange  wall.Listener
	requestSetSelection wall.Listener
	keyboards           wall.List
	cursorMode          CursorMode
	grabbedToplevel     *Toplevel
	grabX, grabY        float64
	grabGeobox          wall.Box
	resizeEdges         uint32

	outputLayout wall.OutputLayout
	// ditto shim idea
	outputs   wall.List
	newOutput wall.Listener
}

func NewServer() (s *Server, e error) {
	s = new(Server)

	s.display = wall.NewDisplay()
	s.backend = s.display.GetEventLoop().BackendAutocreate(nil)

	return
}

func main() {
	server, err := NewServer()
	if err != nil {
		panic(fmt.Sprintf("failed to create server: %s", err))
	}

	fmt.Printf("%p", server)
}
