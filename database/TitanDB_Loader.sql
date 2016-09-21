/******************************
** File: TitanDB_Loader.sql
** Name: Titan Heros Loader
** Desc: Loads the Database with Heros
** Auth: Patrick Butler Monterde
** Date: 09/19/2016
**************************
** Change History
**************************
** PR   Date	     Author    Description
** --   --------   -------   ------------------------------------
** 1    09/19/2016 PBM       Created
*******************************/

DELETE FROM  item;
DELETE FROM  penalty;
DELETE FROM  hero;

ALTER TABLE item AUTO_INCREMENT = 1;
ALTER TABLE penalty AUTO_INCREMENT = 1;
ALTER TABLE hero AUTO_INCREMENT = 1;

SET @heroid = NULL;
SET @total_items = NULL;

# Generates Random Number

#Inserting Gandalf

INSERT INTO hero (


    hero_name,
    player_name,
    player_lastname,
    token,
    userpass,
    energy,
    twitter,
    email,
    title,
    race,
    isAdmin,
    hero_level,
    hclass,
    ttl,
    userhost,
    online,
    xpos,
    ypos
) VALUES (
    'Gandalf',
    'Pat',
    'Butler',
    '123123123123',
    'userPass',
    200,
    '@twitter Handle',
    'megahero@emc.com',
    'The Gray',
    'Human',
    0,
    30,
    'Wizard',
    10,
    'atsome.isp.com',
    1,
    100,
    101

);

# Getting Gandalf HeroID to build up the Tables Item and Penalty tables

SET @heroid = (SELECT hero_id FROM hero WHERE hero_name = 'Gandalf');

SELECT @heroid AS 'Hero ID';


INSERT INTO penalty (

    hero_id,
    logout,
    quit,
    message,
    quest

) VALUES (

    @heroid,
    randomizer(1,100),
    randomizer(1,100),
    randomizer(1,100),
    randomizer(1,100)

);


INSERT INTO item (

    hero_id,
    weapon,
    tunic,
    shield,
    leggings,
    ring,
    gloves,
    boots,
    energy,
    helm,
    charm,
    Amulet,
    Total

) VALUES (

    @heroid,
    randomizer(1,100),
    randomizer(1,100),
    randomizer(1,100),
    randomizer(1,100),
    randomizer(1,100),
    randomizer(1,100),
    randomizer(1,100),
    randomizer(1,100),
    randomizer(1,100),
    randomizer(1,100),
    randomizer(1,100),
    0
);

# Calculate Total Items for Gandalf
SET @total_items = (SELECT  weapon + tunic + shield +
                           leggings + ring + gloves + boots +
                           energy + helm + charm + Amulet
                     FROM item WHERE hero_id = @heroid);

UPDATE item SET Total = @total_items WHERE hero_id = @heroid;



# Inserting Galadrial
SET @heroid = NULL;
SET @total_items = NULL;

INSERT INTO hero (


    hero_name,
    player_name,
    player_lastname,
    token,
    userpass,
    energy,
    twitter,
    email,
    title,
    race,
    isAdmin,
    hero_level,
    hclass,
    ttl,
    userhost,
    online,
    xpos,
    ypos
) VALUES (
    'Galadrial',
    'Concha',
    'Velasco',
    '43534534534534',
    'userPass2',
    220,
    '@twitter Handle2',
    'megahero2@emc.com',
    'Elf Queen',
    'High Elf',
    0,
    50,
    'Healer',
    20,
    'atsome.isp.com',
    1,
    103,
    105

);

# Getting Galadrial HeroID to build up the Tables Item and Penalty tables

SET @heroid = (SELECT hero_id FROM hero WHERE hero_name = 'Galadrial');

SELECT @heroid AS 'Hero ID';


INSERT INTO penalty (

    hero_id,
    logout,
    quit,
    message,
    quest

) VALUES (

    @heroid,
    randomizer(1,100),
    randomizer(1,100),
    randomizer(1,100),
    randomizer(1,100)

);


INSERT INTO item (

    hero_id,
    weapon,
    tunic,
    shield,
    leggings,
    ring,
    gloves,
    boots,
    energy,
    helm,
    charm,
    Amulet,
    Total

) VALUES (

    @heroid,
    randomizer(1,100),
    randomizer(1,100),
    randomizer(1,100),
    randomizer(1,100),
    randomizer(1,100),
    randomizer(1,100),
    randomizer(1,100),
    randomizer(1,100),
    randomizer(1,100),
    randomizer(1,100),
    randomizer(1,100),
    0
);

# Calculate Total Items for Hero
SET @total_items = (SELECT  weapon + tunic + shield +
                           leggings + ring + gloves + boots +
                           energy + helm + charm + Amulet
                     FROM item WHERE hero_id = @heroid);

UPDATE item SET Total = @total_items WHERE hero_id = @heroid;


# ------------------------------


#Insering Legolas

SET @heroid = NULL;
SET @total_items = NULL;

INSERT INTO hero (


    hero_name,
    player_name,
    player_lastname,
    token,
    userpass,
    energy,
    twitter,
    email,
    title,
    race,
    isAdmin,
    hero_level,
    hclass,
    ttl,
    userhost,
    online,
    xpos,
    ypos
) VALUES (
    'Legolas',
    'Magdy',
    'The Great',
    '4353453232334234534534',
    'userPass3',
    250,
    '@twitter Handle3',
    'megahero2@emc.com',
    'The Green',
    'High Elf',
    0,
    70,
    'Ranger',
    25,
    'atsome.isp.com',
    1,
    112,
    115

);

# Getting the HeroID to build up the Tables Item and Penalty tables

SET @heroid = (SELECT hero_id FROM hero WHERE hero_name = 'Legolas');

SELECT @heroid AS 'Hero ID';


INSERT INTO penalty (

    hero_id,
    logout,
    quit,
    message,
    quest

) VALUES (

    @heroid,
    randomizer(1,100),
    randomizer(1,100),
    randomizer(1,100),
    randomizer(1,100)

);


INSERT INTO item (

    hero_id,
    weapon,
    tunic,
    shield,
    leggings,
    ring,
    gloves,
    boots,
    energy,
    helm,
    charm,
    Amulet,
    Total

) VALUES (

    @heroid,
    randomizer(1,100),
    randomizer(1,100),
    randomizer(1,100),
    randomizer(1,100),
    randomizer(1,100),
    randomizer(1,100),
    randomizer(1,100),
    randomizer(1,100),
    randomizer(1,100),
    randomizer(1,100),
    randomizer(1,100),
    0
);

# Calculate Total Items for Hero
SET @total_items = (SELECT  weapon + tunic + shield +
                           leggings + ring + gloves + boots +
                           energy + helm + charm + Amulet
                     FROM item WHERE hero_id = @heroid);

UPDATE item SET Total = @total_items WHERE hero_id = @heroid;




Select * from hero;
select * from penalty;
SELECT * from item;
select @total_items;



