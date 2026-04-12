package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type Marketplace struct {
	ID   int64   `json:"id" db:"id"`
	Name string  `json:"name" db:"name"`
	Logo *string `json:"logo,omitempty" db:"logo"`

	Status          string          `json:"status" db:"status"`
	CommissionRate  decimal.Decimal `json:"commissionRate" db:"commission_rate"`
	IntegrationType string          `json:"integrationType" db:"integration_type"`

	APIURL      *string `json:"apiUrl,omitempty" db:"api_url"`
	APIKey      *string `json:"apiKey,omitempty" db:"api_key"`
	APISecret   *string `json:"apiSecret,omitempty" db:"api_secret"`
	APIEndpoint *string `json:"apiEndpoint,omitempty" db:"api_endpoint"`

	CreatedAt time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time  `json:"updatedAt" db:"updated_at"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" db:"deleted_at"`
}
