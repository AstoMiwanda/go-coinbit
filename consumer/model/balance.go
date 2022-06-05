package model

import (
	"github.com/google/uuid"
	"time"
)

type Balance struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	BalanceAPI
}

type BalanceAPI struct {
	WalletID *uuid.UUID `json:"wallet_id,omitempty" gorm:"primaryKey"`
	Amount   *float64   `json:"amount,omitempty"`
}
