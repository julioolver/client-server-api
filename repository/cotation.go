package repository

import (
	"database/sql"

	"github.com/julioolver/client-server-api/model"
)

type QuotationRepository struct {
	db *sql.DB
}

func NewCotationRepository(db *sql.DB) *QuotationRepository {
	return &QuotationRepository{db: db}
}

func (repository *QuotationRepository) GetCotation() (*model.Cotation, error) {
	return nil, nil
}

func (repository *QuotationRepository) Create(cotation *model.Cotation) error {
	return nil
}
