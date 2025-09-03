package service

import (
	"context"
	"log"
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
	clientCtx, cancelClient := context.WithTimeout(ctx, time.Millisecond*200)

	quote, err := client.GetCotation(clientCtx)

	cancelClient()

	if err != nil {
		return nil, err
	}

	cotationModel := model.NewCotation(quote.Base, quote.Bid)

	ctxDB, cancelDB := context.WithTimeout(ctx, time.Millisecond*10)
	defer cancelDB()

	if err := quotationService.Repository.Create(ctxDB, cotationModel); err != nil {
		log.Printf("[WARN] não foi possível salvar cotação %s: %v", cotationModel.Id, err)
	}

	return cotationModel, nil
}
