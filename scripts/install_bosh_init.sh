#!/usr/bin/env bash

mkdir -p $HOME/bin
wget https://s3.amazonaws.com/bosh-init-artifacts/bosh-init-0.0.4-linux-amd64 -O $HOME/bin/bosh-init
chmod +x $HOME/bin/bosh-init