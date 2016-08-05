#!/bin/bash 

set -e
set -u
set -o pipefail

__dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
__proj_dir="$(dirname "$__dir")"
__proj_name="$(basename $__proj_dir)"

. "${__dir}/common.sh"

# NOTE: these variables control the docker-compose image.
export PLUGIN_SRC="${__proj_dir}"
export LOG_LEVEL="${LOG_LEVEL:-"7"}"
export PROJECT_NAME="${__proj_name}"

TEST_TYPE="${TEST_TYPE:-"large"}"

docker_folder="${__proj_dir}/scripts/docker/${TEST_TYPE}"

_docker_project () {
  echo ${docker_folder}
  cd "${docker_folder}" && "$@"
}

_debug "building docker compose images"
_docker_project docker-compose build
_debug "running test: ${TEST_TYPE}"
_docker_project docker-compose up
test_res=`docker-compose ps -q | xargs docker inspect -f '{{ .Name }} exited with status {{ .State.ExitCode }}' | awk '{print $5}'`
echo "exit code from large_jenkins $test_res"
_debug "stopping docker compose images"
_docker_project docker-compose stop
_docker_project docker-compose rm -f
exit $test_res