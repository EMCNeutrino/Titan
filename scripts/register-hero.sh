#!/bin/bash

echo "HERO_FIRST_NAME=${HERO_FIRST_NAME}"
echo "HERO_LAST_NAME=${HERO_LAST_NAME}"
echo "HERO_NAME=${HERO_NAME}"
echo "HERO_EMAIL=${HERO_EMAIL}"
echo "HERO_TWITTER=${HERO_TWITTER}"
echo "HERO_TOKEN=${HERO_TOKEN}"
echo "HERO_ENGINE_IP=${HERO_ENGINE_IP}"

python /home/ubuntu/Hero/clients/python/Hero-Bot-Client.py
