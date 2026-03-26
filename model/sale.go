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
	CreatedAt       time.Time       `json:"created_at" db:"created_at"`

	PaymentMethod *PaymentMethod `json:"payment_method,omitempty"`
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
	SalesOrderID int64           `json:"sales_order_id" db:"sales_order_id"`
	ProductID    int64           `json:"product_id" db:"product_id"`
	Quantity     int             `json:"quantity" db:"quantity"`
	UnitPrice    decimal.Decimal `json:"unit_price" db:"unit_price"`

	SalesOrder *SalesOrder `json:"sales_order,omitempty"`
	Product    *Product    `json:"product,omitempty"`
}

func (i *SalesOrderItem) Total() decimal.Decimal {
	return i.UnitPrice.Mul(decimal.NewFromInt(int64(i.Quantity)))
}
