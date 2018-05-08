package binance

import (
	"github.com/bloc4ain/cryptocgo"
	binance "github.com/bloc4ain/go-binance"
)

// Market implements official Binance API
type Market struct {
}

// Name returns market's official name
func (m Market) Name() string {
	return "Binance"
}

// Symbols returns currently traded symbols on the market
func (m Market) Symbols() ([]cryptocgo.Symbol, error) {
	info, err := binance.GetExchangeInfo()

	if err != nil {
		return nil, err
	}

	res := make([]cryptocgo.Symbol, 0)

	for _, s := range info.Symbols {
		res = append(res, cryptocgo.Symbol{
			BaseAsset:  s.BaseAsset,
			QuoteAsset: s.QuoteAsset,
		})
	}

	return res, nil
}
