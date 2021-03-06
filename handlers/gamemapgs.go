package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"github.com/triggity/overtrack/models"
)

type GameMapsHandler struct {
	dao *models.GameMapController
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
	out := struct {
		Maps []models.GameMap `json:"maps"`
	}{maps}
	s, _ := json.Marshal(out)
	w.Write([]byte(s))
}

func (g *GameMapsHandler) GetByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	maps, err := g.dao.GetByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	out := struct {
		Map models.GameMap `json:"map"`
	}{maps}
	s, _ := json.Marshal(out)
	w.Write([]byte(s))
}
