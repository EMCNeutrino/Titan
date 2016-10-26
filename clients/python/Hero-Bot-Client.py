# !/bin/python
# Hero Bot Client: Register Hero with Game-controller

import json
import logging.config
import os
import random
import uuid
import sys
import requests

import settings

# Logging Initialization
logging.config.dictConfig(settings.NEUTRINO_HEROS_LOGGING)
logger = logging.getLogger("root")

#SYSTEM Variables
hero_engine_ip = os.environ['HERO_API']
auth_token = os.environ['HERO_ADMIN_TOKEN']
headers = {'X-Auth-Token': auth_token}

# URL of the engine API
engine_url = hero_engine_ip + '/hero'
print('Engine API URL: ' + engine_url)

#Array of Heros
heros = []

def get_hero_race():
    """
    Hero Race generates the Hero race
    """
    race_list = ["Elf", "Dwarf", "Goblin", "Morlock", "Half-Gnome", "Hobbit", "Gnome", "Orc", "Shade", "White Walker",
                 "Hobbgoblin", "Barbarian", "Half-elf", "Hafling", "Half-Orc", "Half-Giant", "Giant", "Drow",
                 "Half-Drow", "Minotaur", "Human", "Dragonborn", "Eladrin", "Ogre", "Half-Ogre", "Vampire", "Troll"]

    races = len(race_list) - 1

    race_id = random.randint(0, races)
    # random on race

    #logger.info("Hero Race: {0}".format(race_list[race_id]))

    return race_list[race_id]


def get_hero_class():
    """
    Hero Class Generates the Hero Class
    """

    hclass_list = ["Barbarian", "Bard", "Cleric", "Druid", "Figter", "Mage", "Wizard", "Monk", "Mistic", "Paladin",
                   "Ranger",
                   "Sorcerer", "Thief", "Ninja", "Warlock", "Defiler", "Shaman", "Invoker", "Psion"]

    hclass = len(hclass_list) - 1

    hclass_id = random.randint(0, hclass)
    # random on race

    #logger.info("Hero Class: {0}".format(hclass_list[hclass_id]))

    return hclass_list[hclass_id]


def get_hero_name():
    """
    Hero Name generates the Hero name
    """

    hero_name = ["James", "Mary", "John", "Patricia", "Robert", "Linda", "Michael", "Barbara", "William", "Elizabeth",
                 "David", "Jennifer", "Richard", "Maria", "Charles", "Susan", "Joseph", "Margaret", "Thomas", "Dorothy",
                 "Christopher", "Lisa", "Daniel", "Nancy", "Paul", "Karen", "Mark", "Betty", "Donald", "Helen",
                 "George", "Sandra", "Kenneth", "Donna", "Steven", "Carol", "Edward", "Ruth", "Brian", "Sharon",
                 "Ronald", "Michelle", "Anthony", "Laura", "Kevin", "Sarah", "Jason", "Kimberly", "Matthew",
                 "Deborah", "Gary", "Jessica", "Timothy", "Shirley", "Jose", "Cynthia", "Larry", "Angela",
                 "Jeffrey", "Melissa", "Frank", "Brenda", "Scott", "Amy", "Eric", "Anna", "Stephen", "Rebecca",
                 "Andrew", "Virginia", "Raymond", "Kathleen", "Gregory", "Pamela", "Joshua", "Martha", "Jerry",
                 "Debra", "Dennis", "Amanda", "Walter", "Stephanie", "Patrick", "Carolyn", "Peter", "Christine",
                 "Harold", "Marie", "Douglas", "Janet", "Henry", "Catherine", "Carl", "Frances", "Arthur", "Ann",
                 "Ryan", "Joyce", "Roger", "Diane"]

    hname = len(hero_name) - 1

    name_id = random.randint(0, hname)

    #logger.info("Hero Name: {0}".format(hero_name[name_id]))

    return hero_name[name_id]


def get_hero_lname():
    """
    Hero Name generates the Hero Last Name
    """

    hero_lname = ["Smith","Johnson","Williams","Jones","Brown","Davis","Miller","Wilson","Moore","Taylor","Anderson",
                  "Thomas","Jackson","White","Harris","Martin","Thompson","Garcia","Martinez","Robinson","Clark",
                  "Rodriguez","Lewis","Lee","Walker","Hall","Allen","Young","Hernandez","King","Wright","Lopez",
                  "Hill","Scott","Green","Adams","Baker","Gonzalez","Nelson","Carter","Mitchell","Perez","Roberts",
                  "Turner","Phillips","Campbell","Parker","Evans","Edwards","Collins","Stewart","Sanchez","Morris",
                  "Rogers","Reed","Cook","Morgan","Bell","Murphy","Bailey","Rivera","Cooper","Richardson","Cox",
                  "Howard","Ward","Torres","Peterson","Gray","Ramirez","James","Watson","Brooks","Kelly","Sanders",
                  "Price","Bennett","Wood","Barnes","Ross","Henderson","Coleman","Jenkins","Perry","Powell","Long",
                  "Patterson","Hughes","Flores","Washington","Butler","Simmons","Foster","Gonzales","Bryant",
                  "Alexander","Russell","Griffin","Diaz","Hayes"]

    hlname = len(hero_lname) - 1

    lname_id = random.randint(0, hlname)

    #logger.info("Hero Last Name: {0}".format(hero_lname[lname_id]))

    return hero_lname[lname_id]


def get_hero_title():
    """
    Hero Name generates the Hero title
    """

    title_noum = ["Dream", "Dreamer", "Dreams","Rainbow","Dreaming","Flight","Wings","Mist",
                          "Sky","Wind","Winter","Misty","Cloud","Fairy","Dragon","End",
                          "Beginning","Tale","Tales","Emperor","Prince","Princess","Willow","Birch","Petals",
                          "Destiny","Theft","Thief","Legend","Prophecy",
                          "Spark","Sparks","Stream","Streams","Waves",
                          "Sword","Darkness","Swords","Silence","Kiss",
                          "Butterfly","Shadow","Ring","Rings","Emerald",
                          "Storm","Storms","Mists","World","Worlds",
                          "Alien","Lord","Lords","Ship","Ships","Star",
                          "Stars","Force","Visions","Vision","Magic",
                          "Wizards","Wizard","Heart","Heat","Twins",
                          "Twilight","Moon","Moons","Planet","Shores",
                          "Pirates","Courage","Time","Academy",
                          "School","Rose","Roses","Stone","Stones",
                          "Sorcerer","Shard","Shards","Slave","Slaves",
                          "Servant","Servants","Serpent","Serpents",
                          "Snake","Soul","Souls","Savior","Spirit",
                          "Spirits","Voyage","Voyages","Voyager","Voyagers",
                          "Return","Legacy","Birth","Healer","Healing",
                          "Year","Years","Death","Dying","Luck","Elves",
                          "Tears","Touch","Son","Sons","Child","Children",
                          "Illusion","Sliver","Destruction","Crying","Weeping",
                          "Gift","Word","Words","Thought","Thoughts","Scent",
                          "Ice","Snow","Night","Silk","Guardian","Angel",
                          "Angels","Secret","Secrets","Search","Eye","Eyes",
                          "Danger","Game","Fire","Flame","Flames","Bride",
                          "Husband","Wife","Time","Flower","Flowers",
                          "Light","Lights","Door","Doors","Window","Windows",
                          "Bridge","Bridges","Ashes","Memory","Thorn",
                          "Thorns","Name","Names","Future","Past",
                          "History","Nothingness","Someone",
                          "Person","Man","Woman","Boy","Girl",
                          "Way","Mage","Witch","Witches","Curse",
                          "Talisman","Mirror","Mirrors","Jewel","Jewels",
                          "Firefly","Fireflies","Cross","Fantasy","Card",
                          "Cards","God","Gods","Goddess","Evil",
                          "Warrior","Warriors","Virgin","War","Battle",
                          "Journey","Brother","Brothers","Sister","Sisters",
                          "Grave","Voice","Voices","Fox","Lion","Wolf",
                          "Wolves","Bunny","Bird","Birds","Hero","Crown",
                          "Sand","Sun","Sunlight","Sands","Hope","Lust",
                          "Passion","Love","Hate","Wrath","Anger","Revenge",
                          "Need","Crossing","Passage","Rite","Rites","Crystal",
                          "Crystals","Cave","Caves","Cavern","Caverns","Island",
                          "Land","Lands","Ocean","Sea","Oceans","Seas","Desert",
                          "Deserts","Forest","Forests","Woods","Jungle","Jungles",
                          "Exile","Empire","Kingdom","City","Village","Garden","Fall",
                          "Spring","Eve","Dawn","Oath","Gold","Cities","Ally","Hand",
                          "Sign","Omen","Court","Courts","Unicorn","Pegasus","Griffon",
                          "Farewell","Tower","Towers","Heaven","Hell","Heavens",
                          "Hunt","Rebirth","Path","Dagger","Daggers","Knife","Knives",
                          "Tooth","Teeth","Bite","Promise","Price","Queen","King",
                          "Queens","Kings","Arrow","Arrows","Song","Singer","Dance",
                          "Dancer","Dances","Dancers","Torture","Lament","Pain",
                          "Assassin","Spy","Spies","Agony","Mistress","Tree",
                          "Trees","Master","Honor","Loyalty","Trust","Apprentice",
                          "Prison","Prisoner","Phoenix","Breath","Horn","Claw","Talon",
                          "Tail","Speaker","Eclipse","Scream","Screams","Revolution",
                          "Thrusting","Wraith","Phantom","Sheep"]

    title_adverb = ["Misty","Lost","Only","Last","First", "Final","Missing","Shadowy","Seventh",
                   "Dark","Darkest","Silver","Silvery", "Black","White","Hidden","Entwined","Invisible",
                   "Next","Seventh","Red","Green","Blue", "Purple","Grey","Bloody","Emerald","Diamond",
                   "Frozen","Sharp","Delicious","Dangerous", "Deep","Twinkling","Dwindling","Missing","Absent",
                   "Vacant","Cold","Hot","Burning","Forgotten", "No","All","Which","What","Hard","Soft",
                   "Playful","Final","Evil","Scarlet","Chaste","Virgin", "Strange","Silent", "Legendary",
                   "Golden","Magic", "Mystic","Majestic","Magical","Mysterious","Eternal",
                   "Winged","Outer","Inner","Silken","Mystical","Crying","Weeping","Lonely","Crushed","Searching",
                   "Desperate","Yearning","Quick","Invincible","New","Old","Ancient","Aging","Dying","Living",
                   "Vengeful","Loving","Crystal","Crystalline","Wooden","Metal","Metallic","Marble","Stony","Rocky",
                   "Great","Royal", "Noble","Wet","Dry","Bleeding","Piercing","Singing", "Dancing","Painful",
                   "Wandering","Loyal","Trusting","Open","Closed","Locked","Free","Chained","Caged",
                   "Empty","Wilted","Lunar","Solar","Screaming","Dead","Shaking","Thrusting","Frantic"]


    noum_size  = len(title_noum) - 1
    adverb_size= len(title_adverb) - 1

    noum_id = random.randint(0, noum_size)
    adverb_id = random.randint(0, adverb_size)

    full_title = "The {0} of the {1}".format(title_noum[noum_id], title_adverb[adverb_id])

    #logger.info("Hero Title: {0}".format(full_title))

    return full_title


def Load_Heros(heros_number):
    """
    Load Heros
    """
    for i in range(0, heros_number):

        myhero = hero()
        myhero.hero_name = get_hero_name()
        #Load the Hero information
        myhero.player_name = get_hero_name()
        myhero.player_lastname = "Bot-O-Matic"
        myhero.token = uuid.uuid4()
        myhero.twitter =  "@{0}.{1}".format(myhero.player_name, myhero.player_lastname)

        random_number = random.randint(1,1000000000)

        myhero.email  = "{0}.{1}{2}@dell.com".format(myhero.player_name,myhero.player_lastname, random_number)
        myhero.title  = get_hero_title()
        myhero.race   = get_hero_race()
        myhero.isAdmin = 0
        myhero.hero_level = random.randint(1,5)
        myhero.hero_level = 0
        myhero.hclass = get_hero_class()
        myhero.ttl = 0
        myhero.userhost = ""
        myhero.xpos = random.randint(1,500)
        myhero.ypos = random.randint(1,500)
        myhero.next_level = 0

        myhero.weapon   = random.randint(1,10)
        myhero.tunic    = random.randint(1,10)
        myhero.shield   = random.randint(1,10)
        myhero.leggings = random.randint(1,10)
        myhero.ring     = random.randint(1,10)
        myhero.gloves   = random.randint(1,10)
        myhero.boots    = random.randint(1,10)
        myhero.helm     = random.randint(1,10)
        myhero.charm    = random.randint(1,10)
        myhero.amulet   = random.randint(1,10)

        myhero.set_total_equipment()

        heros.append(myhero)


def Delete_Hero_by_HeroID(heros_id):
    return "Hello"

    heros.remove(heros_id)


def Get_hero(heros):

    selected = random.randint(1,len(heros))

    return heros[selected]

def register_hero(player_first_name, player_last_name, hero_name, hero_email, hero_class):
    """
    Register a hero
    :param player_first_name:The Hero's first name
    :param player_last_name:The Hero's last name
    :param hero_name:The Hero's name
    :param hero_email:The Hero's email
    :return:The Hero authentication Token
    """

    logger.info("Player: {0} {1} | Hero: {2} | Email: {3}".format(player_first_name, player_last_name, hero_name, hero_email, hero_class))

    # POST with parameters
    data = {'firstName': player_first_name, 'lastName': player_last_name, 'email': hero_email,
            'heroName': hero_name,'heroClass': hero_class }
    response = requests.post(engine_url, headers={'X-Auth-Token': auth_token}, data = json.dumps(data))


    if response.status_code == 200:
        result = json.loads(response.text)
        token = result['token']
        print ('Hero: \'' + hero_name + '\' joined the game!')
        print (hero_name + ' authentication Code: ' + token)
        return token
    else:
        print ('Hero: \'' + hero_name + '\' joining game Failed!')
        print ("HTTP Code: " + str(response.status_code) + " | Response: " + response.text)


def activate_hero(hero_name, token):
    """
    Activate a hero
    :param hero_name:The Hero's name
    :param token:The Hero's authentication token
    """
    url = engine_url + '/' + hero_name + '/activate'
    print url
    response = requests.get(url, headers={'X-Auth-Token': token})

    if response.status_code == 200:

        print ('Hero:' + hero_name + 'is activated')

    else:

        print ('Hero: ' + hero_name + ' activation Failed!')
        print ("HTTP Code: " + str(response.status_code) + " | Response: " + response.text)


class hero:
    """Class that contains the Hero information """

    def __init__(self):
        self.heroID = 0
        self.hero_name = ""
        self.player_name = ""
        self.player_lastname = ""
        self.token = ""
        self.twitter = ""
        self.email = ""
        self.title = ""
        self.race = ""
        self.isAdmin = 0
        self.hero_level = 0
        self.hclass = ""
        self.ttl = 0
        self.userhost = ""
        self.hero_online = 0
        self.xpos = 0
        self.ypos = 0
        self.next_level = 0
        self.weapon = 0
        self.tunic = 0
        self.shield = 0
        self.leggings = 0
        self.ring = 0
        self.gloves = 0
        self.boots = 0
        self.helm = 0
        self.charm = 0
        self.amulet = 0
        self.total_equipment = 0

    def set_total_equipment(self):

        total_equipment = self.weapon + self.tunic + self.shield + self.leggings + self.ring + self.gloves + self.boots + self.helm + self.charm + self.amulet
        #logger.info("Total Equipment: {0}".format(total_equipment))


def main():

    heroesCount = int(sys.argv[1])
    Load_Heros(heroesCount)

    logger.info("Heros created: {0}".format(len(heros)))

    #selected = random.randint(1, len(heros))
    for myhero in heros:

        logger.info("Heros selected: {0}, {1}, {2}, {3}, {4}".format(myhero.player_name, myhero.player_lastname, myhero.hero_name, myhero.email, myhero.hclass))

        try:
            # register a Hero
            token = register_hero(myhero.player_name, myhero.player_lastname, myhero.hero_name, myhero.email, myhero.hclass)

            # activate a Hero
            if token is not None:
                logger.info("Activate hero : {0}".format(myhero.hero_name))
                activate_hero(myhero.hero_name, token)
        except Exception as err:
            logger.error("error: {0}".format(err))



if __name__ == "__main__":
    main()
