package main

import (
	"strings"

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
	for _, path := range ds.ListExistingUpmuPaths() {
		trimPath := strings.TrimPrefix(path, "/upmu/")
		instance := &FrequencyDistiller{basefreq: 120}
		registration := &distil.Registration{
			Instance:   instance,
			UniqueName: "freq_" + strings.Replace(trimPath, "/", "_", -1),
			InputPaths: []string{path + "/L1ANG"},
			OutputPaths: []string{"/demo/" + trimPath + "/freq_1s",
				"/demo/" + trimPath + "/freq_c37"},
		}
		ds.RegisterDistillate(registration)
	}

	ds.StartEngine()
}
