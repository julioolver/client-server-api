package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/julioolver/client-server-api/service"
)

type responseEnvelope struct {
	Currency QuotationResponse `json:"USDBRL"`
}

type QuotationResponse struct {
	Code       string `json:"code"`
	Codein     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
}

func GetCotation(ctx context.Context) (service.Quote, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)

	if err != nil {
		return service.Quote{}, err
	}

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return service.Quote{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return service.Quote{}, fmt.Errorf("status %d", resp.StatusCode)
	}

	var data responseEnvelope

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return service.Quote{}, err
	}

	bidConverted, err := strconv.ParseFloat(data.Currency.Bid, 64)

	if err != nil {
		bidConverted = 0
	}

	return service.Quote{
		Base:      data.Currency.Bid,
		Bid:       bidConverted,
		Quote:     "",
		FetchedAt: time.Now(),
	}, nil
}
