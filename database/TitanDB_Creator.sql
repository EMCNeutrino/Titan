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
*******************************/

/* Creating the TitanDB */
CREATE DATABASE IF NOT EXISTS titandb;

/* Create User and assign Permisions */
CREATE USER IF NOT EXISTS 'titanuser'@localhost IDENTIFIED BY 'Neutrin0R0cks!';
GRANT ALL ON titandb.* TO 'titanuser' IDENTIFIED BY 'Neutrin0R0cks!';

USE titandb;
GO

/* Drop Tables */
DROP TABLE IF EXISTS `item`;
DROP TABLE IF EXISTS `penalty`;
DROP TABLE IF EXISTS `hero`;

/* Drop Functions */
DROP FUNCTION IF EXISTS `randomizer`;

CREATE TABLE hero
(
    id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    hero_name TEXT,
    token TEXT,
    userpass TEXT,
    energy INT,
    twitter TEXT,
    email TEXT,
    title TEXT,
    race TEXT,
    is_admin BOOL,
    level INT,
    class TEXT,
    ttl INT,
    i_online BOOL,
    xpos INT,
    ypos INT
);

CREATE TABLE item
(
    id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    hero_id INT,
    weapon INT,
    tunic INT,
    shield INT,
    leggings INT,
    ring INT,
    gloves INT,
    boots INT,
    energy INT,
    helm INT,
    charm INT,
    Amulet INT,
    Total INT,
    CONSTRAINT item_hero_hero_id_fk FOREIGN KEY (hero_id) REFERENCES hero (id)
);


CREATE TABLE penalty
(
    penalty_id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    hero_id INT,
    logout INT,
    quit INT,
    message INT,
    quest INT,
    CONSTRAINT penalty_hero_hero_id_fk FOREIGN KEY (hero_id) REFERENCES hero (id)
);


/* Add comments to each Table */
ALTER TABLE item COMMENT = 'Holds the items owned by the hero';
ALTER TABLE penalty COMMENT = 'Penalties table stores the penalties accumulated by a hero';
ALTER TABLE hero COMMENT = 'Contains the Hero information';


/* Functions */


/* Random Function */
/* http://stackoverflow.com/questions/14798640/creating-a-random-number-using-mysql */
CREATE FUNCTION randomizer(
    pmin INTEGER,
    pmax INTEGER
)
RETURNS INTEGER(11)
DETERMINISTIC
NO SQL
SQL SECURITY DEFINER
RETURN floor(pmin+RAND()*(pmax-pmin));
