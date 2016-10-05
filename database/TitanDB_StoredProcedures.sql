CREATE DEFINER=`titanuser`@`%` PROCEDURE `hero_add`()
BEGIN

	#Set Hero Variables

	SET @hero_name = (SELECT CONCAT(generate_fname(),' ', generate_lname()) as hero_name);
	SET @player_name = (generate_fname());
	SET @player_lastname = (generate_lname());
	SET @email = (CONCAT(@player_name, '.', @player_lastname,'@dell.com'));
	SET @twitter = (CONCAT('@', @player_name, '.', @player_lastname));
	SET @token = (SELECT UUID());
	SET @userpass = (SELECT UUID());
	SET @hero_level = (randomizer(1,10));
	SET @next_level_calc = (SELECT ROUND(600*(POW(1.16, @hero_level+1))) AS Level_Time);
	SET @next_level_calc2 = (SEC_TO_TIME((SELECT ROUND(600*(POW(1.16, @hero_level))) AS Level_Time)));
	SET @next_level  = (SELECT ADDTIME(NOW(), @next_level_calc2));

#SELECT @hero_level, @next_level_calc, @next_level_calc2, (Now()), @next_level;

INSERT INTO `titandb`.`hero`
(
	`hero_name`,
	`player_name`,
	`player_lastname`,
	`token`,
	`userpass`,
	`energy`,
	`twitter`,
	`email`,
	`title`,
	`race`,
	`isAdmin`,
	`hero_level`,
	`hclass`,
	`ttl`,
	`userhost`,
	`hero_online`,
	`xpos`,
	`ypos`,
	`next_level`

)

VALUES

(
	@hero_name,
	@player_name,
	@player_lastname,
	@token,
	@userpass,
	randomizer(1,100),
	@twitter,
	@email,
	generate_title(),
	generate_race(),
	0,
	@hero_level,
	generate_class(),
	@next_level_calc,
	0,
	1,
	randomizer(1,100),
	randomizer(1,100),
	@next_level

);

SET @heroid = NULL;
SET @total_items = NULL;

# Getting HeroID to build up the Item and Penalty tables

SET @heroid = (SELECT hero_id FROM hero WHERE token = @token);

#Insert into Penalty Table

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

#insert into the Item Table

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



END