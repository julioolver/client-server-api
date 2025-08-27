package service

import (
	"github.com/julioolver/client-server-api/di"
	"github.com/julioolver/client-server-api/model"
)

type QuotationService struct {
	Repository model.QuotationRepository
}

func NewQuotationService(quotationDi di.QuotationDi) *QuotationService {
	return &QuotationService{Repository: quotationDi.QuotationRepo}
}

func (quotationService *QuotationService) getCotation() (*model.Cotation, error) {
}
