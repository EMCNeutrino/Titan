#!/bin/bash

unset ${!OS_*}

set -u

NEUTRINO_VIP=${1:-10.246.151.96}
AUTH_URL=https://${NEUTRINO_VIP}:35357/v3
DOMAIN=default
ADMIN_USER=admin
ADMIN_PASSWORD=admin123
ADMIN_PROJECT=admin
FLAVOR_TINY=m1.tiny

ARGS="--insecure \
      --os-auth-url $AUTH_URL \
      --os-username $ADMIN_USER \
      --os-password $ADMIN_PASSWORD \
      --os-tenant-name $ADMIN_PROJECT \
      --os-user-domain-name $DOMAIN \
      --os-project-domain-name $DOMAIN"

# nova flavor-create FLAVOR_NAME FLAVOR_ID RAM_IN_MB ROOT_DISK_IN_GB NUMBER_OF_VCPUS
FLAVOR=`nova ${ARGS} flavor-list | grep -w "|*[ ]$FLAVOR_TINY[ ]*|" | awk '{print $4}'| grep -x $FLAVOR_TINY`

if [ -z $FLAVOR ]
then
    echo "`date`: Create flavor '$FLAVOR_TINY'"
    set -x
    nova $ARGS flavor-create $FLAVOR_TINY auto 512 0 1
else
    echo "`date`: flavor '$FLAVOR_TINY' exists"
fi
