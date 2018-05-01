package routing

import (
	"html/template"
	"net/http"

	"github.com/bloc4ain/cryptocgo/webui/market"
)

type indexPage struct {
	MarketCards []market.Card
}

var cards = []market.Card{
	market.BinanceCard{},
	market.BittrexCard{},
	market.KuCoinCard{},
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("webui/templates/index.gohtml")

	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	t.Execute(w, &indexPage{cards})
}
