package cryptocgo

type Update struct {
	Market string
	ID     uint64
}

type SymbolUpdate struct {
	Update
	Symbol Symbol
}

type SymbolProcessor interface {
	Init(Controller) error
	Sync([]SymbolUpdate) error
	Process(SymbolUpdate) error
}

type sp struct {
}

func (sp sp) Init(c Controller) {
	c.SyncSymbol(Symbol{BaseAsset: "BTC", QuoteAsset: "TRX"})
	c.SyncMarket("binance", Symbol{BaseAsset: "BTC", QuoteAsset: "TRX"})
	c.SyncMarket("kucoin", Symbol{BaseAsset: "BTC", QuoteAsset: "TRX"})
}

type TickerUpdate struct {
	Update
	Updates Ticker
}

type TickerProcessor interface {
	Init(Controller) error
	Process(TickerUpdate) error
}

type OrderBookUpdate struct {
	Update
	OrderBook OrderBook
}

type OrderBookProcessor interface {
	Init(Controller) error
	Process(OrderBookUpdate) error
}

type TradeUpdate struct {
	Update
	Trade Trade
}

type TradeProcessor interface {
	Init(Controller) error
	Process(TradeUpdate) error
}
