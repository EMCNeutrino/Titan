#/bin/bash

ARG=$1
COUNT=${ARG:-1}

set -ue

source ./openrc

set -x
heat --insecure stack-create -f hero-client.yaml \
     -e lib/env.yaml \
     -P "key_name=$HEAT_KEY_ID" \
     -P "net_id=$HEAT_NET_ID" \
     -P "name=$HEAT_PREFIX" \
     -P "image=$HEAT_IMAGE_ID_HERO" \
     -P "image_trusty=$HEAT_IMAGE_ID_TRUSTY" \
     -P "flavor=$HEAT_FLAVOR_ID" \
     -P "public_network=$HEAT_PUBLIC_NETWORK" \
     -P "count=$COUNT" \
     -P "security_group=$HEAT_SECGROUP_ID" \
     -P "hero_api=$HEAT_HERO_API" \
     -P "hero_token=$HEAT_HERO_TOKEN" \
     -P "neutrino_vip=$NEUTRINO_VIP" \
     ${HEAT_PREFIX}-clients
set +x
