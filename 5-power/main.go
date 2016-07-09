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
		instance := &RealReactivePowerDistiller{}

		if ds.StreamFromPath(path + "/C1ANG") != nil {
			registration1 := &distil.Registration{
				Instance:   instance,
				UniqueName: "pq1_" + strings.Replace(trimPath, "/", "_", -1),
				InputPaths: []string{path + "/C1ANG", path + "/L1ANG", path + "/C1MAG", path + "/L1MAG"},
				OutputPaths: []string{"/demo/" + trimPath + "/L1REAL", "/demo/" + trimPath + "/L1REACTIVE"},
			}
			ds.RegisterDistillate(registration1)
		}
		if ds.StreamFromPath(path + "/C2ANG") != nil {
			registration2 := &distil.Registration{
				Instance:   instance,
				UniqueName: "pq2_" + strings.Replace(trimPath, "/", "_", -1),
				InputPaths: []string{path + "/C2ANG", path + "/L2ANG", path + "/C2MAG", path + "/L2MAG"},
				OutputPaths: []string{"/demo/" + trimPath + "/L2REAL", "/demo/" + trimPath + "/L2REACTIVE"},
			}
			ds.RegisterDistillate(registration2)
		}
		if ds.StreamFromPath(path + "/C3ANG") != nil {
			registration3 := &distil.Registration{
				Instance:   instance,
				UniqueName: "pq3_" + strings.Replace(trimPath, "/", "_", -1),
				InputPaths: []string{path + "/C3ANG", path + "/L3ANG", path + "/C3MAG", path + "/L3MAG"},
				OutputPaths: []string{"/demo/" + trimPath + "/L3REAL", "/demo/" + trimPath + "/L3REACTIVE"},
			}
			ds.RegisterDistillate(registration3)
		}
	}

	ds.StartEngine()
}
