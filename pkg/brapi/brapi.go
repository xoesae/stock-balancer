package brapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/xoesae/stock-balancer/internal/entity"
	"log"
	"net/http"
	"time"
)

type Brapi struct {
	Token  string
	Url    string
	client *http.Client
}

func (b Brapi) GetStockDetails(stock entity.Stock) (entity.Stock, error) {
	var body []byte

	url := fmt.Sprintf("%s/quote/%s?token=%s", b.Url, stock.Ticker, b.Token)

	req, err := http.NewRequest("GET", url, bytes.NewBuffer(body))
	if err != nil {
		log.Fatal(err)
	}

	res, err := b.client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	var response BrapiResponse
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		log.Fatal("Erro ao decodificar resposta JSON:", err)
		return entity.Stock{}, err
	}

	if len(response.Results) == 0 {
		return entity.Stock{}, fmt.Errorf("not found")
	}

	data := response.Results[0]

	return entity.Stock{
		Ticker:       stock.Ticker,
		IdealRatio:   stock.IdealRatio,
		CurrentPrice: data.RegularMarketPrice,
		Amount:       stock.Amount,
		UpdatedAt:    time.Now(),
	}, nil
}
