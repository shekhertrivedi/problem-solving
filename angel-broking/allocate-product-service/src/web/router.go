package web

import (
	"encoding/json"
	"fmt"
	"hystrix-go/hystrix"
	"io/ioutil"
	"log"
	"net/http"
	"problem-solving/angel-broking/allocate-product-service/src/model"
	"problem-solving/angel-broking/allocate-product-service/src/util"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

const (
	paymentServiceBaseURL           = "http://localhost:8003/paymentservicestub"
	paymentServiceSuffix            = "/orders/%v/%v/pay"
	availableQuantityServiceBaseURL = "http://localhost:8002/availablequantity"
	availableQuantityServiceSuffix  = "/products/%v/updatequantity?change=%v&quantity=%v"
	addProductServiceBaseURL        = "http://localhost:8000/addproduct"
	addProductServiceSuffix         = "/users/%v/removecart"
)

// NewRouter create new routes
func NewRouter() {
	router := mux.NewRouter()
	router = router.PathPrefix("/allocateproduct").Subrouter()

	router.HandleFunc("/test", allocateProductTest)

	router.HandleFunc("/allocatecart", allocate).Methods(http.MethodPost)

	srv := &http.Server{
		Handler:      router,
		Addr:         "localhost:8001",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func allocate(w http.ResponseWriter, r *http.Request) {

	// get the body
	cart, err := getRequestBody(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(getErrorObjectByError(err))
		return
	}

	// make payment for cart at once using payment stub

	ps, err := makePayment(cart)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(getErrorObjectByError(err))
		return
	}
	if !ps.Status {
		log.Printf("payment service request failed %v %v %v \n", ps.OrderID, ps.Status, cart)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(getErrorObjectByError(fmt.Errorf("payment service request failed %v %v", ps.OrderID, ps.Status)))
		return
	}

	// call available quantity service for allocation
	isAllocated, err := allocateUnitsSingleThread(cart)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(getErrorObjectByError(err))
		return
	}

	if !isAllocated {
		// start payment return procedure
		w.WriteHeader(http.StatusConflict)
		w.Write(getErrorObjectByError(fmt.Errorf("Shopping cart products are not available %v", cart)))
		return
	}

	// Put the Order details to Shipping Service
	// Either through synchronous REST call OR
	// preferably by posting details to a queue

	// mark the items as success or failure
	removeCart(cart)
	w.WriteHeader(http.StatusOK)

}

func removeCart(cart model.Cart) {

	url := addProductServiceBaseURL + fmt.Sprintf(addProductServiceSuffix, cart.UserID)

	circuitBreaker := func() error {
		resp, err := util.PerformRequest(http.MethodGet, url, nil)
		if err != nil {
			log.Printf("Error occured while removing the cart for user %v", cart.UserID)
			return err
		}

		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			log.Printf("Cart removal API failure for user %v %v", cart.UserID, resp.StatusCode)
			return fmt.Errorf("Cart removal API failure for user %v %v", cart.UserID, resp.StatusCode)
		}
		return nil
	}

	hErr := hystrix.Do("addproductservice", circuitBreaker, nil)
	if hErr != nil {
		log.Printf("Failed to remove the cart for user %v. Need to go via retry", cart.UserID)
		// TODO: Add asynchronous retry mechanism to purge the order details
	}

}

func allocateUnitsMultipleThreads(cart model.Cart) bool {
	// ASSUMPTION: If any item of the cart is not allocated successfully, cancel the entire purchase
	var wg sync.WaitGroup
	wg.Add(len(cart.Items))
	responses := make([]bool, len(cart.Items))
	for i, val := range cart.Items {

		go func() {
			defer wg.Done()
			usr := model.UpdateStockResponse{}
			url := availableQuantityServiceBaseURL + fmt.Sprintf(availableQuantityServiceSuffix, val.ProductID, "2", val.Quantity)
			resp, err := util.PerformRequest(http.MethodGet, url, nil)
			if err != nil {
				log.Printf("Error occured while updating stock %v \n", err)
				responses[i] = false
				return
			}

			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				log.Printf("Error occured while updating stock %v \n", resp.StatusCode)
				responses[i] = false
				return
			}

			body, err := ioutil.ReadAll(resp.Body)
			if err = json.Unmarshal(body, &usr); err != nil {
				log.Printf("Error occured while updating stock URL: %v Response: %v error: %v \n", url, string(body), err)
				responses[i] = false
				return
			}
			responses[i] = usr.Status
		}()
	}
	wg.Wait()

	for _, v := range responses {
		if !v {
			return false
		}
	}
	return true
}

func allocateUnitsSingleThread(cart model.Cart) (bool, error) {
	// ASSUMPTION: If any item of the cart is not allocated successfully, cancel the entire purchase
	for _, val := range cart.Items {

		usr := model.UpdateStockResponse{}
		url := availableQuantityServiceBaseURL + fmt.Sprintf(availableQuantityServiceSuffix, val.ProductID, "2", val.Quantity)
		// flag ensures that if any one of the product is unavailable, then return from function with false,nil
		flag := false

		circuitBreaker := func() error {
			resp, err := util.PerformRequest(http.MethodGet, url, nil)
			if err != nil {
				log.Printf("Error occured while updating stock %v \n", err)
				return fmt.Errorf("Error occured while updating stock %v", err)
			}

			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				log.Printf("Error occured while updating stock %v \n", resp.StatusCode)
				return fmt.Errorf("Error occured while updating stock %v", resp.StatusCode)
			}

			body, err := ioutil.ReadAll(resp.Body)
			if err = json.Unmarshal(body, &usr); err != nil {
				log.Printf("Error occured while updating stock URL: %v Response: %v error: %v \n", url, string(body), err)
				return fmt.Errorf("Error occured while updating stock URL: %v Response: %v error: %v", url, string(body), err)
			}
			if !usr.Status {
				flag = true
				return nil
			}
			return nil
		}

		hErr := hystrix.Do("availablequantityservice", circuitBreaker, nil)
		if hErr != nil {
			return false, hErr
		}
		if flag {
			return false, nil
		}
	}
	return true, nil
}

func makePayment(cart model.Cart) (model.PaymentStatus, error) {

	orderID := cart.UserID + "_" + cart.Items[0].ProductID

	ps := model.PaymentStatus{OrderID: orderID}
	// to be calculated by consulting the pricing service STUB
	amount := 100
	url := paymentServiceBaseURL + fmt.Sprintf(paymentServiceSuffix, orderID, amount)

	circuitBreaker := func() error {
		resp, err := util.PerformRequest(http.MethodGet, url, nil)
		if err != nil {
			log.Println("Error occured while calling payment service ", err)
			return fmt.Errorf("Error occured while calling payment service %v", err)
		}

		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			log.Printf("payment service request failed %v", resp.StatusCode)
			return fmt.Errorf("payment service request failed %v", resp.StatusCode)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err = json.Unmarshal(body, &ps); err != nil {
			log.Printf("payment service request failed URL: %v Response: %v error: %v", url, string(body), err)
			return fmt.Errorf("payment service request failed URL: %v Response: %v error: %v", url, string(body), err)
		}
		return nil
	}

	hErr := hystrix.Do("paymentserviceSTUB", circuitBreaker, nil)
	if hErr != nil {
		return ps, hErr
	}
	return ps, nil
}

func getRequestBody(r *http.Request) (model.Cart, error) {
	var cart model.Cart
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error occured while extracting allocate body %v \n", err)
		return cart, err
	}
	log.Println("Allocation request body got: ", string(data))

	err = json.Unmarshal(data, &cart)
	if err != nil {
		log.Printf("Error occured while unmarshaling the payload %v \n", err)
		return cart, err
	}
	return cart, nil
}

func allocateProductTest(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("test response from allocate product service"))
}

func getErrorObjectByError(err error) []byte {
	b, _ := json.Marshal(model.ErrorObj{Error: err.Error()})
	return b
}
