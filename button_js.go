// +build js

package ui

import "github.com/gopherjs/gopherjs/js"

// Button is a Control that represents a button that the user can
// click to perform an action. A Button has a text label that should
// describe what the button does.
type Button struct {
	element
	onClicked func(*Button)
}

// NewButton creates a new Button with the given text as its label.
func NewButton(text string) *Button {
	b := new(Button)
	b.e = js.Global.Get("document").Call("createElement", "button")
	b.SetText(text)
	b.e.Call("addEventListener", "click", func() {
		b.onClicked(b)
	})
	return b
}

// Text returns the Button's text.
func (b *Button) Text() string {
	return b.e.Get("textContent").String()
}

// SetText sets the Button's text to text.
func (b *Button) SetText(text string) {
	b.e.Set("textContent", text)
}

// OnClicked registers f to be run when the user clicks the Button.
// Only one function can be registered at a time.
func (b *Button) OnClicked(f func(*Button)) {
	b.onClicked = f
}
