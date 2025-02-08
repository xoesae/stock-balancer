package cli

import (
	"flag"
	"fmt"
	"github.com/xoesae/stock-balancer/internal/entity"
	"strconv"
	"strings"
)

func Run() {
	stockService := getStockService()
	portfolioService := getPortfolioService()

	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Println("Investment not set!")
		return
	}

	arg := flag.Arg(0)
	value, _ := strconv.ParseFloat(strings.TrimSpace(arg), 64)

	stocks, err := stockService.GetAll()
	if err != nil {
		fmt.Println(err)
	}

	stocks, err = stockService.UpdateAll(stocks)
	if err != nil {
		fmt.Println(err)
	}

	result := portfolioService.BalancePortfolio(stocks, value)
	soma := 0.0

	for ticker, amount := range result.Buys {
		fmt.Printf("[%s] Comprar %d\n", ticker, amount)
		var st entity.Stock
		for _, s := range stocks {
			if s.Ticker == ticker {
				st = s
			}
		}
		soma += float64(amount) * st.CurrentPrice
	}

	fmt.Printf("Resta: %.2f | Soma %.2f\n", result.Remaining, soma)
}
