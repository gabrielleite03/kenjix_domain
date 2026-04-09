package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type Purchase struct {
	ID            int64           `json:"id" db:"id"`
	InvoiceNumber *string         `json:"invoiceNumber,omitempty" db:"invoice_number"`
	InvoiceType   string          `json:"invoiceType" db:"invoice_type"`
	SupplierID    int64           `json:"supplierId" db:"supplier_id"`
	Status        string          `json:"status" db:"status"`
	Total         decimal.Decimal `json:"total" db:"total"`
	CreatedAt     time.Time       `json:"createdAt" db:"created_at"`
	Items         []PurchaseItem  `json:"items,omitempty"`
}

type PurchaseItem struct {
	ID           int64           `json:"id" db:"id"`
	PurchaseID   int64           `json:"purchaseId" db:"purchase_id"`
	ProductID    int64           `json:"productId" db:"product_id"`
	Quantity     decimal.Decimal `json:"quantity" db:"quantity"`
	CostPrice    decimal.Decimal `json:"costPrice" db:"cost_price"`
	Total        decimal.Decimal `json:"total" db:"total"`
	CostCenterID *int64          `json:"costCenterId,omitempty" db:"cost_center_id"`
}
