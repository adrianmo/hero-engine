package main

// Hero struct contains information about each Hero
type Hero struct {
	HeroID		int64  	`json:"heroid"`
	UserName	string 	`json:"username"`
	UserLastName 	string 	`json:"userlastname"`
	HeroName     	string	`json:"heroname"`
	token    	string 	`json:"token"`
	Twitter  	string	`json:"twitter"`
	Email    	string	`json:"email"`
	Title    	string	`json:"title"`
	HRace    	string	`json:"herorace"`
	IsAdmin  	bool	`json:"isadmin"`
	Level    	int	`json:"level"`
	HClass   	string	`json:"heroclass"`
	TTL      	int64	`json:"energy"`
	Userhost 	string	`json:"userhost"`
	Online   	bool	`json:"online"`
	Xpos		int64	`json:"xpos"`
	Ypos		int64	`json:"ypos"`
	Penalties 	[]Penalty `json:"penalties"`
	Items     	[]Item    `json:"Items"`

}

type Penalty struct {
	PenaltyID  int64     `json:"penaltyid"`
	Logout  int64        `json:"logout"`
	Quest   int64        `json:"quest"`
	Quit    int64        `json:"quit"`
	Message int64        `json:"message"`
}


type Item struct {

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
	Total	 int64	`json:"total"`
}

