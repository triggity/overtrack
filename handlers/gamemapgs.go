package handlers

import (
	"encoding/json"
	"net/http"

	"gopkg.in/olivere/elastic.v5"

	"github.com/gorilla/mux"
	"github.com/triggity/overtrack/models"
)

type GameMapsHandler struct {
	dao models.GameMapDao
}

func NewGameTypesHandler(client *elastic.Client) *GameMapsHandler {
	return &GameMapsHandler{
		models.NewGameMapDao(client),
	}
}

func (g *GameMapsHandler) List(w http.ResponseWriter, r *http.Request) {
	maps, err := g.dao.List(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	s, _ := json.Marshal(maps)
	w.Write([]byte(s))
}

func (g *GameMapsHandler) GetByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	maps, err := g.dao.GetByName(r.Context(), vars["name"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	s, _ := json.Marshal(maps)
	w.Write([]byte(s))
}
