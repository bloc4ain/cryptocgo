package market

// Card represents cryptocurrency market brief info
type Card interface {
	Title() string
	Top() []string
	Updates() <-chan []string
}

// BinanceCard represents Binance market brief info
type BinanceCard struct {
}

// Title returns Binance title
func (c BinanceCard) Title() string {
	return "Binance"
}

// Top returns Binance top most expensive currencies
func (c BinanceCard) Top() []string {
	return []string{
		"TRX/BTC",
		"ETH/BTC",
		"ONT/BTC",
	}
}

// Updates subscribes for top most expensive currencies updates
func (c BinanceCard) Updates() <-chan []string {
	return make(chan []string)
}

// BittrexCard represents Bittrex market brief info
type BittrexCard struct {
}

// Title returns Bittrex title
func (c BittrexCard) Title() string {
	return "Bittrex"
}

// Top returns Bittrex top most expensive currencies
func (c BittrexCard) Top() []string {
	return []string{
		"TRX/BTC",
		"ETH/BTC",
		"ONT/BTC",
	}
}

// Updates subscribes for top most expensive currencies updates
func (c BittrexCard) Updates() <-chan []string {
	return make(chan []string)
}

// KuCoinCard represents KuCoin market brief info
type KuCoinCard struct {
}

// Title returns KuCoin title
func (c KuCoinCard) Title() string {
	return "KuCoin"
}

// Top returns KuCoin top most expensive currencies
func (c KuCoinCard) Top() []string {
	return []string{
		"TRX/BTC",
		"ETH/BTC",
		"ONT/BTC",
	}
}

// Updates subscribes for top most expensive currencies updates
func (c KuCoinCard) Updates() <-chan []string {
	return make(chan []string)
}
