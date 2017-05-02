package main

import "gopkg.in/distil.v4"

func main() {
	// Get a handle to BTrDB. go-distil is implemented as a library
	// so there is no other distillate service to connect to
	ds := distil.NewDISTIL()

	// Construct an instance of your distillate. If you had parameters for
	// the distillate you would maybe have a custom constructor. You could
	// also load the parameters from a config file, or some heuristic
	// algorithm, which we will show in the next few examples
	instance := &NopDistiller{}

	// Now we add this distillate to the DISTIL engine. If you add multiple
	// distillates, they will all get computed in parallel.
	ds.RegisterDistillate(&distil.Registration{
		// The class that implements your algorithm
		Instance: instance,
		// A unique name FOR THIS INSTANCE of the distillate. If you
		// are autogenerating distillates, take care to never produce
		// the same name here. We would normally use a UUID but opted
		// for this so as to be more human friendly. If the program
		// is restarted, this is how it knows where to pick up from.
		UniqueName: "demo_noop_distillate",
		// These are inputs to the distillate that will be loaded
		// and presented to Process()
		InputPaths: []string{"upmu/rpu/hunter_1224/L1MAG"},
		// These are the output paths for the distillate. They must
		// also be strictly unique.
		OutputPaths: []string{"distil/rpu/hunter_1224/L1MAG"},
		// These are the units for the output paths
		OutputUnits: []string{"degrees"},
	})

	//Now we tell the DISTIL library to keep all the registered distillates
	//up to date. The program will not exit.
	ds.StartEngine()
}
