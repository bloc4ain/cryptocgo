package binance

import (
	"log"

	"github.com/bloc4ain/cryptocgo/order"
	binance "github.com/bloc4ain/go-binance"
	r "gopkg.in/gorethink/gorethink.v4"
)

// OrderBookSource manages binance order book state
type OrderBookSource struct {
	symbol  string
	updates chan order.Book
	done    chan struct{}
	err     error
}

// Book returns current book value
func (obs *OrderBookSource) Book() (*order.Book, error) {
	book := new(order.Book)
	err := r.DB("binance").Table("orderBooks").GetAllByIndex("symbol", obs.symbol).ReadOne(book, dbSession)
	return book, err
}

// Updates subscribes for book updates if not subscribed and returns book stream
func (obs *OrderBookSource) Updates() (<-chan order.Book, error) {
	if obs.updates != nil {
		return obs.updates, nil
	}

	query, err := r.DB("binance").Table("orderBooks").GetAllByIndex("Symbol", obs.symbol).Changes().Run(dbSession)

	if err != nil {
		obs.err = err
		return nil, err
	}

	obs.updates = make(chan order.Book)

	go func() {
		defer query.Close()
		defer close(obs.updates)
		defer close(obs.done)

		var update struct {
			NewVal *order.Book `gorethink:"new_val"`
			OldVal *order.Book `gorethink:"old_val"`
		}

	Loop:
		for {
			select {
			case <-obs.done:
				break Loop
			default:
				if query.Next(&update) {
					obs.updates <- *update.NewVal
				} else {
					break Loop
				}
			}
		}

		if query.Err() != nil {
			obs.err = query.Err()
		}
	}()

	return obs.updates, nil
}

// Close ubsubscribes for book updates
func (obs *OrderBookSource) Close() {
	obs.done <- struct{}{}
}

// Error returns last thrown error
func (obs *OrderBookSource) Error() error {
	return obs.err
}

// NewOrderBookSource returns new instance
func NewOrderBookSource(symbol string) *OrderBookSource {
	return &OrderBookSource{
		symbol: symbol,
		done:   make(chan struct{}),
	}
}

type orderBookSynch struct {
	symbol string
	stream *binance.DiffDepthStream
	update struct {
		NewVal *binance.DiffDepth `gorethink:"new_val"`
		OldVal *binance.DiffDepth `gorethink:"old_val"`
	}
	close chan struct{}
}

func (obs *orderBookSynch) start() error {
	log.Println("Starting book synch for symbol", obs.symbol)

	var err error
	obs.stream, err = binance.OpenDiffDepthStream(obs.symbol)

	if err != nil {
		return err
	}

	// Start synch updates
	go func() {
		for {
			select {
			case <-obs.close:
				return
			default:
				if !obs.synchUpdate() {
					return
				}
			}
		}
	}()

	return obs.synchBook()
}

func (obs *orderBookSynch) stop() error {
	log.Println("Stopping book synch for symbol", obs.symbol)
	obs.close <- struct{}{}
	return obs.stream.Close()
}

// synchUpdate fetches update from binance server and inserts it in database
func (obs *orderBookSynch) synchUpdate() bool {
	update, err := obs.stream.Read()

	if err != nil {
		log.Printf("Symbol [%s] synch error: %s\n", obs.symbol, err)
		return false
	}

	log.Println("Received order book update#", update.FinalUpdateID)
	err = r.DB(dbName).Table("orderBookUpdates").Insert(update).Exec(dbSession)

	if err != nil {
		log.Printf("Symbol [%s] synch error: %s\n", obs.symbol, err)
		return false
	}

	return true
}

// synchUpdates fetches update from binance server and inserts it in database
func (obs *orderBookSynch) synchBook() error {
	book, err := binance.GetOrderBook(obs.symbol, "1000")

	if err != nil {
		return err
	}

	var count int
	err = r.DB(dbName).Table("orderBooks").GetAllByIndex("Symbol", obs.symbol).Count().ReadOne(&count, dbSession)

	if err != nil {
		return err
	}

	if count == 0 {
		err = r.DB(dbName).Table("orderBooks").Insert(book).Exec(dbSession)
	} else {
		err = r.DB(dbName).Table("orderBooks").GetAllByIndex("Symbol", obs.symbol).Update(book).Exec(dbSession)
	}

	if err != nil {
		return err
	}

	log.Println("Retrieved latest order book #", book.LastUpdateID)

	query, err := r.DB(dbName).Table("orderBookUpdates").Filter(r.Row.Field("Symbol").Eq(obs.symbol)).Changes().Run(dbSession)

	if err != nil {
		return err
	}

	go func() {
		defer query.Close()
		defer obs.stop()

		for {
			if !query.Next(&obs.update) {
				break
			}

			if obs.update.NewVal == nil {
				continue
			}

			log.Println("Processing order book update#", obs.update.NewVal.FinalUpdateID)
			if ok := mergeBookUpdate(book, obs.update.NewVal); !ok {
				log.Println("Ignoring order book update#", obs.update.NewVal.FinalUpdateID)
				continue
			}

			err := r.DB(dbName).Table("orderBooks").GetAllByIndex("Symbol", obs.symbol).Update(book).Exec(dbSession)
			log.Println("Processed order book update#", obs.update.NewVal.FinalUpdateID)

			if err != nil {
				log.Printf("Symbol [%s] synch error: %s\n", obs.symbol, query.Err())
				break
			}
		}

		if query.Err() != nil {
			log.Printf("Symbol [%s] synch error: %s\n", obs.symbol, query.Err())
			return
		}
	}()

	return nil
}

func mergeBookUpdate(book *binance.OrderBook, update *binance.DiffDepth) bool {
	if update.FinalUpdateID <= book.LastUpdateID {
		return false
	}
	book.Asks = mergeOrders(book.Asks, update.Asks)
	book.Bids = mergeOrders(book.Bids, update.Bids)
	book.LastUpdateID = update.FinalUpdateID
	return true
}

func mergeOrders(orders []binance.Order, update []binance.Order) []binance.Order {
	prices := make(map[float64]float64)

	for _, o := range update {
		prices[o.Price] = o.Quantity
	}

	result := make([]binance.Order, 0)

	for _, o := range orders {
		q, ok := prices[o.Price]
		if !ok {
			result = append(result, o)
			continue
		}
		if q == 0 {
			delete(prices, o.Price)
			continue
		}
		o.Quantity = q
		result = append(result, o)
	}

	for p, q := range prices {
		result = append(result, binance.Order{Price: p, Quantity: q})
	}

	return result
}
