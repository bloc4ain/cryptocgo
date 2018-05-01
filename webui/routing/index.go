package routing

import (
	"html/template"
	"log"
	"net/http"

	"github.com/bloc4ain/cryptocgo/order"

	"github.com/bloc4ain/cryptocgo/binance"
	"github.com/bloc4ain/cryptocgo/webui/market"
	"github.com/gorilla/websocket"
)

type indexPage struct {
	Books []*order.Book
}

var cards = []market.Card{
	market.BinanceCard{},
	market.BittrexCard{},
	market.KuCoinCard{},
}

var upgrader = &websocket.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("webui/templates/index.gohtml")

	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	book, err := binance.NewOrderBookSource("TRXBTC").Book()

	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	page := indexPage{Books: make([]*order.Book, 0)}
	page.Books = append(page.Books, book)

	t.Execute(w, &page)
}

func ws(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	go func() {
		src := binance.NewOrderBookSource("TRXBTC")
		stream, err := src.Updates()
		if err != nil {
			return
		}
		for change := range stream {
			err := ws.WriteJSON(change)
			if err != nil {
				break
			}
		}
		ws.Close()
	}()
}

func wsHandler() http.HandlerFunc {
	log.Println("Starting websocket server")
	return func(w http.ResponseWriter, r *http.Request) {
		ws, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}

		src := binance.NewOrderBookSource("TRXBTC")
		stream, err := src.Updates()
		defer ws.Close()

		if err != nil {
			return
		}

		for change := range stream {
			err := ws.WriteJSON(change)
			if err != nil {
				break
			}
		}
	}
}
