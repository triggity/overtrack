package main

import (
	"flag"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/triggity/overtrack"
	"gopkg.in/olivere/elastic.v5"
)

var (
	address = flag.String("address", ":8000", "address to listen on")
)

func main() {
	flag.Parse()
	client, err := elastic.NewClient(
		elastic.SetURL("http://0.0.0.0:9200"),
		elastic.SetBasicAuth("elastic", "changeme"),
		elastic.SetSniff(false),
	)
	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	overtrack.Server(r, client)
	w := log.New().Writer()
	defer w.Close()
	loggedHandler := handlers.LoggingHandler(w, r)

	log.Infof("starting overtrack at %s", *address)
	log.Fatal(http.ListenAndServe(*address, loggedHandler))

}
