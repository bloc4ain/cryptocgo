package storage

import "github.com/bloc4ain/cryptocgo"

type Storage interface {
	AddSymbol(market string, symbol []cryptocgo.Symbol) error
	DeleteSymbol(market string, symbol []cryptocgo.Symbol) error
	GetSymbols(market string) ([]cryptocgo.Symbol, error)

	AddTicker(market string, ticker struct{}) error
	RemoveTicker(market string, ticker struct{}) error
	UpdateTicker(market string, ticker struct{}) error
	GetTickers(market string) ([]struct{}, error)

	AddOrderBook(market string) error
	RemoveOrderBook(market string) error
	UpdateOrderBook(market string) error
	GetOrderBooks(market string) error

	PushOrders(market string, symbol cryptocgo.Symbol, orders []struct{}) error
}
