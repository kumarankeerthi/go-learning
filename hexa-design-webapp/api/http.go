package api

import (
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
	h.chatServ.ViewChat("k", "h")
}

func (h *handler) SendMessage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("sedning messages")
}
