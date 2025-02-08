package entity

import "time"

type Stock struct {
	Ticker       string    `json:"ticker"`
	IdealRatio   float64   `json:"ideal_ratio"`
	CurrentPrice float64   `json:"current_price"`
	Amount       int       `json:"amount"`
	UpdatedAt    time.Time `json:"updated_at"`
}
