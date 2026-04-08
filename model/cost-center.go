package model

import "github.com/shopspring/decimal"

type CostCenter struct {
	ID          string               `json:"id" db:"id"`
	Name        string               `json:"name" db:"name"`
	Code        string               `json:"code" db:"code"`
	Description string               `json:"description" db:"description"`
	Active      bool                 `json:"active" db:"active"`
	Properties  []CostCenterProperty `json:"properties,omitempty"`
}

type CostCenterPropertyType string

const (
	CostCenterPropertyTypeIndex CostCenterPropertyType = "index"
	CostCenterPropertyTypeValue CostCenterPropertyType = "value"
)

type CostCenterProperty struct {
	ID           int64                  `json:"id" db:"id"`
	CostCenterID string                 `json:"costCenterId" db:"cost_center_id"`
	Name         string                 `json:"name" db:"name"`
	Value        decimal.Decimal        `json:"value" db:"value"`
	Type         CostCenterPropertyType `json:"type" db:"type"`
}
