package market

import (
	"github.com/bloc4ain/cryptocgo/binance"
	"github.com/bloc4ain/cryptocgo/order"
)

// Card represents cryptocurrency market brief info
type Card interface {
	Title() string
	OrderBook() *order.Book
	Updates() <-chan *order.Book
}

// BinanceCard represents Binance market brief info
type BinanceCard struct {
}

// Title returns Binance title
func (c BinanceCard) Title() string {
	return "Binance"
}

// OrderBook returns Binance top most expensive currencies
func (c BinanceCard) OrderBook() *order.Book {
	src := binance.NewOrderBookSource("TRXBTC")
	book, _ := src.Book()
	return book
}

// Updates subscribes for top most expensive currencies updates
func (c BinanceCard) Updates() <-chan *order.Book {
	return make(chan *order.Book)
}

// BittrexCard represents Bittrex market brief info
type BittrexCard struct {
}

// Title returns Bittrex title
func (c BittrexCard) Title() string {
	return "Bittrex"
}

// OrderBook returns Bittrex top most expensive currencies
func (c BittrexCard) OrderBook() *order.Book {
	return nil
}

// Updates subscribes for top most expensive currencies updates
func (c BittrexCard) Updates() <-chan *order.Book {
	return make(chan *order.Book)
}

// KuCoinCard represents KuCoin market brief info
type KuCoinCard struct {
}

// Title returns KuCoin title
func (c KuCoinCard) Title() string {
	return "KuCoin"
}

// OrderBook returns KuCoin top most expensive currencies
func (c KuCoinCard) OrderBook() *order.Book {
	return nil
}

// Updates subscribes for top most expensive currencies updates
func (c KuCoinCard) Updates() <-chan *order.Book {
	return make(chan *order.Book)
}
