#!/bin/bash

set -e
set -u
set -o pipefail

# get the directory the script exists in
__dir="$(cd ../../scripts/docker/large && pwd)"
__proj_dir="$(cd ../../ && pwd)"

export PLUGIN_SRC="${__proj_dir}"

# verifies dependencies and starts influxdb
. "${__proj_dir}/examples/tasks/.setup.sh"

# start the influxdb container
(cd $__dir && docker-compose up)

# clean up containers on exit
function finish {
  (cd $__dir && docker-compose down)
}
trap finish EXIT INT TERM


