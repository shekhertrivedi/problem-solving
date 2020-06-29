package web

import (
	"encoding/json"
	"fmt"
	"hystrix-go/hystrix"
	"io/ioutil"
	"log"
	"net/http"
	"problem-solving/angel-broking/add-product-service/src/model"
	repo "problem-solving/angel-broking/add-product-service/src/repository"
	"problem-solving/angel-broking/add-product-service/src/util"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

const (
	availableQuantityServiceBaseURL = "http://localhost:8002/availablequantity"
	availableQuantityServiceSuffix  = "/products/%v/getquantity"
)

// NewRouter create new routes
func NewRouter() {
	router := mux.NewRouter()
	router = router.PathPrefix("/addproduct").Subrouter()

	router.HandleFunc("/test", addProductTest)

	router.HandleFunc("/products/{productId}/getproductdetails", getproductdetails)

	router.HandleFunc("/users/{userId}/products/{productId}/addtocart", addToCart)

	router.HandleFunc("/users/{userId}/products/{productId}/removefromcart", removeFromCart)

	router.HandleFunc("/users/{userId}/getcart", getCart)

	router.HandleFunc("/users/{userId}/removecart", removeCart)

	//router.HandleFunc("/products/{productId}/quantity/{quantity}/add", getQuantity)

	srv := &http.Server{
		Handler:      router,
		Addr:         "localhost:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

	//http.ListenAndServe(":8000", router)
}

func removeCart(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userId"]
	flag := repo.RemoveCart(userID)

	if flag {
		w.WriteHeader(http.StatusOK)
	}

}

func getCart(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userId"]
	cart := repo.GetCart(userID)

	log.Println(fmt.Sprintf("Request to get cart details. %v", userID))

	b, err := json.Marshal(cart)
	if err != nil {
		log.Println(fmt.Sprintf("Error occured while marshaling cart response. %v, %v", userID, err))
		w.Write(getErrorObjectByError(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(b)
	w.WriteHeader(http.StatusOK)
}

func removeFromCart(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productID := vars["productId"]
	userID := vars["userId"]
	quantity := 1
	var err error
	if r.URL.Query().Get("quantity") != "" {
		quantity, err = strconv.Atoi(r.URL.Query().Get("quantity"))
		if err != nil {
			log.Println(fmt.Sprintf("Error occured while fetching quantity. %v, %v, %v", userID, productID, err))
			w.WriteHeader(http.StatusBadRequest)
			w.Write(getErrorObjectByError(err))
			return
		}
	}

	log.Println(fmt.Sprintf("Request to remove items from cart. %v, %v, %v", userID, productID, quantity))

	cart := model.Cart{UserID: userID}
	cart.Items = []*model.Item{&model.Item{ProductID: productID, Quantity: quantity}}
	err = repo.RemoveFromCart(userID, &model.Item{ProductID: productID, Quantity: quantity})
	if err != nil {
		log.Println(fmt.Sprintf("Error occured while removing items from cart. %v, %v, %v, %v", userID, productID, quantity, err))
		w.WriteHeader(http.StatusBadRequest)
		w.Write(getErrorObjectByError(err))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func addToCart(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productID := vars["productId"]
	userID := vars["userId"]
	quantity := 1
	var err error
	if r.URL.Query().Get("quantity") != "" {
		quantity, err = strconv.Atoi(r.URL.Query().Get("quantity"))
		if err != nil {
			log.Println(fmt.Sprintf("Error occured while adding to cart. %v, %v, %v", userID, productID, err))
			w.WriteHeader(http.StatusBadRequest)
			w.Write(getErrorObjectByError(err))
			return
		}
	}

	log.Println(fmt.Sprintf("Request to add items to cart. %v, %v, %v", userID, productID, quantity))

	// check whether the quantity to be added is present or not
	//isAvailable, err := checkProductAvailability(productID, quantity)
	isAvailable, err := checkProductAvailabilityWithHystrix(productID, quantity)

	if err != nil {
		log.Println(fmt.Sprintf("Error occured while checking availability. %v, %v, %v", userID, productID, err))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(getErrorObjectByError(fmt.Errorf("Error occured while checking availability. %v, %v, %v", userID, productID, err)))
		return
	}

	if !isAvailable {
		log.Println(fmt.Sprintf("Product is not available at this moment. %v, %v, %v", userID, productID, err))
		w.WriteHeader(http.StatusConflict)
		w.Write(getErrorObjectByError(fmt.Errorf("Product is not available at this moment. %v, %v, %v", userID, productID, err)))
		return
	}

	cart := model.Cart{UserID: userID}
	cart.Items = []*model.Item{&model.Item{ProductID: productID, Quantity: quantity}}
	repo.AddToCart(cart)
	w.WriteHeader(http.StatusOK)
}

func checkProductAvailabilityWithHystrix(productID string, quantity int) (bool, error) {

	isAvailable := false
	pd := model.Product{}
	url := availableQuantityServiceBaseURL + fmt.Sprintf(availableQuantityServiceSuffix, productID)

	circuitBreaker := func() error {

		resp, err := util.PerformRequest(http.MethodGet, url, nil)
		if err != nil {
			log.Printf("checkProductAvailability: Error occured while checking availability %v", err)
			return fmt.Errorf("Error occured while checking availability %v %v %v", productID, quantity, err)
		}

		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			log.Printf("checkProductAvailability: Invalid response code while checking availability %v %v %v", productID, quantity, err)
			return fmt.Errorf("checkProductAvailability: Invalid response code while checking availability %v %v %v", productID, quantity, err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err = json.Unmarshal(body, &pd); err != nil {
			log.Printf("checkProductAvailability: unmarshaling failed URL: %v Response: %v error: %v", url, string(body), err)
			return fmt.Errorf("checkProductAvailability: unmarshaling failed URL: %v Response: %v error: %v", url, string(body), err)
		}
		if pd.Stock < quantity {
			return nil
		}
		isAvailable = true
		return nil
	}

	hErr := hystrix.Do("availablequantityservice", circuitBreaker, nil)
	if hErr != nil {
		return false, hErr
	}
	return isAvailable, nil
}

func checkProductAvailability(productID string, quantity int) (bool, error) {

	pd := model.Product{}

	url := availableQuantityServiceBaseURL + fmt.Sprintf(availableQuantityServiceSuffix, productID)
	resp, err := util.PerformRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Printf("checkProductAvailability: Error occured while checking availability %v", err)
		return false, fmt.Errorf("Error occured while checking availability %v %v %v", productID, quantity, err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("checkProductAvailability: Invalid response code while checking availability %v %v %v", productID, quantity, err)
		return false, fmt.Errorf("checkProductAvailability: Invalid response code while checking availability %v %v %v", productID, quantity, err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err = json.Unmarshal(body, &pd); err != nil {
		log.Printf("checkProductAvailability: unmarshaling failed URL: %v Response: %v error: %v", url, string(body), err)
		return false, fmt.Errorf("checkProductAvailability: unmarshaling failed URL: %v Response: %v error: %v", url, string(body), err)
	}

	if pd.Stock < quantity {
		return false, nil
	}
	return true, nil

}

func getproductdetails(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productID := vars["productId"]

	log.Println(fmt.Sprintf("Getting product details. %v", productID))

	productDetails := repo.GetProductDetails(productID)
	b, err := json.Marshal(productDetails)
	if err != nil {
		log.Println(fmt.Sprintf("Error occured while getting product details. %v, %v", productDetails, err))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(getErrorObjectByError(fmt.Errorf("Error occured while getting product details. %v, %v", productDetails, err)))
		return
	}
	w.Write(b)
	w.WriteHeader(http.StatusOK)
}

func addProductTest(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("test response from add product service"))
	w.WriteHeader(http.StatusOK)
}

func getErrorObjectByError(err error) []byte {
	b, _ := json.Marshal(model.ErrorObj{Error: err.Error()})
	return b
}
