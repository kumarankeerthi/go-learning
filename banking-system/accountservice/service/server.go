package service

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/kumarankeerthi/go-learning/banking-system/accountservice/core"
	"github.com/kumarankeerthi/go-learning/banking-system/common/monitoring"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))
	r.With(monitoring.Monitor("accountService", "CreateAccount", "POST /acount")).Post("/account", s.h.CreateAccount)
	r.Get("/metrics", promhttp.Handler().ServeHTTP)
	http.Handle("/", r)
	fmt.Println(http.ListenAndServe(":"+port, nil))
}
