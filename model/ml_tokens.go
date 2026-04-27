package model

import "time"

type MLToken struct {
	ID       int64   `db:"id" json:"id"`
	UserID   int64   `db:"user_id" json:"user_id"`
	Nickname *string `db:"nickname" json:"nickname,omitempty"`

	AccessToken  string `db:"access_token" json:"access_token"`
	RefreshToken string `db:"refresh_token" json:"refresh_token"`

	ExpiresAt time.Time `db:"expires_at" json:"expires_at"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
