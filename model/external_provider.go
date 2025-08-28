package model

import (
	"context"
	"time"
)

type Quote struct {
	Base      string
	Quote     string
	Bid       float64
	FetchedAt time.Time
}

type ExternalProvider interface {
	GetExternalCotation(ctx context.Context) (Quote, error)
}
