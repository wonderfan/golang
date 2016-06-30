#!/usr/bin/env bash

sudo apt-get install git

wget https://get.docker.com/builds/Linux/x86_64/docker-latest.tgz

tar -xvzf docker-latest.tgz

sudo mv docker/* /usr/bin/

sudo docker daemon &
