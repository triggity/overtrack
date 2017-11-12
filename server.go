package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gopkg.in/olivere/elastic.v5"

	"github.com/triggity/overtrack/handlers"
)

func Server(router *mux.Router, client *elastic.Client, db *sqlx.DB) {

	mapsHandler := handlers.NewGameTypesHandler(client, db)
	userHandler := handlers.NewUserHandler(client, db)
	gameHandler := handlers.NewGameHandler(client, db)

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
