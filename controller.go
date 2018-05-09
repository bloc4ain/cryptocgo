package cryptocgo

type Controller interface {
	SyncSymbol(...Symbol)
	SyncMarket(string, ...Symbol)
}
