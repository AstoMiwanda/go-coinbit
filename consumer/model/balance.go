package model

import (
	"github.com/google/uuid"
	"time"
)

type Balance struct {
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	IsThreshold *bool      `json:"is_threshold,omitempty"`
	BalanceAPI
}

type BalanceAPI struct {
	WalletID *uuid.UUID `json:"wallet_id,omitempty" gorm:"primaryKey"`
	Amount   *float64   `json:"amount,omitempty"`
}
