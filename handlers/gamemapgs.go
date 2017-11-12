package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"github.com/triggity/overtrack/models"
)

type GameMapsHandler struct {
	dao *models.GameMapDao
}

func NewGameTypesHandler(db *sqlx.DB) *GameMapsHandler {
	return &GameMapsHandler{
		models.NewGameMapDao(db),
	}
}

func (g *GameMapsHandler) List(w http.ResponseWriter, r *http.Request) {
	maps, err := g.dao.List()
	if err != nil {
		log.Info("failed to list maps ", err)
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
	maps, err := g.dao.GetByName(vars["name"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	s, _ := json.Marshal(maps)
	w.Write([]byte(s))
}
