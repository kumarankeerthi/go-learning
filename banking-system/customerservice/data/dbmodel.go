package data

// Customer type
type Customer struct {
	ID        int64   `json:"id" gorm:"type:serial;primary_key"`
	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
	AddressId   int64 `json:"address" gorm:"ForeignKey:ID"`
}

// Address type
type Address struct {
	ID      int64  `json:"id" gorm:"type:serial;primary_key"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zipCode"`
}