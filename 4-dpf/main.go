package main

import (
	"fmt"
	"strings"
)
import "gopkg.in/distil.v1"
import "os"

func main() {
	// Use default connection params, this makes the resulting executable
	// portable to different installations
	ds := distil.NewDISTIL(distil.FromEnvVars())

	path := os.Getenv("LOC")
	if path == "" {
		fmt.Println("You need the $LOC variable")
		os.Exit(1)
	}

	if ds.StreamFromPath(path+"/C1ANG") != nil {
		instance1 := &DisplacementPFDistiller{basefreq: 120}
		registration1 := &distil.Registration{
			Instance:    instance1,
			UniqueName:  "dpf1_" + strings.Replace(path, "/", "_", -1),
			InputPaths:  []string{path + "/C1ANG", path + "/L1ANG"},
			OutputPaths: []string{path + "/L1DPF"},
		}
		ds.RegisterDistillate(registration1)
	}
	if ds.StreamFromPath(path+"/C2ANG") != nil {
		instance2 := &DisplacementPFDistiller{basefreq: 120}
		registration2 := &distil.Registration{
			Instance:    instance2,
			UniqueName:  "dpf2_" + strings.Replace(path, "/", "_", -1),
			InputPaths:  []string{path + "/C2ANG", path + "/L2ANG"},
			OutputPaths: []string{path + "/L2DPF"},
		}
		ds.RegisterDistillate(registration2)
	}
	if ds.StreamFromPath(path+"/C3ANG") != nil {
		instance3 := &DisplacementPFDistiller{basefreq: 120}
		registration3 := &distil.Registration{
			Instance:    instance3,
			UniqueName:  "dpf3_" + strings.Replace(path, "/", "_", -1),
			InputPaths:  []string{path + "/C3ANG", path + "/L3ANG"},
			OutputPaths: []string{path + "/L3DPF"},
		}
		ds.RegisterDistillate(registration3)
	}

	ds.StartEngine()
}
