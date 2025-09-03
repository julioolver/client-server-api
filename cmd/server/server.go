package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julioolver/client-server-api/db"
	"github.com/julioolver/client-server-api/repository"
	"github.com/julioolver/client-server-api/service"
)

var svc *service.QuotationService

func main() {
	sqlDB, err := db.ConnectDB()

	if err != nil {
		log.Fatal(err)
	}

	defer sqlDB.Close()

	repo := repository.NewCotationRepository(sqlDB)
	svc = service.NewQuotationService(repo)

	http.HandleFunc("/cotacao", HandleGetCotacao)

	log.Println("HTTP em :8087")
	if err := http.ListenAndServe(":8087", nil); err != nil {
		log.Fatal(err)
	}
}

func HandleGetCotacao(w http.ResponseWriter, r *http.Request) {
	cot, err := svc.GetCotation(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	fmt.Println(cot)
}
