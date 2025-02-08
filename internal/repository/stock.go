package repository

import (
	"encoding/json"
	"fmt"
	"github.com/xoesae/stock-balancer/internal/entity"
	"os"
)

type StockRepository interface {
	GetAll() ([]entity.Stock, error)
	Save([]entity.Stock) error
}

type JsonStockRepository struct {
	DataFile string
}

func (r JsonStockRepository) GetAll() ([]entity.Stock, error) {
	var stocks []entity.Stock
	file, err := os.ReadFile(r.DataFile)
	if err != nil {
		return stocks, err
	}

	var data struct {
		Stocks []entity.Stock `json:"stocks"`
	}

	err = json.Unmarshal(file, &data)

	return data.Stocks, err
}

func (r JsonStockRepository) Save(stocks []entity.Stock) error {
	var data struct {
		Stocks []entity.Stock `json:"stocks"`
	}
	data.Stocks = stocks

	serializedData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("error on serialize: %w", err)
	}

	if err := os.WriteFile(r.DataFile, serializedData, 0644); err != nil {
		return fmt.Errorf("error on write file: %w", err)
	}

	return nil
}
