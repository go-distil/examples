package main

import (
	"strings"
	"os"
	"gopkg.in/distil.v1"
)

// This example extends #1 to add a few new things:
//	- loading DB params from the environment variables
//  - a more complex distillate that requires lead samples
//  - automatic generation of algorithm instances

func main() {
	// Use default connection params, this makes the resulting executable
	// portable to different installations
	ds := distil.NewDISTIL(distil.FromEnvVars())

	// Clearly you could have more advanced logic here, but this serves as
	// a good example. Register a frequency distillate for L1ANG of
	// every PMU that has a nonempty L1MAG stream.

	// TODO: read the ref location for each stream
	for _, path := range ds.ListExistingUpmuPaths() {
		var pathref = path
		trimPath := strings.TrimPrefix(path, "/upmu/")
		instance := &AngleDifferenceDistiller{}
		registration := &distil.Registration{
			Instance:   instance,
			UniqueName: "anglediff_" + strings.Replace(trimPath, "/", "_", -1),
			InputPaths: []string{path + "/L1ANG", pathref + "/L1ANG"},
			OutputPaths: []string{"/demo/" + trimPath + "/L1DIFF"},
		}
		ds.RegisterDistillate(registration)
	}

	ds.StartEngine()
}
