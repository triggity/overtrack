package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/triggity/overtrack"
)

var (
	a = flag.String("a", "d", "e")
)

func main() {
	flag.Parse()
	fmt.Printf("hello there! %s", *a)
	r := mux.NewRouter()
	overtrack.Server(r)
	log.Fatal(http.ListenAndServe(":8000", r))

}
