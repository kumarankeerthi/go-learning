package service

import (
	"fmt"

	"github.com/kumarankeerthi/go-learning/banking-system/accountservice/core"
)

type Server struct {
	h *Handler
}

func CreateServer(as core.Service) *Server {
	handler := CreateHandler(as)
	return &Server{
		h: handler,
	}
}

func (s *Server) Start(port string) {
	fmt.Println("Account Service started")
}
