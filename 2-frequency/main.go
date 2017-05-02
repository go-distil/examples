package main

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/distil.v4"
)

// This example extends #1 to add a few new things:
//	- loading DB params from the environment variables
//  - a more complex distillate that requires lead samples
//  - automatic generation of algorithm instances

func main() {
	// Use default connection params, this makes the resulting executable
	// portable to different installations
	ds := distil.NewDISTIL()

	// This will register a distillate that processes a path
	// read from an environment variable
	path := os.Getenv("REF_PMU_PATH")
	if path == "" {
		fmt.Println("Missing $REF_PMU_PATH")
		os.Exit(1)
	}
	ds.RegisterDistillate(&distil.Registration{
		// The class that implements your algorithm
		Instance: &FrequencyDistiller{basefreq: 120},
		// A unique name FOR THIS INSTANCE of the distillate. If you
		// are autogenerating distillates, take care to never produce
		// the same name here. We would normally use a UUID but opted
		// for this so as to be more human friendly. If the program
		// is restarted, this is how it knows where to pick up from.
		UniqueName: "freq_" + strings.Replace(path, "/", "_", -1),
		// These are inputs to the distillate that will be loaded
		// and presented to Process()
		InputPaths: []string{path + "/L1ANG"},
		// These are the output paths for the distillate. They must
		// also be strictly unique.
		OutputPaths: []string{path + "/freq_1s", path + "/freq_c37"},
		// The units for the distillate
		OutputUnits: []string{"hz", "hz"},
	})

	ds.StartEngine()
}
