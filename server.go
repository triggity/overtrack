package overtrack

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/triggity/overtrack/handlers"
)

func Server(router *mux.Router) {

	routes := []struct {
		Route   string
		Handler http.HandlerFunc
		Name    string
	}{
		{"/", handlers.Home, "home"},
		{"/version", handlers.Version, "version"},
	}
	for _, r := range routes {
		router.HandleFunc(r.Route, prometheus.InstrumentHandlerFunc(r.Name, r.Handler))
	}
	router.Handle("/metrics", promhttp.Handler())
}
