package main

import (
	"log"

	"github.com/bloc4ain/cryptocgo/app"
	"github.com/bloc4ain/cryptocgo/markets/binance"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
}

func main() {
	app.AddMarket(binance.Market{})
	app.Run()
}
