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
