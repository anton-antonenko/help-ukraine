#!/bin/bash

apt-get update
apt-get -y install git docker-ce docker build-essential

git clone git@github.com:anton-antonenko/help-ukraine.git
cd help-ukraine
./goinstall.sh
make build

chmod +x ./build/darwin-amd64/help-ukraine

./build/darwin-amd64/help-ukraine