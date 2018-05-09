package app

import (
	"github.com/bloc4ain/cryptocgo"
)

var markets = make(map[string]cryptocgo.Market)

// var processors = make([]cryptocgo.Processor, 0)

// AddMarket registers market in the system
func AddMarket(m cryptocgo.Market) {
	if m == nil {
		panic("Cannot add nil market")
	}
	if _, exists := markets[m.Name()]; exists {
		panic(`Market with name "` + m.Name() + `" already exists`)
	}
	markets[m.Name()] = m
}

// MarketList returns list of all registered markets in the system
func MarketList() []cryptocgo.Market {
	list := make([]cryptocgo.Market, 0)
	for _, m := range markets {
		list = append(list, m)
	}
	return list
}

// AddProcessor registers processor in the system
// func AddProcessor(p cryptocgo.Processor) {
// 	if p == nil {
// 		panic("Cannot add nil processor")
// 	}
// }

// // ProcessorList returns list of all registered processors in the system
// func ProcessorList() []cryptocgo.Processor {
// 	r := make([]cryptocgo.Processor, len(processors))
// 	copy(r, processors)
// 	return r
// }

// Run starts the controller process
func Run() {
	syncSymbols()
}
