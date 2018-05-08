package storage

import (
	r "gopkg.in/gorethink/gorethink.v4"
)

const dbName = "binance"
const dbAddress = "localhost:28015"

var dbSession *r.Session
