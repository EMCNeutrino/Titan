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
    rowCnt, err := res.RowsAffected()
    if err != nil {
      log.Error(err)
    }
    log.Printf("ID = %d, affected = %d\n", lastID, rowCnt)
  }

  return nil
}

func loadFromDB() (*Game, error) {
  db, err := getDBConnection()
  if err != nil {
    return nil, err
  }
  defer db.Close()

  return nil, nil
}
