package webui

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/bloc4ain/cryptocgo/webui/routing"
)

type indexPage struct {
	Title string
	Body  string
}

var host = flag.String("webui_host", "localhost", "Host where the webui will listen on")
var port = flag.String("webui_port", "4040", "Post where the webui will listen on")

// Start runs webui http server
func Start() {
	flag.Parse()
	address := fmt.Sprintf("%s:%s", *host, *port)
	router := routing.NewRouter()
	log.Fatal(http.ListenAndServe(address, router))
}
