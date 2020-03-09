package core

type Repository interface {
	CreateAccount(account Account) error
	SendMoney(txn Transaction) (int64, error)
	GetTransactions(accountID int64) ([]Transaction, error)
}
