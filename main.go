package main

import (
	"github.com/joho/godotenv"
	"github.com/xoesae/stock-balancer/cmd/cli"
)

func main() {
	godotenv.Load()

	cli.Run()
}
