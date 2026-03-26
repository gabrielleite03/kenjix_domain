package model

import "time"

type StockMovementType string

const (
	StockMovementIn         StockMovementType = "IN"
	StockMovementOut        StockMovementType = "OUT"
	StockMovementAdjustment StockMovementType = "ADJUSTMENT"
)

type Stock struct {
	ProductID   int64     `json:"product_id" db:"product_id"`
	WarehouseID int64     `json:"warehouse_id" db:"warehouse_id"`
	Quantity    int       `json:"quantity" db:"quantity"`
	Active      bool      `json:"active" db:"active"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

func (s *Stock) IsActive() bool {
	return s.Active
}

func (s *Stock) Deactivate() {
	s.Active = false
}

type StockMovement struct {
	ID          int64             `json:"id" db:"id"`
	ProductID   int64             `json:"product_id" db:"product_id"`
	WarehouseID int64             `json:"warehouse_id" db:"warehouse_id"`
	Type        StockMovementType `json:"type" db:"type"`
	Quantity    int               `json:"quantity" db:"quantity"`
	CreatedAt   time.Time         `json:"created_at" db:"created_at"`
	Reason      string            `json:"reason" db:"reason"`
}
