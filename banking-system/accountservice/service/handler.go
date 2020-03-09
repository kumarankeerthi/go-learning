package service

import (
	"net/http"

	"github.com/kumarankeerthi/go-learning/banking-system/accountservice/core"
)

type Handler struct {
	service core.Service
}

func CreateHandler(accSer core.Service) *Handler {
	return &Handler{
		service: accSer,
	}
}

func (h *Handler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	h.service.CreateAccount()
}
