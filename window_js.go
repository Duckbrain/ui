// +build js

package ui

import (
	"fmt"

	"github.com/gopherjs/gopherjs/js"
)

var windows = make([]*Window, 0)

// Window is a Control that represents a top-level window.
// A Window contains one child Control that occupies the
// entirety of the window. Though a Window is a Control,
// a Window cannot be the child of another Control.
type Window struct {
	element
	body         *js.Object
	child        Control
	onClosing    func(w *Window) bool
	dragX, dragY int64
	posX, posY   int64
}

// NewWindow creates a new Window.
func NewWindow(title string, width int, height int, hasMenubar bool) *Window {
	w := new(Window)

	w.body = js.Global.Get("document").Call("createElement", "div")
	w.e = js.Global.Get("document").Call("createElement", "div")
	titleBar := js.Global.Get("document").Call("createElement", "div")

	w.body.Get("style").Set("width", width)
	w.e.Get("style").Set("height", height)
	w.e.Get("classList").Call("add", "window-body")
	w.body.Get("classList").Call("add", "window")
	w.body.Call("appendChild", titleBar)
	w.body.Call("appendChild", w.e)

	titleBar.Get("classList").Call("add", "window-title")
	titleBar.Call("addEventListener", "mousedown", func(e *js.Object) {
		w.dragX = e.Get("clientX").Int64() - w.posX
		w.dragY = e.Get("clientY").Int64() - w.posY
		js.Global.Call("addEventListener", "mousemove", w.drag, true)
	})
	js.Global.Call("addEventListener", "mouseup", func(e *js.Object) {
		js.Global.Call("removeEventListener", "mousemove", w.drag, true)
	})

	js.Global.Set("onbeforeunload", func() interface{} {
		for _, win := range windows {
			if win.onClosing != nil {
				if !win.onClosing(win) {
					return "Are you sure you want to close all windows?"
				}
			}
		}
		return nil
	})

	js.Global.Get("document").Get("body").Call("appendChild", w.body)

	return w
}

// Title returns the Window's title.
func (w *Window) Title() string {
	return w.e.Get("title").String()
}

// SetTitle sets the Window's title to title.
func (w *Window) SetTitle(title string) {
	w.e.Set("title", title)
}

// OnClosing registers f to be run when the user clicks the Window's
// close button. Only one function can be registered at a time.
// If f returns true, the window is destroyed with the Destroy method.
// If f returns false, or if OnClosing is never called, the window is not
// destroyed and is kept visible.
func (w *Window) OnClosing(f func(*Window) bool) {
	w.onClosing = f
}

// SetChild sets the Window's child to child. If child is nil, the Window
// will not have a child.
func (w *Window) SetChild(child Control) {
	if w.child != nil {
		w.e.Call("removeChild", elements[w.child.LibuiControl()])
	}
	w.child = child
	w.e.Call("appendChild", elements[child.LibuiControl()])
}

// Margined returns whether the Window has margins around its child.
func (w *Window) Margined() bool {
	return w.e.Get("classList").Call("contains", "window-margined").Bool()
}

// SetMargined controls whether the Window has margins around its
// child. The size of the margins are determined by the OS and its
// best practices.
func (w *Window) SetMargined(margined bool) {
	w.e.Get("classList").Call("toggle", "window-margined", margined)
}

func (w *Window) drag(e *js.Object) {
	w.posX = e.Get("clientX").Int64() - w.dragX
	w.posY = e.Get("clientY").Int64() - w.dragY
	w.body.Get("style").Set("left", fmt.Sprintf("%vpx", w.posX))
	w.body.Get("style").Set("top", fmt.Sprintf("%vpx", w.posY))
}
