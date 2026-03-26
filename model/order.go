package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type PurchaseStatus struct {
	ID          int64   `json:"id" db:"id"`
	Name        string  `json:"name" db:"name"`
	Description *string `json:"description,omitempty" db:"description"`
	Active      bool    `json:"active" db:"active"`
}

func (ps *PurchaseStatus) IsActive() bool {
	return ps.Active
}
func (ps *PurchaseStatus) Deactivate() {
	ps.Active = false
}

type FiscalNumberType struct {
	ID          int64   `json:"id" db:"id"`
	Name        string  `json:"name" db:"name"`
	Description *string `json:"description,omitempty" db:"description"`
	Active      bool    `json:"active" db:"active"`
}

func (fnt *FiscalNumberType) IsActive() bool {
	return fnt.Active
}
func (fnt *FiscalNumberType) Deactivate() {
	fnt.Active = false
}

type PurchaseOrder struct {
	ID                 int64     `json:"id" db:"id"`
	PurchaseStatusID   int64     `json:"purchase_status_id" db:"purchase_status_id"`
	FiscalNumberTypeID int64     `json:"fiscal_number_type_id" db:"fiscal_number_type_id"`
	SupplierID         int64     `json:"supplier_id" db:"supplier_id"`
	FiscalNumber       *string   `json:"fiscal_number,omitempty" db:"fiscal_number"`
	Status             string    `json:"status" db:"status"`
	CreatedAt          time.Time `json:"created_at" db:"created_at"`
	Active             bool      `json:"active" db:"active"`

	PurchaseStatus   *PurchaseStatus   `json:"purchase_status,omitempty"`
	FiscalNumberType *FiscalNumberType `json:"fiscal_number_type,omitempty"`
	Supplier         *Supplier         `json:"supplier,omitempty"`
}

func (po *PurchaseOrder) IsActive() bool {
	return po.Active
}
func (po *PurchaseOrder) Deactivate() {
	po.Active = false
}

type PurchaseOrderItem struct {
	PurchaseOrderID int64           `json:"purchase_order_id" db:"purchase_order_id"`
	ProductID       int64           `json:"product_id" db:"product_id"`
	Quantity        int             `json:"quantity" db:"quantity"`
	UnitPrice       decimal.Decimal `json:"unit_price" db:"unit_price"`
	Active          bool            `json:"active" db:"active"`

	PurchaseOrder *PurchaseOrder `json:"purchase_order,omitempty"`
	Product       *Product       `json:"product,omitempty"`
}

func (poi *PurchaseOrderItem) IsActive() bool {
	return poi.Active
}

func (poi *PurchaseOrderItem) Deactivate() {
	poi.Active = false
}

func (i *PurchaseOrderItem) Total() decimal.Decimal {
	return i.UnitPrice.Mul(decimal.NewFromInt(int64(i.Quantity)))
}
