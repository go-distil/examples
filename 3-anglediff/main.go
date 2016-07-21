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

	refPath := os.Getenv("refpmupath")
	trimRefPath := strings.TrimPrefix(refPath, "/upmu/")
	sref := strings.Split("trimRefPath", "/")
	refUtilityID := sref[0]


	// Clearly you could have more advanced logic here, but this serves as
	// a good example. Register a frequency distillate for L1ANG of
	// every PMU that has a nonempty L1MAG stream.
	for _, path := range ds.ListExistingUpmuPaths() {
		if strings.Compare(path, refPath) != 0 {
			trimPath := strings.TrimPrefix(path, "/upmu/")
			s := strings.Split("trimPath", "/")
			pmuUtilityID := s[0]
			if strings.Compare(refUtilityID, pmuUtilityID) == 0 {
				if ds.StreamFromPath(path + "/L1ANG") != nil {
					instance := &AngleDifferenceDistiller{}
					registration1 := &distil.Registration{
						Instance:   instance,
						UniqueName: "anglediff1_" + strings.Replace(trimRefPath, "/", "_", -1) + "_" + strings.Replace(trimPath, "/", "_", -1),
						InputPaths: []string{path + "/L1ANG", refPath + "/L1ANG"},
						OutputPaths: []string{"/demo/" + trimPath + "/L1DIFF"},
					}
					ds.RegisterDistillate(registration1)
				}
				if ds.StreamFromPath(path + "/L2ANG") != nil {
					instance := &AngleDifferenceDistiller{}
					registration2 := &distil.Registration{
						Instance:   instance,
						UniqueName: "anglediff2_" + strings.Replace(trimRefPath, "/", "_", -1) + "_" + strings.Replace(trimPath, "/", "_", -1),
						InputPaths: []string{path + "/L2ANG", refPath + "/L2ANG"},
						OutputPaths: []string{"/demo/" + trimPath + "/L2DIFF"},
					}
					ds.RegisterDistillate(registration2)
				}
				if ds.StreamFromPath(path + "/L3ANG") != nil {
					instance := &AngleDifferenceDistiller{}
					registration3 := &distil.Registration{
						Instance:   instance,
						UniqueName: "anglediff3_" + strings.Replace(trimRefPath, "/", "_", -1) + "_" + strings.Replace(trimPath, "/", "_", -1),
						InputPaths: []string{path + "/L3ANG", refPath + "/L3ANG"},
						OutputPaths: []string{"/demo/" + trimPath + "/L3DIFF"},
					}
					ds.RegisterDistillate(registration3)
				}
			}
		}
	}

	ds.StartEngine()
}
