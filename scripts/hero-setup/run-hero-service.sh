#/bin/bash

ARG=$1
COUNT=${ARG:-1}
ACTION=${HEAT_ACTION:-create}

set -ue

source ./openrc

set -x
heat --insecure stack-${ACTION} -f hero-service.yaml \
     -e lib/env.yaml \
     -P "key_name=$HEAT_KEY_ID" \
     -P "net_id=$HEAT_NET_ID" \
     -P "name=$HEAT_PREFIX" \
     -P "image=$HEAT_IMAGE_ID_XENIAL" \
     -P "image_trusty=$HEAT_IMAGE_ID_TRUSTY" \
     -P "flavor=$HEAT_FLAVOR_ID" \
     -P "public_network=$HEAT_PUBLIC_NETWORK" \
     -P "count=$COUNT" \
     -P "security_group=$HEAT_SECGROUP_ID" \
     -P "neutrino_vip=$NEUTRINO_VIP" \
     -P "admin_password=$ADMIN_PASSWORD" \
     ${HEAT_PREFIX}-services
set +x
