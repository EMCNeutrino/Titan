#!/usr/bin/env bash

# Variables
MYSQL_PASSWD=root

echo -e "\n--- Updating packages list ---\n"
apt-get -qq update

echo -e "\n--- Install base packages ---\n"
apt-get install -y python3 python3-dev python3-pip python3-setuptools

echo -e "\n--- Install Docker ---\n"
apt-get install -y docker.io
usermod -aG docker ubuntu

echo -e "\n--- Install MySQL specific packages and settings ---\n"
debconf-set-selections <<< "mysql-server mysql-server/root_password password $MYSQL_PASSWD"
debconf-set-selections <<< "mysql-server mysql-server/root_password_again password $MYSQL_PASSWD"
apt-get -y install mysql-server

echo -e "\n--- Install Redis server ---\n"
apt-get install -y redis-server

echo -e "\n--- Install Python packages for Registration website ---\n"
pip3 install -r /home/ubuntu/src/web/requirements.txt

echo -e "\n--- Install Golang 1.7 ---\n"
cd /home/ubuntu
curl -O https://storage.googleapis.com/golang/go1.7.1.linux-amd64.tar.gz
tar xvf go1.7.1.linux-amd64.tar.gz
chown -R root:root ./go
rm -rf /usr/local/go
mv go /usr/local
mkdir -p ./work/github.com/VxRackNeutrino/Hero/
ln -s /home/ubuntu/src/game-engine ./work/github.com/VxRackNeutrino/Hero/game-engine
cat <<EOT >> ./.profile
export GOPATH=/home/ubuntu/work
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
EOT
