package binance

import (
	"errors"
	"fmt"
)

var synchronizers map[string]*orderBookSynch

func init() {
	synchronizers = make(map[string]*orderBookSynch)
}

// StartSync starts single symbol synchronization from binance servers to local database
func StartSync(symbol string) error {
	if symbol == "" {
		return errors.New("Sync aborted due to empty symbol arg")
	}

	synch := &orderBookSynch{symbol: symbol}
	err := synch.start()

	if err != nil {
		return err
	}

	synchronizers[symbol] = synch
	return nil
}

// StopSync stops synchronization for given symbol
func StopSync(symbol string) error {
	if symbol == "" {
		return errors.New("Sync aborted due to empty symbol arg")
	}
	synch, ok := synchronizers[symbol]
	if !ok {
		return fmt.Errorf("No synchronization for symbol %s is running", symbol)
	}
	delete(synchronizers, symbol)
	return synch.stop()
}
