package main

import "time"

// Hero struct contains information about each Hero
type HeroDB struct {
	HeroID		int64  	`json:"heroid"`
	UserName	string 	`json:"username"`
	UserLastName 	string 	`json:"userlastname"`
	HeroName     	string	`json:"heroname"`
	Token    	string 	`json:"token"`
	Twitter  	string	`json:"twitter"`
	Email    	string	`json:"email"`
	Title    	string	`json:"title"`
	HRace    	string	`json:"herorace"`
	IsAdmin  	bool	`json:"isadmin"`
	Level    	int	`json:"herolevel"`
	HClass   	string	`json:"heroclass"`
	TTL      	int	`json:"TTL"`
	Userhost 	string	`json:"userhost"`
	Enabled   	bool	`json:"enabled"`
	Xpos		int	`json:"xpos"`
	Ypos		int	`json:"ypos"`
	NextLevelAt	time.Time  `json:"nextlevelat"`
	Weapon   int	`json:"weapon"`
	Tunic    int	`json:"tunic"`
	Shield   int	`json:"shield"`
	Leggings int	`json:"leggings"`
	Ring     int	`json:"ring"`
	Gloves   int	`json:"gloves"`
	Boots    int	`json:"boots"`
	Helm	 int	`json:"helm"`
	Charm	 int	`json:"charm"`
	Amulet	 int	`json:"amulet"`
	TotalEquipment	 int	`json:"totalequipment"`
	HeroCreatedAt	time.Time  `json:"herocreatedat"`
}

type WorldEvent struct {

	WorldEventID 	int64	`json:"worldeventid"`
	EventType		string	`json:"eventtype"`
	EventText		string	`json:"eventtext"`
	EventTime		time.Time `json:"eventtime"`
}

type HeroWorldEvent struct {

	HeroWorldEventID 	int64	`json:"heroworldeventid"`
	WorldEventID 		int64	`json:"worldeventid"`
	HeroID				int64  	`json:"heroid"`

}