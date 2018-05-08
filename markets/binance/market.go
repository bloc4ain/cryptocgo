package binance

import (
	"log"

	"github.com/bloc4ain/cryptocgo"
	"github.com/bloc4ain/cryptocgo/markets"
	binance "github.com/bloc4ain/go-binance"
)

// Market implements official Binance API
type Market struct {
}

// Name returns market's official name
func (m Market) Name() string {
	return "Binance"
}

// Online returns value indicating whether connection to market API is established
func (m Market) Online() bool {
	return binance.Ping() == nil
}

// Symbols returns currently traded symbols on the market
func (m Market) Symbols() []cryptocgo.Symbol {
	symbols := make([]cryptocgo.Symbol, 0)
	info, err := binance.GetExchangeInfo()

	if err != nil {
		log.Printf("Cannot load symbols from Binance: %s", err)
		return symbols
	}

	for _, s := range info.Symbols {
		symbols = append(symbols, cryptocgo.Symbol{
			BaseAsset:  s.BaseAsset,
			QuoteAsset: s.QuoteAsset,
		})
	}

	return symbols
}

func (m Market) Ticker(s cryptocgo.Symbol) {}

func (m Market) OrderBook(s cryptocgo.Symbol) {}

func (m Market) Trades(i int) []string {
	return nil
}

func init() {
	markets.Add(&Market{})
}
