package brapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/xoesae/stock-balancer/entity"
	"log"
	"net/http"
)

type Brapi struct {
	Token  string
	Url    string
	client *http.Client
}

func (b Brapi) GetStockDetails(stock entity.Stock) (StockData, error) {
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
		return StockData{}, err
	}

	if len(response.Results) == 0 {
		return StockData{}, fmt.Errorf("not found")
	}

	return response.Results[0], nil
}
