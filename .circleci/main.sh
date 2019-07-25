#!/bin/bash -e

PROJ_PATH=`pwd`
GO_VERSION="1.12.6"

# First let's install Go 1.11
echo "Installing Go 1.12..."
cd /tmp

curl -O https://dl.google.com/go/go${GO_VERSION}.linux-amd64.tar.gz
sudo tar -xf go${GO_VERSION}.linux-amd64.tar.gz

# Uninstall older version of Go
sudo rm -rf /usr/local/go
sudo mv go /usr/local

export GOROOT=/usr/local/go
#export GOPATH=$PROJ_PATH
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH

cd $PROJ_PATH

./.circleci/bring-gamma.sh

./test.sh
