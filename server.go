package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/triggity/overtrack/handlers"
)

func Server(router *mux.Router, db *sqlx.DB) {

	mapsHandler := handlers.NewGameTypesHandler(db)
	userHandler := handlers.NewUserHandler(db)
	gameHandler := handlers.NewGameHandler(db)
	herosHandler := handlers.NewHerosHandler(db)

	routes := []struct {
		Route   string
		Handler http.HandlerFunc
		Name    string
	}{
		{"/", handlers.Home, "home"},
		{"/version", handlers.Version, "version"},
		{"/v1/maps", mapsHandler.List, "getMaps"},
		{"/v1/maps/{id}", mapsHandler.GetByName, "getMap"},
		{"/v1/users", userHandler.List, "getUsers"},
		{"/v1/users/{id}", userHandler.GetByName, "getUser"},
		{"/v1/users/{id}/games", gameHandler.GetByUser, "getGames"},
		{"/v1/users/{user}/games/{id}", gameHandler.GetGame, "getGame"},
		{"/v1/heros", herosHandler.List, "getHeros"},
		{"/v1/heros/{id}", herosHandler.GetByID, "getHero"},
	}
	for _, r := range routes {
		router.HandleFunc(r.Route, prometheus.InstrumentHandlerFunc(r.Name, r.Handler))
	}
	router.Handle("/metrics", promhttp.Handler())
}
