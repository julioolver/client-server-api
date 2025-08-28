package di

import (
	"log"

	"github.com/julioolver/client-server-api/db"
	"github.com/julioolver/client-server-api/repository"
	"github.com/julioolver/client-server-api/service"
)

type QuotationDi struct {
	service *service.QuotationService
}

func (QuotationDi) NewQuotationDi() *QuotationDi {
	db, err := db.ConnectDB()

	if err != nil {
		log.Fatal(err)
	}

	di := QuotationDi{
		service: service.NewQuotationService(repository.NewCotationRepository(db)),
	}

	return &di
}
