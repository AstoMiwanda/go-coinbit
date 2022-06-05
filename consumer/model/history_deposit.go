package model

import "time"

type HistoryDeposit struct {
	Amount    *float64   `json:"amount,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	Time      *time.Time `json:"time,omitempty"`
}
