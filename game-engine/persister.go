package main

import (
  "database/sql"
  "strconv"
  "time"

  log "github.com/Sirupsen/logrus"

  _ "github.com/go-sql-driver/mysql"
)

// GetDBConnection builds and returns the database connection
func GetDBConnection(databaseURL string) (*sql.DB, error) {

  db, err := sql.Open("mysql", databaseURL)
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

// SaveToDB persists the Heros in the Database
func SaveToDB(g *Game) error {

  db, err := GetDBConnection(g.databaseURL)
  if err != nil {
    return err
  }
  defer db.Close()

  for _, hero := range g.heroes {
    stmt, err := db.Prepare("INSERT INTO hero " +
      "(player_name, player_lastname, hero_name, email, twitter, hclass, hero_online, token, isAdmin, hero_level, ttl, xpos, ypos, " +
      " ring, amulet, charm, weapon, helm, tunic, gloves, shield, leggings, boots " +
      ") " +
      "VALUES( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? ) " +
      "ON DUPLICATE KEY UPDATE " +
      "hero_online=VALUES(hero_online), hero_level=VALUES(hero_level), ttl=VALUES(ttl), xpos=VALUES(xpos), ypos=VALUES(ypos), " +
      "ring=VALUES(ring), amulet=VALUES(amulet), charm=VALUES(charm), weapon=VALUES(weapon), " +
      "helm=VALUES(helm), tunic=VALUES(tunic), gloves=VALUES(gloves), shield=VALUES(shield), " +
      "leggings=VALUES(leggings), boots=VALUES(boots);")
    if err != nil {
      log.Error(err)
    }
    ttl := int(hero.NextLevelAt.Sub(time.Now()).Seconds())

    _, err = stmt.Exec(hero.FirstName, hero.LastName, hero.HeroName, hero.Email, hero.Twitter, hero.HeroClass, hero.Enabled, hero.Token,
      false, hero.Level, ttl, hero.Xpos, hero.Ypos,
      hero.Ring, hero.Amulet, hero.Charm, hero.Weapon, hero.Helm, hero.Tunic, hero.Gloves, hero.Shield, hero.Leggings, hero.Boots)
    if err != nil {
      log.Error(err)
    }
  }

  return nil
}

// LoadFromDB loads the Heros in the hero table and adds them to the realm
func LoadFromDB(g *Game) error {

  db, err := GetDBConnection(g.databaseURL)
  if err != nil {
    return err
  }
  defer db.Close()

  rows, err := db.Query("SELECT " +
    "hero_id, " +
    "COALESCE(hero_name, '') AS hero_name, " +
    "COALESCE(player_name, '') AS player_name," +
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
    return err
  }
  defer rows.Close()

  for rows.Next() {
    hero := &Hero{}
    var ttl int

    err = rows.Scan(&hero.HeroID, &hero.HeroName, &hero.FirstName, &hero.LastName, &hero.Token, &hero.Twitter, &hero.Email,
      &hero.Title, &hero.HRace, &hero.IsAdmin, &hero.Level, &hero.HeroClass, &ttl, &hero.Userhost, &hero.Enabled,
      &hero.Xpos, &hero.Ypos, &hero.Weapon, &hero.Tunic, &hero.Shield, &hero.Leggings, &hero.Ring, &hero.Gloves,
      &hero.Boots, &hero.Helm, &hero.Charm, &hero.Amulet)

    if err != nil {
      log.Error(err)
      continue
    }

    hero.TotalEquipment = hero.Weapon + hero.Tunic + hero.Shield + hero.Leggings + hero.Ring + hero.Gloves + hero.Boots + hero.Helm + hero.Charm + hero.Amulet
    hero.NextLevelAt = time.Now().Add(time.Duration(ttl) * time.Second)
    g.heroes = append(g.heroes, *hero)

  }
  err = rows.Err()
  if err != nil {
    return err
  }

  return nil
}

// SaveEventToDB adds a world event for a specific hero
func (g *Game) saveEventToDB(message string, heroes []*Hero) error {
  db, err := GetDBConnection(g.databaseURL)
  if err != nil {
    return err
  }
  defer db.Close()

  tx, err := db.Begin()
  if err != nil {
    return err
  }
  defer func() {
    if err != nil {
      tx.Rollback()
      return
    }
    err = tx.Commit()
    if err != nil {
      tx.Rollback()
      return
    }
  }()

  r, err := tx.Exec("INSERT INTO worldevent (event_text) VALUES (?)", message)
  if err != nil {
    return err
  }

  eventID, err := r.LastInsertId()
  if err != nil {
    return err
  }

  for _, hero := range heroes {
    if _, err = tx.Exec("INSERT INTO heroworldevent (hero_id, worldevent_id ) VALUES (?, ?)", hero.HeroID, eventID); err != nil {
      return err
    }
  }

  return nil
}

// Get Item by Hero ID, retries an item's level value for the specified Hero ID
//Parameters: Hero ID, Type of Item, and the SQL.DB connection.
//Return: The Item level value
func Get_Item_By_HeroID(heroID int64, item_type string, conn *sql.DB) int {

  log.Info("Get_item_level_for_user ")

  var current_item_level int

  //Check what is the level of the current Item, update value if needed, msg player
  var query = "SELECT " + item_type + " FROM hero WHERE hero_id=?"

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
