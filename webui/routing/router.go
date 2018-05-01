package routing

import (
	"net/http"

	"github.com/gorilla/mux"
)

// NewRouter returns webgui router
func NewRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", indexHandler)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("webui/static")))

	// r.HandleFunc("/active", activeIndexHandler)
	// r.HandleFunc("/completed", completedIndexHandler)
	// r.HandleFunc("/new", newHandler)
	// r.HandleFunc("/toggle/{id}", toggleHandler)
	// r.HandleFunc("/delete/{id}", deleteHandler)
	// r.HandleFunc("/clear", clearHandler)

	// // Add handler for websocket server
	// r.Handle("/ws/all", newChangesHandler(allChanges))
	// r.Handle("/ws/active", newChangesHandler(activeChanges))
	// r.Handle("/ws/completed", newChangesHandler(completedChanges))

	// Add handler for static files

	return r
}
