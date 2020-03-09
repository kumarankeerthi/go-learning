package main

import (
	"fmt"

	"github.com/kumarankeerthi/go-learning/banking-system/accountservice/cmd"
	"github.com/kumarankeerthi/go-learning/banking-system/accountservice/core"
	"github.com/kumarankeerthi/go-learning/banking-system/accountservice/data"
	"github.com/kumarankeerthi/go-learning/banking-system/accountservice/service"
)

func main() {
	fmt.Println("Account Service")
	config := cmd.DefaultConfig()
	repo := data.CreateRepostiory(config)
	accountService := core.CreateAccountService(repo)

	server := service.CreateServer(accountService)
	server.Start(config.ServicePort)
}
