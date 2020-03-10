// Package core has all the core logic of customerservice
package core
import(
	"context"
)
// Repository inteface is to interact with backend db
type Repository interface {
	AddCustomer(c Customer,ctx context.Context) (string, error)
	GetCustomerByName(fname string, lname string) (Customer, error)
	UpdateCutomerDetails(cid string, c Customer) error
}
