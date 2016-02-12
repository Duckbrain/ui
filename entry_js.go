// +build js

package ui

import "github.com/gopherjs/gopherjs/js"

// Entry is a Control that represents a space that the user can
// type a single line of text into.
type Entry struct {
	onChanged func(*Entry)
	element
}

// NewEntry creates a new Entry.
func NewEntry() *Entry {
	e := new(Entry)
	e.e = js.Global.Get("document").Call("createElement", "input")

	return e
}

// Text returns the Entry's text.
func (e *Entry) Text() string {
	return e.e.Get("value").String()
}

// SetText sets the Entry's text to text.
func (e *Entry) SetText(text string) {
	e.e.Set("value", text)
}

// OnChanged registers f to be run when the user makes a change to
// the Entry. Only one function can be registered at a time.
func (e *Entry) OnChanged(f func(*Entry)) {
	//TODO This will need to use keydown events to emulate ui
	e.onChanged = f
}

// ReadOnly returns whether the Entry can be changed.
func (e *Entry) ReadOnly() bool {
	return e.e.Get("readOnly").Bool()
}

// SetReadOnly sets whether the Entry can be changed.
func (e *Entry) SetReadOnly(ro bool) {
	e.e.Set("readOnly", ro)
}
