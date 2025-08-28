package service

import (
	"context"
	"time"

	"github.com/julioolver/client-server-api/client"
	"github.com/julioolver/client-server-api/model"
)

type QuotationService struct {
	Repository model.QuotationRepository
}

func NewQuotationService(repo model.QuotationRepository) *QuotationService {
	return &QuotationService{Repository: repo}
}

func (quotationService *QuotationService) GetCotation(ctx context.Context) (*model.Cotation, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*400)

	quote, err := client.GetCotation(ctx)

	cancel()

	if err != nil {
		return nil, err
	}

	cotation := model.NewCotation(quote.Base, quote.Bid)

	err = quotationService.Repository.Create(cotation)

	if err != nil {
		return nil, err
	}

	return cotation, nil
}
