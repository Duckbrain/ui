// +build js

package ui

import "github.com/gopherjs/gopherjs/js"

// Box is a Control that holds a group of Controls horizontally
// or vertically. If horizontally, then all controls have the same
// height. If vertically, then all controls have the same width.
// By default, each control has its preferred width (horizontal)
// or height (vertical); if a control is marked "stretchy", it will
// take whatever space is left over. If multiple controls are marked
// stretchy, they will be given equal shares of the leftover space.
// There can also be space between each control ("padding").
type Box struct {
	element
	children []Control
}

// NewHorizontalBox creates a new horizontal Box.
func NewHorizontalBox() *Box {
	b := new(Box)
	b.e = js.Global.Get("document").Call("createElement", "div")
	b.e.Get("classList").Call("add", "box", "box-horizontal")
	return b
}

// NewVerticalBox creates a new vertical Box.
func NewVerticalBox() *Box {
	b := new(Box)
	b.e = js.Global.Get("document").Call("createElement", "div")
	b.e.Get("classList").Call("add", "box", "box-vertical")
	return b
}

// Append adds the given control to the end of the Box.
func (b *Box) Append(child Control, stretchy bool) {
	b.e.Call("appendChild", elements[child.LibuiControl()])
	b.children = append(b.children, child)
}

// Delete deletes the nth control of the Box.
func (b *Box) Delete(n int) {
	b.children = append(b.children[:n], b.children[n+1:]...)
	b.e.Call("removeChild", b.e.Get("childNodes").Index(n))
}

// Padded returns whether there is space between each control
// of the Box.
func (b *Box) Padded() bool {
	return b.e.Get("classList").Call("contains", "box-padded").Bool()
}

// SetPadded controls whether there is space between each control
// of the Box. The size of the padding is determined by the OS and
// its best practices.
func (b *Box) SetPadded(padded bool) {
	b.e.Get("classList").Call("toggle", "box-padded", padded)
}
