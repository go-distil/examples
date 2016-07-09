package main

import (
	"math"

	"gopkg.in/distil.v1"
)

//This is our distillate algorithm
type RealReactivePowerDistiller struct {
	// This line is required. It says this struct inherits some useful
	// default methods.
	distil.DistillateTools

}

// This is our main algorithm. It will automatically be called with chunks
// of data that require processing by the engine.
func (d *RealReactivePowerDistiller) Process(in *distil.InputSet, out *distil.OutputSet) {
	/* Output 0 is real_power.
	 * Output 1 is reactive_power.
  */
	var ns int = in.NumSamples(0)
	var i int
	for i = 0; i < ns; i++ {
		var time int64 = in.Get(0, i).T

		var angI = in.Get(0, i).V
		var angV = in.Get(1, i).V
		var magI = in.Get(2, i).V
		var magV = in.Get(3, i).V

		var p = magV*magI*math.Cos((angV * math.Pi / 180)-(angI * math.Pi / 180))
		var q = magV*magI*math.Sin((angV * math.Pi / 180)-(angI * math.Pi / 180))

		if !math.IsNaN(p) {
			out.Add(0, time, p)
		}

		if !math.IsNaN(q) {
			out.Add(1, time, q)
		}

	}
}
