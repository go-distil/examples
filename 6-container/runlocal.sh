#!/bin/bash

/usr/bin/docker run -it \
  -e DISTIL_MONGO_ADDR="cm1.smartgrid.store:27017" \
  -e DISTIL_BTRDB_ADDR="cm1.smartgrid.store:4410" \
  -e SOURCECODE="github.com/immesys/examples/2-frequency" \
  -e REF_PMU_PATH="/REFSET/LBNL/a6_bus1" \
  btrdb/distiller
