package main

import (
	"strings"

	"github.com/go-distil/distil"
)

const BTrDB = "btrdb.bts.sensor.guru:4410"
const Mongo = "btrdb.bts.sensor.guru"

func main() {
	// This example extends #1 to add two new things:
	//  - a more complex distillate that requires lead samples
	//  - automatic generation of algorithm instances

	ds := distil.NewDISTIL(BTrDB, Mongo)

	// Clearly you could have better logic here, but this serves as
	// a good example. Register a frequency distillate for L1ANG of
	// every PMU
	for _, path := range ds.ListUpmuPaths()[:2] {
		trimPath := strings.TrimPrefix(path, "/upmu/")
		instance := &FrequencyDistiller{basefreq: 120}
		registration := &distil.Registration{
			Instance:    instance,
			UniqueName:  "freq_" + strings.Replace(trimPath, "/", "_", -1),
			InputPaths:  []string{path + "/L1ANG"},
			OutputPaths: []string{"/demo/" + trimPath + "/freq"},
		}
		ds.RegisterDistillate(registration)
	}

	// // Construct an instance of your distillate. If you had parameters for
	// // the distillate you would maybe have a custom constructor. You could
	// // also load the parameters from a config file, or some heuristic
	// // algorithm, which we will show in the next few examples
	// instance := &NopDistiller{}
	//
	// // Now we add this distillate to the DISTIL engine. If you add multiple
	// // distillates, they will all get computed in parallel.
	// ds.RegisterDistillate(&distil.Registration{
	// 	// The class that implements your algorithm
	// 	Instance: instance,
	// 	// A unique name FOR THIS INSTANCE of the distillate. If you
	// 	// are autogenerating distillates, take care to never produce
	// 	// the same name here. We would normally use a UUID but opted
	// 	// for this so as to be more human friendly. If the program
	// 	// is restarted, this is how it knows where to pick up from.
	// 	UniqueName: "demo_noop_distillate",
	// 	// These are inputs to the distillate that will be loaded
	// 	// and presented to Process()
	// 	InputPaths: []string{"/upmu/a6_bus1/L1MAG"},
	// 	// These are the output paths for the distillate. They must
	// 	// also be strictly unique.
	// 	OutputPaths: []string{"/godistil/a6_bus1/L1MAG"},
	// })
	//
	// //Now we tell the DISTIL library to keep all the registered distillates
	// //up to date. The program will not exit.
	// ds.StartEngine()
}
