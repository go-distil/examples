package main

import (
	"strings"
	"os"
	"fmt"
	"gopkg.in/distil.v1"
)

// This example extends #1 to add a few new things:
//	- loading DB params from the environment variables
//  - a more complex distillate that requires lead samples
//  - automatic generation of algorithm instances

func main() {
	BTrDB := os.Getenv("BTRDB_ADDR")
	Mongo := os.Getenv("MONGO_ADDR")
	if BTrDB == "" || Mongo == "" {
		fmt.Println("Missing BTRDB_ADDR or MONGO_ADDR")
		os.Exit(1)
	}

	// Use default connection params, this makes the resulting executable
	// portable to different installations
	ds := distil.NewDISTIL(BTrDB, Mongo)

	// Clearly you could have more advanced logic here, but this serves as
	// a good example. Register a frequency distillate for L1ANG of
	// every PMU that has a nonempty L1MAG stream.

	path := os.Getenv("REF_PMU_PATH")
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
		OutputPaths: []string{path + "/freq_1s", path + "/freq_c37",},
	})

	ds.StartEngine()
}
