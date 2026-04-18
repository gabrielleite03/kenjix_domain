package model

import (
	"time"
)

type StockMovementType string

const (
	StockMovementIn         StockMovementType = "IN"
	StockMovementOut        StockMovementType = "OUT"
	StockMovementAdjustment StockMovementType = "ADJUSTMENT"
)

type Stock struct {
	ID             int64          `json:"id" db:"id"`
	Product        Product        `json:"product" db:"product_id"`
	WarehousePlace WarehousePlace `json:"warehousePlace" db:"warehouse_place_id"`
	PurchaseItem   PurchaseItem   `json:"purchaseItem" db:"purchase_item_id"`
	Quantity       int            `json:"quantity" db:"quantity"`
	Active         bool           `json:"active" db:"active"`
	UpdatedAt      time.Time      `json:"updatedAt" db:"updated_at"`
}

func (s *Stock) IsActive() bool {
	return s.Active
}

func (s *Stock) Deactivate() {
	s.Active = false
}

func (s *Stock) Add(quantity int) {
	s.Quantity += quantity
}

func (s *Stock) Remove(quantity int) {
	s.Quantity -= quantity
}

type StockMovement struct {
	ID               int64             `json:"id" db:"id"`
	ProductID        int64             `json:"productId" db:"product_id"`
	WarehousePlaceID int64             `json:"warehousePlaceId" db:"warehouse_place_id"`
	PurchseItemID    int64             `json:"purchaseItemId" db:"purchase_item_id"`
	Type             StockMovementType `json:"type" db:"type"`
	Quantity         int               `json:"quantity" db:"quantity"`

	// rastreabilidade
	ReferenceID   *int64  `json:"referenceId,omitempty" db:"reference_id"`     //puchaseID
	ReferenceType *string `json:"referenceType,omitempty" db:"reference_type"` //PURCHASE

	Reason    string    `json:"reason,omitempty" db:"reason"` //Inventário corrigido
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}

type StockMovementEager struct {
	ID int64 `json:"id" db:"id"`

	Product        Product        `json:"product" db:"product"`
	WarehousePlace WarehousePlace `json:"warehouse_place" db:"warehouse_place"`
	PurchaseItem   PurchaseItem   `json:"purchase_item" db:"purchase_item"`

	Type     StockMovementType `json:"type" db:"type"`
	Quantity int               `json:"quantity" db:"quantity"`

	ReferenceID   *int64  `json:"reference_id,omitempty" db:"reference_id"`
	ReferenceType *string `json:"reference_type,omitempty" db:"reference_type"`

	Reason *string `json:"reason,omitempty" db:"reason"`

	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
