package domain

// Product represents the product table
type Product struct {
	ID         int64
	Name       string
	SKU        string
	Price      float64
	Active     bool
	CategoryID *int64
}
