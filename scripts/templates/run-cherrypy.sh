#/bin/bash

ARG=$1
COUNT=${ARG:-2}

set -ue

source ./openrc

set -x
heat --insecure stack-create -f cherrypy.yaml \
     -e lib/env.yaml \
     -P "key_name=$HEAT_KEY_ID" \
     -P "net_id=$HEAT_NET_ID" \
     -P "name=$HEAT_PREFIX" \
     -P "image=$HEAT_IMAGE_ID" \
     -P "flavor=$HEAT_FLAVOR_ID" \
     -P "public_network=$HEAT_PUBLIC_NETWORK" \
     -P "count=$COUNT" \
     -P "security_group=$HEAT_SECGROUP_ID" \
     ${HEAT_PREFIX}-cherrypy
set +x
echo "Show floating IPs"
sleep 5
heat --insecure output-show ${HEAT_PREFIX}-cherrypy fip
