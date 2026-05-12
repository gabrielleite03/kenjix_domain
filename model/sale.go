package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type PaymentMethod struct {
	ID     int64  `json:"id" db:"id"`
	Name   string `json:"name" db:"name"`
	Active bool   `json:"active" db:"active"`
}

func (pm *PaymentMethod) IsActive() bool {
	return pm.Active
}

func (pm *PaymentMethod) Deactivate() {
	pm.Active = false
}

type SalesOrder struct {
	ID              int64           `json:"id" db:"id"`
	Price           decimal.Decimal `json:"price" db:"price"`
	Discount        decimal.Decimal `json:"discount" db:"discount"`
	Status          string          `json:"status" db:"status"`
	PaymentMethodID int64           `json:"payment_method_id" db:"payment_method_id"`
	Active          bool            `json:"active" db:"active"`

	CustomerName     *string `json:"customer_name,omitempty" db:"customer_name"`
	CustomerDocument *string `json:"customer_document,omitempty" db:"customer_document"`

	MarketplaceID   *int64  `json:"marketplace_id,omitempty" db:"marketplace_id"`
	ExternalOrderID *string `json:"external_order_id,omitempty" db:"external_order_id"`
	ExternalPackID  *string `json:"external_pack_id,omitempty" db:"external_pack_id"`

	PaymentStatus  *string `json:"payment_status,omitempty" db:"payment_status"`
	DeliveryStatus *string `json:"delivery_status,omitempty" db:"delivery_status"`

	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" db:"updated_at"`

	PaymentMethod *PaymentMethod `json:"payment_method,omitempty"`

	Marketplace *Marketplace `json:"marketplace,omitempty"`

	Invoice *SalesOrderInvoice `json:"invoice,omitempty"`
}

func (s *SalesOrder) Total() decimal.Decimal {
	return s.Price.Sub(s.Discount)
}

func (s *SalesOrder) IsActive() bool {
	return s.Active
}

func (s *SalesOrder) Deactivate() {
	s.Active = false
}

type SalesOrderItem struct {
	SalesOrderID   int64 `json:"sales_order_id" db:"sales_order_id"`
	PurchaseItemID int64 `json:"purchase_item_id" db:"purchase_item_id"`

	Quantity  int             `json:"quantity" db:"quantity"`
	UnitPrice decimal.Decimal `json:"unit_price" db:"unit_price"`

	SalesOrder *SalesOrder `json:"sales_order,omitempty"`

	PurchaseItem *PurchaseItem `json:"purchase_item,omitempty"`
}

func (i *SalesOrderItem) Total() decimal.Decimal {
	return i.UnitPrice.Mul(decimal.NewFromInt(int64(i.Quantity)))
}
