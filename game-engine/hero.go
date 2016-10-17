package main

import "time"

// Hero struct contains information about each Hero
type Hero struct {
  id          int64
  FirstName   string `json:"first_name"`
  LastName    string `json:"last_name"`
  Email       string `json:"email"`
  Twitter     string `json:"twitter"`
  HeroName    string `json:"hero_name"`
  HeroClass   string `json:"hero_class"`
  Enabled     bool   `json:"enabled"`
  TTL         int    `json:"ttl"`
  token       string
  Level       int `json:"level"`
  nextLevelAt time.Time
  CreatedAt   time.Time  `json:"created_at"`
  Equipment   *Equipment `json:"equipment"`
  Xpos        int        `json:"x_pos"`
  Ypos        int        `json:"y_pos"`
}

type Equipment struct {
  Ring     int `json:"ring"`
  Amulet   int `json:"amulet"`
  Charm    int `json:"charm"`
  Weapon   int `json:"weapon"`
  Helm     int `json:"helm"`
  Tunic    int `json:"tunic"`
  Gloves   int `json:"gloves"`
  Shield   int `json:"shield"`
  Leggings int `json:"leggings"`
  Boots    int `json:"boots"`
  Total    int `json:"total"`
}

type Event struct {
  Type string    `json:"type"`
  Text string    `json:"text"`
  Time time.Time `json:"time"`
}

func (h *Hero) getItemLevel(itemType string) int {
  switch itemType {
  case "weapon":
    return h.Equipment.Weapon
  case "tunic":
    return h.Equipment.Tunic
  case "shield":
    return h.Equipment.Shield
  case "leggings":
    return h.Equipment.Leggings
  case "ring":
    return h.Equipment.Ring
  case "gloves":
    return h.Equipment.Gloves
  case "boots":
    return h.Equipment.Boots
  case "helm":
    return h.Equipment.Helm
  case "charm":
    return h.Equipment.Charm
  case "amulet":
    return h.Equipment.Amulet
  }
  return -1
}

// updateItem updates the Item level value for a specified item for a  Hero
func (h *Hero) updateItem(itemType string, itemLevel int) {
  switch itemType {
  case "weapon":
    h.Equipment.Weapon = itemLevel
  case "tunic":
    h.Equipment.Tunic = itemLevel
  case "shield":
    h.Equipment.Shield = itemLevel
  case "leggings":
    h.Equipment.Leggings = itemLevel
  case "ring":
    h.Equipment.Ring = itemLevel
  case "gloves":
    h.Equipment.Gloves = itemLevel
  case "boots":
    h.Equipment.Boots = itemLevel
  case "helm":
    h.Equipment.Helm = itemLevel
  case "charm":
    h.Equipment.Charm = itemLevel
  case "amulet":
    h.Equipment.Amulet = itemLevel
  }
}

func (h *Hero) getTTL() int {
  return int(h.nextLevelAt.Sub(time.Now()).Seconds())
}

func (h *Hero) getTotalItems() int {
  return h.Equipment..Weapon + h.Equipment..Tunic + h.Equipment..Shield + h.Equipment..Leggings + h.Equipment..Ring + h.Equipment..Gloves + h.Equipment..Boots + h.Equipment..Helm + h.Equipment..Charm + h.Equipment..Amulet
}
