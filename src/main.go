package main

import (
	"github.com/ewaldhorn/dommie/dom"
	"github.com/ewaldhorn/tinycanvas/tinycanvas"
)

const (
	canvasWidth  = 150
	canvasHeight = 150
)

var (
	mainCanvas *tinycanvas.TinyCanvas
	ctx        *tinycanvas.Context2D
)

// ----------------------------------------------------------------------------
// bootstrap is a JavaScript-side defined function, called by Wasm in the main
// Go function
//
//export bootstrapApp
func bootstrapApp()

//export startAnimation
func startAnimation()

// ----------------------------------------------------------------------------
func main() {
	startup()

	// ready, now create the canvas etc.
	createMainCanvas()
	setupAnimation()

	// prevent the app for closing before the page does
	ch := make(chan struct{})
	<-ch

	// all done!
	dom.Log("All done.")
}

// ----------------------------------------------------------------------------
// sets up the initial stuff
func startup() {
	dom.Log("Starting TinyClock...")
	setCallbacks()
	dom.Hide("loading")
	bootstrapApp()
}

// ----------------------------------------------------------------------------
// create the main canvas we'll draw on, also get the graphics context
func createMainCanvas() {
	mainCanvas = tinycanvas.NewTinyCanvas(canvasWidth, canvasHeight)
	ctx = mainCanvas.GetContext()
}

// ----------------------------------------------------------------------------
func setCallbacks() {
	setVersionCallback()
	setRefreshEffectCallback()
}
