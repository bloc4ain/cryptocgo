package cryptocgo

import (
	"fmt"
)

// Symbol holds symbol assets
type Symbol struct {
	BaseAsset  string
	QuoteAsset string
}

func (s Symbol) String() string {
	return fmt.Sprintf("Symbol[Base:%s;Quote:%s]", s.BaseAsset, s.QuoteAsset)
}

// Market is interface that wraps trading market data synchronization
type Market interface {
	// Name returns market's official name
	Name() string

	// Online returns value indicating whether connection to market API is established
	Online() bool

	// Symbols returns currently traded symbols on the market
	Symbols() []Symbol

	Ticker(Symbol)

	OrderBook(Symbol)

	Trades(int) []string
}
