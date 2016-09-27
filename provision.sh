#!/usr/bin/env bash

# Variables
MYSQL_PASSWD=root

echo -e "\n--- Updating packages list ---\n"
apt-get -qq update

echo -e "\n--- Install base packages ---\n"
apt-get install -y python3 python3-dev python3-pip python3-setuptools

echo -e "\n--- Install MySQL specific packages and settings ---\n"
debconf-set-selections <<< "mysql-server mysql-server/root_password password $MYSQL_PASSWD"
debconf-set-selections <<< "mysql-server mysql-server/root_password_again password $MYSQL_PASSWD"
apt-get -y install mysql-server

echo -e "\n--- Install Redis server ---\n"
apt-get install -y redis-server

echo -e "\n--- Install Python packages for Registration website ---\n"
pip3 install -r /home/ubuntu/src/web/requirements.txt