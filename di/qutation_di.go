package di

import (
	"log"

	"github.com/julioolver/client-server-api/db"
	"github.com/julioolver/client-server-api/repository"
)

type QuotationDi struct {
	QuotationRepo *repository.QuotationRepository
}

func (QuotationDi) NewQuotationDi() *QuotationDi {
	db, err := db.ConnectDB()

	if err != nil {
		log.Fatal(db)
	}

	var di QuotationDi

	*di.QuotationRepo = *repository.NewCotationRepository(db)

	return &di
}
