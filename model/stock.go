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
	ID               int64     `json:"id" db:"id"`
	ProductID        int64     `json:"productId" db:"product_id"`
	WarehousePlaceID int64     `json:"warehousePlaceId" db:"warehouse_place_id"`
	Quantity         int       `json:"quantity" db:"quantity"`
	Active           bool      `json:"active" db:"active"`
	UpdatedAt        time.Time `json:"updatedAt" db:"updated_at"`
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
	Type             StockMovementType `json:"type" db:"type"`
	Quantity         int               `json:"quantity" db:"quantity"`

	// rastreabilidade
	ReferenceID   *int64  `json:"referenceId,omitempty" db:"reference_id"`     //puchaseID
	ReferenceType *string `json:"referenceType,omitempty" db:"reference_type"` //PURCHASE

	Reason    string    `json:"reason,omitempty" db:"reason"` //Inventário corrigido
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}
