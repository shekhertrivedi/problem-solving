package model

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

// PaymentStatus payment response
type PaymentStatus struct {
	Status  bool   `json:"status,omitempty"`
	OrderID string `json:"orderId,omitempty"`
}

// UpdateStockResponse update stock response
type UpdateStockResponse struct {
	Status    bool   `json:"status,omitempty"`
	ProductID string `json:"productId,omitempty"`
	Stock     int    `json:"stock,omitempty"`
}
