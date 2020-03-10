package core

type Service interface {
	CreateAccount(account Account) error
	SendMoney(txn Transaction) (int64, error)
	GetTransactions(accountID int64) ([]Transaction, error)
}
