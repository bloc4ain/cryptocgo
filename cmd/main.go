package main

import (
	"fmt"
	"log"

	"github.com/bloc4ain/cryptocgo/markets"
	_ "github.com/bloc4ain/cryptocgo/markets/binance"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	for _, m := range markets.List() {
		fmt.Println(m.Name())
		fmt.Println(m.Symbols())
	}
}
