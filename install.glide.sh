#!/usr/bin/env bash

sudo add-apt-repository ppa:masterminds/glide -y
sudo apt-get update -q
sudo apt-get install glide -y
glide install