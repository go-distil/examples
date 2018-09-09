package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	btrdb "gopkg.in/btrdb.v4"
	"gopkg.in/distil.v4"
)

// This example extends #1 to add a few new things:
//	- loading DB params from the environment variables
//  - a more complex distillate that requires lead samples
//  - automatic generation of algorithm instances

func main() {
	// Use default connection params, this makes the resulting executable
	// portable to different installations
	ds := distil.NewDISTIL()

	// Get a connection to BTrDB
	bdb := ds.BTrDBConn()

	// Look up streams that have the name L1MAG in a collection that is a
	// subset of the environment variable COLLECTION_PREFIX
	streams, err := bdb.LookupStreams(context.Background(), os.Getenv("COLLECTION_PREFIX"), true,
		btrdb.OptKV("name", "L1ANG"), nil)
	if err != nil {
		panic(err)
	}

	//For each stream, register a distillate
	for _, st := range streams {
		// Get the stream's collection
		col, err := st.Collection(context.Background())
		if err != nil {
			panic(err)
		}
		// Get it's tags
		tags, err := st.Tags(context.Background())
		if err != nil {
			panic(err)
		}
		// DISTIL refers to streams by "path" which is a concatenation
		// of the collection and the name (as it would appear in mrplotter)
		path := col + "/" + tags["name"]
		fmt.Printf("registering distillate for path=%q\n", path)
		ds.RegisterDistillate(&distil.Registration{
			// The class that implements your algorithm
			Instance: &FrequencyDistiller{basefreq: 120},
			// A unique name FOR THIS INSTANCE of the distillate. If you
			// are autogenerating distillates, take care to never produce
			// the same name here. To be safe we are going to use the input
			// stream uuid prefaced by a unique identifier for this algorithm
			UniqueName: fmt.Sprintf("freq6_%s", strings.Trim(st.UUID().String(), "-")),
			// These are inputs to the distillate that will be loaded
			// and presented to Process()
			InputPaths: []string{path},
			// These are the output paths for the distillate. They must
			// also be strictly unique.
			OutputPaths: []string{col + "/freq_1s", col + "/freq_c37"},
			// The units for the distillate
			OutputUnits: []string{"hz", "hz"},
		})
	}

	ds.StartEngine()
}
