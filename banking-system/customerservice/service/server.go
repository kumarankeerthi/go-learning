package service

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/kumarankeerthi/go-learning/banking-system/common/monitoring"
	"github.com/kumarankeerthi/go-learning/banking-system/common/tracing"
	"github.com/kumarankeerthi/go-learning/banking-system/customerservice/core"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Server struct {
	handler *Handler
}

func CreateServer(serv core.Service) *Server {
	h := CreateHandler(serv)
	return &Server{handler: h}
}

// StartServer function will start the server
func (s *Server) Start(port string) {
	log.Println("Starting HTTP server for customerservice!")

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	r.With(tracing.Trace("AddCustomer")).
		With(monitoring.Monitor("customerservice", "AddCustomer", "GET /customer")).
		Post("/customer", s.handler.AddCustomer)

	r.Get("/health", s.handler.Health)
	r.Get("/metrics", promhttp.Handler().ServeHTTP)

	http.Handle("/", r)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Println("Failed to start HTTP server for customerservice!")
		log.Println("Error : ", err)
	}
}

func test(name string) func(http.Handler) http.Handler {
	fmt.Println(name)
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("before monitor")
			next.ServeHTTP(w, r)
			fmt.Println("after monitor")
		})
	}
}
func myMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("before monitor")
		next.ServeHTTP(w, r)
		fmt.Println("after monitor")
	})
}
