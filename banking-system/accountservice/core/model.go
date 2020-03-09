package core

type Account struct {
	ID         int64   `json:"id"`
	Type       string  `json:"type"`
	Balance    float64 `json:"balance"`
	Status     string  `json:"status"`
	CustomerID int64   `json:"customerId"`
}

type Transaction struct {
	ID          int64   `json:"id"`
	FromAccount int64   `json:"fromAccount"`
	ToAccount   int64   `json:ToAccount"`
	Amount      float64 `json:"amount"`
	Remark      string  `json:"remark"`
}
