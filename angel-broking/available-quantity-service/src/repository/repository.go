package repo

import (
	"context"
	"fmt"
	"log"
	"problem-solving/angel-broking/available-quantity-service/src/model"
	"sync"
)

// Map of <productID,ProductDetails>
var productDetails map[string]*model.Product

// Distributed lock using Redis or Zookeeper to be used for distributed systems
var mutex sync.Mutex

func init() {
	log.Println("Intializing repo..!!")
	productDetails = make(map[string]*model.Product)
	productDetails["123"] = &model.Product{ID: "123", Name: "Test Product", Stock: 10}

}

// GetStock for a product
func GetStock(productID string) (*model.Product, error) {
	if val, ok := productDetails[productID]; ok {
		return val, nil
	}
	return &model.Product{}, fmt.Errorf("Product is not present in store. %v", productID)
}

// AddStock add stock
func AddStock(ctx context.Context, productID string, quantity int) error {
	if val, ok := productDetails[productID]; ok {
		mutex.Lock()
		val.Stock += quantity
		mutex.Unlock()
		return nil
	}
	return fmt.Errorf("Product is not present in store. %v", productID)

}

// SubtractStock subtract stock
func SubtractStock(ctx context.Context, productID string, quantity int) (model.UpdateStockResponse, error) {
	result := model.UpdateStockResponse{ProductID: productID}
	if val, ok := productDetails[productID]; ok {
		mutex.Lock()
		if quantity > val.Stock {
			result.Status = false
			return result, nil
		}
		val.Stock -= quantity
		mutex.Unlock()
		result.Status = true
		return result, nil
	}
	return result, fmt.Errorf("Product is not present in store. %v", productID)
}
