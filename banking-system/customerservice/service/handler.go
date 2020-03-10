package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kumarankeerthi/go-learning/banking-system/common/tracing"
	"github.com/kumarankeerthi/go-learning/banking-system/customerservice/core"
)

type Handler struct {
	service core.Service
}
type HealthCheck struct {
	Health  string
	Message string
}

func CreateHandler(s core.Service) *Handler {
	return &Handler{service: s}
}
func (h *Handler) AddCustomer(w http.ResponseWriter, r *http.Request) {
	reqContext := r.Context()
	span := tracing.TraceFuncCall("Hanlder", reqContext)
	defer span.Finish()
	var c core.Customer
	json.NewDecoder(r.Body).Decode(&c)
	fmt.Println("Handler : AddCustomer")
	result, _ := h.service.AddCustomer(c, reqContext)
	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte("{\"Result\":" + result + "}"))

}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	var hc = HealthCheck{

		Health:  "OK",
		Message: "Up and running",
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(hc)

}
