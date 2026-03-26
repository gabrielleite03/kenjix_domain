package model

import "github.com/shopspring/decimal"

type Product struct {
	ID          int64           `json:"id" db:"id"`
	Name        string          `json:"name" db:"name"`
	SKU         string          `json:"sku" db:"sku"`
	Price       decimal.Decimal `json:"price" db:"price"`
	Marca       string          `json:"marca" db:"marca"`
	Description string          `json:"description" db:"description"`
	Active      bool            `json:"active" db:"active"`
	CategoryID  *int64          `json:"category_id,omitempty" db:"category_id"`

	Properties []ProductProperty `json:"properties,omitempty"`
	Images     []ProductImage    `json:"images,omitempty"`
	Videos     []ProductVideo    `json:"videos,omitempty"`
}

func (p *Product) IsActive() bool {
	return p.Active
}

func (p *Product) Deactivate() {
	p.Active = false
}

type ProductProperty struct {
	ID        int64  `json:"id" db:"id"`
	ProductID int64  `json:"product_id" db:"product_id"`
	Name      string `json:"name" db:"name"`
	Value     string `json:"value" db:"value"`
}

type ProductImage struct {
	ID        int64  `json:"id" db:"id"`
	ProductID int64  `json:"product_id" db:"product_id"`
	URL       string `json:"url" db:"url"`
	Position  int    `json:"position" db:"position"`
	IsPrimary bool   `json:"is_primary" db:"is_primary"`
}

type ProductVideo struct {
	ID        int64   `json:"id" db:"id"`
	ProductID int64   `json:"product_id" db:"product_id"`
	URL       string  `json:"url" db:"url"`
	Provider  *string `json:"provider,omitempty" db:"provider"`
}
