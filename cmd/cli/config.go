package cli

import (
	"github.com/xoesae/stock-balancer/repository"
	"os"
)

func getStockRepository() repository.StockRepository {
	repository := repository.JsonStockRepository{
		DataFile: os.Getenv("STOCK_FILE"),
	}

	return repository
}

func getBrapiConfig() (string, string) {
	return os.Getenv("BRAPI_TOKEN"), os.Getenv("BRAPI_URL")
}
