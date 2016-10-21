#!/bin/bash

cat >> /home/ubuntu/hero-env-from-user << EOF
export HERO_FIRST_NAME=${HERO_FIRST_NAME}
export HERO_LAST_NAME=${HERO_LAST_NAME}
export HERO_NAME=${HERO_NAME}
export HERO_EMAIL=${HERO_EMAIL}
export HERO_TWITTER=${HERO_TWITTER}
export HERO_TOKEN=${HERO_TOKEN}
EOF

python /home/ubuntu/Hero/clients/python/Hero-Bot-Client.py > register-hero.log 2>&1 &
