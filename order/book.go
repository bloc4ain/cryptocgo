package order

// Source is used to load book order data
type Source interface {
	// Book returns current book value
	Book() (book Book, err error)

	// Updates subscribes for book updates
	Updates() (<-chan Book, error)

	// Close ubsubscribes for book updates
	Close()

	// Error returns last thrown error
	Error() error
}

// Order represent market order
type Order struct {
	Price    float64
	Quantity float64
}

func (o Order) String() {

}

// Book represents all market orders for single symbol
type Book struct {
	Symbol string
	Bids   []Order
	Asks   []Order
}

func (b Book) String() {

}
