// +build js

package ui

func ensureMainThread() {
	// do nothing; Browsers don't care which thread we're on because there is always only one
}
