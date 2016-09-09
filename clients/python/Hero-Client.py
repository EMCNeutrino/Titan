#!/bin/python
# Mars Team Client Example written in Python
# Requires the following library to install: sudo pip install websocket-client
# if you encounter errors with a Six import:
# you can try: pip remove six; pip install six
# Windows users: you may need to install the Microsoft Visual C++ Compiler for Python 2.7
# Windows users. please use this link: http://aka.ms/vcpython27
import json

import requests
import websocket
from websocket import create_connection

# Global Variables
hero_name = 'Gandalf'                       # Hero's name
hero_title= 'The Grey'                      # Hero's title
hero_auth = ''                              # The Hero Authentication Token
server_url = 'http://localhost:8080/api'    # URL of the Game Controller API
server_ws = 'ws://localhost:8080/ws'        # URL of the Game Controller Websocket


# Server Method Calls ------------------------------------------------

def register_hero(hero_name):
    """
    Registers the Team in the Server
    :param hero_name:The Hero's name
    :return:The Hero authentication Token
    """

    url = server_url + "/join/" + hero_name
    print('Server API URL: ' + url)
    payload = ''

    # POST with form-encoded data
    response = requests.post(url, data=payload)

    team_auth = response.text
    # print ('Hero Authentication Code:' + team_auth )

    if response.status_code == 200:
        print ('Hero: \'' + hero_name + '\' joined the game!')
        print (hero_name + ' authentication Code: ' + team_auth)
    else:
        print ('Hero: \'' + hero_name + '\' joining game Failed!')
        print ("HTTP Code: " + str(response.status_code) + " | Response: " + response.text)

    return team_auth


# Hero Method Calls ------------------------------------------------
def hero_quest_request(hero_name, hero_auth):
    """
    Sets the team shield up
    curl -i -H 'X-Auth-Token: 1335aa6af5d0289f' -X POST http://localhost:8080/api/Quest/enable
    :param hero_name:The team name
    :param hero_auth:The team authentication token
    :return: nothing
    """
    url = server_url + '/shield/enable'
    auth_header = {'X-Auth-Token': hero_auth}
    quest_request = requests.post(url, headers=auth_header)
    if quest_request.status_code == 200:
        print ('Server: Team: ' + hero_name + ' Shield is UP!')
    else:
        print ('Server: Team: ' + hero_name + ' Quest Reuquest! request Failed!')
        print ("HTTP Code: " + str(quest_request.status_code) + " | Response: " + quest_request.text)


def hero_battle_request(team_name, team_auth):
    """
    Sets the team shield up
    curl -i -H 'X-Auth-Token: 1335aa6af5d0289f' -X POST http://localhost:8080/api/Quest/enable
    :param team_name:The team name
    :param team_auth:The team authentication token
    :return: nothing
    """
    url = server_url + '/shield/enable'
    auth_header = {'X-Auth-Token': team_auth}
    shield_up = requests.post(url, headers=auth_header)
    if shield_up.status_code == 200:
        print ('Server: Team: ' + team_name + ' Shield is UP!')
    else:
        print ('Server: Team: ' + team_name + ' Shield UP! request Failed!')
        print ("HTTP Code: " + str(shield_up.status_code) + " | Response: " + shield_up.text)



def hero_shield_down(team_name, team_auth):
    """
    Sets the team shield Down
    curl -i -H 'X-Auth-Token: 1335aa6af5d0289f' -X POST http://localhost:8080/api/shield/disable
    :param team_name:The team name
    :param team_auth:The team authentication token
    :return: nothing
    """
    url = server_url + '/shield/disable'
    auth_header = {'X-Auth-Token': team_auth}
    shield_down = requests.post(url, headers=auth_header)
    if shield_down.status_code == 200:
        print ('Server: Team: ' + team_name + ' Shield is DOWN!')
    else:
        print ('Server: Team: ' + team_name + ' Shield DOWN! request Failed!')
        print ("HTTP Code: " + str(shield_down.status_code) + " | Response: " + shield_down.text)


# Client Logic ------------------------------------------------

def data_recording(parsed_json):
    """
    Saves the Mars sensor data in data repository
    :param parsed_json:Readings from Mars Sensors
    :return:Nothing
    """
    print("\nData Recording: Saving Data row for persistence. Time: " + str(parsed_json['startedAt']))


def hero_strategy(parsed_json):
    """
  Contains the Team's strategy.
  :param parsed_json: Readings from the Mars Sensors
  :return:Nothing
  """
    # The Strategy for this client is to have the shield up constantly until it is depleted.
    # Then wait until is charged again to a 10% and enable it again

    # Get the Team List
    teams_list = parsed_json['teams']

    # Find this team
    for team in teams_list:
        if team['name'] == hero_name:
            if team['shield'] <> True and team['energy'] > 10:
                # Check if Shield is up and shield energy is larger than 10%
                print("\nGameMove: Hero: {0} Action: Shield UP!| Energy: {1}".format(hero_name, str(team['energy'])))
                hero_quest_request(hero_name, hero_auth)
            else:
               print("\nHero: {0} Action: None| Energy: {1}".format(hero_name, str(team['energy'])))


# Register the Team

hero_auth = register_hero(hero_name)


# Create the WebSocket for Listening
ws = websocket.create_connection(server_ws)

while True:

    json_string = ws.recv()  # Receives the the json information

    # Received '{"running":false,"startedAt":"2015-08-04T00:44:40.854306651Z","readings":{"solarFlare":false,"temperature":-3.
    # 960996217958863,"radiation":872},"teams":[{"name":"TheBorgs","energy":100,"life":0,"shield":false},{"name":"QuickFandang
    # o","energy":100,"life":0,"shield":false},{"name":"InTheBigMessos","energy":32,"life":53,"shield":false},{"name":"MamaMia
    # ","energy":100,"life":100,"shield":false}]}'

    parsed_json = json.loads(json_string)

    # Check if the game has started
    print("Game Status: " + str(parsed_json['running']))

    if not parsed_json['running']:
        print('Waiting for the Game Start')
    else:
        data_recording(parsed_json)
        hero_strategy(parsed_json)

ws.close()

print "Good bye!"
