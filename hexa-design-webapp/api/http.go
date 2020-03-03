package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kumarankeerthi/go-learning/hexa-design-webapp/chat"
)

type RequestHandler interface {
	ViewChat(w http.ResponseWriter, r *http.Request)
	SendMessage(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	chatServ chat.ChatService
}

func CreateHandler(cs chat.ChatService) RequestHandler {
	return &handler{chatServ: cs}
}

func (h *handler) ViewChat(w http.ResponseWriter, r *http.Request) {
	fmt.Println(("viewing messages"))
	var c chat.Chat
	json.NewDecoder(r.Body).Decode(&c)
	messages, _ := h.chatServ.ViewChat(c.Sender, c.Recipent)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}

func (h *handler) SendMessage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("sedning messages")
	var c chat.Chat
	json.NewDecoder(r.Body).Decode(&c)
	h.chatServ.SendMessage(c)
}
