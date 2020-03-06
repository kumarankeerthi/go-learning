package core

// Customer type
type Customer struct {
	ID        int64   `json:"id"`
	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
	Address   Address `json:"address"`
}

// Address type
type Address struct {
	ID      int64  `json:"id"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zipCode"`
}
