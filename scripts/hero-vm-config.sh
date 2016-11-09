#!/bin/bash
set -x

# Activate hero

export HERO_PORTAL_IP=10.246.152.15
export HERO_NAME=larry
# Hero token you received after registration
export HERO_TOKEN=48fdafd90a385c7e

curl -s -H "X-Auth-Token: ${HERO_TOKEN}" http://${HERO_PORTAL_IP}/hero/${HERO_NAME}/activate
