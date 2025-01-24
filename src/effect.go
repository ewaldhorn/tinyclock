package main

import (
	"math"
	"syscall/js"
	"time"
)

// ----------------------------------------------------------------------------
func setupAnimation() {
	startAnimation()
}

// ----------------------------------------------------------------------------
func render() {
	ctx.Save() //a
	ctx.ClearRect(0, 0, canvasWidth, canvasHeight)
	ctx.Translate(75, 75)
	ctx.Call("scale", 0.4, 0.4)
	ctx.Rotate(-math.Pi / 2)
	ctx.StrokeStyle("black")
	ctx.FillStyle("white")
	ctx.LineWidth(8)
	ctx.Set("lineCap", "round")

	// hour marks
	ctx.Save() //b
	for range 12 {
		ctx.BeginPath()
		ctx.Rotate(math.Pi / 6)
		ctx.Call("moveTo", 100, 0)
		ctx.LineTo(120, 0)
		ctx.Stroke()
	}
	ctx.Restore() //b

	// minutes
	ctx.Save() //c
	ctx.LineWidth(5)
	for i := range 60 {
		if i%5 != 0 {
			ctx.BeginPath()
			ctx.Call("moveTo", 117, 0)
			ctx.LineTo(120, 0)
			ctx.Stroke()
		}
		ctx.Rotate(math.Pi / 30)
	}
	ctx.Restore() //c

	now := time.Now()
	sec := float64(now.Second())
	min := float64(now.Minute())
	hour := float64(now.Hour() % 12)

	ctx.FillStyle("black")

	// TODO support
	// canvas.innerText = `The time is: ${hr}:${min}`;

	// Write Hours
	ctx.Save() //d
	ctx.Rotate(math.Pi/6.0*hour + (math.Pi/360.0)*min + (math.Pi/21600)*sec)
	ctx.LineWidth(14)
	ctx.BeginPath()
	ctx.Call("moveTo", -20, 0)
	ctx.LineTo(80, 0)
	ctx.Stroke()
	ctx.Restore() //d

	// Minutes
	// Write Minutes
	ctx.Save() //e
	ctx.Rotate((math.Pi/30)*min + (math.Pi/1800)*sec)
	ctx.LineWidth(10)
	ctx.BeginPath()
	ctx.Call("moveTo", -28, 0)
	ctx.LineTo(112, 0)
	ctx.Stroke()
	ctx.Restore() //e

	// Seconds
	ctx.Save() //f
	ctx.Rotate((sec * math.Pi) / 30)
	ctx.StrokeStyle("#D40000")
	ctx.FillStyle("#D40000")
	ctx.LineWidth(6)
	ctx.BeginPath()
	ctx.Call("moveTo", -30, 0)
	ctx.LineTo(83, 0)
	ctx.Stroke()
	ctx.BeginPath()
	ctx.Arc(0, 0, 10, 0, math.Pi*2, true)
	ctx.Fill()
	ctx.BeginPath()
	ctx.Arc(95, 0, 10, 0, math.Pi*2, true)
	ctx.Stroke()
	ctx.FillStyle("rgb(0 0 0 / 0%)")
	ctx.Arc(0, 0, 3, 0, math.Pi*2, true)
	ctx.Fill()
	ctx.Restore() //f

	ctx.BeginPath()
	ctx.LineWidth(14)
	ctx.StrokeStyle("#325FA2")
	ctx.Arc(0, 0, 142, 0, math.Pi*2, true)
	ctx.Stroke()

	ctx.Restore() //a
}

// ----------------------------------------------------------------------------
// Allows JS to call into Wasm to refresh the effect.
func setRefreshEffectCallback() {
	js.Global().Set("refreshEffect", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		render()
		return nil
	}))
}
