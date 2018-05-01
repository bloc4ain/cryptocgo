package main

import (
	"github.com/bloc4ain/cryptocgo/binance"
	"github.com/bloc4ain/cryptocgo/webui"
)

func main() {

	// src := binance.NewOrderBookSource("TRXBTC")
	// stream, err := src.Updates()

	// if err != nil {
	// 	panic(err)
	// }

	// for {
	// 	book := <-stream
	// 	fmt.Println("Got new book for symbol", book.Symbol)
	// }

	binance.StartSync("TRXBTC")
	webui.Start()
}
