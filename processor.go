package cryptocgo

// Processor wraps raw market data processing functions
type Processor interface {
	ProcessSymbol(Symbol) error
	ProcessTicker(Ticker) error
	ProcessOrderBook(OrderBook) error
	ProcessTrade(Trade) error
}
