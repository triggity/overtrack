package main

import (
	"flag"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"github.com/triggity/overtrack"
)

var (
	address = flag.String("address", ":8000", "address to listen on")
)

func main() {
	flag.Parse()

	r := mux.NewRouter()
	overtrack.Server(r)
	w := log.New().Writer()
	defer w.Close()
	loggedHandler := handlers.LoggingHandler(w, r)

	log.Infof("starting overtrack at %s", *address)
	log.Fatal(http.ListenAndServe(*address, loggedHandler))

}
