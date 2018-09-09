package main

import (
	"math"

	"gopkg.in/distil.v4"
)

//This is our distillate algorithm
type FrequencyDistiller struct {
	// This line is required. It says this struct inherits some useful
	// default methods.
	distil.DistillateTools

	// For the frequency distillate, we make use of a rebase stage, to do
	// that, we need to know the intended frequency of the stream
	basefreq int64
}

// We will now show off all the tunables that a distillate MAY implement if
// it requires them. The first is "lead time". Some algorithms require some
// data outside of the changed range in order to compute. Frequency is one
// of those types of algorithms. By implementing LeadNanos() you can tell
// the DISTIL engine to load some extra data for you, which will be available
// at negative indices in the Process() method. The default implementation
// (in DistillateTools) returns 0
func (d *FrequencyDistiller) LeadNanos() int64 {
	return 1000000000
}

// PadSnap is a rebase stage that will adjust the incoming data to strictly
// appear on a timebase of the given frequency (hence the 'rebase'). Any
// values that do not appear with exactly the right time are snapped to
// the nearest time, and any duplicates are dropped. In addition, any missing
// values are replaced by NaN (hence pad). The advantage of this is that it
// simplifies calculations that refer to values across time, you can rest
// assured that a value 1s ago is exactly basefreq samples away, even if
// there were holes in the data or if there were duplicates. Note that in
// the presence of duplicate data there is ZERO GUARANTEE as to WHICH of the
// multiple duplicate values you receive. In general this makes algorithms
// that compare across time quite useless, as the real time difference
// between the points a fixed interval apart will experience extreme jitter.
// The default implementation (in DistillateTools) returns RebasePassthrough
func (d *FrequencyDistiller) Rebase() distil.Rebaser {
	return distil.RebasePadSnap(d.basefreq)
}

// A little utility function used by the algorithm
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
func (d *FrequencyDistiller) Process(in *distil.InputSet, out *distil.OutputSet) {
	/* Output 0 is freq_1s.
	 * Output 1 is freq_c37.
	 */
	var ns int = in.NumSamples(0)
	var i int
	for i = 0; i < ns; i++ {
		var time int64 = in.Get(0, i).T

		var v1s float64 = angwrap(in.Get(0, i).V-in.Get(0, i-120).V)/360.0 + 60.0
		if !math.IsNaN(v1s) {
			out.Add(0, time, v1s)
		}

		var p1 = in.Get(0, i).V
		var p2 = in.Get(0, i-1).V
		var p3 = in.Get(0, i-2).V
		var p4 = in.Get(0, i-3).V
		var v1 = angwrap(p1 - p2)
		var v2 = angwrap(p2 - p3)
		var v3 = angwrap(p3 - p4)
		var c37 = 60.0 + (((6.0*(v1) + 3.0*(v2) + 1.0*(v3)) / 10) * (120.0 / 360.0))
		if !math.IsNaN(c37) {
			out.Add(1, time, c37)
		}
	}
}
