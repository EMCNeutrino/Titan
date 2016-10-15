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
	TTL      	int64	`json:"TTL"`
	Userhost 	string	`json:"userhost"`
	Online   	bool	`json:"online"`
	Xpos		int64	`json:"xpos"`
	Ypos		int64	`json:"ypos"`
	NextLevel	time.Time  `json:"nextlevel"`
	ItemID   int64  `json:"int64"`
	Weapon   int64	`json:"weapon"`
	Tunic    int64	`json:"tunic"`
	Shield   int64	`json:"shield"`
	Leggings int64	`json:"leggings"`
	Ring     int64	`json:"ring"`
	Gloves   int64	`json:"gloves"`
	Boots    int64	`json:"boots"`
	Energy   int64	`json:"energy"`
	Helm	 int64	`json:"helm"`
	Charm	 int64	`json:"charm"`
	Amulet	 int64	`json:"amulet"`
	TotalEquipment	 int64	`json:"totalequipment"`
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