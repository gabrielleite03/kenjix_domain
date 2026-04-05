package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type ExpenseStatus string

const (
	ExpenseStatusPaid    ExpenseStatus = "paid"
	ExpenseStatusPending ExpenseStatus = "pending"
)

type ExpenseCategory struct {
	ID        int64     `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type Expense struct {
	ID          int64           `json:"id" db:"id"`
	Description string          `json:"description" db:"description"`
	CategoryID  int64           `json:"category_id" db:"category_id"`
	Amount      decimal.Decimal `json:"amount" db:"amount"`
	Date        time.Time       `json:"date" db:"date"`
	Status      string          `json:"status" db:"status"`
	CreatedAt   time.Time       `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at" db:"updated_at"`

	Category    *ExpenseCategory    `json:"category,omitempty"`
	Attachments []ExpenseAttachment `json:"attachments,omitempty"`
}

type ExpenseAttachment struct {
	ID        int64     `json:"id" db:"id"`
	ExpenseID string    `json:"expense_id" db:"expense_id"`
	URL       string    `json:"url" db:"url"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
