package model

import (
	"time"

	"github.com/shopspring/decimal"
)

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

	Properties          []ProductProperty    `json:"properties,omitempty"`
	Images              []ProductImage       `json:"images,omitempty"`
	Videos              []ProductVideo       `json:"videos,omitempty"`
	ProductMarketplaces []ProductMarketplace `json:"product_marketplaces,omitempty"`
	IsKit               bool                 `json:"is_kit" db:"is_kit"` // indica se é um kit (produto pai)
	KitComponents       []ProductKit         `json:"kit_components,omitempty"`
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

type ProductMarketplace struct {
	ID            int64 `db:"id"`
	ProductID     int64 `db:"product_id"`
	MarketplaceID int64 `db:"marketplace_id"`

	ExternalID *string `db:"external_id"` // pode ser null
	ProductURL string  `db:"product_url"`

	Price       *decimal.Decimal `db:"price"` // nullable
	ListingType *string          `db:"listing_type"`
	Status      *string          `db:"status"`

	Active bool `db:"active"`

	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type ProductKit struct {
	ID                 int64 `db:"id"`
	ProductID          int64 `db:"product_id"`           // KIT (produto pai)
	ComponentProductID int64 `db:"component_product_id"` // produto componente
	Quantity           int   `db:"quantity"`
}
