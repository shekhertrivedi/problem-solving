package model

// Product product details
type Product struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Stock int    `json:"stock,omitempty"`
}

// AvailableQuantity available product quantity
type AvailableQuantity struct {
	ProductID   string `json:"productId,omitempty"`
	IsAvailable bool   `json:"isAvailable,omitempty"`
	Quantity    int    `json:"quantity,omitempty"`
}

// ErrorObj error details
type ErrorObj struct {
	Error string `json:"error,omitempty"`
}

// UpdateStockResponse update stock response
type UpdateStockResponse struct {
	Status    bool   `json:"status,omitempty"`
	ProductID string `json:"productId,omitempty"`
	Stock     int    `json:"stock,omitempty"`
}
