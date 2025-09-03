package model

import (
	"time"

	"github.com/google/uuid"
)

type Cotation struct {
	Code      string    `json:"code"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Bid       float64   `json:"bid"`
	CreatedAt time.Time `json:"timestamp"`
}

func NewCotation(name string, bid float64) *Cotation {
	return &Cotation{
		Id:        uuid.NewString(),
		Name:      name,
		Bid:       bid,
		CreatedAt: time.Now(),
	}
}
