package model

type QuotationRepository interface {
	GetCotation() (*Cotation, error)
	Create(cotation *Cotation) error
}
