/******************************
** File: TitanDB_Creator.sql
** Name: Titan's game Database
** Desc: Creates the Titan's game database
** Auth: Patrick Butler Monterde
** Date: 09/19/2016
**************************
** Change History
**************************
** PR   Date	     Author    Description
** --   --------   -------   ------------------------------------
** 1    09/19/2016 PBM       Created
** 2    10/04/2016 PBM		 Updated Create statement and required fields
*******************************/

# Creating the TitanDB
CREATE DATABASE IF NOT EXISTS titandb;

# Create User and assign Permisions
CREATE USER 'titanuser'@localhost IDENTIFIED BY  'Neutrin0R0cks!';
GRANT ALL ON titandb.* TO 'titanuser' IDENTIFIED BY 'Neutrin0R0cks!';

# Drop Tables
DROP TABLE IF EXISTS item;
DROP TABLE IF EXISTS  penalty;
DROP TABLE IF EXISTS  heroworldevent;
DROP TABLE IF EXISTS  hero;
DROP TABLE IF EXISTS  worldevent;

#Drop Functions
DROP FUNCTION IF EXISTS randomizer;



# Tables

CREATE TABLE `hero` (

  `hero_id` 		int(11) NOT NULL AUTO_INCREMENT,
  `hero_name` 		text,
  `player_name` 	text,
  `player_lastname` text,
  `token` 			text,
  `userpass` 		text,
  `energy` 			int(11) DEFAULT NULL,
  `twitter` 		text,
  `email` 			text,
  `title` 			text,
  `race` 			text,
  `isAdmin` 		tinyint(1) DEFAULT NULL,
  `hero_level` 		int(11) DEFAULT NULL,
  `hclass` 			text,
  `ttl` 			int(11) DEFAULT NULL,
  `userhost` 		text,
  `hero_online` 	tinyint(1) DEFAULT NULL,
  `xpos` 			int(11) DEFAULT NULL,
  `ypos` 			int(11) DEFAULT NULL,
  `next_level` 		DATETIME DEFAULT NULL,
  `hero_created` 	TIMESTAMP DEFAULT NOW(),
  
  
  PRIMARY KEY (`hero_id`)
  
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COMMENT='Contains the Hero information';

CREATE TABLE `item` (
  
  `item_id` int(11) NOT NULL AUTO_INCREMENT,
  `hero_id` int(11) DEFAULT NULL,
  `weapon` 	int(11) DEFAULT NULL,
  `tunic` 	int(11) DEFAULT NULL,
  `shield` 	int(11) DEFAULT NULL,
  `leggings` int(11) DEFAULT NULL,
  `ring` 	int(11) DEFAULT NULL,
  `gloves` 	int(11) DEFAULT NULL,
  `boots` 	int(11) DEFAULT NULL,
  `energy` 	int(11) DEFAULT NULL,
  `helm` 	int(11) DEFAULT NULL,
  `charm` 	int(11) DEFAULT NULL,
  `Amulet` 	int(11) DEFAULT NULL,
  `Total` 	int(11) DEFAULT NULL,
  
  PRIMARY KEY (`item_id`),
  KEY `item_hero_hero_id_fk` (`hero_id`),
  
  CONSTRAINT `item_hero_hero_id_fk` FOREIGN KEY (`hero_id`) REFERENCES `hero` (`hero_id`)
  
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COMMENT='Holds  the items owned by the hero';

CREATE TABLE `penalty` (

  `penalty_id` 	int(11) NOT NULL AUTO_INCREMENT,
  `hero_id` 	int(11) DEFAULT NULL,
  `logout` 		int(11) DEFAULT NULL,
  `quit` 		int(11) DEFAULT NULL,
  `message` 	int(11) DEFAULT NULL,
  `quest` 		int(11) DEFAULT NULL,
  
  PRIMARY KEY (`penalty_id`),
  KEY `penalty_hero_hero_id_fk` (`hero_id`),
  
  CONSTRAINT `penalty_hero_hero_id_fk` FOREIGN KEY (`hero_id`) REFERENCES `hero` (`hero_id`)
  
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COMMENT='Penalties table stores the penalties accumulated by a hero';

CREATE TABLE `worldevent` (

  `worldevent_id` 	int(11) NOT NULL AUTO_INCREMENT,
  `event_text` 		text,
  `event_time` 		DATETIME DEFAULT NOW(),
  
  PRIMARY KEY (`worldevent_id`)
  
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COMMENT='WorldEvent table contains the events happening in the world ';


CREATE TABLE `heroworldevent` (

  `heroworldevent_id` 	int(11) NOT NULL AUTO_INCREMENT,
  `hero_id` 			int(11) DEFAULT NULL,
  `worldevent_id` 		int(11) DEFAULT NULL,
  
  PRIMARY KEY (`heroworldevent_id`),
  KEY `heroworldevent_hero_hero_id_fk` (`hero_id`),
  KEY `heroworldevent_worldevent_worldevent_id_fk` (`worldevent_id`),
  
  CONSTRAINT `heroworldevent_hero_hero_id_fk` FOREIGN KEY (`hero_id`) REFERENCES `hero` (`hero_id`),
  CONSTRAINT `heroworldevent_worldevent_worldevent_id_fk` FOREIGN KEY (`worldevent_id`) REFERENCES `worldevent` (`worldevent_id`)
  
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COMMENT='HeroWorldEvent is a Junction table that maps heros to world events ';


# Functions

DROP FUNCTION IF EXISTS randomizer;

DELIMITER $$
CREATE DEFINER=`titanuser`@`%` FUNCTION `randomizer`(
    pmin INTEGER,
    pmax INTEGER
) RETURNS int(11)
    NO SQL
    DETERMINISTIC
RETURN floor(pmin+RAND()*(pmax-pmin))$$
DELIMITER ;

DROP function if exists generate_fname;

DELIMITER $$
CREATE FUNCTION generate_fname () RETURNS varchar(255)
BEGIN
	RETURN ELT(FLOOR(1 + (RAND() * (100-1))), "James","Mary","John","Patricia","Robert","Linda","Michael","Barbara","William","Elizabeth","David","Jennifer","Richard","Maria","Charles","Susan","Joseph","Margaret","Thomas","Dorothy","Christopher","Lisa","Daniel","Nancy","Paul","Karen","Mark","Betty","Donald","Helen","George","Sandra","Kenneth","Donna","Steven","Carol","Edward","Ruth","Brian","Sharon","Ronald","Michelle","Anthony","Laura","Kevin","Sarah","Jason","Kimberly","Matthew","Deborah","Gary","Jessica","Timothy","Shirley","Jose","Cynthia","Larry","Angela","Jeffrey","Melissa","Frank","Brenda","Scott","Amy","Eric","Anna","Stephen","Rebecca","Andrew","Virginia","Raymond","Kathleen","Gregory","Pamela","Joshua","Martha","Jerry","Debra","Dennis","Amanda","Walter","Stephanie","Patrick","Carolyn","Peter","Christine","Harold","Marie","Douglas","Janet","Henry","Catherine","Carl","Frances","Arthur","Ann","Ryan","Joyce","Roger","Diane");
END$$

DELIMITER ;

DROP function if exists generate_lname;
DELIMITER $$
CREATE FUNCTION generate_lname () RETURNS varchar(255)
BEGIN
	RETURN ELT(FLOOR(1 + (RAND() * (100-1))), "Smith","Johnson","Williams","Jones","Brown","Davis","Miller","Wilson","Moore","Taylor","Anderson","Thomas","Jackson","White","Harris","Martin","Thompson","Garcia","Martinez","Robinson","Clark","Rodriguez","Lewis","Lee","Walker","Hall","Allen","Young","Hernandez","King","Wright","Lopez","Hill","Scott","Green","Adams","Baker","Gonzalez","Nelson","Carter","Mitchell","Perez","Roberts","Turner","Phillips","Campbell","Parker","Evans","Edwards","Collins","Stewart","Sanchez","Morris","Rogers","Reed","Cook","Morgan","Bell","Murphy","Bailey","Rivera","Cooper","Richardson","Cox","Howard","Ward","Torres","Peterson","Gray","Ramirez","James","Watson","Brooks","Kelly","Sanders","Price","Bennett","Wood","Barnes","Ross","Henderson","Coleman","Jenkins","Perry","Powell","Long","Patterson","Hughes","Flores","Washington","Butler","Simmons","Foster","Gonzales","Bryant","Alexander","Russell","Griffin","Diaz","Hayes");
END$$
DELIMITER ;



DROP function if exists generate_race;
DELIMITER $$
CREATE FUNCTION generate_race () RETURNS varchar(255)
BEGIN
	RETURN ELT(FLOOR(1 + (RAND() * (27-1))), "Elf","Dwarf","Goblin","Morlock","Half-Gnome","Hobbit","Gnome","Orc","Shade","White Walker","Hobbgoblin","Barbarian","Half-elf","Hafling","Half-Orc","Half-Giant","Giant","Drow","Half-Drow","Minotaur","Human","Dragonborn","Eladrin","Ogre","Half-Ogre","Vampire","Troll");
END$$
DELIMITER ;

DROP function if exists generate_class;
DELIMITER $$
CREATE FUNCTION generate_class () RETURNS varchar(255)
BEGIN
	RETURN ELT(FLOOR(1 + (RAND() * (19-1))), "Barbarian","Bard","Cleric","Druid","Figter","Mage","Wizard","Monk","Mistic","Paladin","Ranger","Sorcerer","Thief","Ninja","Warlock","Defiler","Shaman","Invoker","Psion");
END$$
DELIMITER ;


# Random Fantasy Title Generator by Jellyn
# view-source:http://core.binghamton.edu/~jandrews/random/randomtitles.html

DROP function if exists generate_title;
DELIMITER $$
CREATE FUNCTION generate_title () RETURNS varchar(255)
BEGIN
	SET @noum =  ELT(FLOOR(1 + (RAND() * (337-1))), "Dream", "Dreamer", "Dreams","Rainbow",
													"Dreaming","Flight","Wings","Mist",
													"Sky","Wind","Winter","Misty",
													"Cloud","Fairy","Dragon","End",
													"Beginning","Tale","Tales","Emperor",
													"Prince","Princess","Willow","Birch","Petals",
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
													"Thrusting","Wraith","Phantom","Sheep");

    SET @adjective = ELT(FLOOR(1 + (RAND() * (122-1))), "Misty","Lost","Only","Last","First",
														"Final","Missing","Shadowy","Seventh",
														"Dark","Darkest","Silver","Silvery",
														"Black","White","Hidden","Entwined","Invisible",
														"Next","Seventh","Red","Green","Blue",
														"Purple","Grey","Bloody","Emerald","Diamond",
														"Frozen","Sharp","Delicious","Dangerous",
														"Deep","Twinkling","Dwindling","Missing","Absent",
														"Vacant","Cold","Hot","Burning","Forgotten",
														"No","All","Which","What","Hard","Soft",
														"Playful","Final","Evil","Scarlet","Chaste","Virgin",
														"Strange","Silent", "Legendary","Golden","Magic",
														"Mystic","Majestic","Magical","Mysterious","Eternal",
														"Winged","Outer","Inner","Silken","Mystical",
														"Crying","Weeping","Lonely","Crushed","Searching",
														"Desperate","Yearning","Quick","Invincible","New",
														"Old","Ancient","Aging","Dying","Living","Vengeful",
														"Loving","Crystal","Crystalline","Wooden","Metal",
														"Metallic","Marble","Stony","Rocky","Great","Royal",
														"Noble","Wet","Dry","Bleeding","Piercing","Singing",
														"Dancing","Painful","Wandering","Loyal","Trusting",
														"Open","Closed","Locked","Free","Chained","Caged",
														"Empty","Wilted","Lunar","Solar","Screaming","Dead",
														"Shaking","Thrusting","Frantic");
	RETURN CONCAT('The ', @adjective, ' of ', @noum);
	
END$$
DELIMITER ;






