package cli

import (
	"github.com/xoesae/stock-balancer/internal/repository"
	"github.com/xoesae/stock-balancer/internal/service"
	"github.com/xoesae/stock-balancer/pkg/brapi"
	"os"
)

func getBrapi() brapi.Brapi {
	return brapi.NewBrapiClient(
		os.Getenv("BRAPI_TOKEN"),
		os.Getenv("BRAPI_URL"),
	)
}

func getStockRepository() repository.StockRepository {
	return repository.JsonStockRepository{
		DataFile: os.Getenv("STOCK_FILE"),
	}
}

func getStockService() service.StockService {
	return service.StockService{
		Repository: getStockRepository(),
		Api:        getBrapi(),
	}
}

func getPortfolioService() service.PortfolioService {
	return service.PortfolioService{
		Repository: getStockRepository(),
	}
}
