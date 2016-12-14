#!/bin/bash

set -e
set -u
set -o pipefail

# get the directory the script exists in
#__dir="$(cd ../../scripts/docker/large && pwd)"
__dir=$(cd $(dirname ${BASH_SOURCE[0]})/../scripts/docker/large && pwd)
__proj_dir="$(cd $(dirname ${BASH_SOURCE[0]})/../ && pwd)"
__proj_name="$(basename $__proj_dir)"

export PLUGIN_SRC="${__proj_dir}"

# verifies dependencies and starts influxdb
. "${__proj_dir}/examples/.setup.sh"

# downloads plugins, starts snap, load plugins and start a task
__id=$(docker run -e SNAP_VERSION=latest -d -v ${PLUGIN_SRC}:/${__proj_name} --net=host mkrolik/snap-pytest)
# clean up containers on exit
function finish {
  (docker kill ${__id})
}
trap finish EXIT INT TERM

docker exec -it ${__id} bash -c "PLUGIN_PATH=/etc/snap/plugins /${__proj_name}/examples/file-psutil.sh && printf \"\n\nhint: type 'snaptel task list'\ntype 'exit' when your done\n\n\" && bash"
