package overtrack

import (
	"net/http"

	"gopkg.in/olivere/elastic.v5"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/triggity/overtrack/handlers"
)

func Server(router *mux.Router, client *elastic.Client) {

	mapsHandler := handlers.NewGameTypesHandler(client)
	userHandler := handlers.NewUserHandler(client)
	gameHandler := handlers.NewGameHandler(client)

	routes := []struct {
		Route   string
		Handler http.HandlerFunc
		Name    string
	}{
		{"/", handlers.Home, "home"},
		{"/version", handlers.Version, "version"},
		{"/v1/maps", mapsHandler.List, "getMaps"},
		{"/v1/maps/{name}", mapsHandler.GetByName, "getMap"},
		{"/v1/users", userHandler.List, "getUsers"},
		{"/v1/users/{id}", userHandler.GetByName, "getUser"},
		{"/v1/users/{id}/games", gameHandler.GetByUser, "getGames"},
	}
	for _, r := range routes {
		router.HandleFunc(r.Route, prometheus.InstrumentHandlerFunc(r.Name, r.Handler))
	}
	router.Handle("/metrics", promhttp.Handler())
}
