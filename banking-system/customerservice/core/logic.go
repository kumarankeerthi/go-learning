// Package core has all the core logic of customerservice
package core

import (
	"fmt"
)

// CustomerService has the implementation of Service interface
type CustomerService struct {
	repo Repository
}

// CreateCutomerService will create a CustomerService object given a Repository
func CreateCutomerService(repo Repository) *CustomerService {
	return &CustomerService{repo: repo}
}

// AddCustomer will have the logic to validate the request and then invoke the repo for the operation
func (cs *CustomerService) AddCustomer(c Customer) (string, error) {
	// add validation here
	// in this sample project, it will be a pass thro from trasport to db layer
	fmt.Println("CustomerService : AddCustomer")
	return cs.repo.AddCustomer(c)
}

// GetCustomerByName will have the logic to validate the request and then invoke the repo for the operation
func (cs *CustomerService) GetCustomerByName(fname string, lname string) (Customer, error) {
	// add validation here
	// in this sample project, it will be a pass thro from trasport to db layer
	fmt.Println("CustomerService : GetCustomerByName")
	return cs.repo.GetCustomerByName(fname, lname)
}

// UpdateCutomerDetails will have the logic to validate the request and then invoke the repo for the operation
func (cs *CustomerService) UpdateCutomerDetails(cid string, c Customer) error {
	// add validation here
	// in this sample project, it will be a pass thro from trasport to db layer
	fmt.Println("CustomerService : UpdateCutomerDetails")
	return cs.repo.UpdateCutomerDetails(cid, c)
}
