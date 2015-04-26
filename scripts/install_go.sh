#!/usr/bin/env bash

set -ex 

DIR=$( cd "$( dirname "$0" )" && pwd )

source $DIR/go.env

GO_ARCHIVE_URL=https://storage.googleapis.com/golang/go1.3.3.linux-amd64.tar.gz
GO_ARCHIVE=$HOME/$(basename $GO_ARCHIVE_URL)

echo "Downloading go..."
mkdir -p $(dirname $GOROOT)
rm -rf $GOROOT

wget -q $GO_ARCHIVE_URL -O $GO_ARCHIVE
tar xf $GO_ARCHIVE -C $(dirname $GOROOT)
chmod -R a+w $GOROOT

rm -f $GO_ARCHIVE