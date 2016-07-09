package main

import (
	"math"

	"gopkg.in/distil.v1"
)

//This is our distillate algorithm
type AngleDifferenceDistiller struct {
	// This line is required. It says this struct inherits some useful
	// default methods.
	distil.DistillateTools
}

func angwrap(d float64) float64 {
	if d > 180 {
		return d - 360
	} else if d < -180 {
		return d + 360
	} else {
		return d
	}
}

// This is our main algorithm. It will automatically be called with chunks
// of data that require processing by the engine.
func (d *AngleDifferenceDistiller) Process(in *distil.InputSet, out *distil.OutputSet) {
	// Output 0 is angdiff.
	var ns int = in.NumSamples(0)
	var i int
	for i = 0; i < ns; i++ {
		var time int64 = in.Get(0, i).T
		var ang1 = in.Get(0, i).V
		var ang2 = in.Get(1, i).V
		var angdiff = angwrap(ang1 - ang2)

		if !math.IsNaN(angdiff) {
			out.Add(0, time, angdiff)
		}
	}
}
