package main

import (
	"syscall/js"
)

// ----------------------------------------------------------------------------
func setupAnimation() {
	startAnimation()
}

// ----------------------------------------------------------------------------
func render() {
}

// ----------------------------------------------------------------------------
// Allows JS to call into Wasm to refresh the effect.
func setRefreshEffectCallback() {
	js.Global().Set("refreshEffect", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		render()
		return nil
	}))
}
