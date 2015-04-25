#!/usr/bin/env bash

set -ex

DIR=$( cd "$( dirname "$0" )" && pwd )

echo "Creating base docker image..."
# wget https://s3.amazonaws.com/bosh-warden-stemcells/bosh-stemcell-2776-warden-boshlite-ubuntu-trusty-go_agent.tgz
# tar zxf $DIR/bosh-stemcell-2776-warden-boshlite-ubuntu-trusty-go_agent.tgz
# cat image | sudo docker import - bosh:stemcell

echo "Installing go..."
source $DIR/go.env
# $DIR/install_go.sh

echo "Installing agent..."
bosh_agent_path=github.com/cloudfoundry/bosh-agent
go get -d $bosh_agent_path
(
	cd $GOPATH/src/$bosh_agent_path
	bin/build
)

echo "Building new image with latest agent..."
cp -r $DIR/docker $HOME
mv $GOPATH/src/$bosh_agent_path/out/bosh-agent $HOME/docker/assets/

cd $HOME/docker
docker build .