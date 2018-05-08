package app

import (
	"log"
	"sync"
	"time"

	"github.com/bloc4ain/cryptocgo"
)

type marketSymbols struct {
	market  string
	symbols []cryptocgo.Symbol
}

func syncSymbols() {
	log.Printf("Symbol sync started for %d markets", len(markets))

	var wg sync.WaitGroup
	var ms = make([]marketSymbols, 0)
	var mx = new(sync.Mutex)
	var st = time.Now()

	for _, m := range markets {
		wg.Add(1)

		go func(m cryptocgo.Market) {
			defer wg.Done()

			if s, err := m.Symbols(); err == nil {
				mx.Lock()
				ms = append(ms, marketSymbols{market: m.Name(), symbols: s})
				mx.Unlock()
				log.Printf("Symbol sync for market [%s] completed", m.Name())
			} else {
				log.Printf("Could not fetch symbols for market [%s]: %s", m.Name(), err)
			}
		}(m)
	}

	wg.Wait()
	log.Printf("Sync completed in %s", shortDuration(time.Since(st)))
}
