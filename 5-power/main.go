package main

import "strings"
import "gopkg.in/distil.v1"
import "os"
import "fmt"

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

	path := os.Getenv("LOC")

	if ds.StreamFromPath(path + "/C1ANG") != nil {
		instance1 := &RealReactivePowerDistiller{}
		registration1 := &distil.Registration{
			Instance:   instance1,
			UniqueName: "pq1_" + strings.Replace(path, "/", "_", -1),
			InputPaths: []string{path + "/C1ANG", path + "/L1ANG", path + "/C1MAG", path + "/L1MAG"},
			OutputPaths: []string{path + "/L1P", path + "/L1Q"},
		}
		ds.RegisterDistillate(registration1)
	}
	if ds.StreamFromPath(path + "/C2ANG") != nil {
		instance2 := &RealReactivePowerDistiller{}
		registration2 := &distil.Registration{
			Instance:   instance2,
			UniqueName: "pq2_" + strings.Replace(path, "/", "_", -1),
			InputPaths: []string{path + "/C2ANG", path + "/L2ANG", path + "/C2MAG", path + "/L2MAG"},
			OutputPaths: []string{path + "/L2P", path + "/L2Q"},
		}
		ds.RegisterDistillate(registration2)
	}
	if ds.StreamFromPath(path + "/C3ANG") != nil {
		instance3 := &RealReactivePowerDistiller{}
		registration3 := &distil.Registration{
			Instance:   instance3,
			UniqueName: "pq3_" + strings.Replace(path, "/", "_", -1),
			InputPaths: []string{path + "/C3ANG", path + "/L3ANG", path + "/C3MAG", path + "/L3MAG"},
			OutputPaths: []string{path + "/L3P", path + "/L3Q"},
		}
		ds.RegisterDistillate(registration3)
	}


	ds.StartEngine()
}
