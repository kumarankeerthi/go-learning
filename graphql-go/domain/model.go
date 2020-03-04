package domain

import "time"

type Customer struct {
	ID        string `json:"id,omitempty"`
	FirstName string `json:"firstName"`
	LastName  string `json:lastName"`
	Email     string `json:email"`
}

type Ticket struct {
	ID         string    `json:id,omitempty"`
	TravelDate time.Time `json:"date"`
	Customer   Customer  `json:customer"`
	Route      Route     `json:"route"`
}

type Route struct {
	ID          string `json:id,omitempty"`
	Source      string `json:"source"`
	Destination string `json:"destination"`
}

var Customers []Customer = []Customer{
	Customer{
		FirstName: "Keerthi",
		LastName:  "Kumaran",
		Email:     "test@gmail.com",
		ID:        "1234",
	},
	Customer{
		FirstName: "Maruti",
		LastName:  "Gorti",
		Email:     "test@gmail.com",
		ID:        "1234",
	},
}

var Routes []Route = []Route{
	Route{
		ID:          "1234",
		Source:      "RDU",
		Destination: "BLR",
	},
}

var Tickets []Ticket = []Ticket{
	Ticket{
		ID:         "1234",
		TravelDate: time.Now(),
		Customer:   Customers[0],
		Route:      Routes[0],
	},
	Ticket{
		ID:         "2222",
		TravelDate: time.Now(),
		Customer:   Customers[1],
		Route:      Routes[0],
	},
}
