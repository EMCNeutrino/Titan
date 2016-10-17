package main

import "time"

// Hero struct contains information about each Hero
type Hero struct {
  HeroID         int64     `json:"heroid"`
  FirstName      string    `json:"username"`
  LastName       string    `json:"userlastname"`
  HeroName       string    `json:"heroname"`
  Token          string    `json:"token"`
  Twitter        string    `json:"twitter"`
  Email          string    `json:"email"`
  Title          string    `json:"title"`
  HRace          string    `json:"herorace"`
  IsAdmin        bool      `json:"isadmin"`
  Level          int       `json:"herolevel"`
  HeroClass      string    `json:"heroclass"`
  TTL            int       `json:"TTL"`
  Userhost       string    `json:"userhost"`
  Enabled        bool      `json:"enabled"`
  Xpos           int       `json:"xpos"`
  Ypos           int       `json:"ypos"`
  NextLevelAt    time.Time `json:"nextlevelat"`
  Weapon         int       `json:"weapon"`
  Tunic          int       `json:"tunic"`
  Shield         int       `json:"shield"`
  Leggings       int       `json:"leggings"`
  Ring           int       `json:"ring"`
  Gloves         int       `json:"gloves"`
  Boots          int       `json:"boots"`
  Helm           int       `json:"helm"`
  Charm          int       `json:"charm"`
  Amulet         int       `json:"amulet"`
  TotalEquipment int       `json:"totalequipment"`
  HeroCreatedAt  time.Time `json:"herocreatedat"`
}

type WorldEvent struct {
  WorldEventID int64     `json:"worldeventid"`
  EventType    string    `json:"eventtype"`
  EventText    string    `json:"eventtext"`
  EventTime    time.Time `json:"eventtime"`
}

type HeroWorldEvent struct {
  HeroWorldEventID int64 `json:"heroworldeventid"`
  WorldEventID     int64 `json:"worldeventid"`
  HeroID           int64 `json:"heroid"`
}

func (h *Hero) getItemLevel(itemType string) int {
  switch itemType {
  case "weapon":
    return h.Weapon
  case "tunic":
    return h.Tunic
  case "shield":
    return h.Shield
  case "leggings":
    return h.Leggings
  case "ring":
    return h.Ring
  case "gloves":
    return h.Gloves
  case "boots":
    return h.Boots
  case "helm":
    return h.Helm
  case "charm":
    return h.Charm
  case "amulet":
    return h.Amulet
  }
  return -1
}

// updateItem updates the Item level value for a specified item for a  Hero
func (h *Hero) updateItem(itemType string, itemLevel int) {
  switch itemType {
  case "weapon":
    h.Weapon = itemLevel
  case "tunic":
    h.Tunic = itemLevel
  case "shield":
    h.Shield = itemLevel
  case "leggings":
    h.Leggings = itemLevel
  case "ring":
    h.Ring = itemLevel
  case "gloves":
    h.Gloves = itemLevel
  case "boots":
    h.Boots = itemLevel
  case "helm":
    h.Helm = itemLevel
  case "charm":
    h.Charm = itemLevel
  case "amulet":
    h.Amulet = itemLevel
  }
}
