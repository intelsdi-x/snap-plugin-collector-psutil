#!/bin/bash 

#http://www.apache.org/licenses/LICENSE-2.0.txt
#
#
#Copyright 2016 Intel Corporation
#
#Licensed under the Apache License, Version 2.0 (the "License");
#you may not use this file except in compliance with the License.
#You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
#Unless required by applicable law or agreed to in writing, software
#distributed under the License is distributed on an "AS IS" BASIS,
#WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#See the License for the specific language governing permissions and
#limitations under the License.

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