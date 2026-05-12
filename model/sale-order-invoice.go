package model

import "time"

type SalesOrderInvoice struct {
	ID           int64 `db:"id" json:"id"`
	SalesOrderID int64 `db:"sales_order_id" json:"salesOrderId"`

	NFeKey *string `db:"nfe_key" json:"nfeKey,omitempty"`
	Number *int64  `db:"number" json:"number,omitempty"`
	Series *int    `db:"series" json:"series,omitempty"`

	Status       string  `db:"status" json:"status"`
	StatusCode   *string `db:"status_code" json:"statusCode,omitempty"`
	StatusReason *string `db:"status_reason" json:"statusReason,omitempty"`

	Protocol     *string    `db:"protocol" json:"protocol,omitempty"`
	IssuedAt     *time.Time `db:"issued_at" json:"issuedAt,omitempty"`
	AuthorizedAt *time.Time `db:"authorized_at" json:"authorizedAt,omitempty"`
	CancelledAt  *time.Time `db:"cancelled_at" json:"cancelledAt,omitempty"`

	XMLPath *string `db:"xml_path" json:"xmlPath,omitempty"`
	PDFPath *string `db:"pdf_path" json:"pdfPath,omitempty"`

	CreatedAt time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt time.Time `db:"updated_at" json:"updatedAt"`
}
