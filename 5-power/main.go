package main

import "strings"
import "gopkg.in/distil.v1"
import "os"

func main() {
	// Use default connection params, this makes the resulting executable
	// portable to different installations
	ds := distil.NewDISTIL(distil.FromEnvVars())

	path := os.Getenv("LOC")

	if ds.StreamFromPath(path+"/C1ANG") != nil {
		instance1 := &RealReactivePowerDistiller{basefreq: 120}
		registration1 := &distil.Registration{
			Instance:    instance1,
			UniqueName:  "pq1_" + strings.Replace(path, "/", "_", -1),
			InputPaths:  []string{path + "/C1ANG", path + "/L1ANG", path + "/C1MAG", path + "/L1MAG"},
			OutputPaths: []string{path + "/L1P", path + "/L1Q"},
		}
		ds.RegisterDistillate(registration1)
	}
	if ds.StreamFromPath(path+"/C2ANG") != nil {
		instance2 := &RealReactivePowerDistiller{basefreq: 120}
		registration2 := &distil.Registration{
			Instance:    instance2,
			UniqueName:  "pq2_" + strings.Replace(path, "/", "_", -1),
			InputPaths:  []string{path + "/C2ANG", path + "/L2ANG", path + "/C2MAG", path + "/L2MAG"},
			OutputPaths: []string{path + "/L2P", path + "/L2Q"},
		}
		ds.RegisterDistillate(registration2)
	}
	if ds.StreamFromPath(path+"/C3ANG") != nil {
		instance3 := &RealReactivePowerDistiller{basefreq: 120}
		registration3 := &distil.Registration{
			Instance:    instance3,
			UniqueName:  "pq3_" + strings.Replace(path, "/", "_", -1),
			InputPaths:  []string{path + "/C3ANG", path + "/L3ANG", path + "/C3MAG", path + "/L3MAG"},
			OutputPaths: []string{path + "/L3P", path + "/L3Q"},
		}
		ds.RegisterDistillate(registration3)
	}

	ds.StartEngine()
}
