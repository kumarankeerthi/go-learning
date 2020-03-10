// Package core has all the core logic of customerservice
package core

import "context"

// Service inteface is expose the customerservice functionalities via different transport, in this case http
type Service interface {
	AddCustomer(c Customer,ctx context.Context) (string, error)
	GetCustomerByName(fname string, lname string) (Customer, error)
	UpdateCutomerDetails(cid string, c Customer) error
}
