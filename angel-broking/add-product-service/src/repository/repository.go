package repo

import (
	"fmt"
	"log"
	"problem-solving/angel-broking/add-product-service/src/model"
)

var productDetails map[string]*model.Product
var cartDetails map[string]*model.Cart

func init() {
	log.Println("Intializing repo..!!")
	productDetails = make(map[string]*model.Product)
	productDetails["123"] = &model.Product{ID: "123", Name: "Test Product"}

	cartDetails = make(map[string]*model.Cart)
}

// GetProductDetails get details by product ID
func GetProductDetails(productID string) *model.Product {
	return productDetails[productID]
}

// AddToCart add items to cart
func AddToCart(cart model.Cart) {
	// flag identifies if the product to be added pre-exist into the cart or not
	// if the product pre-exists, then we update the quantity
	// if doesn't then add it as a new item
	flag := false
	if val, ok := cartDetails[cart.UserID]; ok {
		for _, item := range val.Items {
			if item.ProductID == cart.Items[0].ProductID {
				item.Quantity += cart.Items[0].Quantity
				flag = true
			}
		}
		if !flag {
			val.Items = append(val.Items, cart.Items[0])
		}

		return
	}
	cartDetails[cart.UserID] = &cart
}

// RemoveFromCart remove from cart
func RemoveFromCart(userID string, item *model.Item) error {
	cart := cartDetails[userID]
	for i, v := range cart.Items {
		if v.ProductID == item.ProductID {
			if v.Quantity < item.Quantity {
				return fmt.Errorf("Quantity not present in Cart %v %v", item.ProductID, item.Quantity)
			} else if v.Quantity > item.Quantity {
				cart.Items[i] = &model.Item{ProductID: cart.Items[i].ProductID, Quantity: cart.Items[i].Quantity - item.Quantity}
				return nil
			} else {
				//cart.Items = append(cart.Items[:i], cart.Items[i+1:])
				cart.Items[i] = cart.Items[len(cart.Items)-1]
				//cart.Items[len(cart.Items)-1]
				cart.Items = cart.Items[:len(cart.Items)-1]
				return nil
			}
		}
	}
	return fmt.Errorf("Product not present in Cart %v %v", item.ProductID, item.Quantity)
}

// GetCart get cart details by user ID
func GetCart(userID string) *model.Cart {
	return cartDetails[userID]
}

// RemoveCart remove cart by user ID
func RemoveCart(userID string) bool {
	if _, ok := cartDetails[userID]; ok {
		delete(cartDetails, userID)
		return true
	}
	return false
}
