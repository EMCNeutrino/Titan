#!/bin/bash

set -ue

unset ${!OS_*}
export OS_AUTH_URL=https://${VIP:-10.246.151.32}:35357/v3
export OS_IDENTITY_API_VERSION=3
export OS_USERNAME=heroadmin
export OS_PASSWORD=admin123
export OS_PROJECT_NAME=project_hero
export OS_TENANT_NAME=project_hero
export OS_USER_DOMAIN_NAME=hero
export OS_PROJECT_DOMAIN_NAME=hero

PUBLIC_NETWORK=public
PRIVATE_NETWORK=hero-network
KEYPAIR=hero-keypair
FLAVOR=m1.tiny
IMAGE=hero-ubuntu
HERO_VM_NAME=${HERO_VM_NAME:-hero-vm-01}
HERE_CONFIG=../hero-vm-config.sh

PROJECT_ID=`openstack --insecure project list | grep $OS_PROJECT_NAME | awk '{print $2}'`
NET_ID=`neutron --insecure net-list --tenant-id $PROJECT_ID | grep $PRIVATE_NETWORK | awk '{print $2}'`

NOVA="nova --insecure"

set -x
$NOVA boot \
  --flavor $FLAVOR \
  --image $IMAGE \
  --key-name $KEYPAIR \
  --nic net-id=$NET_ID \
  --user-data $HERE_CONFIG \
  $HERO_VM_NAME 
set +x

FIXED_IP=`$NOVA interface-list $HERO_VM_NAME | grep -v "+-----" | grep -v "| Port State[ ]*|" | awk '{print $8}'`
FLOATING_IP=`$NOVA floating-ip-list | grep "|[ ]*$FIXED_IP[ ]*|" | awk '{print $4}'`

if [ -z $FLOATING_IP ]
then
    echo "`date`: allocate a floating ip for from '$PUBLIC_NETWORK'"
    FLOATING_IP=`$NOVA floating-ip-create | grep $PUBLIC_NETWORK | awk '{print $4}'`
fi

if [ ! -z $FLOATING_IP ]
then
    echo "`date`: Associate floating ip '$FLOATING_IP' with $HERO_VM_NAME. The allocated ip is released upon error"
    set -x
    $NOVA add-floating-ip  $HERO_VM_NAME $FLOATING_IP
    set +x
    if [ $? != 0 ]
    then
        echo "`date`: Release floating ip '$FLOATING_IP'"
        set -x
        $NOVA floating-ip-delete $FLOATING_IP
        set +x
    fi
else
    echo "`date`: Cannot obtain floating ip for '$HERO_VM_NAME'. Fix the issue and re-run the script."
    exit 1
fi

