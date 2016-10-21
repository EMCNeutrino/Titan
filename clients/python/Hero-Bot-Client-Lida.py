#!/usr/bin/python
# Reqister hero user with game controller

import json
import requests
import sys
import os

hero_engine_ip = os.environ['HERO_ENGINE_IP']
hero_first_name = os.environ['HERO_FIRST_NAME']
hero_last_name = os.environ['HERO_LAST_NAME']
hero_name = os.environ['HERO_NAME']
hero_email = os.environ['HERO_EMAIL']
hero_class = 'chief'
auth_token = '1234'
headers = {'X-Auth-Token': auth_token}

# URL of the engine API
engine_url = 'http://' + hero_engine_ip + '/hero'
print('Engine API URL: ' + engine_url)


def register_hero(hero_first_name, hero_last_name, hero_name, hero_email):
    """
    Register a hero
    :param hero_first_name:The Hero's first name
    :param hero_last_name:The Hero's last name
    :param hero_name:The Hero's name
    :param hero_email:The Hero's email
    :return:The Hero authentication Token
    """

    # POST with parameters
    response = requests.post(engine_url, headers={'X-Auth-Token': auth_token},
                             params={'firstName': hero_first_name, 'lastName': hero_last_name, 'heroName': hero_name,
                                     'heroClass': hero_class, 'email': hero_email})

    team_auth = json.dumps(response.text)

    if response.status_code == 200:
        print ('Hero: \'' + hero_name + '\' joined the game!')
        print (hero_name + ' authentication Code: ' + team_auth)
    else:
        print ('Hero: \'' + hero_name + '\' joining game Failed!')
        print ("HTTP Code: " + str(response.status_code) + " | Response: " + response.text)

    return team_auth


def activate_hero(hero_name, token):
    """
    Activate a hero
    :param hero_name:The Hero's name
    :param token:The Hero's authentication token
    """
    url = engine_url + '/' + hero_name + '/activate'
    response = requests.post(url, headers={'X-Auth-Token': token})

    if response.status_code == 200:

        print ('Hero:' + hero_name + 'is activated')

    else:

        print ('Hero: ' + hero_name + ' activation Failed!')
        print ("HTTP Code: " + str(response.status_code) + " | Response: " + response.text)


if __name__ == '__main__':
    # register a Hero
    token = register_hero(hero_first_name, hero_last_name, hero_name, hero_email)
    # activate a Hero
    activate_hero(hero_name, token)
