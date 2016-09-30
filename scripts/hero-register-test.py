#!/usr/bin/python
# Reqister hero user with game controller

import json
import requests
import sys
import os

# Global Variables
hero_name = 'Gandalf'                       # Hero's name
hero_title= 'The Grey'                      # Hero's title
hero_auth = ''                              # The Hero Authentication Token
game_controller_ip = os.environ['GAME_CONTROLLER_IP'] 
# URL of the Game Controller API
server_url = 'http://' + game_controller_ip + '/api'   


# Server Method Calls ------------------------------------------------

def register_hero(hero_name, hero_title):
    """
    Registers the Team in the Server
    :param hero_name:The Hero's name
    :param hero_title:The Hero's title
    :return:The Hero authentication Token
    """

    url = server_url + "/heros/"
    print('Game Controller API URL: ' + url)
    raw_payload = {'name' : hero_name, 'title' : hero_title}
    json_payload = json.dumps(raw_payload)

    # POST with form-encoded data
    response = requests.post(url, data=json_payload)

    team_auth = json.dumps(response.text)

    if response.status_code == 200:
        print ('Hero: \'' + hero_name + '\' joined the game!')
        print (hero_name + ' authentication Code: ' + team_auth)
    else:
        print ('Hero: \'' + hero_name + '\' joining game Failed!')
        print ("HTTP Code: " + str(response.status_code) + " | Response: " + response.text)

    return team_auth

if __name__ == '__main__':
    name = os.environ['HERO_NAME']
    title = os.environ['HERO_TITLE']
    register_hero(name, title)

