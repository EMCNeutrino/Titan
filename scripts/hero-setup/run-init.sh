#/bin/bash

set -ue

source ./openrc

set -x

heat --insecure stack-create -f init.yaml \
     -e lib/env.yaml \
     -P "image_id_trusty=$HEAT_IMAGE_ID_TRUSTY" \
     -P "image_id_xenial=$HEAT_IMAGE_ID_XENIAL" \
     -P "image_location_trusty=$HEAT_IMAGE_LOCATION_TRUSTY" \
     -P "image_location_xenial=$HEAT_IMAGE_LOCATION_XENIAL" \
     -P "public_network=$HEAT_PUBLIC_NETWORK" \
     -P "prefix=$HEAT_PREFIX" \
     -P "dns_servers=$HEAT_DNS_SERVERS" \
     ${HEAT_PREFIX}-init
