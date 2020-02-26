package main

import (
	"errors"
	"fmt"
)

func passwholenumber(n int) (int, error) {
	if n < 0 {
		return -1, errors.New("not a whole number")
	}
	return n, nil
}

func main() {

	n, err := passwholenumber(-1)
	if err == nil {
		fmt.Println(n, "is a whole number")
	} else {
		fmt.Println("number passed is not a whole number")
	}
}
