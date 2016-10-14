package main

import (
  "database/sql"
  "time"

  log "github.com/Sirupsen/logrus"

  _ "github.com/go-sql-driver/mysql"
)

func getDBConnection() (*sql.DB, error) {
  db, err := sql.Open("mysql", "root:root@/titandb")
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

func saveToDB(g *Game) error {
  db, err := getDBConnection()
  if err != nil {
    return err
  }
  defer db.Close()

  for _, hero := range g.heroes {
    stmt, err := db.Prepare("INSERT INTO hero " +
      "(name, email, class, enabled, token, is_admin, level, ttl, xpos, ypos) " +
      "VALUES( ?, ?, ?, ?, ?, ?, ?, ?, ?, ? ) " +
      "ON DUPLICATE KEY UPDATE " +
      "enabled=VALUES(enabled), level=VALUES(level), ttl=VALUES(ttl), xpos=VALUES(xpos), ypos=VALUES(ypos);")
    if err != nil {
      log.Error(err)
    }
    ttl := int(hero.nextLevelAt.Sub(time.Now()).Seconds())
    res, err := stmt.Exec(hero.Name, hero.Email, hero.Class, hero.Enabled, hero.token, false, hero.Level, ttl, hero.Xpos, hero.Ypos)
    if err != nil {
      log.Error(err)
    }
    lastID, err := res.LastInsertId()
    if err != nil {
      log.Error(err)
    }

    // Update Equipment
    stmt, err = db.Prepare("INSERT INTO equipment " +
      "(hero_id, ring, amulet, charm, weapon, helm, tunic, gloves, shield, leggings, boots) " +
      "VALUES( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? ) " +
      "ON DUPLICATE KEY UPDATE " +
      "ring=VALUES(ring), amulet=VALUES(amulet), charm=VALUES(charm), weapon=VALUES(weapon), " +
      "helm=VALUES(helm), tunic=VALUES(tunic), gloves=VALUES(gloves), shield=VALUES(shield), " +
      "leggings=VALUES(leggings), boots=VALUES(boots);")
    if err != nil {
      log.Error(err)
    }

    res, err = stmt.Exec(lastID, hero.Equipment.Ring, hero.Equipment.Amulet, hero.Equipment.Charm,
      hero.Equipment.Weapon, hero.Equipment.Helm, hero.Equipment.Tunic, hero.Equipment.Gloves,
      hero.Equipment.Shield, hero.Equipment.Leggings, hero.Equipment.Boots)
    if err != nil {
      log.Error(err)
    }
  }

  return nil
}

func loadFromDB() (*Game, error) {
  db, err := getDBConnection()
  if err != nil {
    return nil, err
  }
  defer db.Close()

  game := &Game{
    startedAt:        time.Now(),
    heroes:           []Hero{},
    joinChan:         make(chan JoinRequest),
    activateHeroChan: make(chan ActivateHeroRequest),
    exitChan:         make(chan []byte),
    adminToken:       "1234",
  }

  rows, err := db.Query("SELECT name, email, class, enabled, token, level, ttl, xpos, ypos, " +
    "IFNULL(ring, 0), IFNULL(amulet, 0), IFNULL(charm, 0), IFNULL(weapon, 0), IFNULL(helm, 0), " +
    "IFNULL(tunic, 0), IFNULL(gloves, 0), IFNULL(shield, 0), IFNULL(leggings, 0), IFNULL(boots, 0) " +
    "FROM hero LEFT JOIN equipment ON hero.id=equipment.hero_id")
  if err != nil {
    return nil, err
  }
  defer rows.Close()

  for rows.Next() {
    hero := &Hero{
      Equipment: Equipment{},
    }
    var ttl int
    err := rows.Scan(&hero.Name, &hero.Email, &hero.Class, &hero.Enabled,
      &hero.token, &hero.Level, &ttl, &hero.Xpos, &hero.Ypos,
      &hero.Equipment.Ring, &hero.Equipment.Amulet, &hero.Equipment.Charm,
      &hero.Equipment.Weapon, &hero.Equipment.Helm, &hero.Equipment.Tunic,
      &hero.Equipment.Gloves, &hero.Equipment.Shield, &hero.Equipment.Leggings,
      &hero.Equipment.Boots)
    if err != nil {
      log.Error(err)
    }
    hero.nextLevelAt = time.Now().Add(time.Duration(ttl) * time.Second)
    game.heroes = append(game.heroes, *hero)
  }
  err = rows.Err()
  if err != nil {
    return nil, err
  }

  return game, nil
}
