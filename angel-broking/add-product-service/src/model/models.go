package model

// Product product details
type Product struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Stock int    `json:"stock,omitempty"`
}

// ErrorObj error details
type ErrorObj struct {
	Error string `json:"error,omitempty"`
}

// Cart shopping cart
type Cart struct {
	UserID string  `json:"userId,omitempty"`
	Items  []*Item `json:"items,omitempty"`
}

// Item shopping cart item
type Item struct {
	ProductID string `json:"productId,omitempty"`
	Quantity  int    `json:"quantity,omitempty"`
}
