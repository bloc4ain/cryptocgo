package main

import (
	"github.com/bloc4ain/cryptocgo/webui"
)

func main() {
	// binance.StartSync("TRXBTC")

	// src := binance.NewOrderBookSource("TRXBTC")
	// stream, err := src.Updates()

	// if err != nil {
	// 	panic(err)
	// }

	// for {
	// 	book := <-stream
	// 	fmt.Println("Got new book for symbol", book.Symbol)
	// }

	webui.Start()
}
