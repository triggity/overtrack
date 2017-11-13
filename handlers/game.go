package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/triggity/overtrack/models"
)

type GameHandler struct {
	dao *models.GameController
}

func NewGameHandler(db *sqlx.DB) *GameHandler {
	return &GameHandler{
		models.NewGameController(db),
	}
}

func (g *GameHandler) GetByUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	maps, err := g.dao.GetByUser(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	s, _ := json.Marshal(struct {
		Games []models.Game `json:"games"`
	}{maps})
	w.Write([]byte(s))
}

func (g *GameHandler) GetGame(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user, err := strconv.Atoi(vars["user"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	game, err := g.dao.GetByGame(user, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	out := struct {
		Game models.Game `json:"game"`
	}{game}
	s, _ := json.Marshal(out)
	w.Write([]byte(s))
}
