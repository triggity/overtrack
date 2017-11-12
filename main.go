package main

import (
	"flag"
	"net/http"

	"github.com/jmoiron/sqlx"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

var (
	address = flag.String("address", ":8000", "address to listen on")
)

func main() {
	flag.Parse()
	db, err := sqlx.Open("postgres", "string")
	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	Server(r, nil, db)
	w := log.New().Writer()
	defer w.Close()
	loggedHandler := handlers.LoggingHandler(w, r)

	log.Infof("starting overtrack at %s", *address)
	log.Fatal(http.ListenAndServe(*address, loggedHandler))

}
