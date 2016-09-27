#!/bin/bash

sudo apt-get update

# Python3
sudo apt-get install -y python3 python3-dev python3-pip python3-setuptools

# Mysql Server 5.7
sudo apt-get install -y mysql-server

# Redis
sudo apt-get install -y redis

# Python dependencies for Django website
sudo pip3 install -r web/requirements.txt
