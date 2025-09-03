package model

import "context"

type QuotationRepository interface {
	GetCotation() (*Cotation, error)
	Create(ctx context.Context, cotation *Cotation) error
}
