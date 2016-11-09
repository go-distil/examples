#!/bin/bash

/usr/bin/docker run \
  --name distil-mpa-noop \
  -e DISTIL_MONGO_ADDR="cm1.smartgrid.store:4410" \
  -e DISTIL_BTRDB_ADDR="cm1.smartgrid.store:4410" \
  -e SOURCECODE="github.com/go-distil/examples/2-frequency" \
  -e REF_PMU_PATH="/REFSET/LBNL/a6_bus1" \
  btrdb/distiller
