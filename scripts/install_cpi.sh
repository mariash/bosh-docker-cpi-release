#!/usr/bin/env bash

set -e

go build -o $(dirname $0)/../release/blobs/docker-cpi $(dirname $0)/../cpi

