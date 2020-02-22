package controller

import (
	"fmt"
	"net/http"
)

func GetClaims(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("in getClaims handler function")

}

func GetClaimsOfEmployee(w http.ResponseWriter, r *http.Request) {
	fmt.Print("in getClaimsOfEmployee handle function")
}
