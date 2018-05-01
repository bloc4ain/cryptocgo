package binance

import (
	r "gopkg.in/gorethink/gorethink.v4"
)

const dbName = "binance"
const dbAddress = "localhost:28015"

var tables = [...]string{
	"orderBooks",
	"orderBookUpdates",
}

var indexes = map[string][]string{
	"orderBooks": {"Symbol"},
}

var dbSession *r.Session

// Connect to database
func init() {
	var err error
	dbSession, err = r.Connect(r.ConnectOpts{
		Address: dbAddress,
	})
	if err != nil {
		panic(err)
	}
}

// Create missing tables
func init() {
	list := make([]string, 0)
	err := r.Expr(tables).Difference(r.DB(dbName).TableList()).ReadAll(&list, dbSession)

	if err != nil {
		panic(err)
	}

	for _, t := range list {
		err := r.DB(dbName).TableCreate(t).Exec(dbSession)
		if err != nil {
			panic(err)
		}
	}
}

// Drop orphan tables
func init() {
	list := make([]string, 0)
	err := r.DB(dbName).TableList().Difference(tables).ReadAll(&list, dbSession)

	if err != nil {
		panic(err)
	}

	for _, t := range list {
		err := r.DB(dbName).TableDrop(t).Exec(dbSession)
		if err != nil {
			panic(err)
		}
	}
}

// Create indexes
func init() {
	for t, is := range indexes {
		list := make([]string, 0)
		err := r.Expr(is).Difference(r.DB(dbName).Table(t).IndexList()).ReadAll(&list, dbSession)

		if err != nil {
			panic(err)
		}

		for _, i := range list {
			if err := r.DB(dbName).Table(t).IndexCreate(i).Exec(dbSession); err != nil {
				panic(err)
			}
			if err := r.DB(dbName).Table(t).IndexWait().Exec(dbSession); err != nil {
				panic(err)
			}
		}
	}
}
