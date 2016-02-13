// +build js

package ui

import "github.com/gopherjs/gopherjs/js"

// Main initializes package ui, runs f to set up the program,
// and executes the GUI main loop. f should set up the program's
// initial state: open the main window, create controls, and set up
// events. It should then return, at which point Main will
// process events until Quit is called, at which point Main will return
// nil. If package ui fails to initialize, Main returns an appropriate
// error.
func Main(f func()) error {
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

	js.Global.Get("document").Call("addEventListener", "DOMContentLoaded", f)
	return nil
}

// Quit queues a return from Main. It does not exit the program.
// It also does not immediately cause Main to return; Main will
// return when it next can. Quit must be called from the GUI thread.
func Quit() {
	js.Global.Call("close")
}

// QueueMain queues f to be executed on the GUI thread when
// next possible. It returns immediately; that is, it does not wait
// for the function to actually be executed. QueueMain is the only
// function that can be called from other goroutines, and its
// primary purpose is to allow communication between other
// goroutines and the GUI thread. Calling QueueMain after Quit
// has been called results in undefined behavior.
func QueueMain(f func()) {
	go f()
}

// no need to lock this; this API is only safe on the main thread
var shouldQuitFunc func() bool

// OnShouldQuit schedules f to be exeucted when the OS wants
// the program to quit or when a Quit menu item has been clicked.
// Only one function may be registered at a time. If the function
// returns true, Quit will be called. If the function returns false, or
// if OnShouldQuit is never called. Quit will not be called and the
// OS will be told that the program needs to continue running.
func OnShouldQuit(f func() bool) {
	js.Global.Set("onbeforeunload", func() interface{} {
		if f() {
			return nil
		} else {
			return "Are you sure you want to leave?"
		}
	})
}
