package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	
)

type Employee struct {
	ID        string   `json:"id,omitempty"`
	FirstName string   `json:"firstname,omitempty"`
	LastName  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var employee []Employee

func main() {
	router := mux.NewRouter()

	{

		employee = append(employee, Employee{ID: "1", FirstName: "Tom", LastName: "Smith", Address: &Address{City: "Cary", State: "NC"}})
		employee = append(employee, Employee{ID: "1", FirstName: "Tom", LastName: "Smith", Address: &Address{City: "Cary", State: "NC"}})
		employee = append(employee, Employee{ID: "1", FirstName: "Tom", LastName: "Smith", Address: &Address{City: "Cary", State: "NC"}})
		employee = append(employee, Employee{ID: "1", FirstName: "Tom", LastName: "Smith", Address: &Address{City: "Cary", State: "NC"}})
		employee = append(employee, Employee{ID: "1", FirstName: "Tom", LastName: "Smith", Address: &Address{City: "Cary", State: "NC"}})

	}
	router.HandleFunc("/employee", GetEmployees).Methods("GET")
	router.HandleFunc("/employee/{id}", GetEmployee).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))

}

func GetEmployee(w http.ResponseWriter, r *http.Request) {
	// log.Printf("inside api")
	// json.NewEncoder(w).Encode(employee)
}

func GetEmployees(w http.ResponseWriter, r *http.Request) {
	log.Printf("inside api")
	json.NewEncoder(w).Encode(employee)
}
