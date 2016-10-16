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

	log.Info("Inserting World Event: HeroID: " + strconv.FormatInt(heroID, 10) + " | Event: " + worldEvent)

	var query = "INSERT INTO worldevent (event_text) VALUES ('" + worldEvent + "')"

	log.Info("[DAL] World Event Query: " + query)

	statement, err := conn.Exec(query)

	last_worldevent_Id, err := statement.LastInsertId()

	if err != nil {
		log.Errorln("[DAL] Insert_Item_for_Hero: Worldevent Insert failed: %s", err)
	}

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
