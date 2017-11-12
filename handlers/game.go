package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"gopkg.in/olivere/elastic.v5"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/triggity/overtrack/models"
)

type GameHandler struct {
	dao *models.GameDao
}

func NewGameHandler(client *elastic.Client, db *sqlx.DB) *GameHandler {
	return &GameHandler{
		models.NewGameDao(client, db),
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
	maps, err := g.dao.GetByUser(r.Context(), id)
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
