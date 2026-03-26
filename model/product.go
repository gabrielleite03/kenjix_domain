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
	Category    *Category       `json:"category,omitempty"`
}

func (p *Product) IsActive() bool {
	return p.Active
}

func (p *Product) Deactivate() {
	p.Active = false
}
