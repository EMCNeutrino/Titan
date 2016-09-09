package main

// Hero struct contains information about each Hero
type Hero struct {
	Name     string	`json:"name"`
	token    string
	Energy   int64	`json:"energy"`
	Life     int64	`json:"life"`
	Shield   bool	`json:"shield"`
	// New Hero Fields
	Twitter  string `json:"twitter"`
	Email    string `json:"email"`
	Title    string `json:"title"`
	IsAdmin  bool	`json:"isadmin"`
	Level    int	`json:"level"`
	HClass   string	`json:"heroclass"`
	TTL      int64	`json:"energy"`
	Userhost string	`json:"userhost"`
	Online   bool	`json:"online"`
	Penalties []Penalty `json:"penalties"`
	Items     []Item    `json:"Items"`

}

type Penalty struct {
	Logout  int64        `json:"logout"`
	Quest   int64        `json:"quest"`
	Quit    int64        `json:"quit"`
	Message int64        `json:"message"`
}


type Item struct {

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
	Total	 int64	`json:"total"`
}

