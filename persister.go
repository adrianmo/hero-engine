package main

import (
  "database/sql"
  "time"

  log "github.com/Sirupsen/logrus"

  _ "github.com/go-sql-driver/mysql"
  "fmt"
)

// GetDBConnection builds and returns the database connection
func GetDBConnection(databaseURL string) (*sql.DB, error) {

  db, err := sql.Open("mysql", databaseURL+"?parseTime=true")
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
      "(player_name, player_lastname, hero_name, email, twitter, hclass, hero_online, token, hero_level, ttl, xpos, ypos, " +
      " ring, amulet, charm, weapon, helm, tunic, gloves, shield, leggings, boots " +
      ") " +
      "VALUES( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? ) " +
      "ON DUPLICATE KEY UPDATE " +
      "hero_online=VALUES(hero_online), hero_level=VALUES(hero_level), ttl=VALUES(ttl), xpos=VALUES(xpos), ypos=VALUES(ypos), " +
      "ring=VALUES(ring), amulet=VALUES(amulet), charm=VALUES(charm), weapon=VALUES(weapon), " +
      "helm=VALUES(helm), tunic=VALUES(tunic), gloves=VALUES(gloves), shield=VALUES(shield), " +
      "leggings=VALUES(leggings), boots=VALUES(boots);")
    if err != nil {
      log.Error(err)
    }

    ttl := int(hero.nextLevelAt.Sub(time.Now()).Seconds())
    res, err := stmt.Exec(hero.FirstName, hero.LastName, hero.HeroName, hero.Email, hero.Twitter, hero.HeroClass, hero.Enabled, hero.token,
      hero.Level, hero.HeroClass, hero.HeroTitle, ttl, hero.Xpos, hero.Ypos,
      hero.Equipment.Ring, hero.Equipment.Amulet, hero.Equipment.Charm, hero.Equipment.Weapon, hero.Equipment.Helm, hero.Equipment.Tunic, hero.Equipment.Gloves, hero.Equipment.Shield, hero.Equipment.Leggings, hero.Equipment.Boots)
    if err != nil {
      log.Error(err)
    }

    lastID, err := res.LastInsertId()
    if err != nil {
      log.Error(err)
    } else {
      hero.id = lastID
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
    "hero_level,  " +
    "COALESCE(hclass, '') AS hclass , +" +
    "COALESCE(race, '') AS race , +" +
    "COALESCE(title, '') AS title , +" +
    " ttl, hero_online, xpos, ypos, " +
    "IFNULL(weapon, 0), IFNULL(tunic, 0), IFNULL(shield, 0), IFNULL(leggings, 0), IFNULL(ring, 0), " +
    "IFNULL(gloves, 0), IFNULL(boots, 0), IFNULL(helm, 0), IFNULL(charm, 0) , IFNULL(amulet, 0) " +
    "total_equipment FROM hero")

  if err != nil {
    return err
  }
  defer rows.Close()

  for rows.Next() {
    hero := &Hero{Equipment: &Equipment{}}
    var ttl int

    err = rows.Scan(&hero.id, &hero.HeroName, &hero.FirstName, &hero.LastName, &hero.token, &hero.Twitter, &hero.Email,
      &hero.Level, &hero.HeroClass, &hero.HeroRace, &hero.HeroTitle, &ttl, &hero.Enabled,
      &hero.Xpos, &hero.Ypos, &hero.Equipment.Weapon, &hero.Equipment.Tunic, &hero.Equipment.Shield, &hero.Equipment.Leggings, &hero.Equipment.Ring, &hero.Equipment.Gloves,
      &hero.Equipment.Boots, &hero.Equipment.Helm, &hero.Equipment.Charm, &hero.Equipment.Amulet)

    if err != nil {
      log.Error(err)
      continue
    }

    hero.nextLevelAt = time.Now().Add(time.Duration(ttl) * time.Second)
    g.heroes = append(g.heroes, hero)

    //Message Realm
    var message = fmt.Sprintf("%s, %s, of the %s race has joined Bacelona's Fantasy Realm. Next Level in %d seconds.",hero.HeroName, hero.HeroTitle, hero.HeroRace, hero.getTTL())
    g.sendEvent(message, hero)

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

  r, err := tx.Exec("INSERT INTO worldevent (event_text) VALUES (?)", message)
  if err != nil {
    return err
  }

  eventID, err := r.LastInsertId()
  if err != nil {
    return err
  }

  for _, hero := range heroes {
    if _, err = tx.Exec("INSERT INTO heroworldevent (hero_id, worldevent_id ) VALUES (?, ?)", hero.id, eventID); err != nil {
      return err
    }
  }

  err = tx.Commit()
  if err != nil {
    tx.Rollback()
    return err
  }

  return nil
}

func (g *Game) GetEventsForHeroFromDB(heroID int64) ([]Event, error) {
  db, err := GetDBConnection(g.databaseURL)
  if err != nil {
    return nil, err
  }
  defer db.Close()

  rows, err := db.Query("SELECT w.event_text, w.event_time FROM heroworldevent h INNER JOIN worldevent w ON h.worldevent_id=w.worldevent_id WHERE h.hero_id=? ORDER BY w.event_time DESC", heroID)

  if err != nil {
    return nil, err
  }
  defer rows.Close()

  var events []Event
  for rows.Next() {
    event := &Event{}
    err = rows.Scan(&event.Text, &event.Time)
    if err != nil {
      log.Error(err)
      continue
    }
    events = append(events, *event)
  }
  err = rows.Err()
  if err != nil {
    return nil, err
  }

  return events, nil
}
