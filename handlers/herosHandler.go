package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/triggity/overtrack/models"
)

type HerosHandler struct {
	dao *models.HeroDao
}

func NewHerosHandler(db *sqlx.DB) *HerosHandler {
	return &HerosHandler{
		models.NewHeroDao(db),
	}
}

func (h *HerosHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	hero, err := h.dao.GetHero(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	out := struct {
		Hero models.Hero `json:"hero"`
	}{hero}
	s, _ := json.Marshal(out)
	w.Write([]byte(s))
}

func (h *HerosHandler) List(w http.ResponseWriter, r *http.Request) {
	heros, err := h.dao.List()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	out := struct {
		Heros []models.Hero `json:"heros"`
	}{heros}
	s, _ := json.Marshal(out)
	w.Write([]byte(s))
}
