package service

import (
	"github.com/xoesae/stock-balancer/internal/entity"
	"github.com/xoesae/stock-balancer/internal/repository"
	"time"
)

type MarketApi interface {
	GetStockDetails(stock entity.Stock) (entity.Stock, error)
}

type StockService struct {
	Repository repository.StockRepository
	Api        MarketApi
}

func (s *StockService) GetAll() ([]entity.Stock, error) {
	return s.Repository.GetAll()
}

func (s *StockService) UpdateAll(stocks []entity.Stock) ([]entity.Stock, error) {
	var updatedStocks []entity.Stock

	for _, stock := range stocks {
		if stock.UpdatedAt.After(time.Now().Add(-2 * time.Hour)) {
			updatedStocks = append(updatedStocks, stock)
			continue
		}

		data, err := s.Api.GetStockDetails(stock)
		if err != nil {
			updatedStocks = append(updatedStocks, data)
			continue
		}

		updatedStocks = append(updatedStocks, data)
	}

	err := s.Repository.Save(updatedStocks)
	if err != nil {
		return updatedStocks, err
	}

	return updatedStocks, nil
}
