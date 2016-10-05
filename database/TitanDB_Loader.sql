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
** 2    10/04/2016 PBM       Updated with Stored Procedure Generator
*******************************/

DELETE FROM  item;
DELETE FROM  penalty;
DELETE FROM  hero;

ALTER TABLE item AUTO_INCREMENT = 1;
ALTER TABLE penalty AUTO_INCREMENT = 1;
ALTER TABLE hero AUTO_INCREMENT = 1;

# Change the number in the procedure to N heros
Call create_heros(10);


SELECT * FROM hero;
SELECT * FROM penalty;
SELECT * FROM item;
