package model

import "time"

type Quotation struct {
	Bid       float64   `json:"bid"`
	CreatedAt time.Time `json:"created_at"`
}
