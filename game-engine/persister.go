package main

import (
	"database/sql"
	"time"

	log "github.com/Sirupsen/logrus"

	_ "github.com/go-sql-driver/mysql"
	"encoding/json"

)

func getDBConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", "titanuser:Neutrin0R0cks!@/titandb")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

func SaveToDB(g *Game) error {

	log.Info("[Persister] SaveToDB Called ------------------------------------------------------- ")
	db, err := getDBConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	for _, hero := range g.heroes {
		stmt, err := db.Prepare("INSERT INTO hero " +
			"(hero_name, email, hclass, hero_online, token, isAdmin, hero_level, ttl, xpos, ypos, " +
			" ring, amulet, charm, weapon, helm, tunic, gloves, shield, leggings, boots " +
			") " +
			"VALUES( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? ) " +
			"ON DUPLICATE KEY UPDATE " +
			"hero_online=VALUES(hero_online), hero_level=VALUES(hero_level), ttl=VALUES(ttl), xpos=VALUES(xpos), ypos=VALUES(ypos), " +
			"ring=VALUES(ring), amulet=VALUES(amulet), charm=VALUES(charm), weapon=VALUES(weapon), " +
			"helm=VALUES(helm), tunic=VALUES(tunic), gloves=VALUES(gloves), shield=VALUES(shield), " +
			"leggings=VALUES(leggings), boots=VALUES(boots);")
		if err != nil {
			log.Error(err)
		}
		ttl := int(hero.NextLevelAt.Sub(time.Now()).Seconds())

		_, err = stmt.Exec(hero.HeroName, hero.Email, hero.HClass, hero.Enabled, hero.Token, false, hero.Level, ttl, hero.Xpos, hero.Ypos,
			hero.Ring, hero.Amulet, hero.Charm, hero.Weapon, hero.Helm, hero.Tunic, hero.Gloves, hero.Shield, hero.Leggings, hero.Boots)
		if err != nil {
			log.Error(err)
		}

	}

	return nil
}

func LoadFromDB() (*Game, error) {

	log.Info("[Persister] LoadFromDB Called ------------------------------------------------------- ")

	db, err := getDBConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	game := &Game{
		startedAt:        time.Now(),
		heroes:           []HeroDB{},
		joinChan:         make(chan JoinRequest),
		activateHeroChan: make(chan ActivateHeroRequest),
		exitChan:         make(chan []byte),
		adminToken:       "1234",
	}

	rows, err := db.Query("SELECT hero_id, COALESCE(hero_name, '') AS hero_name, COALESCE(player_name, '') AS player_name," +
		"COALESCE(player_lastname, '') AS player_lastname, " +
		"COALESCE(token, '') AS token, " +
		"COALESCE(twitter, '') AS twiter, " +
		"COALESCE(email, 'NoEmail') AS email, " +
		"COALESCE(title, '') AS title, " +
		"COALESCE(race, '') AS race, " +
		"isadmin, hero_level,  " +
		"COALESCE(hclass, '') AS hclass , ttl, " +
		"COALESCE(userhost, '') AS userhost, hero_online, xpos, ypos, " +
		"IFNULL(weapon, 0), IFNULL(tunic, 0), IFNULL(shield, 0), IFNULL(leggings, 0), IFNULL(ring, 0), " +
		"IFNULL(gloves, 0), IFNULL(boots, 0), IFNULL(helm, 0), IFNULL(charm, 0) , IFNULL(amulet, 0) " +
		"total_equipment FROM hero")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		hero := &HeroDB{}
		var ttl int

		err := rows.Scan(&hero.HeroID, &hero.HeroName, &hero.UserName, &hero.UserLastName, &hero.Token, &hero.Twitter, &hero.Email,
			&hero.Title, &hero.HRace, &hero.IsAdmin, &hero.Level, &hero.HClass, &ttl, &hero.Userhost, &hero.Enabled,
			&hero.Xpos, &hero.Ypos, &hero.Weapon, &hero.Tunic, &hero.Shield, &hero.Leggings, &hero.Ring, &hero.Gloves,
			&hero.Boots, &hero.Helm, &hero.Charm, &hero.Amulet)

		if err != nil {
			log.Error(err)
		}

		hero.TotalEquipment = hero.Weapon + hero.Tunic + hero.Shield + hero.Leggings + hero.Ring + hero.Gloves + hero.Boots + hero.Helm + hero.Charm + hero.Amulet
		hero.NextLevelAt = time.Now().Add(time.Duration(ttl) * time.Second)

		if(hero.HeroName != "") {  //Fixes the extra record with empty information.

			game.heroes = append(game.heroes, *hero)

			hero2json, _ := json.Marshal(hero)
			log.Info(string(hero2json))

			Hero_Joined_World_Notification(hero, db)
		}

	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return game, nil
}

func Hero_Joined_World_Notification(hero *HeroDB, conn *sql.DB) {

	log.Info("[Persister] Hero_Joined_World_Notification Called ------------------------------------------------------- ")


	var message = hero.HeroName + ", " + hero.Title + ", is now online from " + hero.UserName +
		" " + hero.UserLastName + "(" + hero.Twitter + "). Next Level in " + hero.NextLevelAt.String()

	log.Info(message)

	Insert_World_Event_for_Hero(hero.HeroID, message, conn)



}
