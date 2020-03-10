// Package data contains the db client
package data

import (
	"context"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/kumarankeerthi/go-learning/banking-system/common/tracing"
	"github.com/kumarankeerthi/go-learning/banking-system/customerservice/cmd"
	"github.com/kumarankeerthi/go-learning/banking-system/customerservice/core"
)

// CockroachRepo type will hold the intance of cockroachDB
type CockroachRepo struct {
	crDB *gorm.DB
}

// CreateRepository will create the DBClient
func CreateRepository(cfg *cmd.Config) *CockroachRepo {
	fmt.Println("Creating cockroachdb repository")
	db, err := gorm.Open("postgres", cfg.CockroachdbConnURL)
	if err != nil {
		fmt.Println("error connecting to cockroachdb!!", err.Error())
	}
	fmt.Println("Sucessfully connected to cockroachDB!!!")
	return &CockroachRepo{crDB: db}
	//return nil
}

// AddCustomer will have the logic insert data into cockraodhdb
func (crRepo *CockroachRepo) AddCustomer(c core.Customer, ctx context.Context) (string, error) {
	span := tracing.TraceFuncCall("CockroachDB", ctx)
	defer span.Finish()
	time.Sleep(2 * time.Second)
	fmt.Println("CockroachRepo : AddCustomer")
	dbCustomer := getDBCustomerObj(c)
	dbAddress := getDBAddressObj(c)
	tableExists := crRepo.crDB.HasTable(&Customer{})
	if !tableExists {
		fmt.Println("CockroachRepo : table does not existes- creating now")
		crRepo.crDB.AutoMigrate(&Customer{})
	}
	tableExists = crRepo.crDB.HasTable(&Address{})
	if !tableExists {
		fmt.Println("CockroachRepo : table does not existes- creating now")
		crRepo.crDB.AutoMigrate(&Address{})
	}
	row := new(Address)
	d := crRepo.crDB.Create(&dbAddress).Scan(&row)
	if d.Error != nil {
		fmt.Println("Error inserring address")
	}
	dbCustomer.AddressId = row.ID
	crRepo.crDB.Create(&dbCustomer)
	return "Success", nil
}

// GetCustomerByName will have the logic to fecth data from cockraodhdb
func (crRepo *CockroachRepo) GetCustomerByName(fname string, lname string) (core.Customer, error) {

	fmt.Println("CockroachRepo : GetCustomerByName")
	return core.Customer{}, nil
}

// UpdateCutomerDetails will have the logic to update cockraodhdb
func (crRepo *CockroachRepo) UpdateCutomerDetails(cid string, c core.Customer) error {

	fmt.Println("CockroachRepo : UpdateCutomerDetails")
	return nil
}

func getDBCustomerObj(c core.Customer) Customer {
	return Customer{
		FirstName: c.FirstName,
		LastName:  c.LastName,
	}
}

func getDBAddressObj(c core.Customer) Address {
	return Address{
		City:    c.Address.City,
		State:   c.Address.State,
		ZipCode: c.Address.ZipCode,
	}
}
