package model

import "github.com/shopspring/decimal"

type Product struct {
	ID          int64            `json:"id" db:"id"`
	Name        string           `json:"name" db:"name"`
	SKU         string           `json:"sku" db:"sku"`
	Weight      *decimal.Decimal `json:"weight" db:"weight"`
	Price       decimal.Decimal  `json:"price" db:"price"`
	Marca       string           `json:"marca" db:"marca"`
	EAN         *string          `json:"ean" db:"ean"`
	NCM         *string          `json:"ncm" db:"ncm"`
	Description string           `json:"description" db:"description"`
	Active      bool             `json:"active" db:"active"`
	Volume      decimal.Decimal  `json:"volume" db:"volume"`
	CategoryID  *int64           `json:"category_id,omitempty" db:"category_id"`
	Category    *Category        `json:"category,omitempty"`

	Prices []ProductPrice `json:"prices,omitempty"` // preços por marketplace
	Stocks []Stock        `json:"stocks,omitempty"` // estoque por warehouse

	Properties []ProductProperty `json:"properties,omitempty"`
	Images     []ProductImage    `json:"images,omitempty"`
	Videos     []ProductVideo    `json:"videos,omitempty"`
}

type ProductPrice struct {
	ID            int64           `json:"id" db:"id"`
	ProductID     int64           `json:"product_id" db:"product_id"`
	MarketplaceID int64           `json:"marketplace_id" db:"marketplace_id"`
	Price         decimal.Decimal `json:"price" db:"price"`
	Active        bool            `json:"active" db:"active"`
	Volume        decimal.Decimal `json:"volume" db:"volume"`
	Product       *Product        `json:"product,omitempty"`
	Marketplace   *Marketplace    `json:"marketplace,omitempty"`
}

func (p *Product) IsActive() bool {
	return p.Active
}

func (p *Product) Deactivate() {
	p.Active = false
}

func (p *ProductPrice) Activate() {
	p.Active = true
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
