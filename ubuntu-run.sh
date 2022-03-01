#!/bin/bash

apt-get install git docker-ce docker-ce-cli containerd.io build-essential

./goinstall.sh

git clone git@github.com:anton-antonenko/help-ukraine.git
cd help-ukraine
make build

chmod +x ./build/darwin-amd64/help-ukraine

./build/darwin-amd64/help-ukraine