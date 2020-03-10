package data

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/kumarankeerthi/go-learning/banking-system/accountservice/cmd"
	"github.com/kumarankeerthi/go-learning/banking-system/accountservice/core"
)

type CockroachRepo struct {
	crDB *gorm.DB
}

func CreateRepostiory(cfg *cmd.Config) *CockroachRepo {
	fmt.Println("Creating cockroachDB connection")
	db, err := gorm.Open("postgres", cfg.CockroachdbConnURL)
	if err != nil {
		fmt.Println("Error establishing connection to DB")
	}
	fmt.Println("Sucessfully connected to cockroachDB!!!")
	return &CockroachRepo{
		crDB: db,
	}
}

func (crRepo *CockroachRepo) CreateAccount(account core.Account) error {
	fmt.Println("Repo : CreateAccount")
	return nil
}

func (crRepo *CockroachRepo) SendMoney(txn core.Transaction) (int64, error) {
	fmt.Println("Repo : SendMoney")
	return 123, nil
}

func (crRepo *CockroachRepo) GetTransactions(accountID int64) ([]core.Transaction, error) {
	fmt.Println("Repo : GetTransactions")
	return nil, nil
}
