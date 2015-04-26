#!/usr/bin/env bash

set -ex

DIR=$( cd "$( dirname "$0" )" && pwd )

echo "Creating base docker image..."

(
	cd $HOME
	wget https://s3.amazonaws.com/bosh-warden-stemcells/bosh-stemcell-2776-warden-boshlite-ubuntu-trusty-go_agent.tgz
	tar zxf bosh-stemcell-2776-warden-boshlite-ubuntu-trusty-go_agent.tgz
	cat image | sudo docker import - bosh:stemcell
)

cd $DIR/docker
docker build .