package model

type Marketplace struct {
	ID              int64   `json:"id" db:"id"`
	Name            string  `json:"name" db:"name"`
	Logo            string  `json:"logo" db:"logo"`
	ComissionRate   float64 `json:"comission_rate" db:"comission_rate"`
	IntegrationType string  `json:"integration_type" db:"integration_type"`
	ApiKey          string  `json:"api_key" db:"api_key"`
	ApiSecret       string  `json:"api_secret" db:"api_secret"`
	ApiEndpoint     string  `json:"api_endpoint" db:"api_endpoint"`
	Active          bool    `json:"active" db:"active"`
}
