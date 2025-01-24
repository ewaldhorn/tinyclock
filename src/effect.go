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
	prepareClockFace()
	renderHourMarks()
	renderMinuteMarks()

	now := time.Now()
	sec := float64(now.Second())
	min := float64(now.Minute())
	hour := float64(now.Hour() % 12)

	// TODO support
	// canvas.innerText = `The time is: ${hr}:${min}`;

	ctx.FillStyle("black")
	renderHours(hour, min, sec)
	renderMinutes(min, sec)
	renderSeconds(sec)
	renderClockOutline()
}

// ----------------------------------------------------------------------------
func prepareClockFace() {
	ctx.Save() //a
	ctx.ClearRect(0, 0, canvasWidth, canvasHeight)
	ctx.Translate(75, 75)
	ctx.Scale(0.4, 0.4)
	ctx.Rotate(-math.Pi / 2)
	ctx.StrokeStyle("black")
	ctx.FillStyle("white")
	ctx.LineWidth(8)
	ctx.LineCap("round")
}

// ----------------------------------------------------------------------------
func renderHourMarks() {
	ctx.Save() //b
	for range 12 {
		ctx.BeginPath()
		ctx.Rotate(math.Pi / 6)
		ctx.MoveTo(100, 0)
		ctx.LineTo(120, 0)
		ctx.Stroke()
	}
	ctx.Restore() //b
}

// ----------------------------------------------------------------------------
func renderMinuteMarks() {
	ctx.Save() //c
	ctx.LineWidth(5)
	for i := range 60 {
		if i%5 != 0 {
			ctx.BeginPath()
			ctx.MoveTo(117, 0)
			ctx.LineTo(120, 0)
			ctx.Stroke()
		}
		ctx.Rotate(math.Pi / 30)
	}
	ctx.Restore() //c
}

// ----------------------------------------------------------------------------
func renderHours(hour, min, sec float64) {
	ctx.Save() //d
	ctx.Rotate(math.Pi/6.0*hour + (math.Pi/360.0)*min + (math.Pi/21600)*sec)
	ctx.LineWidth(14)
	ctx.BeginPath()
	ctx.MoveTo(-20, 0)
	ctx.LineTo(80, 0)
	ctx.Stroke()
	ctx.Restore() //d
}

// ----------------------------------------------------------------------------
func renderMinutes(min, sec float64) {
	ctx.Save() //e
	ctx.Rotate((math.Pi/30)*min + (math.Pi/1800)*sec)
	ctx.LineWidth(10)
	ctx.BeginPath()
	ctx.MoveTo(-28, 0)
	ctx.LineTo(112, 0)
	ctx.Stroke()
	ctx.Restore() //e
}

// ----------------------------------------------------------------------------
func renderSeconds(sec float64) {
	ctx.Save() //f
	ctx.Rotate((sec * math.Pi) / 30)
	ctx.StrokeStyle("#D40000")
	ctx.FillStyle("#D40000")
	ctx.LineWidth(6)
	ctx.BeginPath()
	ctx.MoveTo(-30, 0)
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
}

// ----------------------------------------------------------------------------
func renderClockOutline() {
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
