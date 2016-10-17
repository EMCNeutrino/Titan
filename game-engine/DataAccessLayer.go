package main

import (

	"database/sql"
	"strconv"
	log "github.com/Sirupsen/logrus"
)


//Insert World Event for Hero, adds a world event notification for a specific hero
//Parameters: Hero ID, World Event Message and the SQL.DB connection.
//Return: None
func Insert_World_Event_for_Hero(heroID int64, worldEvent string, conn *sql.DB) {

	log.Info("[DAL] Insert_World_Event_for_Hero : HeroID: " + strconv.FormatInt(heroID, 10) + " | Event: " + worldEvent)

	var query = "INSERT INTO worldevent (event_text) VALUES ('" + worldEvent + "')"

	log.Info("[DAL] World Event Query: " + query)

	statement, err := conn.Exec(query)

	if err != nil {
		log.Errorln("[DAL] Insert_Item_for_Hero: Worldevent Insert failed: %s", err)
	}

	last_worldevent_Id, err := statement.LastInsertId()

	if last_worldevent_Id != 0 {

		Insert_Hero_World_Event(heroID, int(last_worldevent_Id), conn)
	}
}

//Insert Hero World Event, populates the table that maps Heros with their World events
//Parameters: Hero ID, World Event ID and the SQL.DB connection.
//Return: None
func Insert_Hero_World_Event(heroID int64, woldEvent_id int, conn *sql.DB) {

	var query = "INSERT INTO heroworldevent (hero_id, worldevent_id ) VALUES (" + strconv.FormatInt(heroID, 10) + ", " + strconv.Itoa(woldEvent_id) + ")"

	log.Info("[DAL] Heroworldevent Query: " + query)

	_, err := conn.Exec(query)

	if err != nil {
		log.Errorln("[DAL] Insert_Item_for_Hero: Heroworldevent Insert failed: %s", err)
	}
}

//Get Item by Hero ID, retries an item's level value for the specified Hero ID
//Parameters: Hero ID, Type of Item, and the SQL.DB connection.
//Return: The Item level value
func Get_Item_By_HeroID(heroID int64, item_type string, conn *sql.DB) int {

	log.Info("Get_item_level_for_user ")

	var current_item_level int

	//Check what is the level of the current Item, update value if needed, msg player
	var query = "SELECT " + item_type + " FROM hero WHERE hero_id = ?"

	log.Info("Select Query: " + query)

	stmt, err := conn.Prepare(query)
	if err != nil {
		log.Errorln("DB: Get_Item_By_HeroID: Prepare Query failed: %s", err)
	}

	err = stmt.QueryRow(heroID).Scan(&current_item_level)

	if err != nil {
		log.Errorln("DB: Get_Item_By_HeroID: QueryRow failed: %s", err)
	}

	log.Info("Item Value: " + strconv.Itoa(current_item_level))

	defer stmt.Close()

	return current_item_level
}

//Update Item for Hero, updates the Item level value for a specified item for a  Hero (using the Hero ID)
//Parameters: Hero ID, Type of Item, The new item level and the SQL.DB connection.
//Return: None
func Update_Item_for_Hero(heroID int64, item_type string, item_level int, conn *sql.DB) {

	log.Info("Updating Item: " + item_type + " with Level: " + strconv.Itoa(item_level) +" from Hero ID: " + strconv.FormatInt(heroID, 10))

	var query = "UPDATE hero SET " + item_type + "=" + strconv.Itoa(item_level) +" WHERE hero_id="+ strconv.FormatInt(heroID, 10)

	_, err := conn.Exec(query)

	if err != nil {
		log.Errorln("DB: Insert_Item_for_Hero: Item Insert failed: %s", err)
	}


	log.Info("Executed: " + query)
}