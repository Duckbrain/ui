// +build js

package ui

import "github.com/gopherjs/gopherjs/js"

// Label is a Control that represents a line of text that cannot be
// interacted with. TODO rest of documentation.
type Label struct {
	element
}

// NewButton creates a new Label with the given text as its label.
func NewLabel(text string) *Label {
	l := new(Label)
	l.e = js.Global.Get("document").Call("createElement", "label")
	l.SetText(text)
	return l
}

// Text returns the Label's text.
func (l *Label) Text() string {
	return l.e.Get("textContent").String()
}

// SetText sets the Label's text to text.
func (l *Label) SetText(text string) {
	l.e.Set("textContent", text)
}
