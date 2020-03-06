// Package core has all the core logic of customerservice
package core

// Repository inteface is to interact with backend db
type Repository interface {
	AddCustomer(c Customer) (string, error)
	GetCustomerByName(fname string, lname string) (Customer, error)
	UpdateCutomerDetails(cid string, c Customer) error
}
