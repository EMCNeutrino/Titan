#!/bin/bash

sudo apt-get update
sudo apt-get install -y python3 python3-dev python3-pip python3-setuptools

# Mysql Server 5.7
sudo apt-get install -y mysql-server

# Python
sudo pip3 install -r web/requirements.txt
