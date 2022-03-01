#!/bin/bash

apt-get update
apt-get -y install git docker-ce docker build-essential

git clone https://github.com/anton-antonenko/help-ukraine.git
cd help-ukraine
./goinstall.sh
source /home/ubuntu/.bashrc
make build

chmod +x ./build/linux-amd64/help-ukraine

./build/linux-amd64/help-ukraine