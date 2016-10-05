#/bin/bash

set -ue

source ./openrc

set -x

heat --insecure stack-create -f init.yaml \
     -e lib/env.yaml \
     -P "image_id=$HEAT_IMAGE_ID" \
     -P "image_location=$HEAT_IMAGE_LOCATION" \
     -P "public_network=$HEAT_PUBLIC_NETWORK" \
     -P "prefix=$HEAT_PREFIX" \
     -P "dns_servers=$HEAT_DNS_SERVERS" \
     ${HEAT_PREFIX}-init
