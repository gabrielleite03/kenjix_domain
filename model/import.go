package model

import (
	"time"

	"github.com/shopspring/decimal"
)

const (
	ImportCostFreight   = "FREIGHT"
	ImportCostTax       = "TAX"
	ImportCostInsurance = "INSURANCE"
	ImportCostCustoms   = "CUSTOMS"
	ImportCostStorage   = "STORAGE"
	ImportCostOther     = "OTHER"
)

type ImportProcess struct {
	ID              int64           `json:"id" db:"id"`
	PurchaseOrderID int64           `json:"purchase_order_id" db:"purchase_order_id"`
	Incoterm        string          `json:"incoterm" db:"incoterm"`
	ExchangeRate    decimal.Decimal `json:"exchange_rate" db:"exchange_rate"`
	Status          string          `json:"status" db:"status"`
	ArrivalDate     *time.Time      `json:"arrival_date,omitempty" db:"arrival_date"`

	PurchaseOrder *PurchaseOrder `json:"purchase_order,omitempty"`
}

type ImportCost struct {
	ID              int64           `json:"id" db:"id"`
	ImportProcessID int64           `json:"import_process_id" db:"import_process_id"`
	Type            string          `json:"type" db:"type"`
	Description     *string         `json:"description,omitempty" db:"description"`
	Amount          decimal.Decimal `json:"amount" db:"amount"`
	Currency        string          `json:"currency" db:"currency"`

	ImportProcess *ImportProcess `json:"import_process,omitempty"`
}

func (c *ImportCost) IsForeignCurrency() bool {
	return c.Currency != "BRL"
}

type ImportCostAllocation struct {
	ImportCostID    int64           `json:"import_cost_id" db:"import_cost_id"`
	ProductID       int64           `json:"product_id" db:"product_id"`
	AllocatedAmount decimal.Decimal `json:"allocated_amount" db:"allocated_amount"`

	ImportCost *ImportCost `json:"import_cost,omitempty"`
	Product    *Product    `json:"product,omitempty"`
}

func (a *ImportCostAllocation) IsAllocated() bool {
	return a.AllocatedAmount.GreaterThan(decimal.Zero)
}
