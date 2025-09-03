package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

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

func (repository *QuotationRepository) Create(ctx context.Context, cotation *model.Cotation) error {
	time.Sleep(time.Second * 5)
	_, err := repository.db.ExecContext(ctx, "INSERT INTO quotations (id, bid, name) VALUES (?, ?, ?)", cotation.Id, cotation.Bid, cotation.Name)

	if err != nil {
		return fmt.Errorf("insert quotation: %w", err)
	}

	fmt.Println("entrou aqui")
	return nil
}
