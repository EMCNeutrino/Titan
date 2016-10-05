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

DELIMITER $$
CREATE DEFINER=`titanuser`@`%` FUNCTION `randomizer`(
    pmin INTEGER,
    pmax INTEGER
) RETURNS int(11)
    NO SQL
    DETERMINISTIC
RETURN floor(pmin+RAND()*(pmax-pmin))$$
DELIMITER ;














