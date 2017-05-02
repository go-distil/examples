package main

import (
	"fmt"
	"strings"
)
import "gopkg.in/distil.v4"
import "os"

// This example extends #1 to add a few new things:
//  - loading DB params from the environment variables
//  - a more complex distillate that requires lead samples
//  - automatic generation of algorithm instances

func main() {
	// Use default connection params, this makes the resulting executable
	// portable to different installations
	ds := distil.NewDISTIL()

	// The path leading to the PMU, excluding the final element
	path := os.Getenv("LOC")
	// The order of the phases on the PMU e.g "123"
	order := os.Getenv("LOC_ORDER")

	// The reference PMU to compare to
	refPath := os.Getenv("REF")
	// The order of the phases on the reference.
	refOrder := os.Getenv("REF_ORDER")

	if path == "" || order == "" || refPath == "" || refOrder == "" {
		fmt.Println("You need to specify $LOC, $LOC_ORDER, $REF and $REF_ORDER")
		os.Exit(1)
	}

	// For example if LOC_ORDER = "123" and REF_ORDER = "132" then phase 2 and 3 are swapped
	for i := 0; i < len(order); i++ {
		instance := &AngleDifferenceDistiller{basefreq: 120}
		registration := &distil.Registration{
			Instance:    instance,
			UniqueName:  "anglediff_L" + string(order[i]) + "_" + strings.Replace(path, "/", "_", -1) + "_L" + string(refOrder[i]) + "_" + strings.Replace(refPath, "/", "_", -1),
			InputPaths:  []string{path + "/L" + string(order[i]) + "ANG", refPath + "/L" + string(refOrder[i]) + "ANG"},
			OutputPaths: []string{path + "/anglediff_L" + string(order[i]) + "_rel_" + strings.Replace(refPath, "/", "_", -1) + "_L" + string(refOrder[i])},
			OutputUnits: []string{"degrees"},
		}
		ds.RegisterDistillate(registration)
	}
	ds.StartEngine()
}
