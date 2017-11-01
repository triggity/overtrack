package overtrack

import (
	"github.com/gorilla/mux"

	"github.com/triggity/overtrack/handlers"
)

func Server(router *mux.Router) {

	router.HandleFunc("/", handlers.Home)
	router.HandleFunc("/version", handlers.Version)
}
