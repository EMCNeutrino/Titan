#!/bin/bash

export GAME_CONTROLLER_IP=10.246.154.49
export HERO_NAME=`hostname`
export HERO_TITLE="Neutrino Guru"

# should preinstall python and git to speed up boot process
# install python setup tools
apt-get install -y python-setuptools python-dev build-essential

# install git
apt-get install -y git

git config --global http.sslVerify false

cd /opt
# clone Hero project as ubuntu user
git clone https://github.com/VxRackNeutrino/Hero

python /opt/Hero/scripts/hero-register-test.py
