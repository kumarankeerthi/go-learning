// Package main - entry point for customerservice
package main

import (
	"fmt"

	"github.com/kumarankeerthi/go-learning/banking-system/customerservice/service"
)

var appName = "customerservice"

func main() {
	fmt.Println("App Name : ", appName)
	
	service.StartServer("8080")

}
