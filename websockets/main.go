package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {

	fmt.Println(" starting of websocket example")
	http.HandleFunc("/getData", response)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func response(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("This is a message from the server!!")
}
