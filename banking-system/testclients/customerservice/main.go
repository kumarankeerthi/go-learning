package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Test Client for CustomerService")

	for i := 0; i < 50; i++ {
		go func() {
			respone, err := http.Post("http://127.0.0.1:8500/customer", "application/json", strings.NewReader("test"))
			if err != nil {
				fmt.Println("error invoking service", err.Error())
			}
			fmt.Println(respone.Body)
		}()

	}
	select {}
}
