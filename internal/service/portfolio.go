package service

import (
	"github.com/xoesae/stock-balancer/internal/entity"
	"github.com/xoesae/stock-balancer/internal/repository"
	"math"
	"sort"
)

type buysPerTicker map[string]int

type stockPriority struct {
	stock entity.Stock
	rank  float64
}

type PortfolioService struct {
	Repository repository.StockRepository
}

func (s *PortfolioService) calculatePortfolioValue(stocks []entity.Stock) float64 {
	total := 0.0

	for _, stock := range stocks {
		total += stock.CurrentPrice * float64(stock.Amount)
	}

	return total
}

func (s *PortfolioService) calculateIdealValuePerTicker(stocks []entity.Stock, total float64) map[string]float64 {
	values := make(map[string]float64)

	for _, stock := range stocks {
		values[stock.Ticker] = stock.IdealRatio * total
	}

	return values
}

func (s *PortfolioService) makePriorityList(stocks []entity.Stock, buys buysPerTicker, idealValues map[string]float64, remaining float64) []stockPriority {
	priorities := make([]stockPriority, len(stocks))

	for i, stock := range stocks {
		currentShares := buys[stock.Ticker]
		currentValue := stock.CurrentPrice * (float64(stock.Amount) + float64(currentShares))
		idealValue := idealValues[stock.Ticker]
		buyAmount := idealValue - currentValue

		improvement := 0.0
		if buyAmount > 0 && stock.CurrentPrice > 0 && stock.CurrentPrice <= remaining {
			newValue := currentValue + stock.CurrentPrice
			deviationBefore := math.Abs(idealValue - currentValue)
			deviationAfter := math.Abs(idealValue - newValue)
			improvement = (deviationBefore - deviationAfter) / stock.CurrentPrice

			// Penalization for stocks that have many buys
			improvement = improvement * (1.0 / (1.0 + float64(currentShares)))
		}

		priorities[i] = stockPriority{stock: stock, rank: improvement}
	}

	return priorities
}

func (s *PortfolioService) sortPriorities(priorities []stockPriority) []stockPriority {
	sort.Slice(priorities, func(i, j int) bool {
		// if they have same rank, order by price
		if priorities[i].rank == priorities[j].rank {
			return priorities[i].stock.CurrentPrice < priorities[j].stock.CurrentPrice
		}

		// else, sort by rank
		return priorities[i].rank > priorities[j].rank
	})

	return priorities
}

func (s *PortfolioService) BalancePortfolio(stocks []entity.Stock, investment float64) entity.BalanceResult {
	total := s.calculatePortfolioValue(stocks)
	totalAfterInvestment := total + investment

	idealValues := s.calculateIdealValuePerTicker(stocks, totalAfterInvestment)

	buys := make(buysPerTicker)
	remaining := investment

	// Buy while remaining investment
	for remaining > 0 {
		priorities := s.makePriorityList(stocks, buys, idealValues, remaining)
		priorities = s.sortPriorities(priorities)

		// Try to buy the stock that have the best priority
		bought := false
		for _, priority := range priorities {
			if priority.rank > 0 && priority.stock.CurrentPrice <= remaining {
				buys[priority.stock.Ticker]++
				remaining -= priority.stock.CurrentPrice
				bought = true
				break
			}
		}

		if !bought {
			break
		}
	}

	return entity.BalanceResult{Buys: buys, Remaining: remaining}
}
