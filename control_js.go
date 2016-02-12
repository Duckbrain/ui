// +build js

package ui

import "github.com/gopherjs/gopherjs/js"

var elements = make(map[uintptr]*element)

// Control represents a GUI control. It provdes methods
// common to all Controls.
//
// To create a new Control, implement the control on
// the libui side, then provide access to that control on
// the Go side via an implementation of Control as
// described.
type Control interface {
	// Destroy destroys the Control.
	//
	// Implementations should do any necessary cleanup,
	// then call LibuiControlDestroy.
	Destroy()

	// LibuiControl returns the libui uiControl pointer that backs
	// the Control. This is only used by package ui itself and should
	// not be called by programs.
	LibuiControl() uintptr

	// Handle returns the OS-level handle that backs the
	// Control. On OSs that use reference counting for
	// controls, Handle does not increment the reference
	// count; you are sharing package ui's reference.
	//
	// Implementations should call LibuiControlHandle and
	// document exactly what kind of handle is returned.
	Handle() uintptr

	// Show shows the Control.
	//
	// Implementations should call LibuiControlShow.
	Show()

	// Hide shows the Control. Hidden controls do not participate
	// in layout (that is, Box, Grid, etc. does not reserve space for
	// hidden controls).
	//
	// Implementations should call LibuiControlHide.
	Hide()

	// Enable enables the Control.
	//
	// Implementations should call LibuiControlEnable.
	Enable()

	// Disable disables the Control.
	//
	// Implementations should call LibuiControlDisable.
	Disable()
}

type element struct {
	e *js.Object
}

// Destroy destroys the element.
func (e *element) Destroy() {
	// TODO Call removeChild for element
	// e.e.Get("parentElement").Call
	e.e = nil
}

// LibuiControl returns the libui uiControl pointer that backs
// the Window. This is only used by package ui itself and should
// not be called by programs.
func (e *element) LibuiControl() uintptr {
	elements[e.e.Unsafe()] = e
	return e.e.Unsafe()
}

// Handle returns the OS-level handle associated with this element.
// On Windows this is an HWND of a standard Windows API EDIT
// class (as provided by Common Controls version 6).
// On GTK+ this is a pointer to a GtkEntry.
// On OS X this is a pointer to a NSTextField.
func (e *element) Handle() uintptr {
	return e.e.Unsafe()
}

// Show shows the Entry.
func (e *element) Show() {
	e.e.Set("hidden", false)
}

// Hide hides the Entry.
func (e *element) Hide() {
	e.e.Set("hidden", true)
}

// Enable enables the Entry.
func (e *element) Enable() {
	e.e.Set("disabled", false)
}

// Disable disables the Entry.
func (e *element) Disable() {
	e.e.Set("disabled", true)
}
