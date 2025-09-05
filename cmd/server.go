package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/julioolver/client-server-api/database"
	"github.com/julioolver/client-server-api/model"
)

type App struct {
	db *sql.DB
}

func (app App) GetCotation(w http.ResponseWriter, req *http.Request) {
	ctxReq, cancelReq := context.WithTimeout(req.Context(), time.Millisecond*500)
	defer cancelReq()

	request, err := http.NewRequestWithContext(ctxReq, http.MethodGet, "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)

	if err != nil {
		panic(err)
	}

	resp, err := http.DefaultClient.Do(request)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	var q model.QuotationClient

	json.NewDecoder(resp.Body).Decode(&q)

	fmt.Println(q.USDBRL.Bid)
}

func main() {
	db, err := database.GetConnection()

	if err != nil {
		panic(err)
	}

	defer db.Close()

	app := &App{
		db: db,
	}

	http.HandleFunc("/cotacao", app.GetCotation)

	log.Fatal(http.ListenAndServe(":8083", nil))
}
