package markets

import (
	"log"

	"github.com/bloc4ain/cryptocgo"
)

var markets = make(map[string]cryptocgo.Market)

// Add registers market to the system
func Add(m cryptocgo.Market) {
	if _, exists := markets[m.Name()]; exists {
		log.Printf(`Market with name "%s" already exists`, m.Name())
	}
	markets[m.Name()] = m
}

// List returns list of currently registered markets
func List() []cryptocgo.Market {
	list := make([]cryptocgo.Market, 0)
	for _, m := range markets {
		list = append(list, m)
	}
	return list
}
