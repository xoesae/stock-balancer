package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/xoesae/stock-balancer/internal/entity"
	"os"
	"strconv"
)

func Run() {
	var rootCmd = &cobra.Command{
		Use:   "portfolio",
		Short: "CLI para gerenciar sua carteira de ações",
	}

	var listCmd = &cobra.Command{
		Use:   "list",
		Short: "Lista as ações da carteira",
		Run: func(cmd *cobra.Command, args []string) {
			listStocks()
		},
	}

	var rebalanceCmd = &cobra.Command{
		Use:   "rebalance [valor]",
		Short: "Rebalanceia a carteira com o valor passado",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			value, err := strconv.ParseFloat(args[0], 64)
			if err != nil {
				fmt.Println("Erro: valor inválido. Use um número como 1000.50")
				os.Exit(1)
			}
			fmt.Printf("Rebalanceando carteira com R$ %.2f\n", value)

			balancePortfolio(value)
		},
	}

	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(rebalanceCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func listStocks() {
	stockService := getStockService()
	stocks, err := stockService.GetAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	separator := "+----------+----------+------------------+-----------+---------------------+"

	fmt.Println(separator)
	fmt.Printf("|%-10s| %-9s| %-15s | %-10s| %-20s|\n", "Ticker", "Ideal(%)", "Preço Atual (R$)", "Qtd", "Atualizado em")
	fmt.Println(separator)

	for _, stock := range stocks {
		fmt.Printf("|%-10s| %-9.2f| %-16.2f | %-10d| %-20s|\n",
			stock.Ticker,
			stock.IdealRatio*100,
			stock.CurrentPrice,
			stock.Amount,
			stock.UpdatedAt.Format("2006-01-02 15:04"),
		)
	}

	fmt.Println(separator)
}

func balancePortfolio(value float64) {
	stockService := getStockService()
	portfolioService := getPortfolioService()

	stocks, err := stockService.GetAll()
	if err != nil {
		fmt.Println(err)
	}

	stocks, err = stockService.UpdateAll(stocks)
	if err != nil {
		fmt.Println(err)
	}

	result := portfolioService.BalancePortfolio(stocks, value)
	sum := 0.0

	separator := "+----------+------------+------------+------------+"
	fmt.Println(separator)
	fmt.Printf("|%-10s| %-11s| %-11s| %-11s|\n", "Ticker", "Quantidade", "Preço (R$)", "Total")
	fmt.Println(separator)

	for ticker, amount := range result.Buys {
		var st entity.Stock
		for _, s := range stocks {
			if s.Ticker == ticker {
				st = s
			}
		}
		total := float64(amount) * st.CurrentPrice
		sum += total

		fmt.Printf("|%-10s| %-11d| %-11.2f| %-11.2f|\n",
			st.Ticker,
			amount,
			st.CurrentPrice,
			total,
		)
	}

	fmt.Println(separator)
	fmt.Printf("\nResta: %.2f | Soma %.2f\n", result.Remaining, sum)
}
