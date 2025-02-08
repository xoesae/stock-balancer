package cli

import (
	"flag"
	"fmt"
	"github.com/xoesae/stock-balancer/pkg/brapi"
	"github.com/xoesae/stock-balancer/service"
	"strconv"
	"strings"
)

func Run() {
	stockRepository := getStockRepository()
	stockService := service.StockService{Repository: stockRepository}
	token, url := getBrapiConfig()
	brapiClient := brapi.NewBrapiClient(token, url)

	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Println("Investment not set!")
		return
	}

	arg := flag.Arg(0)
	value, _ := strconv.ParseFloat(strings.TrimSpace(arg), 64)

	stocks, err := stockRepository.GetAll()
	if err != nil {
		fmt.Println(err)
	}

	stocks, err = stockService.UpdateAll(brapiClient, stocks)
	if err != nil {
		fmt.Println(err)
	}

	result := stockService.BalancePortfolio(stocks, value)

	for ticker, amount := range result.Buys {
		fmt.Printf("[%s] Comprar %d\n", ticker, amount)
	}

	fmt.Printf("Resta: %.2f\n", result.Remaining)
}
