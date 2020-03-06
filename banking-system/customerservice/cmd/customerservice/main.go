// Package main - entry point for customerservice
package main

import (
	"fmt"

	"github.com/kumarankeerthi/go-learning/banking-system/customerservice/cmd"
	"github.com/kumarankeerthi/go-learning/banking-system/customerservice/core"
	"github.com/kumarankeerthi/go-learning/banking-system/customerservice/data"
	"github.com/kumarankeerthi/go-learning/banking-system/customerservice/service"
)

var appName = "customerservice"

func main() {
	fmt.Println("App Name : ", appName)
	cfg := cmd.DefaultConfig()
	customerRepository := data.CreateRepository(cfg)
	customerService := core.CreateCutomerService(customerRepository)
	server := service.CreateServer(customerService)
	server.Start(cfg.ServicePort)

}
