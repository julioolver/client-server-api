package di

import (
	"log"

	"github.com/julioolver/client-server-api/db"
	"github.com/julioolver/client-server-api/repository"
	"github.com/julioolver/client-server-api/service"
)

type QuotationDi struct {
	Service *service.QuotationService
}

func NewQuotationDi() *QuotationDi {
	db, err := db.ConnectDB()

	if err != nil {
		log.Fatal(err)
	}

	di := QuotationDi{
		Service: service.NewQuotationService(repository.NewCotationRepository(db)),
	}

	return &di
}
