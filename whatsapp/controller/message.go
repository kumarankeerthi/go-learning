package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Message struct {
	message string
}

func GetMessage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("list of all messages"))
	fmt.Println("completed sending messages")
}

func SendMessage(w http.ResponseWriter, r *http.Request) {
	var s Message
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&s); err != nil {
		w.Write(([]byte("Error processing message")))
	}
	w.Write(([]byte("Message send")))
}
