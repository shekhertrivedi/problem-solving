package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router = router.PathPrefix("/paymentservicestub").Subrouter()

	router.HandleFunc("/orders/{orderId}/{amount}/pay", pay)

	srv := &http.Server{
		Handler:      router,
		Addr:         "localhost:8003",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

// PaymentStatus payment status
type PaymentStatus struct {
	Status  bool   `json:"status,omitempty"`
	OrderID string `json:"orderId,omitempty"`
}

func pay(w http.ResponseWriter, r *http.Request) {
	orderID := mux.Vars(r)["orderId"]
	amount := mux.Vars(r)["amount"]
	log.Println(fmt.Sprintf("Payment request received for Order %v Amount %v", orderID, amount))

	ps := PaymentStatus{Status: true, OrderID: orderID}

	b, _ := json.Marshal(ps)

	w.Write(b)
	w.WriteHeader(http.StatusOK)
}
