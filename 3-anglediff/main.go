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
	order := os.Getenv("LOC_ORDER")

	refPath := os.Getenv("REF")
	refOrder := os.Getenv("REF_ORDER")

	// For example if LOC_ORDER = "123" and REF_ORDER = "132" then phase 2 and 3 are swapped

	for i := 0; i < len(order); i++ {
		instance := &AngleDifferenceDistiller{basefreq: 120}
		registration := &distil.Registration{
			Instance:   instance,
			UniqueName: "anglediff_L" + string(order[i]) + "_" + strings.Replace(path, "/", "_", -1) + "_L" + string(refOrder[i]) + "_" + strings.Replace(refPath, "/", "_", -1),
			InputPaths: []string{path + "/L" + string(order[i]) + "ANG", refPath + "/L" + string(refOrder[i]) + "ANG"},
			OutputPaths: []string{path + "/anglediff_L" + string(order[i]) + "_rel_" + strings.Replace(refPath, "/", "_", -1) + "_L" + string(refOrder[i])},
		}
		ds.RegisterDistillate(registration)
	}
	ds.StartEngine()
}
