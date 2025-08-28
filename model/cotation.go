package model

import "time"

type Cotation struct {
	Code      string  `json:"code"`
	Name      string  `json:"name"`
	Bid       float64 `json:"bid"`
	Timestamp string  `json:"timestamp"`
}

func NewCotation(name string, bid float64) *Cotation {
	return &Cotation{
		Name:      name,
		Bid:       bid,
		Timestamp: time.Now().Format("y-m-d h:m"),
	}
}
