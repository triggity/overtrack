package overtrack

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Server(router *mux.Router) {
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello!"))
	})
}
