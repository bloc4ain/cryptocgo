package cryptocgo

// Market is interface that wraps trading market data synchronization
type Market interface {
	// API
	Name() string
	Symbols() ([]Symbol, error)
	// Ticker(Symbol) (Ticker, error)
	// OrderBook(Symbol) (OrderBook, error)
	// LastTrades() ([]Trade, error)

	// // Events
	// SymbolUpdates(chan<- Symbol) error
	// TickerUpdates(Symbol, chan<- Ticker) error
	// OrderBookUpdates(Symbol, chan<- OrderBook) error
	// TradeStream(chan<- []Trade) error

	// // Controls
	// Metrics()
	// StartSynch() error
	// StopSynch() error
}
