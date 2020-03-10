package core

import "fmt"

type AccountService struct {
	repo Repository
}

func CreateAccountService(repo Repository) *AccountService {
	return &AccountService{
		repo: repo,
	}
}

func (accSer *AccountService) CreateAccount(account Account) error {
	fmt.Println("AccountService : CreateAccount")
	return nil
}
func (accSer *AccountService) SendMoney(Transaction) (int64, error) {
	fmt.Println("AccountService : SendMoney")
	return 123, nil
}
func (accSer *AccountService) GetTransactions(accountID int64) ([]Transaction, error) {
	fmt.Println("AccountService : GetTransactions")
	return nil, nil
}
