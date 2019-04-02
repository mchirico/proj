#!/bin/bash

set -eu

DIR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )
export fly_target=${fly_target:-mce}
echo "Concourse API target ${fly_target}"
echo "Tutorial $(basename $DIR)"

pushd $DIR
  fly -t ${fly_target} set-pipeline -p proj-pipeline -c build-golang-pipeline.yml -n
  fly -t ${fly_target} unpause-pipeline -p proj-pipeline
#  fly -t ${fly_target} trigger-job -w -j tutorial-pipeline/job-hello-world
popd

echo -e "\n\n                  Common commands:"
echo -e "**************************************\n\n"
echo -e "\n"
echo -e "                           fly -t mce watch --job proj-pipeline/unit"
echo -e "                           fly -t mce builds|grep 'proj-pipeline'"
echo -e "                           fly -t mce destroy-pipeline -p proj-pipeline -n"
echo -e "                           fly -t mce workers -d "
echo -e "                            "
echo -e "                           To login to a running container: "
echo -e "                           fly -t mce intercept --job proj-pipeline/unit "
echo -e "\n"
echo -e "\n"

