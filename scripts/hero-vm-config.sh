#!/bin/bash

export HERO_FIRST_NAME=Gandalf
export HERO_LAST_NAME=TheGray
export HERO_NAME=Magician
export HERO_EMAIL="gandalf_thegray@example.com"
# authentication token you received after registration
export HERO_TOKEN=1234
# optional Twitter account
export HERO_TWITTER="Gandalf"

# ----------------------------------------
export HERO_ENGINE_IP=10.246.152.15
/home/ubuntu/Hero/scripts/register-hero.sh
