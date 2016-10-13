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
DROP TABLE IF EXISTS `equipment`;
DROP TABLE IF EXISTS `hero`;

/* Drop Functions */
DROP FUNCTION IF EXISTS `randomizer`;

CREATE TABLE hero
(
    id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    class VARCHAR(255),
    enabled BOOL DEFAULT 0,
    token VARCHAR(255),
    is_admin BOOL DEFAULT 0,
    level INT DEFAULT 0,
    ttl INT,
    xpos INT,
    ypos INT,
    UNIQUE (email)
);

CREATE TABLE equipment
(
    hero_id INT PRIMARY KEY NOT NULL,
    ring INT DEFAULT 0,
    amulet INT DEFAULT 0,
    charm INT DEFAULT 0,
    weapon INT DEFAULT 0,
    helm INT DEFAULT 0,
    tunic INT DEFAULT 0,
    gloves INT DEFAULT 0,
    shield INT DEFAULT 0,
    leggings INT DEFAULT 0,
    boots INT DEFAULT 0,
    CONSTRAINT item_hero_hero_id_fk FOREIGN KEY (hero_id) REFERENCES hero (id)
);

/* Add comments to each Table */
ALTER TABLE hero COMMENT = 'Contains the Hero information';
ALTER TABLE equipment COMMENT = 'Holds the items owned by the hero';


/* Functions */

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
