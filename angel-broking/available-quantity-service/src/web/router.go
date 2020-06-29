package web

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"problem-solving/angel-broking/available-quantity-service/src/model"
	"strconv"
	"time"

	repo "problem-solving/angel-broking/available-quantity-service/src/repository"

	"github.com/gorilla/mux"
)

// NewRouter create new routes
func NewRouter() {
	router := mux.NewRouter()
	router = router.PathPrefix("/availablequantity").Subrouter()

	router.HandleFunc("/test", availableQuantityTest)

	router.HandleFunc("/products/{productId}/getquantity", getQuantity)

	// 1 is to increase and 2 is to decrease
	//?change={up}&quantity={quantity}
	router.HandleFunc("/products/{productId}/updatequantity", updateQuantity)

	srv := &http.Server{
		Handler:      router,
		Addr:         "localhost:8002",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func getQuantity(w http.ResponseWriter, r *http.Request) {

	productID := mux.Vars(r)["productId"]

	stockDetails, err := repo.GetStock(productID)
	if err != nil {
		log.Println(fmt.Sprintf("Error occured while getting Stock details for Product %v, %v", productID, err))
		w.Write(getErrorObjectByError(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Println(fmt.Sprintf("Stock details for Product %v, %v", productID, stockDetails.Stock))

	b, err := json.Marshal(stockDetails)
	if err != nil {
		w.Write(getErrorObjectByError(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(b)
	w.WriteHeader(http.StatusOK)
}

func updateQuantity(w http.ResponseWriter, r *http.Request) {

	productID := mux.Vars(r)["productId"]

	// change represents whether to be added to sbtracted
	// 1- Add  2- Subtract
	change := r.URL.Query().Get("change")
	quantity, err := strconv.Atoi(r.URL.Query().Get("quantity"))
	if err != nil {
		w.Write(getErrorObjectByError(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Println(fmt.Sprintf("Stock update request. %v, %v, %v", productID, change, quantity))

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	switch change {
	case "1":
		err := repo.AddStock(ctx, productID, quantity)
		if err != nil {
			w.Write(getErrorObjectByError(fmt.Errorf("Product not present in store. %v, %v, %v %v", productID, change, quantity, err)))
			w.WriteHeader(http.StatusBadRequest)
		}
		w.WriteHeader(http.StatusOK)

	case "2":
		// status will be returned as true will successfully deducted else false
		resp, err := repo.SubtractStock(ctx, productID, quantity)
		if err != nil {
			w.Write(getErrorObjectByError(fmt.Errorf("Product not present in store. %v, %v, %v %v", productID, change, quantity, err)))
			w.WriteHeader(http.StatusBadRequest)
		}
		b, err := json.Marshal(resp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(b)
		w.WriteHeader(http.StatusOK)

	default:
		log.Println(fmt.Sprintf("Invalid change type. %v, %v, %v", productID, change, quantity))
		w.Write(getErrorObjectByError(fmt.Errorf("Invalid change type. %v, %v, %v", productID, change, quantity)))
		w.WriteHeader(http.StatusBadRequest)
	}

}

func availableQuantityTest(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("test response from available Quantity service"))
}

func getErrorObjectByError(err error) []byte {
	b, _ := json.Marshal(model.ErrorObj{Error: err.Error()})
	return b
}
