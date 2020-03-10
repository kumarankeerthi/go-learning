package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kumarankeerthi/go-learning/banking-system/accountservice/core"
)

func main() {
	fmt.Println("Test Client for AccountService")
	data := SetUpAccountsData()

	for _, d := range data {
		v, _ := json.Marshal(d)
		r := bytes.NewReader(v)

		respone, err := http.Post("http://127.0.0.1:8501/account", "application/json", r)
		if err != nil {
			fmt.Println("error invoking service", err.Error())
		}
		fmt.Println(respone.Body)
	}
}

func SetUpAccountsData() []core.Account {
	a := core.Account{
		Type:       "Checknig",
		Balance:    100,
		CustomerID: 1234,
	}

	var data []core.Account
	var b core.Account

	for i := 0.0; i < 50; i++ {
		b.Balance = a.Balance * i
		b.CustomerID = a.CustomerID
		b.Type = a.Type
		data = append(data, b)
	}
	return data
}
