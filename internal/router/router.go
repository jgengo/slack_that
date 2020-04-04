package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// New ...
func New() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		// handler = handlers.CombinedLoggingHandler(os.Stdout, handler)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	log.Println("http (info) listening...")
	return router
}
